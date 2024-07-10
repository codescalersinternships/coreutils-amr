package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var n bool
	flag.BoolVar(&n, "n", false, "number the output lines")
	flag.Parse()

	args := flag.Args()
	var out *bufio.Reader
	if len(os.Args) == 0 {
		log.Fatalf("provide a file name")
	}
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Can't open file")
	}
	out = bufio.NewReader(file)

	lineNumber := 1
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		}
		if n {
			fmt.Printf("%d. %s\n", lineNumber, line)
			lineNumber++
		} else {
			fmt.Print(line)
		}
	}
}
