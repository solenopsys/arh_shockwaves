package configs

import (
	"fmt"
	"regexp"
	"strings"
)

func PatternMatching(name string, filter string) (bool, error) {
	pattern := strings.Replace(filter, "*", ".*", -1)
	return regexp.MatchString("^"+pattern+"$", name)
}

func PatternMatchingRank(name string, filter string) ([]int, error) {
	pattern := strings.Replace(filter, "*", ".*", -1)
	matched, err := regexp.MatchString("^"+pattern+"$", name)
	if err != nil {
		return nil, err
	}

	ids := []int{}

	if matched {
		for i, _ := range name {

			ids = append(ids, i)

		}
		return ids, err
	} else {
		return nil, err
	}

}

func PatternMatchingSubstringRank(name string, filter string) ([]int, error) {
	pattern := strings.Replace(filter, "*", ".*", -1)

	// Create a regular expression object
	r := regexp.MustCompile("^" + pattern + "$")

	// Define a function to replace each match
	replacer := func(match string) string {
		// Replace the matched digits with asterisks
		return strings.Repeat("#", len(match))
	}

	// Replace the matched groups of four digits with asterisks
	result := r.ReplaceAllStringFunc(name, replacer)

	fmt.Println(result)

	ids := []int{}
	for i, c := range result {
		if string(c) == "#" {
			ids = append(ids, i)
		}
	}
	return ids, nil

}
