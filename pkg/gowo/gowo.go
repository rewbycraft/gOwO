package gowo

import (
	"fmt"
	"math/rand"
	"strings"
)

var prefixes = []string{
	"<3 ",
	"0w0 ",
	"H-hewwo?? ",
	"HIIII! ",
	"Haiiii! ",
	"Huohhhh. ",
	"OWO ",
	"OwO ",
	"UwU ",
}

var suffixes = []string{
	" :3",
	" UwU",
	" ʕʘ‿ʘʔ",
	" >_>",
	" ^_^",
	"..",
	" Huoh.",
	" ^-^",
	" ;_;",
	" ;-;",
	" xD",
	" x3",
	" :D",
	" :P",
	" ;3",
	" XDDD",
	", fwendo",
	" ㅇㅅㅇ",
	" (人◕ω◕)",
	"（＾ｖ＾）",
	" Sigh.",
	" x3",
	" ._.",
	" (• o •)",
	" >_<",
}

var subs = map[string]string{
	"r": "w",
	"l": "w",
	"R": "W",
	"L": "W",
	//  "ow": "OwO",
	"no":   "nu",
	"has":  "haz",
	"have": "haz",
	"you":  "uu",
	"the ": "da ",
	"The ": "Da ",
}

func getRandomFromSlice(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func OwOify(text string, enablePrefix bool, enableSuffix bool) string {
	prefix := ""
	suffix := ""
	if enablePrefix {
		prefix = getRandomFromSlice(prefixes)
	}

	if enableSuffix {
		suffix = getRandomFromSlice(suffixes)
	}

	replacedString := string(text)

	for k, v := range subs {
		replacedString = strings.Replace(replacedString, k, v, -1)
	}

	return fmt.Sprintf("%s%s%s", prefix, replacedString, suffix)
}
