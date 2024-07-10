package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func printDir(dir string, depth, currentDepth int) {
	if currentDepth > depth {
		return
	}
	entries, err := os.ReadDir(dir)
	check(err)
	for _, entry := range entries {
		prefix := "|"
		indent := "――"
		if currentDepth > 0 {
			prefix = "|"
			for i := 0; i < currentDepth; i++ {
				prefix += "    "
			}
			indent = prefix + "――"
		}
		fmt.Println(prefix)
		fmt.Println(indent, entry.Name())

		if entry.IsDir() && currentDepth < depth {
			newPath := filepath.Join(dir, entry.Name())
			printDir(newPath, depth, currentDepth+1)
		}
	}
}

func main() {
	depthFlag := flag.String("depth", "1", "Depth to print")
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatal("Usage: go run main.go -depth=<depth> <directory_path>")
	}

	depth, err := strconv.Atoi(*depthFlag)
	check(err)

	dirPath := flag.Arg(0)

	printDir(dirPath, depth, 0)
}
