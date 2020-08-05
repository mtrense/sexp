package sexp

import (
	"fmt"

	"github.com/mtrense/parsertk/parser"
)

type Expression interface {
	Evaluate(ctx Context) Expression
}

type FunctionCall struct {
	functionName string
	arguments    []Expression
}

func (s *FunctionCall) Evaluate(ctx Context) Expression {
	ref := ctx.Lookup(s.functionName)
	switch r := ref.(type) {
	case *Function:
		return r.body(ctx, s.functionName, s.arguments...)
	}
	return nil
}

type Void struct{}

func (s *Void) Evaluate(ctx Context) Expression {
	return nil
}

type Literal struct {
	typ   parser.NodeType
	value interface{}
}

func String(s string) *Literal {
	return &Literal{
		typ:   TypeStringLiteral,
		value: s,
	}
}

func Number(n float64) *Literal {
	return &Literal{
		typ:   TypeNumberLiteral,
		value: n,
	}
}

func Bool(b bool) *Literal {
	return &Literal{
		typ:   TypeBooleanLiteral,
		value: b,
	}
}

func (s *Literal) Evaluate(ctx Context) Expression {
	return s
}

func (s *Literal) Type() parser.NodeType {
	return s.typ
}

func (s *Literal) Value() interface{} {
	return s.value
}

func (s *Literal) String() (string, bool) {
	if s.typ != TypeStringLiteral {
		return "", false
	}
	return s.value.(string), true
}

func (s *Literal) StringValue() string {
	return fmt.Sprintf("%v", s.value)
}

func (s *Literal) Number() (float64, bool) {
	if s.typ != TypeNumberLiteral {
		return 0, false
	}
	return s.value.(float64), true
}

func (s *Literal) NumberValue() float64 {
	n, _ := s.Number()
	return n
}

func (s *Literal) Bool() (bool, bool) {
	if s.typ != TypeBooleanLiteral {
		return false, false
	}
	return s.value.(bool), true
}

func (s *Literal) BoolValue() bool {
	b, _ := s.Bool()
	return b
}

type Error struct {
	err     error
	message string
}

func (s *Error) Evaluate(ctx Context) Expression {
	return s
}

func (s *Error) Message() string {
	return s.message
}

func DumpExpressions(expressions ...Expression) {
	for _, expression := range expressions {
		switch e := expression.(type) {
		case *Void:
			fmt.Printf("Void\n")
		case *FunctionCall:
			fmt.Printf("Function: %s (%d arguments)\n", e.functionName, len(e.arguments))
			DumpExpressions(e.arguments...)
		case *Literal:
			fmt.Printf("Literal: %v\n", e.value)
		}
	}
}
