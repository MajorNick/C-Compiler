package analyzer

import (
	"C-Compiler/internal/ast"
)

type Analyzer struct {
	program  *ast.Program
	symTable *SymbolTable
	errors   []string
}

func New(program *ast.Program) *Analyzer {
	return &Analyzer{
		program:  program,
		symTable: NewSymbolTable(),
		errors:   []string{},
	}
}

func (a *Analyzer) Analyze() {
	for _, note := range a.program.Stats {
		a.EvaluateStatement(note)
	}
}

func (a *Analyzer) EvaluateStatement(node ast.Node) {

	switch n := node.(type) {
	case *ast.DeclStatement:
		a.EvaluateDeclVar(n)
	case *ast.ExpressionStatement:
		a.EvaluateExpression(n.Expression)
	case *ast.BlockStatement:
		a.EvaluateBlockStatement(n)
	case *ast.IfExpression:
		a.EvaluateIfExpression(n)
	case *ast.ReturnStatement:
		a.EvaluateReturnStatement(n)
	}
}

func (a *Analyzer) EvaluateDeclVar(stmt *ast.DeclStatement) {
	decl, ok := stmt.Statement.(*ast.VariableDecStatement)
	if ok {
		for _, v := range decl.Vars {
			a.symTable.Define(v.Ident, v.Type.Type)
			if v.Value != nil {
				a.EvaluateExpression(v.Value)
			}
		}
	}

	fn, ok := stmt.Statement.(*ast.FunctionLiteral)
	if ok {
		a.symTable.Define("main", fn.Token.Type)

		a.symTable = NewEnclosedSymbolTable(a.symTable)

		for _, arg := range fn.Arguments {
			a.symTable.Define(arg.Ident, arg.Type.Type)
		}

		a.EvaluateBlockStatement(fn.Body)

		a.symTable = a.symTable.Outer
	}
}

func (a *Analyzer) EvaluateExpression(exp ast.Expression) {
	switch e := exp.(type) {
	case *ast.Identifier:
		_, ok := a.symTable.Resolve(e.Value)
		if !ok {
			a.errors = append(a.errors, "Variable "+e.Value+" not defined")
		}
	case *ast.InfixExpression:
		a.EvaluateExpression(e.Left)
		a.EvaluateExpression(e.Right)
	case *ast.PrefixExpression:
		a.EvaluateExpression(e.Right)
	}
}

func (a *Analyzer) EvaluateBlockStatement(block *ast.BlockStatement) {
	a.symTable = NewEnclosedSymbolTable(a.symTable)
	defer func() { a.symTable = a.symTable.Outer }()

	for _, stmt := range block.Statements {
		a.EvaluateStatement(stmt)
	}
}

func (a *Analyzer) EvaluateIfExpression(ie *ast.IfExpression) {
	a.EvaluateExpression(ie.Condition)
	a.EvaluateStatement(ie.Consequence)
	if ie.Alternative != nil {
		a.EvaluateStatement(ie.Alternative)
	}
}

func (a *Analyzer) EvaluateReturnStatement(rs *ast.ReturnStatement) {
	if rs.ReturnValue != nil {
		a.EvaluateStatement(rs.ReturnValue)
	}
}
