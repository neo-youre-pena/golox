package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"lox/parseerror"
	"lox/scanner"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func runFile(file string) {
	dat, err := ioutil.ReadFile(file)
	check(err)
	run(string(dat))
	if parseerror.HadError {
		os.Exit(65)
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		dat, err := reader.ReadBytes('\n')
		check(err)
		run(string(dat))
	}
}

func run(src string) {
	scanner := scanner.New(src)
	tokens := scanner.ScanTokens()
	for _, t := range tokens {
		fmt.Println(t)
	}
}

func main() {
	flag.String("file", "", "the script file to execute")
	flag.Parse()

	args := flag.Args()
	if len(args) > 1 {
		fmt.Println("Usage: ./lox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}
