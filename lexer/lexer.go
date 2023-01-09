package lexer

import (
	"C-Compiler/token"
	"fmt"
)
type Lexer struct{
  source string
  pos int
  readPos int
  ch byte
}
func New(code string)*Lexer{
   l := &Lexer{source: code}
   l.readChar()
   return l
}

func (l *Lexer)readChar(){
  if l.readPos >= len(l.source){
    l.ch = 0
  }else{
    l.ch = l.source[l.readPos]
  }

  l.pos = l.readPos

  l.readPos++
}

func newToken(tokenType token.TokenType,ch byte) token.Token{
  return token.Token{Type: tokenType,Literal : string(ch)}
}

func (l *Lexer)NextToken() token.Token{
  
  l.skipSpaces()
  
 
  var tok token.Token
  switch l.ch{
  
  case '=':
    tok = newToken(token.ASSIGN,l.ch)
  case '{':
    tok = newToken(token.LBRACE,l.ch)
  case '}':
    tok = newToken(token.RBRACE,l.ch)
  case '(':
    tok = newToken(token.LPAREN,l.ch)
  case ')':
    tok = newToken(token.RPAREN,l.ch)
  case ',':
    tok = newToken(token.COMMA,l.ch)
  case ';':
    tok = newToken(token.SEMICOLON,l.ch)
  case '+':
    tok = newToken(token.PLUS,l.ch)
  case '-':
    tok = newToken(token.MINUS,l.ch)
  case '*':
    tok = newToken(token.ASTERISK,l.ch)
  case '/':
   
    l.readChar() 
    // 1 line comment
     
    if l.ch == '/'{
     
      for l.ch != '\n'{
        l.readChar()
      }
      return newToken(token.COMMENT,l.ch)
    }else{
      if l.ch == '*'{
        l.readChar()
      
        for !(l.ch == '*' && l.source[l.readPos]== '/'){
          l.readChar()
          if l.source[l.pos] == 0{
            tok.Literal = ""
            tok.Type = token.EOF 
            l.readChar()
            return tok 
          }
        }
        l.readChar()
        l.readChar()
       
        return newToken(token.COMMENT,'/')
      }
      fmt.Println(l.pos)
    }
  case '>':
    tok = newToken(token.RIGHT,l.ch)
  case '<':
    tok = newToken(token.LEFT,l.ch)
  case '!':
    tok = newToken(token.BANG,l.ch)
  
  case 0:
    
    tok.Literal = ""
    tok.Type = token.EOF
    
  default:
   
    if isLetter(l.ch){
      tok.Literal = l.readIdentifier()
      tok.Type = token.IDENT
      return  tok
    }else{
    if isDigit(l.ch){
          tok.Type = token.INT
          tok.Literal = l.readNumber()
          return tok
        }else{
            tok = newToken(token.ILLEGAL,l.ch)
        }
    }
  }
  l.readChar()
  return tok
}

func (l *Lexer)readIdentifier()string{
  pos := l.pos
  for isLetter(l.ch){
    l.readChar()
  }
  return l.source[pos:l.pos]
}
func isLetter(b byte)bool{
  if (b>='a'&&b<='z')||(b>='A'&&b<='Z')||b == '_'{
    return true
  }else{
    return false
  }
}

func isDigit(b byte) bool{
    return (b>= '0'&&b<= '9')
}
func (l *Lexer)readNumber()string{
  pos := l.pos
  for isDigit(l.ch){
    l.readChar()
  }
  return l.source[pos:l.pos]
}
func (l *Lexer)skipSpaces(){
  for l.ch == '\n' || l.ch == '\t' || l.ch ==' '{
    l.readChar()
    
  }

}

