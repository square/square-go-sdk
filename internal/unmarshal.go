package internal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// PermissiveUnmarshal unmarshals the JSON data into the provided interface using the native reflect package.
// If the input contains unassignable values, the corresponding input field will be ignored. Other
// errors will be returned. It also sets fields tagged with fern:"extra" or fern:"raw" on the struct.
// TODO: Log errors instead of silently ignoring them.
func PermissiveUnmarshal(data []byte, out interface{}) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("out must be a non-nil pointer")
	}
	v = v.Elem()
	t := v.Type()

	if t.Kind() != reflect.Struct {
		// fallback to standard unmarshal for non-structs
		return json.Unmarshal(data, out)
	}

	if err := populateStructFields(data, v, t); err != nil {
		return err
	}
	if err := setFernFields(data, v, t); err != nil {
		return err
	}
	return nil
}

// populateStructFields sets struct fields from JSON using permissive logic.
func populateStructFields(data []byte, v reflect.Value, t reflect.Type) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}
		// Ignore json tags after the first comma (e.g. omitempty)
		jsonTags := strings.Split(jsonTag, ",")
		key := jsonTags[0]
		val, ok := raw[key]
		if !ok {
			continue
		}
		fieldVal := v.Field(i)
		if !fieldVal.CanSet() {
			continue
		}
		handleField(val, fieldVal)
	}
	return nil
}

// setFernFields sets fields tagged with fern:"extra" or fern:"raw" on the struct v of type t.
// "extra" is used to store extra properties that are not part of the JSON schema.
// "raw" is used to store the raw JSON of data.
func setFernFields(data []byte, v reflect.Value, t reflect.Type) error {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fernTag := field.Tag.Get("fern")
		if fernTag == "" {
			continue // Skip fields without a fern tag
		}
		fieldVal := v.Field(i)
		if !fieldVal.CanSet() {
			continue // Skip unsettable fields (e.g., unexported)
		}
		switch fernTag {
		case "extra":
			// Only set if the field is a map
			if fieldVal.Kind() != reflect.Map {
				continue
			}
			extra, err := ExtractExtraPropertiesUnmarshalled(data, v.Interface())
			if err != nil {
				return err
			}
			if extra == nil {
				fieldVal.Set(reflect.Zero(fieldVal.Type()))
			} else {
				fieldVal.Set(reflect.ValueOf(extra))
			}
		case "raw":
			// Only set if the field is a json.RawMessage
			if fieldVal.Type() != reflect.TypeOf(json.RawMessage{}) {
				continue
			}
			fieldVal.Set(reflect.ValueOf(json.RawMessage(data)))
		}
	}
	return nil
}

// handleField sets the value v from the provided JSON data, handling structs, pointers, slices, and maps recursively.
func handleField(data json.RawMessage, v reflect.Value) error {
	t := v.Type()

	switch t.Kind() {
	case reflect.Struct:
		// Recursively unmarshal into struct fields.
		return PermissiveUnmarshal(data, v.Addr().Interface())
	case reflect.Ptr:
		// Allocate and set the pointed-to value.
		elemType := t.Elem()
		ptr := reflect.New(elemType)
		if err := handleField(data, ptr.Elem()); err == nil {
			v.Set(ptr)
		}
	case reflect.Slice:
		// Decode each element recursively.
		elemType := t.Elem()
		var rawSlice []json.RawMessage
		_ = json.Unmarshal(data, &rawSlice)
		slice := reflect.MakeSlice(t, 0, len(rawSlice))
		for _, elemData := range rawSlice {
			elem := reflect.New(elemType).Elem()
			if err := handleField(elemData, elem); err == nil {
				slice = reflect.Append(slice, elem)
			}
		}
		v.Set(slice)
	case reflect.Map:
		// Decode each value recursively.
		keyType := t.Key()
		elemType := t.Elem()
		var rawMap map[string]json.RawMessage
		_ = json.Unmarshal(data, &rawMap)
		m := reflect.MakeMap(t)
		for k, elemData := range rawMap {
			key := reflect.ValueOf(k).Convert(keyType)
			elem := reflect.New(elemType).Elem()
			if err := handleField(elemData, elem); err == nil {
				// Skip invalid map values.
				m.SetMapIndex(key, elem)
			}
		}
		v.Set(m)
	default:
		// Try to unmarshal directly into the value.
		ptr := reflect.New(t)
		if err := json.Unmarshal(data, ptr.Interface()); err == nil {
			v.Set(ptr.Elem())
		}
	}
	return nil
}
