package parser

import (
	"C-Compiler/ast"
	"C-Compiler/lexer"
	"C-Compiler/token"
	"fmt"
)
const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX 
	CALL
)


type(
	prefixParserFn func() ast.Expression
	infixParserFn func(ast.Expression) ast.Expression   //  argument is for left side of expression 
)

type Parser struct{
	l * lexer.Lexer
	curTok  token.Token
	nextTok  token.Token
	errors []string

	prefixParserFns map[token.TokenType]prefixParserFn
	infixParserFns map[token.TokenType]infixParserFn
}




func New(l *lexer.Lexer)*Parser{
	p := Parser{}
	p.l = l
	p.nextToken()
	p.nextToken()
	p.prefixParserFns= make(map[token.TokenType]prefixParserFn)
	p.addPrefixFn(token.IDENT,p.parseIdentifier)
	p.errors= []string{}
	return &p
}
func (p * Parser)nextToken(){
	p.curTok = p.nextTok
	p.nextTok = p.l.NextToken()
}

//errors 
func (p * Parser)Errors()[]string{
	return p.errors
}
func (p * Parser)nextError(t token.TokenType){
	err := fmt.Sprintf("expected %s token, but got %s !",t,p.nextTok.Type)
	p.errors = append(p.errors,err)
}


// end errors 
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
	case token.WHILE:

	case token.IF:
		return p.parseIfStatement()

	default:
		return p.parseExpressionStatement()
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





func (p * Parser) parseIfStatement() *ast.IfStatement{
	stmt := &ast.IfStatement{Token:p.curTok}
	if !p.exceptNext(token.LBRACE){
		// error in syntax
		p.nextError(token.LBRACE)
	}
	// parse expressions
	
	return stmt
}

func (p * Parser)parseExpressionStatement() *ast.ExpressionStatement{
	stmt := &ast.ExpressionStatement{Token: p.curTok}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.nextTokenIs(token.SEMICOLON){
		p.nextToken()
	}
	return stmt
}

func (p * Parser)parseExpression(precendence int) ast.Expression{
	prefix  := p.prefixParserFns[p.curTok.Type]
	if prefix == nil{
		return nil
	}else{
		return prefix()
	}
}

func (p * Parser)parseIdentifier() ast.Expression{
	return &ast.Identifier{Token: p.curTok,Value: p.curTok.Literal}
}


// helper function


func (p * Parser) variableDeclarationAssign(){

}
func (p * Parser) variableDeclaration(){

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

// expression parser FN helper

func (p * Parser) addPrefixFn(t token.TokenType,fn prefixParserFn ){
	p.prefixParserFns[t] = fn
}

func (p * Parser) addInfixFn(t token.TokenType,fn infixParserFn ){
	p.infixParserFns[t] = fn
}