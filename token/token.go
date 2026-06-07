package token

import "fmt"

type Token uint

const (
	Number Token = iota

	OpenParen
	CloseParen

	Bang
	BangEq
	Plus
	Minus
	Mul
	Div
	Eq
	Less
	LessEq
	Greater
	GreaterEq
	ColEq
	And
	Or

	Ident
	If
	Then
	Else
	Eof
)

type Position struct {
	Filename string
	Line     int
	Column   int
}

func (p Position) String() string {
	return fmt.Sprintf("%s:%d:%d", p.Filename, p.Line, p.Column)
}
