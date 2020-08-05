package stdlib

import (
	"fmt"

	"github.com/mtrense/sexp"
)

var (
	Sprintf = func(ctx sexp.Context, name string, arguments ...sexp.Expression) sexp.Expression {
		arg0 := sexp.ResolveArguments(ctx, arguments[0])[0]
		var params []interface{}
		for _, arg := range sexp.ResolveArguments(ctx, arguments[1:]...) {
			switch c := arg.(type) {
			case *sexp.Literal:
				params = append(params, c.Value())
			}
		}
		formatString, _ := arg0.(*sexp.Literal).String()
		return sexp.String(fmt.Sprintf(formatString, params...))
	}
)
