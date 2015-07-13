package fuzzer

import (
	"regexp"
	"testing"
	//	"fmt"
)

func TestRandStrFromRegex(t *testing.T) {

	tests := []string{
		"/([a-zA-Z0-9_]+)\\.(php|py|html)\\?([a-z]+=[a-zA-Z0-9]+)",
		"(?i:(\\!\\=|\\&\\&|\\|\\||>>|<<|>=|<=|<>|<=>|xor|rlike|regexp|isnull)|(?:not\\s+between\\s+0\\s+and)|(?:is\\s+null)|(like\\s+null)|(?:(?:^|\\W)in[+\\s]*\\([\\s\\d\"]+[^()]*_\\))|(?:xor|<>|rlike(?:\\s+binary)?)|(?:regexp\\s+binary))",
		"(^[\"'`´’‘;]+|[\"'`´’‘;]+$)",
		"(/\\*!?|\\*/|[';]--|--[\\s\\r\\n\\v\\f]|(?:--[^-]*?-)|([^\\-&])#.*?[\\s\\r\\n\\v\\f]|;?\\x00)",
		"<script[^>]*>[\\s\\S]*?<\\/script[[\\s\\S]*[\\s\\S]", "<script[^>]*>[\\s\\S]*?<\\/script[^>]*>",
		"(<script[^>]*>[\\s\\S]*?<\\/script[^>]*>|<script[^>]*>[\\s\\S]*?<\\/script[[\\s\\S]]*[\\s\\S]|<script[^>]*>[\\s\\S]*?<\\/script[\\s]*[\\s]|<script[^>]*>[\\s\\S]*?<\\/script|<script[^>]*>[\\s\\S]*?)",
		"(?i)(<script[^>]*>[\\s\\S]*?<\\/script[^>]*>|<script[^>]*>[\\s\\S]*?<\\/script[[\\s\\S]]*[\\s\\S]|<script[^>]*>[\\s\\S]*?<\\/script[\\s]*[\\s]|<script[^>]*>[\\s\\S]*?<\\/script|<script[^>]*>[\\s\\S]*?)",
		/*"(?i)(?:\\x5c|(?:%(?:2(?:5(?:2f|5c)|%46|f)|c(?:0%(?:9v|af)|1%1c)|u(?:221[56]|002f)|%32(?:%46|F)|e0%80%af|1u|5c)|\\/))(?:%(?:2(?:(?:52)?e|%45)|(?:e0%8|c)0%ae|u(?:002e|2024)|%32(?:%45|E))|\\.){2}(?:\\x5c|(?:%(?:2(?:5(?:2f|5c)|%46|f)|c(?:0%(?:9v|af)|1%1c)|u(?:221[56]|002f)|%32(?:%46|F)|e0%80%af|1u|5c)|\\/))"*/
		"[&\\?]_SESSION\\[[^\\]]{5}\\][^=]*?=",
	}

	//fmt.Println(fuzzer.RandStrFromRegex("f((a|u)|n(n|p)y)"))

	for i := 0; i < len(tests); i++ {

		myRegexp := tests[i]

		result := RandStrFromRegex(myRegexp)
		pattern := regexp.MustCompile(myRegexp)

		//		fmt.Println(result)
		if !pattern.MatchString(result) {
			t.Error("String generated : " + result + "\n" + "Regex :" + myRegexp)
		}

	}
}
