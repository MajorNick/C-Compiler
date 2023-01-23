package parser

import (
	"C-Compiler/ast"
	"C-Compiler/lexer"
	"fmt"

	"testing"
)
var dataTypes = map[string]bool{
	"char":true,
	"int":true,
	"short":true,
	"long":true,
	"void*":true,
	"char*":true,
	"int*":true,
	"short*":true,
	"long*":true,
}

func TestVarDecStatement(t *testing.T){
	input := `
	int m = 0;
	char c = 'a';
	long l=10;
	`
	
	l := lexer.New(input)
	p := New(l)
		fmt.Println("FSAS")
	program := p.ParseProgram()

	if program == nil{
		t.Fatalf("nil after ParseProrgram")
	}
	if len(program.Stats) != 3{
		t.Fatalf("Wrong number of statements after ParseProgram")
	}

	tests := []struct{
		expectedType string
		expectedIdent string
	}{
		{"int","m"},
		{"char","c"},
		{"long","l"},
	}


	for i := range tests{
		state := program.Stats[i]
		if !wrapperTestDecStat(t,state,tests[i].expectedIdent,tests[i].expectedType){
				return 
		}
	}


}



func wrapperTestDecStat(t *testing.T,s ast.Statement,name string,tp string)bool {
	if _,ok := dataTypes[s.TokenLiteral()]; !ok{
		t.Errorf("Wrong  type: %v", s.TokenLiteral())
	}
	
	decState,ok := s.(*ast.DeclStatement)
	if !ok {
		t.Errorf("s Isn't Declare Statemen!")
		return false
	}
	if s.TokenLiteral() != tp {
		t.Errorf("Wrong Type: expected %s got : %s ",tp,s.TokenLiteral())
		return false
	}
	if decState.Name.Value != name{
		t.Errorf("decState.Name.Value not '%s'. got=%s", name, decState.Name.Value)
		return false
	} 
	if decState.Name.TokenLiteral() != name{
		t.Errorf("decState.Name.TokenLiteral not '%s'. got=%s", name, decState.Name.TokenLiteral())
		return false
	}
	return true
}

func TestParseReturn(t *testing.T){
	source := ` 
	return 5;
	return 42;	
	`

	l := lexer.New(source)
	p := New(l)

	program := p.ParseProgram()
	
	if len(program.Stats) != 2{
		t.Fatalf("Expected 2 Statement but got: %d",len(program.Stats))
	}

	for _,state := range program.Stats{
		returnStmt,ok := state.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Expected ast.ReturnStatement! got: %T",state)
			continue
		}
		if returnStmt.TokenLiteral() != "return"{
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",returnStmt.TokenLiteral())
		} 
	}	
}

func TestIdentifierExpression(t *testing.T){
	source := "variable;"
	l := lexer.New(source)
	p := New(l)
	program := p.ParseProgram()

	checkParserError(t,p)

	if len(program.Stats) != 1{
		t.Fatalf("expected 1 Statement, But Got %d", len(program.Stats))
	}
	stmt, ok := program.Stats[0].(*ast.ExpressionStatement)
	if !ok{
		t.Fatalf("Program.Stats[0] isn't Expression Statement Type!")
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok{
		t.Fatalf("Is n't Identifier Type!")
	}
	if ident.Value != "variable"{
		t.Fatalf("expected Token Literal: varibale, but got %s",ident.Value)
	}
	if ident.TokenLiteral() != "variable"{
		t.Fatalf("expected Token Literal: varibale, but got %s",ident.TokenLiteral())
	}
}

func TestNumberLiteral(t *testing.T){
	
}






//check errors
func checkParserError(t *testing.T, p * Parser){
	errors := p.Errors()
	if len(errors) == 0{
		return
	}else{	
		for i,v := range errors{
			t.Errorf("%d) %q",i,v)
		}
		
		t.FailNow()
	}
}