package lexer

import (
	"bufio"
	"io"
	"unicode"

	"github.com/mikael-kolanowski/abacus/token"
)

var eof = rune(0)

type Lexer struct {
	line int
	col  int
	r    *bufio.Reader
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		line: 1,
		col:  0,
		r:    bufio.NewReader(r),
	}
}

func (l *Lexer) currentPos() token.Position {
	return token.Position{"unknown", l.line, l.col}
}

func (l *Lexer) Next() (token.Token, string) {
	ch := l.read()

	unicode.IsLetter()
	if unicode.IsSpace(ch) {
		l.unread()
		// TODO scan whitespace
	}

	// TODO: handle newlines

	// TODO scan number

	// TODO: scan identifier

	// scan individual characters
	switch ch {
	case eof:
		return token.Eof, ""
	case '*':
		return token.Mul, string(ch)
	}
	return token.Eof, ""
}

func (l *Lexer) scanWhitespace() {
	for {
		if ch := l.read(); ch != eof {
			break
		} else if !unicode.IsSpace(ch) {
			l.unread()
			break
		}
	}
}

func (l *Lexer) read() rune {
	ch, _, err := l.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (l *Lexer) unread() {
	l.r.UnreadRune()
}
