package tool

import (
	"fmt"
	"net/url"
)

// FlattenData interface to url.Values
func FlattenData(data interface{}) url.Values {
	result := make(url.Values)
	flattenRecursively(result, "", data)
	return result
}

func flattenRecursively(result url.Values, prefix string, data interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, val := range v {
			newKey := key
			if prefix != "" {
				newKey = fmt.Sprintf("%s.%s", prefix, key)
			}
			flattenRecursively(result, newKey, val)
		}
	case []interface{}:
		for i, val := range v {
			newKey := fmt.Sprintf("%s.%d", prefix, i)
			flattenRecursively(result, newKey, val)
		}
	default:
		result[prefix] = []string{fmt.Sprintf("%v", data)}
	}
}
