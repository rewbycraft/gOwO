package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	gowo "github.com/rewbycraft/gOwO/pkg/gowo"
)

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

		fmt.Println(gowo.OwOify(string(text), *prefixPtr, *suffixPtr))
	}
}
