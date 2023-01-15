package ast
import (
	"C-Compiler/token"
)

type Node interface{
	TokenLiteral()string
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
type DeclStatement struct{
	Token token.Token 
	Name *Identifier
	Value Expression
}
type ReturnStatement struct{
	Token token.Token
	ReturnValue Expression
}
func (rs * ReturnStatement)statementNode(){}
func (rs * ReturnStatement)TokenLiteral()string{
	return rs.Token.Literal
}

func (dc * DeclStatement)TokenLiteral() string{
	return dc.Token.Literal
}
//to satisfy interface 
func (dc * DeclStatement)statementNode() {

}
//to satisfy interface 
func (id * Identifier)statementNode(){}
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