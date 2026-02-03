package parser

import (
	"C-Compiler/internal/ast"
	"C-Compiler/internal/lexer"
	"C-Compiler/internal/token"
	"fmt"

	"testing"
)

var dataTypes = map[string]bool{
	"char":   true,
	"int":    true,
	"short":  true,
	"long":   true,
	"void*":  true,
	"char*":  true,
	"int*":   true,
	"short*": true,
	"long*":  true,
}

type decTest struct {
	expectedType         string
	expectedIdent        string
	expectedValueLiteral string
}

func TestVarDecStatement(t *testing.T) {
	input := `
		int m = 0;
		long l=10;
		`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("nil after ParseProrgram")
	}
	if len(program.Stats) != 2 {
		t.Fatalf("Wrong number of statements after ParseProgram")
	}

	tests := []decTest{
		{"int", "m", "0"},
		{"long", "l", "10"},
	}

	for i := range tests {
		state := program.Stats[i]
		if !wrapperTestDecStat(t, state, tests[i]) {
			return
		}
	}
}

func wrapperTestDecStat(t *testing.T, s ast.Statement, test decTest) bool {
	if _, ok := dataTypes[s.TokenLiteral()]; !ok {
		t.Errorf("Wrong  type: %v", s.TokenLiteral())
	}

	decState, ok := s.(*ast.DeclStatement).Statement.(*ast.VariableDecStatement)
	if !ok {
		t.Errorf("s Isn't variable Declare Statemen!")
		return false
	}
	if decState.TokenLiteral() != test.expectedType {
		t.Errorf("Wrong Type: expected %s got : %s ", test.expectedType, s.TokenLiteral())
		return false
	}
	firstVar := decState.Vars[0]

	if firstVar.Ident != test.expectedIdent {
		t.Errorf("variable name  not  '%s'. got=%s", test.expectedIdent, firstVar.Ident)
		return false
	}

	if firstVar.Value.String() != test.expectedValueLiteral {
		t.Errorf("variable value not  '%s'. got=%s", test.expectedValueLiteral, firstVar.Value.String())
		return false
	}

	return true
}

func TestParseReturn(t *testing.T) {
	source := ` 
	return 5;
	return 42;	
	`

	l := lexer.New(source)
	p := New(l)

	program := p.ParseProgram()

	if len(program.Stats) != 2 {
		t.Fatalf("Expected 2 Statement but got: %d", len(program.Stats))
	}

	for _, state := range program.Stats {
		returnStmt, ok := state.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Expected ast.ReturnStatement! got: %T", state)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	source := "variable;"
	l := lexer.New(source)
	p := New(l)
	program := p.ParseProgram()

	checkParserError(t, p)

	if len(program.Stats) != 1 {
		t.Fatalf("expected 1 Statement, But Got %d", len(program.Stats))
	}
	stmt, ok := program.Stats[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Program.Stats[0] isn't Expression Statement Type!")
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("Is n't Identifier Type!")
	}
	if ident.Value != "variable" {
		t.Fatalf("expected Value: varibale, but got %s", ident.Value)
	}
	if ident.TokenLiteral() != "variable" {
		t.Fatalf("expected Token Literal: varibale, but got %s", ident.TokenLiteral())
	}
}

func TestNumberLiteral(t *testing.T) {
	source := "5;"
	l := lexer.New(source)
	p := New(l)

	program := p.ParseProgram()
	checkParserError(t, p)

	if len(program.Stats) != 1 {
		t.Fatalf("expected 1 Expression, But Got %d", len(program.Stats))
	}
	stmt, ok := program.Stats[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Program.Stats[0] isn't Expression Statement Type!")
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("Is not't IntegerLitral Type!")
	}
	if literal.Value != 5 {
		t.Fatalf("expected Value: varibale, but got %d", literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Fatalf("expected Token Literal: varibale, but got %s", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	Tests := []struct {
		input    string
		operator string
		IntValue int64
	}{
		{"!23;", "!", 23},
		{"-7;", "-", 7},
	}
	for _, test := range Tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserError(t, p)
		if len(program.Stats) != 1 {
			t.Fatalf("expected 1 Expression, But Got %d", len(program.Stats))
		}
		stmt, ok := program.Stats[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Program.Stats[0] isn't Expression Statement Type!")
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt Is not't PrefixExpression Type!")
		}
		if exp.Operator != test.operator {
			t.Fatalf("expected exp.Operator  '%s'. got=%s", test.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, test.IntValue) {
			return
		}
	}
}
func testIntegerLiteral(t *testing.T, exp ast.Expression, value int64) bool {
	integer, ok := exp.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("exp isn't ast.IntegerLiteral Type. Got: %T", exp)
		return false
	}
	if integer.Value != value {
		t.Errorf("Wrong integer value. Got %d, expected:%d!", integer.Value, value)
		return false
	}

	return true
}
func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp isn't *ast.Identifier type. Got := %T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("ident.Value expected %s, got %s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.Value expected %s, got %s", value, ident.TokenLiteral())
		return false
	}
	return true
}

func TestParseInfixExpressions(t *testing.T) {
	Tests := []struct {
		source   string
		left     int64
		operator string
		right    int64
	}{
		{"2+2;", 2, "+", 2},
		{"2 * 2;", 2, "*", 2},
		{"2 /2;", 2, "/", 2},
		{"2 == 2;", 2, "==", 2},
		{"2 >= 2;", 2, ">=", 2},
		{"2 <= 2;", 2, "<=", 2},
	}

	for _, test := range Tests {
		l := lexer.New(test.source)
		p := New(l)
		program := p.ParseProgram()

		checkParserError(t, p)
		if len(program.Stats) != 1 {
			t.Fatalf("expected 1 Expression, But Got %d", len(program.Stats))
		}
		stmt, ok := program.Stats[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Program.Stats[0] isn't Expression Statement Type!")
		}
		exp, ok := stmt.Expression.(*ast.InfixExpression)

		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Expression)
		}
		if !testIntegerLiteral(t, exp.Left, test.left) {
			return
		}
		if exp.Operator != test.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", test.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, test.right) {
			return
		}
	}
}

