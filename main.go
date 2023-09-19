package main

import (
	"Ristretto/parser"
	"Ristretto/syntax"
	"github.com/antlr4-go/antlr/v4"
	"github.com/liyue201/gostl/ds/stack"
)

func main() {
	is, _ := antlr.NewFileStream(".\\test.ris")

	lexer := parser.NewRistrettoLexer(is)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parse := parser.NewRistrettoParser(tokens)

	tree := parse.Start_()

	visitor := &syntax.RistrettoVisitor{
		S: stack.New[int](),
	}

	res, _ := visitor.Visit(tree).(string)

	println(res)
}
