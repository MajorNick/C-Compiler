package analyzer

import (
	"C-Compiler/internal/lexer"
	"C-Compiler/internal/parser"
	"testing"
)

func TestVarDeclaration(t *testing.T) {
	input := "int a = 5;"
	l := lexer.New(input)
	p := parser.New(l)
	prog := p.ParseProgram()

	a := New(prog)
	a.Analyze()

	if len(a.errors) != 0 {
		t.Errorf("Expected no errors, got %d: %v", len(a.errors), a.errors)
	}
}

func TestUndefinedVariable(t *testing.T) {
	input := "int a = b;" // 'b' is not defined
	l := lexer.New(input)
	p := parser.New(l)
	prog := p.ParseProgram()

	a := New(prog)
	a.Analyze()

	if len(a.errors) == 0 {
		t.Errorf("Expected error for undefined variable 'b', got 0 errors")
	}
	if len(a.errors) > 0 && a.errors[0] != "Variable b not defined" {
		t.Errorf("Expected error 'Variable b not defined', got '%s'", a.errors[0])
	}
}

func TestMainFunctionScope(t *testing.T) {
	input := `
	int main() {
		int a = 5;
		int b = a;
	}`
	l := lexer.New(input)
	p := parser.New(l)
	prog := p.ParseProgram()

	a := New(prog)
	a.Analyze()

	if len(a.errors) != 0 {
		t.Errorf("Expected no errors in main function, got %v", a.errors)
	}
}

func TestMainFunctionUndefined(t *testing.T) {
	input := `
	int main() {
		int a = c; // c undefined
	}`
	l := lexer.New(input)
	p := parser.New(l)
	prog := p.ParseProgram()

	a := New(prog)
	a.Analyze()

	if len(a.errors) == 0 {
		t.Errorf("Expected error for undefined variable 'c'")
	}
}

func TestUndefinedFunctionCall(t *testing.T) {
	input := `
	int main() {
		int a = 5; 
		int c = 123;
		c=b;
	}`
	l := lexer.New(input)
	p := parser.New(l)
	prog := p.ParseProgram()

	a := New(prog)
	a.Analyze()

	if len(a.errors) == 0 {
		t.Errorf("Expected error for undefined variable 'b'")
	}
}
