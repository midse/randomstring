package randomstring

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const MAX_CHARS int = 255

// Testing purpose
//const MAX_CHARS int = 10

func generatePunct() []string {
	allRanges := appendIntRanges([][]int{rangeInt(33, 47), rangeInt(58, 64), rangeInt(91, 96), rangeInt(123, 126)})

	return intToStringRange(allRanges)
}

var digit = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var lower = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var upper = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var punct = generatePunct()
var any = appendStrRanges([][]string{upper, lower, digit, punct})

// UNUSED
//var salt []string = appendStrRanges([][]string{upper, lower, digit, []string{".", "/"}})
//var binary []string = intToStringRange(rangeInt(0, 255))

var patterns = map[string][]string{
	".":   any,
	"\\d": digit,
	"\\D": appendStrRanges([][]string{upper, lower, punct}),
	"\\w": appendStrRanges([][]string{upper, lower, digit, []string{"_"}}),
	"\\W": appendStrRanges([][]string{upper, lower, digit, punct}), // We should delete '_' from punct
	"\\s": []string{" ", "\t", "\f", "\r"},
	"\\S": appendStrRanges([][]string{upper, lower, digit, punct}), // this might be wrong
}

var octalPattern = regexp.MustCompile("[0-7]")
var wordPattern = regexp.MustCompile("\\w")
var rangePattern = regexp.MustCompile("^(\\d*),(\\d*)$")
var commaPattern = regexp.MustCompile(".*,.*")
var dealCharsPattern = regexp.MustCompile("[\\\\.{\\[*+?\\(]")
var backslashPattern = regexp.MustCompile("(\\\\x[0-9a-f]{2}|\\\\.{1}|.{1})")
var nonCapturingPattern = regexp.MustCompile("\\?([imsU]*:|P<[a-zA-Z0-9]+>).*$")

