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
  ENDLN = "ENDLN"
  COMMENT ="COMMENT"
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

  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  //statements
  VAR  = "VAR"
  VOID = "VOID"
  TRUE = "TRUE"
  FALSE = "FALSE"
  IF = "IF"
  ELSE = "ELSE"
  RETURN = "RETURN"

  LONG = "LONG"
  SHORT = "SHORT"
  INT = "INT"
  CHAR = "CHAR"
  

  WHILE = "WHILE"
  FORLOOP = "FOR"
  STRUCT = "STRUCT"

)
var keywords = map[string]TokenType{
  "void": VOID,
  "for": FORLOOP,
  "while": WHILE,
  "struct": STRUCT,
  "char": CHAR,
  "int": INT,
  "long": LONG,
  "short": SHORT,
}
//func LookForIdent
