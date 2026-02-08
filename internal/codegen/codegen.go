package codegen

import (
	"C-Compiler/internal/ast"
	"bytes"
	"fmt"
)

type Codegen struct {
	code bytes.Buffer
}

func New() *Codegen {
	return &Codegen{}
}

func (c *Codegen) Generate(program *ast.Program) string {
	c.code.WriteString("global main\n")
	c.code.WriteString("section .text\n\n")

	for _, stmt := range program.Stats {
		c.generateStatement(stmt)
	}

	return c.code.String()
}

func (c *Codegen) generateStatement(node ast.Statement) {
	if node == nil {
		return
	}
	switch stmt := node.(type) {
	case *ast.DeclStatement:
		c.generateStatement(stmt.Statement)

	case *ast.FunctionLiteral:

		// TODO: support other functions than main
		c.code.WriteString("main:\n")
		c.code.WriteString("  push rbp\n")
		c.code.WriteString("  mov rbp, rsp\n")

		c.generateBlock(stmt.Body)

		c.code.WriteString("  pop rbp\n")
		c.code.WriteString("  ret\n")

	case *ast.ReturnStatement:
		if stmt.ReturnValue != nil {
			if exprStmt, ok := stmt.ReturnValue.(*ast.ExpressionStatement); ok {
				c.generateExpression(exprStmt.Expression)
			}
		}
		c.code.WriteString("  pop rbp\n")
		c.code.WriteString("  ret\n")

	case *ast.ExpressionStatement:
		c.generateExpression(stmt.Expression)

	case *ast.BlockStatement:
		c.generateBlock(stmt)
	}
}

func (c *Codegen) generateBlock(block *ast.BlockStatement) {
	for _, stmt := range block.Statements {
		c.generateStatement(stmt)
	}
}

// TODO: optimize to directly pop from stack to rbx to avoid extra mov instruction for efficiency

func (c *Codegen) generateExpression(node ast.Expression) {
	if node == nil {
		return
	}
	switch exp := node.(type) {
	case *ast.IntegerLiteral:
		c.code.WriteString(fmt.Sprintf("  mov rax, %d\n", exp.Value))

	case *ast.InfixExpression:
		c.generateExpression(exp.Left)
		c.code.WriteString("  push rax\n")

		c.generateExpression(exp.Right)
		c.code.WriteString("  mov rbx, rax\n")
		c.code.WriteString("  pop rax\n")

		switch exp.Operator {
		case "+":
			c.code.WriteString("  add rax, rbx\n")
		case "-":
			c.code.WriteString("  sub rax, rbx\n")
		case "*":
			c.code.WriteString("  imul rax, rbx\n")
		case "/":
			c.code.WriteString("  cqo\n") // Sign extend rax to rdx:rax
			c.code.WriteString("  idiv rbx\n")
		}
	}
}
