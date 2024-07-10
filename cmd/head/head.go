package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"io"
)

func readLines(reader io.Reader, n int) {
	bufReader := bufio.NewReader(reader)
	for i := 0; i < n; i++ {
		line, err := bufReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading line:", err)
			break
		}
		fmt.Print(line)
	}
}

func main() {
	var n int
	flag.IntVar(&n, "n", 10, "number of lines to print")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		readLines(os.Stdin, n)
	} else {
		for _, fileName := range args {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Printf("Failed to open file: %s\n", fileName)
				continue
			}
			defer file.Close()
			readLines(file, n)
		}
	}
}
