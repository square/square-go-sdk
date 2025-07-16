package unmarshal

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

// UnmarshalMapstructure unmarshals the JSON data into the provided interface using the mapstructure package.
// It makes an attempt with weakly typed input to convert strings to numbers, booleans, etc.
// See https://pkg.go.dev/github.com/mitchellh/mapstructure#DecoderConfig
func UnmarshalMapstructure(data []byte, out interface{}) error {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(data), &jsonMap)
	if err != nil {
		return err
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           out,
		TagName:          "json",
	})
	if err != nil {
		return err
	}
	if err = decoder.Decode(jsonMap); err != nil {
		return err
	}
	return nil
}
