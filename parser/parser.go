package parser

import(
	"C-Compiler/token"
	"C-Compiler/ast"
	"C-Compiler/lexer"
)

type Parser struct{
	l * lexer.Lexer
	curTok  token.Token
	nextTok  token.Token
}

func New(l *lexer.Lexer)*Parser{
	p := Parser{}
	p.l = l
	p.nextToken()
	p.nextToken()
	return &p
}
func (p * Parser)nextToken(){
	p.curTok = p.nextTok
	p.nextTok = p.l.NextToken()
}

func (p* Parser)ParseProgram() *ast.Program{
	return nil
}