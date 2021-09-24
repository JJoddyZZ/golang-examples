package somemap

import "strings"

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, sub := range strings.Fields(s) {
		if _, ok := m[sub]; ok {
			m[sub] += 1
			continue
		}
		m[sub] = 1
	}
	return m
}
