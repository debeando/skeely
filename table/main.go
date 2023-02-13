package table

import (
	"regexp"
	"strconv"
	"strings"
)

func stringToArray(v string) (values []string) {
	for _, value := range strings.Split(v, ",") {
		values = append(values, strings.Trim(value, "`"))
	}

	return values
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func find(e, r string) [][]string {
	return regexp.MustCompile(e).FindAllStringSubmatch(r, -1)
}

func findMatchOne(e, r string) string {
	re := regexp.MustCompile(e)
	match := re.FindStringSubmatch(r)

	for i, name := range re.SubexpNames() {
		if len(match) > 0 && i != 0 && name != "" {
			return match[i]
		}
	}

	return ""
}
