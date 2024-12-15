package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	readFileAndPrintIt()
}

func readFileAndPrintIt() {
	of, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, of)
	fmt.Println()

	// bs := make([]byte, 32*1024)
	// of.Read(bs)
	// of.Close()
	// fmt.Println(string(bs))
}
