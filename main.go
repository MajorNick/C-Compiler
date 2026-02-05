package main

import (
	"C-Compiler/internal/lexer"
	"C-Compiler/internal/parser"
	"fmt"
)

func main() {
	str := `int main(){
		int k  =5;
		return 0;
	}`

	lexer := lexer.New(str)

	parser := parser.New(lexer)
	program := parser.ParseProgram()
	fmt.Println(program)
}
