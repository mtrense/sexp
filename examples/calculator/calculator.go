package calculator

import (
	"github.com/mtrense/sexp"
	"github.com/mtrense/sexp/stdlib"
)

func BuildContext() sexp.Context {
	return sexp.NewStaticRootContext().
		DefineFunction("add", stdlib.Add).
		DefineFunction("+", stdlib.Add).
		DefineFunction("times", stdlib.Times).
		DefineFunction("*", stdlib.Times).
		DefineFunction("format", stdlib.Sprintf)
}
