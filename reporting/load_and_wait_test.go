package reporting

import (
	context "context"
	json "encoding/json"
	http "net/http"
	httptest "net/http/httptest"
	sync "sync"
	testing "testing"
	time "time"

	square "github.com/square/square-go-sdk/v3"
	core "github.com/square/square-go-sdk/v3/core"
	option "github.com/square/square-go-sdk/v3/option"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

// The Reporting API answers a still-processing `/reporting/v1/load` query with an
// HTTP 200 whose body is `{"error": "Continue wait"}`. (*Client).LoadAndWait owns
// the retry loop around that sentinel. These tests drive the loop against a local
// server that scripts the load responses, which also proves the sentinel survives
// the generated client's deserialization (rather than being mistaken for a result).

const (
	continueWaitBody = `{"error":"Continue wait"}`
	resolvedBody     = `{"results":[{}]}`
)

// scriptedServer serves the i-th body in bodies for each /reporting/v1/load call
// (clamping to the last), and returns a func reporting how many calls it received.
func scriptedServer(t *testing.T, bodies []string) (*httptest.Server, func() int) {
	t.Helper()
	var mu sync.Mutex
	var calls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		body := bodies[min(calls, len(bodies)-1)]
		calls++
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)
	return server, func() int {
		mu.Lock()
		defer mu.Unlock()
		return calls
	}
}

func testClient(serverURL string) *Client {
	return NewClient(core.NewRequestOptions(
		option.WithBaseURL(serverURL),
		option.WithToken("test-token"),
	))
}

// fastPoll keeps backoff negligible so attempt-counting tests stay quick.
func fastPoll(maxAttempts int) *LoadAndWaitOptions {
	return &LoadAndWaitOptions{
		MaxAttempts:   maxAttempts,
		InitialDelay:  time.Millisecond,
		MaxDelay:      time.Millisecond,
		BackoffFactor: 1,
	}
}

func TestLoadAndWait_PollsPastContinueWait(t *testing.T) {
	server, calls := scriptedServer(t, []string{continueWaitBody, continueWaitBody, resolvedBody})
	client := testClient(server.URL)

	response, err := client.LoadAndWait(context.Background(), &square.LoadRequest{}, fastPoll(5))
	require.NoError(t, err)
	// The helper must never hand back the raw sentinel.
	_, hasError := response.GetExtraProperties()["error"]
	assert.False(t, hasError)
	assert.NotEmpty(t, response.GetResults())
	assert.Equal(t, 3, calls())
}

func TestLoadAndWait_ReturnsImmediatelyWhenResolved(t *testing.T) {
	server, calls := scriptedServer(t, []string{resolvedBody})
	client := testClient(server.URL)

	response, err := client.LoadAndWait(context.Background(), &square.LoadRequest{}, nil)
	require.NoError(t, err)
	assert.NotEmpty(t, response.GetResults())
	assert.Equal(t, 1, calls())
}

func TestLoadAndWait_ExhaustsMaxAttempts(t *testing.T) {
	server, calls := scriptedServer(t, []string{continueWaitBody}) // never resolves
	client := testClient(server.URL)

	_, err := client.LoadAndWait(context.Background(), &square.LoadRequest{}, fastPoll(3))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "did not complete after 3 attempts")
	assert.Equal(t, 3, calls())
}

func TestLoadAndWait_StopsOnContextCancellation(t *testing.T) {
	server, _ := scriptedServer(t, []string{continueWaitBody}) // would otherwise poll forever
	client := testClient(server.URL)

	// The first load returns quickly; the deadline then fires during the backoff wait.
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	options := &LoadAndWaitOptions{MaxAttempts: 10, InitialDelay: time.Second, MaxDelay: time.Second, BackoffFactor: 1}

	_, err := client.LoadAndWait(ctx, &square.LoadRequest{}, options)
	require.Error(t, err)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

// TestIsContinueWait_SurvivesDeserialization is the crux of the design: the
// generated LoadResponse keeps unrecognized keys in ExtraProperties, so the
// `error` sentinel lands there (and Results stays empty) instead of being
// dropped. If that ever changes, LoadAndWait would mistake "Continue wait" for a
// real result and return it instead of polling.
func TestIsContinueWait_SurvivesDeserialization(t *testing.T) {
	var pending square.LoadResponse
	require.NoError(t, json.Unmarshal([]byte(continueWaitBody), &pending))
	assert.True(t, isContinueWait(&pending))
	assert.Empty(t, pending.GetResults())

	var resolved square.LoadResponse
	require.NoError(t, json.Unmarshal([]byte(resolvedBody), &resolved))
	assert.False(t, isContinueWait(&resolved))
	assert.NotEmpty(t, resolved.GetResults())

	assert.False(t, isContinueWait(nil))
}
