package main

import (
	"bufio"
	"fmt"
	"github.com/pmylund/sortutil"
	"os"
)

// readLines reads the lines of a file into memory
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func printUsage() {
	fmt.Println("Usage: propertysort <filename>")
}

func main() {
	args := os.Args
	if len(args) != 2 || args[1] == "--help" {
		printUsage()
		return
	}

	file := os.Args[1]
	lines, err := readLines(file)

	if err != nil {
		panic(err)
	}

	sortutil.CiAsc(lines)

	err = writeLines(lines, file)
	if err != nil {
		panic(err)
	}
}
