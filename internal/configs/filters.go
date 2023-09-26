package configs

import (
	"regexp"
	"strings"
)

func PatternMatching(name string, filter string) (bool, error) {
	pattern := strings.Replace(filter, "*", ".*", -1)
	return regexp.MatchString("^"+pattern+"$", name)
}
