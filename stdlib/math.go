package stdlib

import "github.com/mtrense/sexp"

var (
	Add = func(ctx sexp.Context, name string, arguments ...sexp.Expression) sexp.Expression {
		sum := 0.0
		for _, arg := range sexp.ResolveArguments(ctx, arguments...) {
			switch c := arg.(type) {
			case *sexp.Literal:
				nmb, ok := c.Number()
				if ok {
					sum += nmb
				}
			}
		}
		return sexp.Number(sum)
	}
	Times = func(ctx sexp.Context, name string, arguments ...sexp.Expression) sexp.Expression {
		product := 1.0
		for _, arg := range sexp.ResolveArguments(ctx, arguments...) {
			switch c := arg.(type) {
			case *sexp.Literal:
				nmb, ok := c.Number()
				if ok {
					product *= nmb
				}
			}
		}
		return sexp.Number(product)
	}
)
