package ast
import (
	"C-Compiler/token"
)

type Node interface{
	tokenLiteral()string
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
	Token token.Token // always IDENT
	Value string
}
type DeclStatement struct{
	Token token.Token 
	Name *Identifier
	Value Expression
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


func (p *Program)tokenLiteral()string{
	if len(p.Stats)>0{
		return p.Stats[0].tokenLiteral()
	}else{
		return ""
	}
}