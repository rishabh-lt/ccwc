package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NoOfBytes(ans []byte) int {
	return len(ans)

}

func NoOfLines(ans []byte) int {
	lines := strings.Split(string(ans), "\n")
	if lines[len(lines)-1] == "" {
		return len(lines) - 1
	}
	return len(lines)
}

func NoOfWords(ans []byte) int {
	words := strings.Fields(string(ans))
	return len(words)
}

func NoOfCharacters(ans []byte) int {
	return utf8.RuneCount(ans)
}

func main() {
	countBytes := flag.Bool("c", false, "The number of bytes in each input file is written to the standard output.")
	countLines := flag.Bool("l", false, "The number of bytes in each input file is written to the standard output.")
	countWords := flag.Bool("w", false, "The number of words in each input file is written to the standard output.")
	countCharacters := flag.Bool("m", false, "The number of characters in each input file is written to the standard output.")
	flag.Parse()

	var ans []byte
	var err error
	var fileName string = ""
	var args = flag.Args()

	stat, err := os.Stdin.Stat()
	check(err)

	if stat.Mode()&os.ModeCharDevice == 0 {
		ans, err = io.ReadAll(os.Stdin)
		check(err)
	} else {
		fileName = args[len(args)-1]
		if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
			fmt.Println("file doesn't exsist")
			return
		}
		ans, err = os.ReadFile(fileName)
		check(err)
	}

	if *countBytes {
		fmt.Printf("%d %s \n", NoOfBytes(ans), fileName)
		return
	}

	if *countLines {
		fmt.Printf("%d %s \n", NoOfLines(ans), fileName)
		return
	}

	if *countWords {
		fmt.Printf("%d %s \n", NoOfWords(ans), fileName)
		return
	}

	if *countCharacters {
		fmt.Printf("%d %s \n", NoOfCharacters(ans), fileName)
		return
	}

	fmt.Printf("%d %d %d %s \n", NoOfLines(ans), NoOfWords(ans), NoOfBytes(ans), fileName)
}