func TestOperatorPrecendence(t *testing.T) {
	Tests := []struct {
		source   string
		expected string
	}{
		{
			"-a*b;",
			"((-a)*b)\n",
		},
		{
			"!-a;",
			"!(-a)\n",
		}, {
			"a+b/c;",
			"(a+(b/c))\n",
		}, {
			"a + b * c + d / e - f;",
			"(((a + (b * c)) + (d / e)) - f)\n",
		}, {

			"c+-a;",
			"(c+(-a))\n",
		},
	}

	for _, test := range Tests {
		l := lexer.New(test.source)
		p := New(l)
		program := p.ParseProgram()
		checkParserError(t, p)
		if test.expected != program.String() {

		}
	}
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	_, ok := expected.(string)
	fmt.Printf("%T", expected)
	fmt.Println(ok)
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, int64(v))
	case string:

		return testIdentifier(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}

	t.Errorf("type can't handled got: %T", exp)
	return false
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{},
	operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.OperatorExpression. got=%T(%s)", exp, exp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}
	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}
	return true
}

func TestBooleanExpressiong(t *testing.T) {
	source := "false;"

	l := lexer.New(source)
	p := New(l)
	program := p.ParseProgram()

	checkParserError(t, p)
	stmt, ok := program.Stats[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Program.Stats[0] isn't Expression Statement Type! Got:%T", stmt)
	}
	b, ok := stmt.Expression.(*ast.Boolean)

	if !ok {
		t.Fatalf("Is not't Boolean Type!Got:%T", b)
	}

	if b.Value == true {
		t.Fatalf("expected Value: \"false\", but got \"%t\"", b.Value)
	}
	if b.TokenLiteral() != "false" {
		t.Fatalf("expected Token Literal: false, but got %s", b.TokenLiteral())
	}
}
func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	b, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("Exp isn't ast.Boolean Type, Got:%T", b)
		return false
	}

	if value != b.Value {
		t.Errorf("Wrong Exp Value! Expected:%t, Got:%t", value, b.Value)
		return false
	}

	if b.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("Wrong Exp.TokenLiteral! Expected:%s, Got:%s", fmt.Sprintf("%t", value), b.TokenLiteral())
		return false
	}

	return true
}

