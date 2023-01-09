package lexer
import(
  "testing"
  "C-Compiler/token"
)

func TestNextToken (t *testing.T){
  source := "{}(),;"
  tests := []struct {
    expected token.TokenType
    expectedLit string
  }{
  {token.LBRACE,"{"},
  {token.RBRACE,"}"},
  {token.LPAREN,"("},
  {token.RPAREN,")"},
  {token.COMMA,","},
  {token.SEMICOLON,";"},
  {token.EOF,""},
  }
  l := New(source)
  for _,v := range tests{
    tok := l.NextToken()
    if (tok.Type != v.expected) || (tok.Literal != v.expectedLit){
      t.Fatalf("WRONG!!! expected: %v, got: %v", v,tok)
    }
  }
}
 func TestComments (t *testing.T){
  source := "nika\n// nika \n nika 1  /* nikaaaaaaa */  fassfs"
  
    
  tests := []struct {
    expected token.TokenType
    expectedLit string
  }{
    {token.IDENT,"nika"},
    // {token.ENDLN,"\n"},
     {token.COMMENT,"\n"},
     {token.IDENT,"nika"},
     {token.INT,"1"},
     {token.COMMENT,"/"},
     {token.IDENT,"fassfs"},
     {token.EOF,""},
  }
  l := New(source)
  for _,v := range tests{
    tok := l.NextToken()
    if (tok.Type != v.expected) || (tok.Literal != v.expectedLit){
      t.Fatalf("WRONG!!! expected: %v, got: %v", v,tok)
    }
  }
}
