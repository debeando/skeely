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

func StringToArrayInt(v string) (values []int) {
	for _, value := range strings.Split(v, ",") {
		if len(value) > 0 {
			values = append(values, StringToInt(strings.Trim(value, " ")))
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

func StringToBool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

func StringIsEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}
	return false
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

func UnduplicateArrayInt(intSlice []int) (list []int) {
	keys := make(map[int]bool)

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func IntInArrayInt(slice []int, value int) bool {
	for index := range slice {
		if slice[index] == value {
			return true
		}
	}

	return false
}
