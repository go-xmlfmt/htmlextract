////////////////////////////////////////////////////////////////////////////
// Program: extractor.go
// Purpose: HTML Extractor
// Authors: Antonio Sun (c) 2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type Extractor struct {
	*html.Tokenizer
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func NewExtractor(i io.Reader) *Extractor {
	return &Extractor{html.NewTokenizer(i)}
}

////////////////////////////////////////////////////////////////////////////
// Method definitions

func (z Extractor) Walk(w io.Writer) error {
	// https://godoc.org/golang.org/x/net/html
	depth := 0
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			err := z.Err()
			if err == io.EOF {
				return nil // finished reading
			}
			return err
		case html.TextToken:
			if depth > 0 {
				// emitBytes should copy the []byte it receives,
				// if it doesn't process it immediately.
				// emitBytes(z.Text())

				text := strings.TrimSpace(string(z.Text()))
				if text != "" {
					z.PrintElmt(w, depth, text)
				}
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			tag := strings.ToLower(string(tn))
			if tt == html.StartTagToken {
				if tag == "body" {
					depth = 1
				}
				z.PrintElmt(w, depth, tag)
				depth++
			} else {
				depth--
			}
		}
	}
}

func (z *Extractor) PrintElmt(w io.Writer, depth int, s string) {
	fmt.Fprintf(w, "%*s%s\n", depth*2, "", s)
}
