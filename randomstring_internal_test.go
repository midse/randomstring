package randomstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeInt(t *testing.T) {
	assert.Equal(t, rangeInt(1, 5), []int{1, 2, 3, 4, 5}, "they should be equal")
	assert.Equal(t, rangeInt(5, 1), []int{}, "a > b empty return")
	assert.Equal(t, rangeInt(1, 1), []int{1}, "a = b should return a")
	assert.Equal(t, rangeInt(0, 1), []int{0, 1}, "should be equal")
	assert.Equal(t, rangeInt(2, 4), []int{2, 3, 4}, "should be equal")
	assert.Equal(t, rangeInt(-5, 0), []int{-5, -4, -3, -2, -1, 0}, "should be equal")
}

/*func TestIntToStringRange(t *testing.T) {
	t.Error("Fail!")
}

func TestAppendIntRanges(t *testing.T) {
	t.Error("Fail!")
}

func TestAppendStrRanges(t *testing.T) {
	t.Error("Fail!")
}

func TestPrepend(t *testing.T) {
	t.Error("Fail!")
}

func TestPrependArray(t *testing.T) {
	t.Error("Fail!")
}

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
