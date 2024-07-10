package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"log"
)

func main() {
	var l, w, c bool
	flag.BoolVar(&l, "l", false, "number of lines")
	flag.BoolVar(&w, "w", false, "number of words")
	flag.BoolVar(&c, "c", false, "number of characters")
	flag.Parse()

	args := flag.Args()
	var reader *bufio.Reader
	if len(args) == 1 {
		file, err := os.Open(args[0])
		if err != nil {
			log.Fatalf("Invalid file")
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	lines, words, chars := 0, 0, 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lines++
		chars += len(line)
		word := strings.Fields(line)
		words += len(word)
	}

	check := !l && !w && !c

	if check || l {
		fmt.Println(lines)
	}
	if check || w {
		fmt.Println(words)
	}
	if check || c {
		fmt.Println(chars)
	}
}
