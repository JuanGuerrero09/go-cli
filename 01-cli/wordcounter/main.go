package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")

	flag.Parse()

  counter, int := count(os.Stdin, *lines, *bytes)

	fmt.Printf("Total %s: %d", counter, int)

}

func count(r io.Reader, countLines bool, countBytes bool) (string, int) {

	scanner := bufio.NewScanner(r)
  counter := "words"

	scanner.Split(bufio.ScanWords)

	if countBytes {
		scanner.Split(bufio.ScanBytes)
    counter = "bytes"
	}

	if countLines {
		scanner.Split(bufio.ScanLines)
    counter = "lines"
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return counter, wc
}
