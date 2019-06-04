package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/darashi/kouotsu/translator"
)

func main() {
	var input string

	if len(os.Args) > 1 {
		input = strings.Join(os.Args[1:], " ")
	} else {
		fmt.Scan(&input)
	}
	r := translator.Translate(input)
	fmt.Println(r)
}
