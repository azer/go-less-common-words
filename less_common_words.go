package lessCommonWords

import (
	"github.com/azer/go-most-common-words"
	"regexp"
	"strings"
)

func Filter(list []string) []string {
	result := []string{}

	for _, el := range list {
		el = strings.ToLower(el)

		if len(el) < 2 || mostCommonWords.Dict[el] {
			continue
		}

		result = append(result, el)
	}

	return result
}

func Keywords(input string) []string {
	return Filter(regexp.MustCompile("[^\\w]").Split(input, -1))
}
