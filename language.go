package sexp

import (
	"regexp"
	"strconv"
	"unicode"

	"github.com/mtrense/parsertk/lexer"
	"github.com/mtrense/parsertk/parser"
)

const (
	TokenTypeStart      lexer.TokenType = "START"
	TokenTypeWhitespace lexer.TokenType = "WS"
	TokenTypeSymbol     lexer.TokenType = "SYMBOL"
	TokenTypeString     lexer.TokenType = "STRING"
	TokenTypeNumber     lexer.TokenType = "NUMBER"
	TokenTypeBoolean    lexer.TokenType = "BOOL"
	TokenTypeEnd        lexer.TokenType = "END"
	TokenTypeEOF        lexer.TokenType = "EOF"
	TokenTypeError      lexer.TokenType = "ERR"
)

var (
	SexpTokens = []lexer.TokenConsumer{
		lexer.ConsumeSingleRune(TokenTypeStart, '('),
		lexer.ConsumeSingleRune(TokenTypeEnd, ')'),
		lexer.ConsumeText(TokenTypeBoolean, "true"),
		lexer.ConsumeText(TokenTypeBoolean, "false"),
		lexer.ConsumeRegexpValidated(regexp.MustCompile("^(0|0\\.\\d+|[1-9]\\d*(\\.\\d+)?)$"), lexer.ConsumeRunes(TokenTypeNumber, "0123456789.")),
		lexer.ConsumeRunes(TokenTypeSymbol, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_*+!?$#%&."),
		lexer.ConsumeCharacterClass(TokenTypeWhitespace, unicode.White_Space),
		lexer.ConsumeString(TokenTypeString),
	}
)

var DebugPrinter = lexer.NewDebugPrintingVisitor(nil)

const (
	TypeRoot           parser.NodeType = "root"
	TypeList           parser.NodeType = "list"
	TypeSymbol         parser.NodeType = "symb"
	TypeStringLiteral  parser.NodeType = "strn"
	TypeNumberLiteral  parser.NodeType = "numb"
	TypeBooleanLiteral parser.NodeType = "bool"
)

func RegisterNodeFactories(p *parser.Parser) {
	p.RegisterFactory(TokenTypeStart, func(cn *parser.Node, tok lexer.Token) *parser.Node {
		return cn.AddChild(TypeList, nil, 0, 0)
	})
	p.RegisterFactory(TokenTypeEnd, func(cn *parser.Node, tok lexer.Token) *parser.Node {
		return cn.Parent()
	})
	p.RegisterFactory(TokenTypeSymbol, func(cn *parser.Node, tok lexer.Token) *parser.Node {
		cn.AddChild(TypeSymbol, tok.Value, 0, 0)
		return nil
	})
	p.RegisterFactory(TokenTypeString, func(cn *parser.Node, tok lexer.Token) *parser.Node {
		cn.AddChild(TypeStringLiteral, tok.Value, 0, 0)
		return nil
	})
	p.RegisterFactory(TokenTypeNumber, func(cn *parser.Node, tok lexer.Token) *parser.Node {
		num, _ := strconv.ParseFloat(tok.Value, 64)
		cn.AddChild(TypeNumberLiteral, num, 0, 0)
		return nil
	})
	p.RegisterFactory(TokenTypeBoolean, func(cn *parser.Node, tok lexer.Token) *parser.Node {
		b, _ := strconv.ParseBool(tok.Value)
		cn.AddChild(TypeBooleanLiteral, b, 0, 0)
		return nil
	})
	p.RegisterFactory(TokenTypeWhitespace, parser.Ignore)
	p.RegisterFactory(TokenTypeEOF, parser.Ignore)
	p.RegisterFactory(TokenTypeError, parser.Ignore)
}
