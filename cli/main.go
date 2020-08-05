package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mtrense/parsertk/lexer"
	"github.com/mtrense/parsertk/parser"
	"github.com/mtrense/sexp"
	"github.com/mtrense/sexp/examples/boilerplate"
	"github.com/mtrense/sexp/examples/calculator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "none"
	commit  = "none"
	app     = &cobra.Command{
		Use:   "sexp",
		Short: "S-Exp Parser examples",
	}
	cmdRun = &cobra.Command{
		Use:   "run",
		Short: "Run examples",
		Run:   executeRun,
	}
	cmdPrint = &cobra.Command{
		Use:   "print",
		Short: "Print AST and info from examples",
		Run:   executePrint,
	}
	cmdVersion = &cobra.Command{
		Use:   "version",
		Short: "Show sexps version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s (ref: %s)\n", version, commit)
		},
	}
)

func init() {
	app.AddCommand(cmdRun, cmdPrint, cmdVersion)

	cmdRun.PersistentFlags().StringP("example", "e", "calculator", "Example context to load")
	cmdRun.MarkPersistentFlagRequired("example")

	viper.SetEnvPrefix("SEXP")
	viper.AutomaticEnv()
}

func main() {
	if err := app.Execute(); err != nil {
		panic(err)
	}
}

func executeRun(cmd *cobra.Command, args []string) {
	var rootCtx sexp.Context
	example, _ := cmd.Flags().GetString("example")
	switch example {
	case "calculator", "calc", "c":
		rootCtx = calculator.BuildContext()
	case "boilerplate", "bp":
		rootCtx = boilerplate.BuildContext()
	}
	results := evaluate(rootCtx, parseToExpressions(readCode())...)
	for pos, result := range results {
		var formatted string
		switch r := result.(type) {
		case *sexp.Literal:
			formatted = fmt.Sprintf("%v", r.Value())
		case *sexp.Error:
			formatted = fmt.Sprintf("%v", r.Message())
		case *sexp.Void:
			formatted = "()"
		}
		fmt.Printf(" [ %d ]\n%s\n", pos, formatted)
	}
}

func executePrint(cmd *cobra.Command, args []string) {
	print(readCode())
}

func evaluate(rootCtx sexp.Context, expressions ...sexp.Expression) []sexp.Expression {
	var results []sexp.Expression
	for _, expression := range expressions {
		result := expression.Evaluate(rootCtx)
		results = append(results, result)
	}
	return results
}

func readCode() string {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func parseToExpressions(code string) []sexp.Expression {
	return sexp.ParseProgram(lexer.StringReader(code))
}

func print(code string) {
	lexer.LexStatic(lexer.StringReader(code), sexp.DebugPrinter.Visit, sexp.TokenTypeEOF, sexp.TokenTypeError, sexp.SexpTokens...)
	fmt.Println()
	p := parser.NewParser(sexp.TypeRoot)
	sexp.RegisterNodeFactories(&p)
	lexer.LexStatic(lexer.StringReader(code), p.Visit, sexp.TokenTypeEOF, sexp.TokenTypeError, sexp.SexpTokens...)
	parser.DumpTree(p.RootNode())
	fmt.Println()
	sexp.DumpExpressions(parseToExpressions(code)...)
}
