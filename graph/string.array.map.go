package graph

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalStringArrayMap(m map[string][]string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		if m == nil {
			w.Write([]byte("null"))
			return
		}

		data, err := json.Marshal(m)
		if err != nil {
			// GraphQL expects valid JSON output
			w.Write([]byte("null"))
			return
		}

		w.Write(data)
	})
}

func UnmarshalStringArrayMap(v interface{}) (map[string][]string, error) {
	if v == nil {
		return nil, nil
	}

	input, ok := v.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected map[string]interface{}, got %T", v)
	}

	result := make(map[string][]string, len(input))

	for key, value := range input {
		rawSlice, ok := value.([]interface{})
		if !ok {
			return nil, fmt.Errorf("expected array for key %s, got %T", key, value)
		}

		strSlice := make([]string, len(rawSlice))
		for i, elem := range rawSlice {
			str, ok := elem.(string)
			if !ok {
				return nil, fmt.Errorf("expected string in array for key %s, got %T", key, elem)
			}
			strSlice[i] = str
		}

		result[key] = strSlice
	}

	return result, nil
}
