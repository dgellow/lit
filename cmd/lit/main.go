package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/dgellow/lit"
)

var (
	flagInputFile = flag.String("input", "stdin", "input file")
	flagEmbedded  = flag.Bool("embedded", false, "generate embeddable HTML")
)

func main() {
	flag.Parse()

	filename := *flagInputFile
	var input io.Reader
	switch filename {
	case "", "stdin":
		input = os.Stdin
	default:
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("failed to open file: %s: %s\n", filename, err)
			os.Exit(2)
		}
		input = file
	}

	doc := lit.NewDocument(filename, input)
	if errs := doc.ReadAndParse(); len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		os.Exit(2)
	}
	doc.Process()

	var b bytes.Buffer
	doc.Write(lit.SideBySideTempl, &b)

	if *flagEmbedded {
		os.Stdout.Write(b.Bytes())
		return
	}

	lit.DocumentTempl.Execute(os.Stdout, b.Bytes())
}
