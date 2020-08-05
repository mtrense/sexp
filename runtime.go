package sexp

type Context interface {
	Lookup(name string) Reference
}

type Reference interface {
}

type StaticContext struct {
	inherit     bool
	parent      Context
	symbolTable map[string]Reference
}

func NewStaticRootContext() *StaticContext {
	return &StaticContext{
		symbolTable: make(map[string]Reference),
	}
}

func (s *StaticContext) AddChild(inherit bool) Context {
	return &StaticContext{
		inherit:     inherit,
		parent:      s,
		symbolTable: make(map[string]Reference),
	}
}

func (s *StaticContext) Lookup(name string) Reference {
	ref, ok := s.symbolTable[name]
	if ok {
		return ref
	}
	if s.parent != nil {
		return s.parent.Lookup(name)
	}
	return nil
}

func (s *StaticContext) Parent() Context {
	return s.parent
}

type FunctionBody = func(ctx Context, name string, arguments ...Expression) Expression

type Function struct {
	body FunctionBody
}

func (s *StaticContext) DefineFunction(name string, def FunctionBody) *StaticContext {
	s.symbolTable[name] = &Function{
		body: def,
	}
	return s
}

func ResolveArguments(ctx Context, arguments ...Expression) []Expression {
	var results = make([]Expression, len(arguments))
	for i, arg := range arguments {
		switch c := arg.(type) {
		case *FunctionCall:
			results[i] = c.Evaluate(ctx)
		default:
			results[i] = c
		}
	}
	return results
}
