package lit

import (
	"fmt"
	"go/scanner"
	"go/token"
)

type tokenizeMode int

type tokenizeError struct {
	position token.Position
	msg      string
}

func newTokenizeError(pos token.Position, msg string) tokenizeError {
	return tokenizeError{
		position: pos,
		msg:      msg,
	}
}

func (err tokenizeError) Error() string {
	return fmt.Sprintf("scanning error: %s:%d:%d: %s",
		err.position.Filename, err.position.Line, err.position.Column, err.msg,
	)
}

type scanChunk struct {
	position token.Position
	token    token.Token
	literal  string
}

func tokenize(filename string, src []byte) ([]scanChunk, []error) {
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile(filename, fset.Base(), len(src))

	var errs []error
	errHandler := func(pos token.Position, msg string) {
		errs = append(errs, newTokenizeError(pos, msg))
	}

	s.Init(file, src, errHandler, scanner.ScanComments)

	var chunks []scanChunk
	for {
		pos, t, l := s.Scan()
		switch t {
		case token.ILLEGAL: // ignore, handled by scanner error handler
		case token.COMMENT:
			chunks = append(chunks, scanChunk{position: fset.Position(pos), token: t, literal: l})
		default:
			chunks = append(chunks, scanChunk{position: fset.Position(pos), token: t, literal: l})
		}

		if t == token.EOF {
			break
		}
	}

	return chunks, errs
}
