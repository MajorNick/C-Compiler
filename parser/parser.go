package parser

import (
	"C-Compiler/ast"
	"C-Compiler/lexer"
	"C-Compiler/token"
	"fmt"
	"strconv"
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

var precendences = map[token.TokenType]int{
	token.EQ : EQUALS,
	token.NOT_EQ : EQUALS,
	token.LTE: LESSGREATER,
	token.GTE: LESSGREATER,
	token.PLUS: SUM,
	token.MINUS: SUM,
	token.SLASH: PRODUCT,
	token.ASTERISK: PRODUCT,
	
}

func (p *Parser)nextPrecendence() int{
	if prec, ok := precendences[p.nextTok.Type];ok{
		return prec
	}else{
		return LOWEST
	}
}
func (p *Parser)curPrecendence() int{
	if prec, ok := precendences[p.curTok.Type];ok{
		return prec
	}else{
		return LOWEST
	}
}

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
	p.errors= []string{}
	// add parser fns
	p.prefixParserFns= make(map[token.TokenType]prefixParserFn)
	p.addPrefixFn(token.INT, p.parseIntegerLiteral)
	p.addPrefixFn(token.IDENT,p.parseIdentifier)
	p.addPrefixFn(token.BANG,p.parsePrefixExpression)
	p.addPrefixFn(token.MINUS,p.parsePrefixExpression)
	// infix parsers
	p.infixParserFns= make(map[token.TokenType]infixParserFn)
	p.addInfixFn(token.MINUS,p.parseInfixExpression)
	p.addInfixFn(token.PLUS,p.parseInfixExpression)
	p.addInfixFn(token.LTE,p.parseInfixExpression)
	p.addInfixFn(token.GTE,p.parseInfixExpression)
	p.addInfixFn(token.EQ,p.parseInfixExpression)
	p.addInfixFn(token.NOT_EQ,p.parseInfixExpression)
	p.addInfixFn(token.ASTERISK,p.parseInfixExpression)
	p.addInfixFn(token.SLASH,p.parseInfixExpression)
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

func (p * Parser) noPrefixParserFn(t token.TokenType) {
	msg := fmt.Sprintf("No Prefix Parser Function found for: %s",t)
	p.errors = append(p.errors, msg)
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
//	case token.INT,token.LONG,token.SHORT,token.CHAR:
//		return p.parseDecStatement()
//	case token.INTP,token.LONGP,token.SHORTP,token.CHARP,token.VOIDP:
//		return p.parseDecStatement()
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
	
	//p.nextToken()
	// currently not parsing expression
	
	for p.nextTokenIs(token.SEMICOLON){
		
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

func (p * Parser) parseExpressionStatement() *ast.ExpressionStatement{
	
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
		p.noPrefixParserFn(p.curTok.Type)
		return nil
	}
		leftExpression := prefix()
	for !p.nextTokenIs(token.SEMICOLON) && precendence < p.nextPrecendence(){
		infix := p.infixParserFns[p.nextTok.Type]
		if infix == nil{
			return leftExpression
		}
		p.nextToken()
		leftExpression = infix(leftExpression)
	}
	
		return leftExpression
	
}

func (p * Parser)parseIdentifier() ast.Expression{
	return &ast.Identifier{Token: p.curTok,Value: p.curTok.Literal}
}

func (p *Parser)parseIntegerLiteral() ast.Expression{
	lit := &ast.IntegerLiteral{Token: p.curTok}
	value, err := strconv.ParseInt(p.curTok.Literal,0,64)

	if err != nil{
		er := fmt.Sprintf("Couldn't Parse %q to Integer",p.curTok.Literal)
		p.errors = append(p.errors, er)
	}
	lit.Value = value
	return lit
}

func (p * Parser) parsePrefixExpression()ast.Expression{
	expression := &ast.PrefixExpression{
		Token: p.curTok,
		Operator: p.curTok.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

// infix Parser 
func (p* Parser) parseInfixExpression(left ast.Expression)ast.Expression{
	expression := &ast.InfixExpression{
		Token: p.curTok,
		Operator: p.curTok.Literal,
		Left: left,
	}
	precendence := p.nextPrecendence()
	
	p.nextToken()
	expression.Right = p.parseExpression(precendence)
	fmt.Println(expression.Left,expression.Right,expression.Operator)

	return expression
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