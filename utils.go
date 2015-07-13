package randomstring

import (
	"math/rand"
	"regexp"
	"strings"
)

// @TODO : check parameters
func rangeInt(a int, b int) []int {
	if a > b {
		return []int{}
	}

	myRange := make([]int, (b-a)+1)

	j := 0
	for i := a; i <= b; i += 1 {
		myRange[j] = i
		j++
	}

	return myRange
}

func intToStringRange(intRange []int) []string {
	strRange := make([]string, len(intRange))

	for i, value := range intRange {
		strRange[i] = string(value)
	}

	return strRange
}

func appendIntRanges(ranges [][]int) []int {
	var appended []int

	for i := 0; i < len(ranges); i += 1 {
		appended = append(appended, ranges[i]...)
	}

	return appended
}

func appendStrRanges(ranges [][]string) []string {
	var appended []string

	for i := 0; i < len(ranges); i += 1 {
		appended = append(appended, ranges[i]...)
	}

	return appended
}

func prepend(toPrepend string, array []string) []string {
	return append([]string{toPrepend}, array...)
}

func prependArray(toPrepend []string, array []string) []string {
	for i := len(toPrepend) - 1; i >= 0; i -= 1 {
		array = prepend(toPrepend[i], array)
	}

	return array
}

func shift(array []string) (string, []string) {
	if len(array) > 1 {
		return array[0], array[1:]
	}

	if len(array) > 0 {
		return array[0], []string{}
	}

	return "", nil
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(32, 126))
	}
	return string(bytes)
}

func deleteFromIndex(index int, array []string) []string {
	return append(array[:index], array[index+1:]...)
}

func deleteKeysFromArray(keys []string, array []string) []string {
	pattern := ".*?"

	for _, value := range keys {
		pattern += value + "|"
	}

	pattern = pattern[:len(pattern)-1]
	pattern += ".*?"

	keyPattern, _ := regexp.Compile(pattern)

	for index := 0; index < len(array); index += 1 {
		if r := keyPattern.FindString(array[index]); r != "" {
			array = deleteFromIndex(index, array)
			index -= 1
		}
	}

	return array
}

func arrayToMap(array []string) map[string][]string {
	rule := make(map[string][]string)

	for _, value := range array {
		if value == "" {
			continue
		}

		//key := strings.Split(strings.Trim(value, " "), ":")
		key := smartSplit(strings.Trim(value, " "), ":")

		if _, ok := rule[key[0]]; ok {
			if len(key) > 1 {
				clean_key := strings.Join(key[1:], "")
				rule[key[0]] = append(rule[key[0]], strings.Trim(clean_key, "\""))
			}
		} else {
			if len(key) > 1 {
				clean_key := strings.Join(key[1:], "")
				rule[key[0]] = []string{strings.Trim(clean_key, "\"")}
			} else {
				rule[key[0]] = []string{}
			}
		}
	}

	return rule
}

func smartSplit(s string, sep string) []string {
	var splitStr []string
	var tmpStr, char string

	inBlock := false

	for i := 0; i < len(s); i++ {
		char = string(s[i])

		if char == "\"" {
			inBlock = !inBlock
		}

		if char != sep {
			tmpStr += char
		} else {
			if inBlock && char == sep {
				tmpStr += char
			} else {
				splitStr = append(splitStr, tmpStr)
				tmpStr = ""
			}
		}
	}

	if tmpStr != "" {
		splitStr = append(splitStr, tmpStr)
	}

	return splitStr

}