func TestOperatorPrecendenceParsing(t *testing.T) {
	Test := []struct {
		source   string
		expected string
	}{
		{"true;",
			"true"},
		{"false;",
			"false"},

		{
			"2>1 == false;",
			"(2 > (1 == false))",
		},
		{
			"1 < 2 == true;",
			"(1 < (2 == true))",
		},
	}

	for _, test := range Test {

		l := lexer.New(test.source)
		p := New(l)
		program := p.ParseProgram()

		checkParserError(t, p)

		if program.String() != test.expected+"\n" {
			t.Errorf("wrong String. Expected: %q, Got: %q", test.expected+"\n", program.String())
		}

	}

}
func TestIfExpression(t *testing.T) {
	source := `if ( m > 5 ){
		m == b;
		}else{
		m == 3;
		}`
	l := lexer.New(source)
	p := New(l)
	program := p.ParseProgram()

	checkParserError(t, p)
	if len(program.Stats) != 1 {
		t.Fatalf("Wrong Program.Stats size expected:1 Got:%d", len(program.Stats))
	}
	stmt, ok := program.Stats[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Program.Stats[0] isn't Expression Statement Type! Got:%T", stmt)
	}

	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt isn't IFExpression  Type! Got:%T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "m", ">", 5) {
		return
	}

	cons, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T", exp.Consequence.Statements[0])
	}

	inf, ok := cons.Expression.(*ast.InfixExpression)

	if !ok {
		t.Fatalf("cons is not ast.infixExpression. got=%T", cons.Expression)
	}

	if !testInfixExpression(t, inf, "m", "==", "b") {
		return
	}
	if exp.Alternative == nil {
		t.Errorf("exp.Alternative.Statements was not nil. got=%+v", exp.Alternative)
	}
	alt, ok := exp.Alternative.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T", exp.Alternative.Statements[0])
	}

	altInf, ok := alt.Expression.(*ast.InfixExpression)

	if !ok {
		t.Fatalf("cons is not ast.infixExpression. got=%T", alt.Expression)
	}

	if !testInfixExpression(t, altInf, "m", "==", 3) {
		return
	}

}

func TestDecStatement(t *testing.T) {
	source := `
	int a=5+23,b=523,c;
	`
	l := lexer.New(source)
	p := New(l)
	program := p.ParseProgram()

	//t.Fatalf("FAS")
	checkParserError(t, p)

	if len(program.Stats) != 1 {
		t.Fatalf("Wrong number of statements.")
	}
	st, ok := program.Stats[0].(*ast.DeclStatement)
	if !ok {
		t.Fatalf("program stat[0] isnt Statement")
	}
	stmt, ok := st.Statement.(*ast.VariableDecStatement)

	if !ok {
		t.Fatalf("st isnt VariableDecStatement")
	}
	if stmt.Vars[0].Ident != "a" {
		t.Fatalf("on 0 index expected a but got %s", stmt.Vars[0].Value.TokenLiteral())
	}
	if stmt.Vars[0].Value.String() != "(5 + 23)" {
		t.Fatalf("Expected value (5 + 23) but got :%s", stmt.Vars[0].Value.String())
	}
	if stmt.Vars[1].Ident != "b" {
		t.Fatalf("on 0 index expected b but got %s", stmt.Vars[0].Value.TokenLiteral())
	}
	if stmt.Vars[1].Value.String() != "523" {
		t.Fatalf("Expected value 523 but got :%s", stmt.Vars[0].Value.String())
	}
	if stmt.Vars[2].Ident != "c" {
		t.Fatalf("on 0 index expected c but got %s", stmt.Vars[0].Value.TokenLiteral())
	}
}

func TestFunctionLiteral(t *testing.T) {
	source := `
	int sum(int a,int b){
		return a+b;
	}
	`
	l := lexer.New(source)
	p := New(l)
	program := p.ParseProgram()
	checkParserError(t, p)
	if len(program.Stats) != 1 {
		t.Fatalf("expected  ! Statement but got %d", len(program.Stats))
	}
	stmt, ok := program.Stats[0].(*ast.DeclStatement)
	if !ok {
		t.Fatalf("program stat[0] isnt Expression Statement Got: %T", program.Stats[0])
	}
	fn, ok := stmt.Statement.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt isn't FunctionLiteral Got: %T", stmt)
	}
	if len(fn.Arguments) != 2 {
		t.Fatalf("expected 2 Arguments but Got: %d", len(fn.Arguments))
	}
	arg1 := fn.Arguments[0]

	if arg1.Ident != "a" {
		t.Fatalf("error in argument parsing expected: a got: %s", arg1.Ident)
	}
	if arg1.Type.Type != token.INT {
		t.Fatalf("error in argument parsing expected: token.INT got: %v", arg1.Type)
	}
	arg2 := fn.Arguments[1]
	if arg2.Ident != "b" {
		t.Fatalf("error in argument parsing expected: b got: %s", arg2.Ident)
	}
	if arg2.Type.Type != token.INT {
		t.Fatalf("error in argument parsing expected: token.INT got: %v", arg2.Type)
	}

	if len(fn.Body.Statements) != 1 {
		t.Fatalf("expected  1 Statement but got %d", len(fn.Body.Statements))
	}
	ret, ok := fn.Body.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("Function's body isn't ast.ExpressionStatement Got:= %T", fn.Body.Statements[0])
	}

	if ret.ReturnValue.String() != "(a + b)" {
		t.Fatalf("Fn's return value error. expected: (a + b). got: %s", ret.ReturnValue.String())
	}

}

// check errors
func checkParserError(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	} else {
		for i, v := range errors {
			t.Errorf("%d) %q", i, v)
		}

		t.FailNow()
	}
}
