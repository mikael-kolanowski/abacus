package lexer

import (
	"io"

	"github.com/mikael-kolanowski/abacus/token"
)

type Position struct {
	Filename string
	Line     int
	Column   int
}

type Lexer struct {
	line int
	col  int
	r    io.Reader
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		line: 1,
		col:  0,
		r:    r,
	}
}

func (l *Lexer) currentPos() token.Position {
	return token.Position{"unknown", l.line, l.col}
}

func (l *Lexer) Next() (token.Position, token.Token, string) {
	return l.currentPos(), token.Eof, ""
}
