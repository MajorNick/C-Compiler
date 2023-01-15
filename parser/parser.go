package parser

import (
	"C-Compiler/ast"
	"C-Compiler/lexer"
	"C-Compiler/token"
	
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
	program := &ast.Program{}
	program.Stats = []ast.Statement{}
	
	for p.curTok.Type != token.EOF{
		
		stmt := p.parseStatement()
		p.nextToken()
		if stmt != nil{
			program.Stats = append(program.Stats,stmt)
		}
	}

	return program
}

func (p *Parser)parseStatement() ast.Statement{
	switch p.curTok.Type{
	case token.INT,token.LONG,token.SHORT,token.CHAR:
		return p.parseDecStatement()
	case token.INTP,token.LONGP,token.SHORTP,token.CHARP,token.VOIDP:
		return p.parseDecStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	}
	
	return nil
}

func (p *Parser)parseDecStatement() *ast.DeclStatement{
	stmt := &ast.DeclStatement{Token: p.curTok}
	if !p.exceptNext(token.IDENT){
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curTok,Value: p.curTok.Literal}
	// currently not parsing expressions
	p.nextToken()
	for !p.curTokenIs(token.SEMICOLON){
		if p.curTokenIs(token.IDENT){
			if p.nextTokenIs(token.ASSIGN){
				//p.nextToken()
				p.variableDeclarationAssign()
			}else{
				if p.nextTokenIs(token.COMMA){
					//p.nextToken()
				p.variableDeclaration()
				}
			}
		}
		p.nextToken()
		
	}
	

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement{
	stmt := &ast.ReturnStatement{Token:p.curTok}
	p.nextToken()
	// currently not parsing expression

	for !p.curTokenIs(token.SEMICOLON){
		p.nextToken()
	} 
	return stmt
}

func (p * Parser) curTokenIs(t token.TokenType)bool{
	return p.curTok.Type == t 
}
func (p * Parser) nextTokenIs(t token.TokenType)bool{
	return p.nextTok.Type == t 
}

func (p * Parser) exceptNext(t token.TokenType) bool{
	if p.nextTokenIs(t){
		p.nextToken()
		return true
	}else{
		return false
	}
}

func (p * Parser) variableDeclarationAssign(){

}
func (p * Parser) variableDeclaration(){

}
