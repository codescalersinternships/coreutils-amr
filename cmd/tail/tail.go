package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"io"
)

func readLines(reader io.Reader, text *[]string) {
	bufReader := bufio.NewReader(reader)
	for {
		line, err := bufReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error reading line: %v", err)
		}
		*text = append(*text, line)
	}
}

func main() {
	var n int
	flag.IntVar(&n, "n", 10, "number of lines to print")
	flag.Parse()

	args := flag.Args()
	var text []string

	if len(args) == 0 {
		readLines(os.Stdin, &text)
	} else {
		for _, fileName := range args {
			file, err := os.Open(fileName)
			if err != nil {
				log.Fatalf("Failed to open file: %s", fileName)
			}
			defer file.Close()
			readLines(file, &text)
		}
	}

	if len(text) < n {
		n = len(text)
	}

	for i := len(text) - n; i < len(text); i++ {
		fmt.Print(text[i])
	}
}