func dealWithChars(char string, chars []string, str [][]string) (string, []string, [][]string) {
	switch char {
	case "\\":
		var tmp string
		var curChar string

		tmp, chars = shift(chars)
		if tmp == "x" {
			curChar, chars = shift(chars)
			tmp += curChar
			curChar, chars = shift(chars)
			tmp += curChar

			ascii, err := hex.DecodeString(tmp[1:])

			if err != nil {
				str = append(str, []string{"\\" + tmp})
			} else {
				str = append(str, []string{string(ascii)})
			}

		} else if octalPattern.MatchString(tmp) {
			fmt.Println("octal parsing not implemented -> treating litterally")

			str = append(str, []string{tmp})
		} else if _, ok := patterns["\\"+tmp]; ok {
			char += tmp
			str = append(str, patterns["\\"+tmp])
		} else {
			/*if wordPattern.MatchString(tmp) {
				fmt.Println("\\" + tmp + " treated as litteral")
			}*/

			str = append(str, []string{tmp})
		}

	case ".":
		str = append(str, patterns["."])
	case "[":
		var tmp []string
		var negate bool

		// To handle the negate case
		if string(chars[0]) == "^" {
			negate = true
			_, chars = shift(chars) // We can remove the "^"
		}

		for { // Infinite loop, yeah.
			char, chars = shift(chars)

			if char == "]" || (char == "" && len(chars) <= 0) {
				break
			}

			// Is this a range ?
			if char == "-" && chars[0] != "]" && len(chars) > 0 && len(tmp) > 0 {
				char, chars = shift(chars)

				start := tmp[len(tmp)-1][0]

				if string(start) != "[" {
					for n := start + 1; n <= char[0]; n++ {
						tmp = append(tmp, string(n))
					}
				} else {
					tmp = append(tmp, char)
				}
			} else if char == "\\" && len(chars) > 1 {
				if value, ok := patterns["\\"+string(chars[0])]; ok {
					tmp = append(tmp, value...)

					_, chars = shift(chars)

				} else {
					char, chars = shift(chars)
					tmp = append(tmp, string(char))
				}
			} else {
				//fmt.Println(char + " will be treated literally inside []")

				tmp = append(tmp, char)
			}
		}

		if negate {
			results := backslashPattern.FindAllStringSubmatch(strings.Join(tmp, ""), -1)

			toRemove := []string{}

			for i := 0; i < len(results); i++ {
				curResult := results[i][1]

				if len(curResult) > 2 && string(curResult[1]) == "x" {
					ascii, _ := hex.DecodeString(curResult[2:])
					curResult = string(ascii)
				}

				toRemove = append(toRemove, curResult)
			}

			filtered := []string{}

			for _, item := range any {
				canAppend := true

				for _, x := range toRemove {
					if item == x {
						canAppend = false
						break
					}
				}

				if canAppend {
					filtered = append(filtered, item)
				}
			}

			tmp = filtered

		}

		if char != "]" {
			fmt.Println("unmatched [] !") // must panic?
		}

		str = append(str, tmp)

	case "*":

		// We don't handle this case --> .*? so we delete the "?"
		if string(chars[0]) == "?" {
			_, chars = shift(chars)
		}

		chars = prependArray(strings.Split("{0,}", ""), chars)
	case "+":
		chars = prependArray(strings.Split("{1,}", ""), chars)
	case "?":
		chars = prependArray(strings.Split("{0,1}", ""), chars)
	case "{":
		var closed bool

		for n := 0; n < len(chars); n++ {
			if chars[n] == "}" {
				closed = true
				break
			}
		}

		if closed {
			var tmp string
			var nbOfChar, min, max int
			var rangeGenerated bool

			for { // Infinite loop, yeah.
				char, chars = shift(chars)

				if (len(chars) <= 0 && char == "") || char == "}" {
					break
				}

				tmp += char
			}

			// @TODO : check if tmp is well formed ?

			if commaPattern.MatchString(tmp) {
				result := rangePattern.FindStringSubmatch(tmp)

				// @TODO : a clean way to deal with this?
				if len(result) > 2 {
					rangeGenerated = true

					if result[1] != "" {
						min, _ = strconv.Atoi(result[1])
					}

					if result[2] != "" {
						max, _ = strconv.Atoi(result[2])
					} else {
						max = MAX_CHARS
					}

					if min > max {
						panic("Bad range !")
					}

					if min == max {
						nbOfChar = min
					} else {
						nbOfChar = min + randInt(min, max+1)
					}

				} else {
					fmt.Println(result)
					panic("Malformed range")
				}

			}

			if !rangeGenerated {
				nbOfChar, _ = strconv.Atoi(tmp)
			}

			if nbOfChar == 0 && len(str)-1 >= 0 {
				str = str[:len(str)-1]
			} else {
				if len(str)-1 >= 0 {
					last := str[len(str)-1]
					for n := 0; n < (nbOfChar - 1); n++ {
						str = append(str, last)
					}
				}
			}
		} else {
			str = append(str, []string{char})
		}

	case "(":
		var closed bool
		var n int

		nbParLeft := 1

		// @TODO : (?iMsu) case
		// @TODO : nest parenthesis parsing improvement
		for n = 0; n < len(chars); n++ {
			if chars[n] == "(" && n > 0 && chars[n-1] != "\\" {
				nbParLeft++
			}

			if chars[n] == ")" && nbParLeft == 1 && n > 0 && chars[n-1] != "\\" {
				closed = true
				break
			}

			if chars[n] == ")" && n > 0 && chars[n-1] != "\\" {
				nbParLeft--
			}

		}

		if !closed {
			str = append(str, []string{char})
		} else {
			// Naive approach
			//choices := strings.Split(strings.Join(chars[:n], ""), "|")

			nbParLeft = 1
			doNotSplit := false
			cleanChoices := []string{}
			lastPos := 0
			endPar := 0

			for i := 0; i < len(chars); i++ {
				if chars[i] == "(" && i > 0 && chars[i-1] != "\\" {
					nbParLeft++
					doNotSplit = true
				}

				if chars[i] == ")" && nbParLeft == 1 && i > 0 && chars[i-1] != "\\" {
					endPar = i
					closed = true
					break
				}

				if chars[i] == ")" && i > 0 && chars[i-1] != "\\" {
					nbParLeft--
					doNotSplit = false
				}

				if chars[i] == "|" && !doNotSplit {
					cleanChoices = append(cleanChoices, strings.Join(chars[lastPos:i], ""))
					lastPos = i + 1
				}
			}

			cleanChoices = append(cleanChoices, strings.Join(chars[lastPos:endPar], ""))

			// Could be "shifted" before ?
			chars = chars[n+1:]

			if nonCapturingPattern.MatchString(cleanChoices[0]) {
				// Quick and dirty... getting last index of > or : in order to remove
				// stuff like ?: ?i: ?<tag>
				index := strings.Index(cleanChoices[0], ">") + strings.Index(cleanChoices[0], ":") + 1
				cleanChoices[0] = cleanChoices[0][index+1:]
			}

			// Random choice
			cleanChoices = strings.Split(cleanChoices[randInt(0, len(cleanChoices))], "")
			chars = prependArray(cleanChoices, chars)

		}

	}

	return char, chars, str
}

func FromRegex(myPattern string) string {
	//fmt.Println(myPattern)

	if _, err := regexp.Compile(myPattern); err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	var char, randomStr string
	var str [][]string

	chars := strings.Split(myPattern, "")

	for { // Infinite loop, yeah.
		char, chars = shift(chars)

		// Break the loop if 'chars' is empty == parsing is over
		if len(chars) <= 0 && char == "" {
			break
		}

		// Special chars are treated in a special way
		if dealCharsPattern.MatchString(char) {
			char, chars, str = dealWithChars(char, chars, str)
		} else {
			str = append(str, []string{char})
		}

	}

	// Generate the final string
	for i := 0; i < len(str); i++ {
		curStr := str[i]

		if len(curStr) > 1 {
			randomStr += string(curStr[randInt(0, len(curStr))])
		} else {
			randomStr += string(curStr[0])
		}
	}

	return randomStr
}
