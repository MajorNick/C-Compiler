package token

type TokenType string

type Token struct{
  Type TokenType
  Literal string
}
const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"
  IDENT = "IDENT"
  INT = "INT"
  // operators
  ASSIGN  = "="
  PLUS  =  "+"
  MINUS =  "-"
  ASTERISK = "*"
  SLASH = "/"
  BANG = "!"
  LEFT = "<"
  RIGHT = ">"
  //

  COMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  //statements
  FUNCTION = "FUNCTION"
  LET = "LET"
  TRUE = "TRUE"
  FALSE = "FALSE"
  IF = "IF"
  ELSE = "ELSE"
  RETURN = "RETURN"
)
