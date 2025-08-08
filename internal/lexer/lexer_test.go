package lexer

import (
	"C-Compiler/internal/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	source := "{}(),;"
	tests := []struct {
		expected    token.TokenType
		expectedLit string
	}{
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(source)
	for _, v := range tests {
		tok := l.NextToken()
		if (tok.Type != v.expected) || (tok.Literal != v.expectedLit) {
			t.Errorf("WRONG!!! expected: %v, got: %v", v, tok)
		}
	}
}
func TestComments(t *testing.T) {
	source := "nika\n// nika \n nika 1  /* nikaaaaaaa */  fassfs"

	tests := []struct {
		expected    token.TokenType
		expectedLit string
	}{
		{token.IDENT, "nika"},
		{token.COMMENT, " nika "},
		{token.IDENT, "nika"},
		{token.NUMBER, "1"},
		{token.COMMENT, " nikaaaaaaa "},
		{token.IDENT, "fassfs"},
		{token.EOF, ""},
	}
	l := New(source)
	for _, v := range tests {
		tok := l.NextToken()
		if (tok.Type != v.expected) || (tok.Literal != v.expectedLit) {
			t.Errorf("WRONG!!! expected: %v, got: %v", v, tok)
		}
	}
}

func TestIntDeclaration(t *testing.T) {
	source := `
  	int b = 10;
   	long l=10;
`

	tests := []struct {
		expected    token.TokenType
		expectedLit string
	}{

		{token.INT, "int"},
		{token.IDENT, "b"},
		{token.ASSIGN, "="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.LONG, "long"},
		{token.IDENT, "l"},
		{token.ASSIGN, "="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
	}
	l := New(source)
	for i, v := range tests {
		tok := l.NextToken()
		if (tok.Type != v.expected) || (tok.Literal != v.expectedLit) {
			t.Errorf("WRONG!!! expected: %v, got: %v, on %d th test", v, tok, i)
		}
	}
}

func TestHalfCode(t *testing.T) {

	source := `    
  int* a;
  int b = 10;
  a = &b;
  return a;`

	tests := []struct {
		expected    token.TokenType
		expectedLit string
	}{
		{token.INTP, "int*"},
		{token.IDENT, "a"},
		{token.SEMICOLON, ";"},
		{token.INT, "int"},
		{token.IDENT, "b"},
		{token.ASSIGN, "="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.AMPERSAND, "&"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.RETURN, "return"},
		{token.IDENT, "a"},
		{token.SEMICOLON, ";"},
	}
	l := New(source)
	for i, v := range tests {
		tok := l.NextToken()
		if (tok.Type != v.expected) || (tok.Literal != v.expectedLit) {
			t.Errorf("WRONG!!! expected: %v, got: %v, on %d th test", v, tok, i)
		}
	}
}
func TestFunc(t *testing.T) {
	source := `    
  int main() {
    // Write C code here
    int* a;
    int b = 10;
    a = &b;
    return a;
}`

	tests := []struct {
		expected    token.TokenType
		expectedLit string
	}{
		{token.INT, "int"},
		{token.MAIN, "main"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.COMMENT, " Write C code here"},
		{token.INTP, "int*"},
		{token.IDENT, "a"},
		{token.SEMICOLON, ";"},
		{token.INT, "int"},
		{token.IDENT, "b"},
		{token.ASSIGN, "="},
		{token.NUMBER, "10"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.AMPERSAND, "&"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.RETURN, "return"},
		{token.IDENT, "a"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
	}
	l := New(source)
	for i, v := range tests {
		tok := l.NextToken()
		if (tok.Type != v.expected) || (tok.Literal != v.expectedLit) {
			t.Errorf("WRONG!!! expected: %v, got: %v, on %d th test", v, tok, i)
		}
	}

}
