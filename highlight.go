package lit

import (
	"bytes"
	"fmt"
	"go/token"
)

func highlight(src []byte, chunks []scanChunk) []byte {
	var content bytes.Buffer
	var current scanChunk
	var next scanChunk
	fmt.Println("")
	for i := range chunks {
		if i == len(chunks)-1 {
			current = next
		} else {
			current, next = chunks[i], chunks[i+1]
		}
		literal, leading, trailing := split(src[current.position.Offset:next.position.Offset])
		content.Write(leading)
		content.Write(highlightChunk(current, literal))
		content.Write(trailing)
	}
	return content.Bytes()
}

func split(s []byte) ([]byte, []byte, []byte) {
	literal := bytes.TrimSpace(s)
	leading := s[:bytes.Index(s, literal)]
	trailing := s[len(leading)+len(literal):]
	return literal, leading, trailing
}

func highlightChunk(c scanChunk, literal []byte) []byte {
	var cl string
	switch c.token {
	case token.COMMENT: // we do not highlight comments as they are rendered separately
	case token.IDENT:
		cl = "ident"
	case token.STRING:
		cl = "string"
	case token.FUNC:
		cl = "func"
	case token.EOF:
		cl = "eof"
	case token.SEMICOLON:
		if len(literal) == 0 {
			cl = "eol"
		} else {
			cl = "semicolon"
		}
	}

	switch {
	case c.token.IsKeyword():
		cl += " is-keyword"
	case c.token.IsLiteral():
		cl += " is-literal"
	case c.token.IsOperator():
		cl += " is-operator"
	}

	return htmlTag(cl+" chunk", literal)
}

func htmlTag(class string, content []byte) []byte {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf(`<div class="%s">%s</div>`, class, string(content[:])))
	return b.Bytes()
}
