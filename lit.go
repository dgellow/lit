package lit

import (
	"bytes"
	"fmt"
	"go/token"
	"io"
	"io/ioutil"
	"strings"
)

type section struct {
	code             []scanChunk
	comments         []scanChunk
	renderedCode     []byte
	renderedComments []byte
}

// Document represents a literate programming source file.
type Document struct {
	filename string
	input    io.Reader
	src      []byte
	sections []section
}

// NewDocument allocates a new `Document`.
func NewDocument(filename string, r io.Reader) *Document {
	return &Document{filename: filename, input: r}
}

// ReadAndParse tries to read the input file, scan it to get the list of tokens, then split it in sections
// composed of comments and some code. A new section is created each time a new comment block (i.e: demarked
// by code or new lines) is found.
func (d *Document) ReadAndParse() []error {
	b, err := ioutil.ReadAll(d.input)
	if err != nil {
		return []error{fmt.Errorf("failed to read file %q: %s", d.filename, err)}
	}
	d.src = b

	chunks, errs := tokenize(d.filename, d.src)
	if len(errs) != 0 {
		return errs
	}

	var ss []section
	var s section
	for i := range chunks {
		if chunks[i].token == token.COMMENT {
			if len(s.code) != 0 {
				ss = append(ss, s)
				s = section{}
			}
			s.comments = append(s.comments, chunks[i])
		} else {
			s.code = append(s.code, chunks[i])
		}
	}
	d.sections = append(ss, s)

	return nil
}

// Render generates information for syntax highlighting of the code and formats comments in a more readable way.
func (d *Document) Render() {
	for i := range d.sections {
		d.sections[i].renderedCode = bytes.TrimSpace(highlight(d.src, d.sections[i].code))
		d.sections[i].renderedComments = bytes.TrimSpace(formatComments(d.sections[i].comments))
	}
}

func formatComments(chunks []scanChunk) []byte {
	var b bytes.Buffer
	for _, c := range chunks {
		b.WriteString(strings.TrimSpace(c.literal) + " ")
	}
	return b.Bytes()
}

// Write the document to the writer `w`.
func (d *Document) Write(w io.Writer) {
	type sect struct {
		Code     string
		Comments string
	}
	type data struct {
		Sections []sect
	}
	var dt data
	for _, s := range d.sections {
		dt.Sections = append(dt.Sections, sect{
			Code:     string(s.renderedCode[:]),
			Comments: string(s.renderedComments[:]),
		})
	}
	sideBySideTempl.Execute(w, dt)
}
