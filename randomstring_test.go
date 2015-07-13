package randomstring

import (
	"regexp"
	"testing"
)

func TestFromRegex(t *testing.T) {

	tests := []string{
		"[a-z0-9_]{3,16}",
		"#?([a-f0-9]{6}|[a-f0-9]{3})",
		"[a-z0-9_\\.]+@[0-9a-z\\.]+\\.[a-z\\.]{2,6}",
	}

	for i := 0; i < len(tests); i++ {

		myRegexp := tests[i]

		result := FromRegex(myRegexp)
		pattern := regexp.MustCompile(myRegexp)

		if !pattern.MatchString(result) {
			t.Error("String generated : " + result + "\n" + "Regex :" + myRegexp)
		}
	}
}
