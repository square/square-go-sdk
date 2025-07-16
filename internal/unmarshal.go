package internal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// PermissiveUnmarshal unmarshals the JSON data into the provided interface using the native reflect package.
// If the input contains unassignable values, the corresponding input field will be ignored. Other
// errors will be returned.
func PermissiveUnmarshal(data []byte, out interface{}) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("out must be a non-nil pointer")
	}
	v = v.Elem()
	t := v.Type()

	switch t.Kind() {
	case reflect.Struct:
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

			// TODO: Parse custom tags to enforce strict validation of input values
			// Ignore json tags after the first comma (e.g. omitempty)
			jsonTags := strings.Split(jsonTag, ",")
			key := jsonTags[0]
			if val, ok := raw[key]; ok {
				fieldVal := v.Field(i)
				if !fieldVal.CanSet() {
					continue
				}
				handleField(val, fieldVal)
			}
		}

		// Set fern:extra and fern:raw fields if present
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fernTag := field.Tag.Get("fern")
			if fernTag == "extra" {
				fieldVal := v.Field(i)
				if fieldVal.CanSet() && fieldVal.Kind() == reflect.Map {
					extra, err := ExtractExtraPropertiesUnmarshalled(data, v.Interface())
					if err == nil {
						if extra == nil {
							// Set to nil explicitly
							fieldVal.Set(reflect.Zero(fieldVal.Type()))
						} else {
							fieldVal.Set(reflect.ValueOf(extra))
						}
					}
				}
			} else if fernTag == "raw" {
				fieldVal := v.Field(i)
				if fieldVal.CanSet() && fieldVal.Type() == reflect.TypeOf(json.RawMessage{}) {
					fieldVal.Set(reflect.ValueOf(json.RawMessage(data)))
				}
			}
		}
	default:
		// fallback to standard unmarshal for non-structs
		return json.Unmarshal(data, out)
	}
	return nil
}

func handleField(data json.RawMessage, v reflect.Value) error {
	t := v.Type()
	switch t.Kind() {
	case reflect.Struct:
		// Pass the correct JSON for this field
		PermissiveUnmarshal(data, v.Addr().Interface())
	case reflect.Ptr:
		elemType := t.Elem()
		ptr := reflect.New(elemType)
		handleField(data, ptr.Elem())
		v.Set(ptr)
	case reflect.Slice:
		elemType := t.Elem()
		var rawSlice []json.RawMessage
		if err := json.Unmarshal(data, &rawSlice); err == nil {
			slice := reflect.MakeSlice(t, 0, len(rawSlice))
			for _, elemData := range rawSlice {
				elem := reflect.New(elemType).Elem()
				handleField(elemData, elem)
				slice = reflect.Append(slice, elem)
			}
			v.Set(slice)
		}
	case reflect.Map:
		keyType := t.Key()
		elemType := t.Elem()
		var rawMap map[string]json.RawMessage
		if err := json.Unmarshal(data, &rawMap); err == nil {
			m := reflect.MakeMap(t)
			for k, elemData := range rawMap {
				key := reflect.ValueOf(k).Convert(keyType)
				elem := reflect.New(elemType).Elem()
				handleField(elemData, elem)
				m.SetMapIndex(key, elem)
			}
			v.Set(m)
		}
	default:
		ptr := reflect.New(t)
		if err := json.Unmarshal(data, ptr.Interface()); err == nil {
			v.Set(ptr.Elem())
		}
	}
	return nil
}
