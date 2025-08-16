package service

import (
	"reflect"
	"testing"
)

// Test helper structs
type Definition struct {
	ID           string   `json:"id"`
	Examples     []string `json:"examples"`
	Meaning      string   `json:"meaning"`
	PartOfSpeech string   `json:"partOfSpeech"`
}

func TestJsonTypeOf(t *testing.T) {
	tests := []struct {
		name     string
		input    reflect.Type
		expected interface{}
	}{
		{
			name:     "string type",
			input:    reflect.TypeOf(""),
			expected: "string",
		},
		{
			name:  "array type",
			input: reflect.TypeOf([3]string{}),
			expected: map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "string",
				},
			},
		},
	}

	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			result := JsonTypeOf(item.input)
			if !reflect.DeepEqual(result, item.expected) {
				t.Errorf("JsonTypeOf() = %v, want %v", result, item.expected)
			}
		})
	}
}
