package pkg

import (
	"regexp"
	"strings"
)

func Capitalize(s string) string {
	re := regexp.MustCompile(`(^.|\s.)`)

	return re.ReplaceAllStringFunc(
		s,
		func(m string) string {
			return strings.ToUpper(m)
		},
	)
}
