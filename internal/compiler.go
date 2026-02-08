package internal

import (
	"C-Compiler/internal/analyzer"
	"C-Compiler/internal/codegen"
	"C-Compiler/internal/lexer"
	"C-Compiler/internal/parser"
)

type Compiler struct {
	source string
}

func CompilerFromSource(source string) *Compiler {
	return &Compiler{source: source}
}

func (c *Compiler) Compile() (string, []string) {
	l := lexer.New(c.source)
	p := parser.New(l)
	prog := p.ParseProgram()
	a := analyzer.New(prog)
	a.Analyze()
	hasErrors, errors := a.Errors()
	if hasErrors {
		return "", errors
	}

	generator := codegen.New()
	assembly := generator.Generate(prog)

	return assembly, nil
}
