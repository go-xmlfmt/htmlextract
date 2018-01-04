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
	z *html.Tokenizer
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func NewExtractor(i io.Reader) *Extractor {
	return &Extractor{z: html.NewTokenizer(i)}
}

////////////////////////////////////////////////////////////////////////////
// Method definitions

func Walk(e *Extractor, w io.Writer) error {
	// https://godoc.org/golang.org/x/net/html
	depth := 0
	for {
		tt := e.z.Next()
		switch tt {
		case html.ErrorToken:
			err := e.z.Err()
			if err == io.EOF {
				return nil // finished reading
			}
			return err
		case html.StartTagToken, html.EndTagToken:
			if tt == html.StartTagToken {
				e.VisitToken(tt, w, &depth)
				depth++
			} else {
				depth--
			}
		default:
			e.VisitToken(tt, w, &depth)
		}
	}
}

func (e *Extractor) VisitToken(tt html.TokenType, w io.Writer, depth *int) {
	verbose(2, ">: %d (%v)", *depth, tt)
	switch tt {
	case html.TextToken:
		if *depth > 0 {
			// emitBytes should copy the []byte it receives,
			// if it doesn't process it immediately.
			// emitBytes(e.z.Text())

			text := strings.TrimSpace(string(e.z.Text()))
			if text != "" {
				e.PrintElmt(w, *depth, text)
			}
		}
	case html.StartTagToken, html.SelfClosingTagToken:
		tn, _ := e.z.TagName()
		tag := string(tn)
		verbose(2, " T: %s", tag)
		if tag == "body" {
			*depth = 0
		}
		e.PrintElmt(w, *depth, tag)
	}
	verbose(2, "<: %d", *depth)
}

func (z *Extractor) PrintElmt(w io.Writer, depth int, s string) {
	fmt.Fprintf(w, "%*s%s\n", depth*2, "", s)
}
