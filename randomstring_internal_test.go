package randomstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeInt(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5}, rangeInt(1, 5), "they should be equal")
	assert.Equal(t, []int{}, rangeInt(5, 1), "a > b empty return")
	assert.Equal(t, []int{1}, rangeInt(1, 1), "a = b should return a")
	assert.Equal(t, []int{0, 1}, rangeInt(0, 1), "should be equal")
	assert.Equal(t, []int{2, 3, 4}, rangeInt(2, 4), "should be equal")
	assert.Equal(t, []int{-5, -4, -3, -2, -1, 0}, rangeInt(-5, 0), "should be equal")
}

func TestIntToStringRange(t *testing.T) {
	assert.Equal(t, []string{"!", "\"", "#"}, intToStringRange([]int{33, 34, 35}), "they should be equal")
	assert.Equal(t, []string{}, intToStringRange([]int{}), "should be equal")
	assert.Equal(t, []string{"1"}, intToStringRange([]int{49}), "should be equal")
}

func TestAppendIntRanges(t *testing.T) {
	assert.Equal(t, rangeInt(1, 10), appendIntRanges([][]int{rangeInt(1, 5), rangeInt(6, 10)}), "they should be equal")
	assert.Equal(t, []int(nil), appendIntRanges([][]int{}), "they should be equal")
}

func TestAppendStrRanges(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, appendStrRanges([][]string{[]string{"a", "b", "c"}, []string{"d", "e"}}), "they should be equal")
	assert.Equal(t, []string(nil), appendStrRanges([][]string{}), "they should be equal")
}

func TestPrepend(t *testing.T) {
	assert.Equal(t, []string{"a"}, prepend("a", []string{}), "should be equal")
	assert.Equal(t, []string{"a", "b", "c"}, prepend("a", []string{"b", "c"}), "should be equal")
}

func TestPrependArray(t *testing.T) {
	assert.Equal(t, []string{"a"}, prependArray([]string{"a"}, []string{}), "should be equal")
	assert.Equal(t, []string{}, prependArray([]string{}, []string{}), "should be equal")
	assert.Equal(t, []string{"a"}, prependArray([]string{}, []string{"a"}), "should be equal")
	assert.Equal(t, []string{"a", "b", "c"}, prependArray([]string{"a", "b"}, []string{"c"}), "should be equal")
	assert.Equal(t, []string{"a", "b", "c"}, prependArray([]string{"a"}, []string{"b", "c"}), "should be equal")
}

/*
func TestShift(t *testing.T) {
	t.Error("Fail!")
}

func TestRandInt(t *testing.T) {
	t.Error("Fail!")
}

func TestRandomString(t *testing.T) {
	t.Error("Fail!")
}

func TestDeleteFromIndex(t *testing.T) {
	t.Error("Fail!")
}

func TestDeleteKeysFromArray(t *testing.T) {
	t.Error("Fail!")
}

func TestArrayToMap(t *testing.T) {
	t.Error("Fail!")
}

func TestSmartSplit(t *testing.T) {
	t.Error("Fail!")
}*/
