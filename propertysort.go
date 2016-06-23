package main

import (
	"bufio"
	"fmt"
	"github.com/derkork/properties"
	"os"
)

func printUsage() {
	fmt.Println("Usage: propertysort <filename>")
}

func handlePanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args
	if len(args) != 2 || args[1] == "--help" {
		printUsage()
		return
	}
	file := os.Args[1]

	p := properties.MustLoadFile(file, properties.ISO_8859_1)
	p.Sort()

	f, err := os.Create(file)
	handlePanic(err)
	defer f.Close()

	writer := bufio.NewWriter(f)
	_, err = p.Write(writer, properties.ISO_8859_1)
	handlePanic(err)
	writer.Flush()

}
