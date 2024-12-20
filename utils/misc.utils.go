package utils

import "fmt"

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func Keys(obj map[string]any) []string {
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}

	return keys
}

func Equal(a, b []any) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", b[i]) {
			return false
		}
	}
	return true
}
