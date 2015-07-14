package randomstring

import "math/rand"

// @TODO : check parameters
func rangeInt(a int, b int) []int {
	if a > b {
		return []int{}
	}

	myRange := make([]int, (b-a)+1)

	j := 0
	for i := a; i <= b; i++ {
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

	for i := 0; i < len(ranges); i++ {
		appended = append(appended, ranges[i]...)
	}

	return appended
}

func appendStrRanges(ranges [][]string) []string {
	var appended []string

	for i := 0; i < len(ranges); i++ {
		appended = append(appended, ranges[i]...)
	}

	return appended
}

func prepend(toPrepend string, array []string) []string {
	return append([]string{toPrepend}, array...)
}

func prependArray(toPrepend []string, array []string) []string {
	for i := len(toPrepend) - 1; i >= 0; i-- {
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
