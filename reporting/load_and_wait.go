package reporting

import (
	context "context"
	fmt "fmt"
	square "github.com/square/square-go-sdk/v4"
	option "github.com/square/square-go-sdk/v4/option"
	time "time"
)

// continueWaitSentinel is the value of the `error` field the Reporting API
// returns on an HTTP 200 while a `/reporting/v1/load` query is still processing.
// It is NOT a failure — the identical request should be re-sent until real
// results arrive. See https://developer.squareup.com/docs/reporting-api/overview.
const continueWaitSentinel = "Continue wait"

// Default polling parameters for LoadAndWait.
const (
	defaultMaxAttempts   = 20
	defaultInitialDelay  = 2 * time.Second
	defaultMaxDelay      = 20 * time.Second
	defaultBackoffFactor = 2.0
)

// LoadAndWaitOptions configures the polling loop in LoadAndWait. The zero value
// is valid: any field left unset falls back to its default.
type LoadAndWaitOptions struct {
	// MaxAttempts is the maximum number of load calls before giving up. Default 20.
	MaxAttempts int
	// InitialDelay is the delay before the first retry. Default 2s.
	InitialDelay time.Duration
	// MaxDelay is the upper bound on the backoff delay. Default 20s.
	MaxDelay time.Duration
	// BackoffFactor multiplies the delay after each attempt. Default 2.
	BackoffFactor float64
}

// LoadAndWait runs a reporting query and transparently polls until it resolves,
// returning the final *square.LoadResponse.
//
// The `/reporting/v1/load` endpoint is asynchronous: a query that is still being
// computed comes back as an HTTP 200 whose body is `{"error": "Continue wait"}`
// rather than the results. Callers are expected to re-send the identical request,
// with backoff, until real results arrive. LoadAndWait owns that retry loop.
//
// pollOptions tunes the backoff (pass nil for defaults); opts are forwarded to
// each underlying Load call. Polling stops early — returning ctx.Err() — if ctx
// is cancelled or its deadline expires. It returns an error if the query has not
// resolved within the configured maximum number of attempts.
func (c *Client) LoadAndWait(
	ctx context.Context,
	request *square.LoadRequest,
	pollOptions *LoadAndWaitOptions,
	opts ...option.RequestOption,
) (*square.LoadResponse, error) {
	settings := resolveLoadAndWaitOptions(pollOptions)
	delay := settings.InitialDelay
	for attempt := 1; attempt <= settings.MaxAttempts; attempt++ {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		response, err := c.Load(ctx, request, opts...)
		if err != nil {
			return nil, err
		}
		if !isContinueWait(response) {
			return response, nil
		}
		if attempt == settings.MaxAttempts {
			break
		}
		if err := sleep(ctx, delay); err != nil {
			return nil, err
		}
		delay = time.Duration(float64(delay) * settings.BackoffFactor)
		if delay > settings.MaxDelay {
			delay = settings.MaxDelay
		}
	}
	return nil, fmt.Errorf(
		"reporting query did not complete after %d attempts (%q)",
		settings.MaxAttempts,
		continueWaitSentinel,
	)
}

// isContinueWait reports whether response is the Reporting API's "still
// processing" sentinel rather than a real result. A `{"error": "Continue wait"}`
// body deserializes into a *square.LoadResponse whose `error` key lands in
// ExtraProperties (it is not a declared field) while the flat `data` field
// stays empty; that is the signal to retry.
func isContinueWait(response *square.LoadResponse) bool {
	if response == nil {
		return false
	}
	value, ok := response.GetExtraProperties()["error"].(string)
	return ok && value == continueWaitSentinel
}

// sleep waits for d, returning early with ctx.Err() if ctx is cancelled first.
func sleep(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// resolveLoadAndWaitOptions applies defaults for any unset (non-positive) field.
func resolveLoadAndWaitOptions(options *LoadAndWaitOptions) LoadAndWaitOptions {
	resolved := LoadAndWaitOptions{
		MaxAttempts:   defaultMaxAttempts,
		InitialDelay:  defaultInitialDelay,
		MaxDelay:      defaultMaxDelay,
		BackoffFactor: defaultBackoffFactor,
	}
	if options == nil {
		return resolved
	}
	if options.MaxAttempts > 0 {
		resolved.MaxAttempts = options.MaxAttempts
	}
	if options.InitialDelay > 0 {
		resolved.InitialDelay = options.InitialDelay
	}
	if options.MaxDelay > 0 {
		resolved.MaxDelay = options.MaxDelay
	}
	if options.BackoffFactor > 0 {
		resolved.BackoffFactor = options.BackoffFactor
	}
	return resolved
}
