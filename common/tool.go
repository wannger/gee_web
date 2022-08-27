package common

import "strings"

// url解析成数组，"*"、":"也可以

func ParsePattern(pattern string) []string {
	newPattern := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range newPattern {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}
