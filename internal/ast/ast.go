package ast

import (
	"C-Compiler/internal/token"
	"bytes"
	//"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Expression interface {
	Node
	expressionNode()
}
type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Stats []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Stats {
		out.WriteString(s.String())
	}
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

type FunctionLiteral struct {
	Token     token.Token // return type of fn
	Arguments []*Variable
	Body      *BlockStatement
}

func (fl *FunctionLiteral) statementNode() {}
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}
func (fl *FunctionLiteral) String() string {
	/*
		var out bytes.Buffer
		args := []string{}
		for _,v := range  fl.Arguments{
			args = append(args, v.String())
		}
		out.WriteString(fl.TokenLiteral())
		out.WriteString("(")
		out.WriteString(strings.Join(args,", "))
		out.WriteString(")")
		out.WriteString(fl.Body.String())
	*/
	return "tmp"
}

type Variable struct {
	Type  token.Token
	Ident string
	Value Statement
}
type DeclStatement struct {
	Token     token.Token
	Statement Statement
}

type VariableDecStatement struct {
	Token token.Token
	Vars  []*Variable
}

func (vds *VariableDecStatement) statementNode() {}
func (vds *VariableDecStatement) TokenLiteral() string {
	return vds.Token.Literal
}
func (vds *VariableDecStatement) String() string {
	return "test"
}

// return
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Statement
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
func (id *ReturnStatement) String() string {
	return "test"
}

// Declare
func (dc *DeclStatement) TokenLiteral() string {
	return dc.Token.Literal
}

// to satisfy interface
func (dc *DeclStatement) statementNode() {
}
func (id *DeclStatement) String() string {
	return "test"
}

// to satisfy interface
func (id *Identifier) expressionNode() {}
func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}
func (id *Identifier) String() string {
	return id.Value
}

func (p *Program) TokenLiteral() string {
	if len(p.Stats) > 0 {
		return p.Stats[0].TokenLiteral()
	} else {
		return ""
	}
}

// If Statement

//expressions

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (exp *ExpressionStatement) statementNode() {}
func (exp *ExpressionStatement) TokenLiteral() string {
	return exp.Token.Literal
}
func (exp *ExpressionStatement) String() string {
	if exp != nil {
		return exp.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string {
	return il.TokenLiteral()
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}
func (pe *PrefixExpression) String() string {
	return "test"
}

type InfixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
	Left     Expression
}

func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (oe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}
func (b *Boolean) String() string {
	return b.Token.Literal
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

type BlockStatement struct {
	Token      token.Token //in most cases it will be {
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
