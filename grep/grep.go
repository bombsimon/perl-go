package grep

import (
	"regexp"
	"strings"
)

// List will return a new list with elements matching passed value.
func List(v string, l []string) []string {
	return searchList(v, l, false)
}

// First will return the first match matching passed value.
func First(v string, l []string) string {
	r := searchList(v, l, true)

	if len(r) > 0 {
		return r[0]
	}

	return ""
}

func searchList(v string, l []string, first bool) []string {
	var re *regexp.Regexp

	if strings.HasPrefix(v, "/") && strings.HasSuffix(v, "/") {
		re = regexp.MustCompile(v[1 : len(v)-1])
	}

	nl := []string{}

	for _, i := range l {
		var isMatch bool

		// If string matches as is set flag to match.
		if v == i {
			isMatch = true
		}

		// If we have a regexp and the string didn't match literally, try regexp
		// parsing.
		if re != nil && !isMatch {
			isMatch = len(re.FindAllString(i, -1)) > 0
		}

		if isMatch {
			nl = append(nl, i)

			// If we only looking for first match (i.e. is in list), return the
			// list now.
			if first {
				return nl
			}
		}
	}

	return nl
}
