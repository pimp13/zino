package main

import (
	"fmt"

	"github.com/pimp13/zino/lexer"
	"github.com/pimp13/zino/token"
)

func main() {
	input := `
	x = 10
	print "hello zino"
	y = x + 5
	`

	l := lexer.NewLexer(input)
	for {
		tok := l.NextToken()
		fmt.Printf("%+v\n", tok)
		if tok.Type == token.EOF {
			break
		}
	}
}
