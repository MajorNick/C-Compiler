package ast
import (
	"C-Compiler/token"
	
)

type Node interface{
	TokenLiteral()string
//	String()string
}

type Expression interface{
	Node
	expressionNode()
}
type Statement interface{
	Node
	statementNode()
}

type Program struct{
	Stats []Statement
}
type Identifier struct{
	Token token.Token 
	Value string
}

// statemet structs

type DeclStatement struct{
	Token token.Token 
	Name *Identifier
	Value Expression
}

// return 
type ReturnStatement struct{
	Token token.Token
	ReturnValue Expression
}
func (rs * ReturnStatement)statementNode(){}
func (rs * ReturnStatement)TokenLiteral()string{
	return rs.Token.Literal
}





 // Declare 
func (dc * DeclStatement)TokenLiteral() string{
	return dc.Token.Literal
}
//to satisfy interface 
func (dc * DeclStatement)statementNode() {

}

//to satisfy interface 
func (id * Identifier)expressionNode(){}
func (id * Identifier)TokenLiteral()string{
	return id.Token.Literal
}


func (p *Program)TokenLiteral()string{
	if len(p.Stats)>0{
		return p.Stats[0].TokenLiteral()
	}else{
		return ""
	}
}




// If Statement

type IfStatement struct{
	Token token.Token
	BoolExpression Expression
}

func (ifs * IfStatement)statementNode(){}
func (ifs * IfStatement)TokenLiteral()string{
	return ifs.Token.Literal
}


//expressions 

type ExpressionStatement struct{
	Token token.Token
	Expression Expression
}
func (exp * ExpressionStatement)statementNode(){}
func (exp * ExpressionStatement)TokenLiteral()string{
	return exp.Token.Literal
}

type IntegerLiteral struct{
	Token token.Token
	Value int64
}
func (il * IntegerLiteral)expressionNode(){}
func (il * IntegerLiteral)TokenLiteral()string{
	return il.Token.Literal
}