package util

import (
	"fmt"
	"net/url"
)

func MapToUrlValues(values *url.Values, prefix string, data map[string]any) {
	keyFormater := "%s"

	if prefix != "" {
		keyFormater = prefix + "[%s]"
	}

	for key, value := range data {
		keyValue := fmt.Sprintf(keyFormater, key)

		switch v := value.(type) {
		case string:
			values.Add(keyValue, v)
		case []string:
		case []int:
		case []int32:
		case []int64:
		case []any:
			for _, s := range v {
				values.Add(keyValue, fmt.Sprintf("%v", s))
			}
		case int, int64, float64, bool:
			values.Add(keyValue, fmt.Sprintf("%v", v))
		case map[string]any:
			MapToUrlValues(values, keyValue, v)
		default:
		}
	}
}
