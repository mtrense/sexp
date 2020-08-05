package boilerplate

import "github.com/mtrense/sexp"

var rootContext *sexp.StaticContext

func BuildContext() sexp.Context {
	return sexp.NewStaticRootContext()
	// rootContext.DefineFunction("add")
}
