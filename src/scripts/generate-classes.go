// This script will generate dynamicly the classes.go file

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	DefineAst("../interpreter/", "expr", []string{
		"Binary   : Left Expr , Operator token.Token , Right Expr ",
		"Grouping : Expression Expr ",
		"Literal  : Value interface {} ",
		"Unary    : Operator token.Token , Right Expr ",
	})
}

func DefineAst(outputDir string, baseName string, types []string) {
	var path = outputDir + baseName + ".go"
	var buffer bytes.Buffer

	buffer.WriteString("package interpreter\n")
	buffer.WriteString("\n")
	buffer.WriteString("import (\n")
	buffer.WriteString("\t\"github.com/neo-youre-pena/golox/src/token\"\n")
	buffer.WriteString(")\n")
	buffer.WriteString("\n")
	buffer.WriteString("type " + capitalize(baseName) + " interface {\n")
	buffer.WriteString("}\n")
	buffer.WriteString("\n")

	buffer.WriteString("type Visitor interface {\n")
	for _, t := range types {
		className := strings.TrimSpace(t[:bytes.IndexByte([]byte(t), ':')])
		buffer.WriteString("\tvisitFor" + className + "(*" + className + ")\n")
	}

	buffer.WriteString("}\n")

	for _, t := range types {
		className := strings.TrimSpace(t[:bytes.IndexByte([]byte(t), ':')])
		fields := t[bytes.IndexByte([]byte(t), ':')+1:]
		defineType(&buffer, className, fields)
		defineAcceptFunc(&buffer, className)
	}

	f, err := os.Create(path)
	check(err)

	_, err = f.WriteString(buffer.String())

	check(err)

	defer f.Close()

	f.Sync()

}

func defineType(buffer *bytes.Buffer, className string, fieldList string) {
	buffer.WriteString("type " + className + " struct {\n")

	fields := bytes.Split([]byte(fieldList), []byte{','})

	for _, field := range fields {
		fmt.Println(string(field))
		buffer.WriteString("\t" + capitalize(string(field)) + "\n")
	}

	buffer.WriteString("}\n")
}

func defineAcceptFunc(buffer *bytes.Buffer, className string) {
	buffer.WriteString("func (c *" + className + ") accept(v Visitor) {\n")
	buffer.WriteString("\tv.visitFor" + className + "(c)\n")
	buffer.WriteString("}\n")
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
