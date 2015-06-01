package lessCommonWords

import (
	"github.com/azer/go-most-common-words"
	"regexp"
	"strings"
)

func Filter(list []string, dictionaries ...map[string]bool) []string {
	result := []string{}

	for _, el := range list {
		el = strings.ToLower(el)

		if len(el) < 2 || mostCommonWords.Dict[el] {
			continue
		}

		skip := false
		for _, dict := range dictionaries {
			if dict[el] {
				skip = true
				break
			}
		}

		if skip {
			continue
		}

		result = append(result, el)
	}

	return result
}

func Keywords(input string, dictionaries ...map[string]bool) []string {
	return Filter(regexp.MustCompile("[^\\w]").Split(input, -1), dictionaries...)
}
