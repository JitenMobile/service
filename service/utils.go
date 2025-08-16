package service

import (
	"reflect"
	"strings"
)

// ------ schema helpers ------

func JsonTypeOf(t reflect.Type) any {
	switch t.Kind() {
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "boolean"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return "number"
	case reflect.Slice, reflect.Array:
		return map[string]any{
			"type":  "array",
			"items": map[string]any{"type": JsonTypeOf(t.Elem())},
		}
	case reflect.Struct:
		props, req := StructProperties(t)
		return map[string]any{
			"type":       "object",
			"properties": props,
			"required":   req,
		}
	default:
		return "string"
	}
}

func StructProperties(t reflect.Type) (map[string]any, []string) {
	props := make(map[string]any)
	var required []string

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if f.PkgPath != "" {
			continue
		}

		tag := f.Tag.Get("json")
		if tag == "-" {
			continue
		}

		name := f.Name
		if tag != "" {
			parts := strings.Split(tag, ",")
			if parts[0] != "" {
				name = parts[0]
			}

			om := false
			for _, p := range parts[1:] {
				if p == "omitempty" {
					om = true
					break
				}
			}
			if !om {
				required = append(required, name)
			}
		} else {
			required = append(required, name)
		}
		props[name] = map[string]any{"type": JsonTypeOf(f.Type)}
	}

	return props, required
}
