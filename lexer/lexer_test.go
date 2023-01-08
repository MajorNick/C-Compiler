package lexer

func TestNextToken (t *testing.T){
  source := '{}(),;'
  tests := []struct {
    expected TokenType
    expectedLit string
  }{
  {token.LBRACE,"{"}
  {token.RBRACE,"}"}
  {token.LPAREN,"("}
  {token.RPAREN,")"}
  {token.COMMA,","}
  {token.SEMICOLON,";"}
  {token.EOF,""}
  }
  l := New(source)
  for _,v := range tests{
    tok := l.NextToken()
    if tok.TokenType != v.expected && tok.Literal = v.expectedLit{
      t.Fatalf("WRONG!!! expected: %v, got: %v", v,tok)
    }
  }
}
