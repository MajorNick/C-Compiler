package analyzer

import (
	"C-Compiler/internal/token"
)

type Symbol struct {
	Name string
	Type token.TokenType
}

type SymbolTable struct {
	Symbols map[string]Symbol
	Outer   *SymbolTable
}

func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{Symbols: s}
}

func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.Outer = outer
	return s
}

func (s *SymbolTable) Define(name string, t token.TokenType) Symbol {
	symbol := Symbol{Name: name, Type: t}
	s.Symbols[name] = symbol
	return symbol
}

func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.Symbols[name]
	if !ok && s.Outer != nil {
		obj, ok = s.Outer.Resolve(name)
		return obj, ok
	}
	return obj, ok
}
