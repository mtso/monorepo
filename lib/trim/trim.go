package trim

import (
	"strings"
)

func trim(str string) string {
	parts := strings.Split(str, "\n")
	if parts[0] == "" {
		parts = parts[1:]
	}
	if parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}
	return strings.Join(parts, "\n")
}
