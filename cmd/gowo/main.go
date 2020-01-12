package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var prefixes = []string{
	"<3 ",
	"0w0",
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

func main() {
	prefixPtr := flag.Bool("prefix", true, "Add a prefix")
	suffixPtr := flag.Bool("suffix", true, "Add a suffix")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	for {
		text, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		prefix := ""
		suffix := ""
		if *prefixPtr {
			prefix = getRandomFromSlice(prefixes)
		}

		if *suffixPtr {
			suffix = getRandomFromSlice(suffixes)
		}

		replacedString := string(text)

		for k, v := range subs {
			replacedString = strings.Replace(replacedString, k, v, -1)
		}

		fmt.Printf("%s%s%s\n", prefix, replacedString, suffix)
	}
}
