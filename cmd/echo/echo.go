package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var n bool
	flag.BoolVar(&n, "n", false, "omits the trailing newline.")
	flag.Parse()

	args := flag.Args()
	out := strings.Join(args, " ")

	if n {
		fmt.Print(out)
	} else {
		fmt.Println(out)
	}
}
