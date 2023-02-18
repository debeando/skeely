package common

import (
	"regexp"
	"strconv"
	"strings"
)

func StringToArray(v string) (values []string) {
	for _, value := range strings.Split(v, ",") {
		if len(value) > 0 {
			values = append(values, strings.Trim(value, "`"))
		}
	}

	return values
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func Find(e, r string) [][]string {
	return regexp.MustCompile(e).FindAllStringSubmatch(r, -1)
}

func FindMatchOne(e, r string, l int) string {
	re := regexp.MustCompile(e)
	match := re.FindStringSubmatch(r)

	if len(match) > 0 {
		return match[l]
	}

	return ""
}

func StringIn(x, y string) bool {
	return strings.Contains(strings.ToUpper(x), y)
}
