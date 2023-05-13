// This script will generate dynamicly the classes.go file

package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	DefineAst("interpreter", "Expr", []string{
		"Binary   : Expr left, Token operator, Expr right",
		"Grouping : Expr expression",
		"Literal  : interface{} value",
		"Unary    : Token operator, Expr right",
	})
}

func DefineAst(outputDir string, baseName string, types []string) {
	var path = outputDir + baseName + ".go"

	f, err := os.Create(path)
	check(err)

	_, err = f.WriteString("package interpreter\n")
	check(err)

	defer f.Close()

	f.Sync()

}
