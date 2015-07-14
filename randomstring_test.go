package randomstring_test

import (
	"regexp"
	"testing"

	"github.com/midse/randomstring"
	"github.com/stretchr/testify/assert"
)

func TestFromRegex(t *testing.T) {

	tests := []string{
		"[a-z0-9_]{3,16}",
		"#?([a-f0-9]{6}|[a-f0-9]{3})",
		"[a-z0-9_.-]+",
		"[a-z0-9_.-]+@[0-9a-z.-]+\\.[a-z.]{2,6}",
		"https?:\\/\\/([\\da-z\\.-]+)\\.([a-z\\.]{2,6})([\\/\\w .-]*)*\\/?",
	}

	for i := 0; i < len(tests); i++ {

		myRegexp := tests[i]

		result := randomstring.FromRegex(myRegexp)
		pattern := regexp.MustCompile(myRegexp)

		assert.Equal(t, true, pattern.MatchString(result), "Regex: "+myRegexp+" Output: "+result)
	}
}
