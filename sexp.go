package sexp

import (
	"github.com/mtrense/parsertk/lexer"
	"github.com/mtrense/parsertk/parser"
)

func ParseProgram(sourceCode lexer.BufferedRuneReader) []Expression {
	p := parser.NewParser(TypeRoot)
	RegisterNodeFactories(&p)
	lexer.LexStatic(sourceCode, p.Visit, TokenTypeEOF, TokenTypeError, SexpTokens...)
	var expressions []Expression
	for _, node := range p.RootNode().Children() {
		expressions = append(expressions, GenerateProgramNode(node))
	}
	return expressions
}

func GenerateProgramNode(cn *parser.Node) Expression {
	switch cn.Type() {
	case TypeBooleanLiteral, TypeNumberLiteral, TypeStringLiteral:
		return &Literal{
			typ:   cn.Type(),
			value: cn.Value(),
		}
	case TypeList:
		if len(cn.Children()) == 0 {
			return &Void{}
		}
		fc := FunctionCall{}
		fc.functionName = cn.Children()[0].Value().(string)
		for _, child := range cn.Children()[1:] {
			fc.arguments = append(fc.arguments, GenerateProgramNode(child))
		}
		return &fc
	}
	return nil
}
