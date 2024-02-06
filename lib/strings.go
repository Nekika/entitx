package lib

import (
	"fmt"
	"strings"
)

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	first := strings.ToUpper(s[:1])
	rest := s[1:]
	return fmt.Sprintf("%s%s", first, rest)
}
