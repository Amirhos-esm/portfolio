package graph

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestMarshalStringArrayMap(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string][]string
		expected string
	}{
		{
			name: "valid map",
			input: map[string][]string{
				"tech": {"go", "graphql"},
				"soft": {"communication"},
			},
			expected: `{"soft":["communication"],"tech":["go","graphql"]}`,
		},
		{
			name:     "nil map",
			input:    nil,
			expected: "null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			MarshalStringArrayMap(tt.input).MarshalGQL(&buf)

			// Since JSON map order is not guaranteed,
			// compare using json decoding
			if tt.expected == "null" {
				if buf.String() != "null" {
					t.Fatalf("expected null, got %s", buf.String())
				}
				return
			}

			var gotMap map[string][]string
			var expectedMap map[string][]string

			if err := json.Unmarshal(buf.Bytes(), &gotMap); err != nil {
				t.Fatalf("failed to unmarshal result: %v", err)
			}

			if err := json.Unmarshal([]byte(tt.expected), &expectedMap); err != nil {
				t.Fatalf("failed to unmarshal expected: %v", err)
			}

			if !reflect.DeepEqual(gotMap, expectedMap) {
				t.Fatalf("expected %v, got %v", expectedMap, gotMap)
			}
		})
	}
}
