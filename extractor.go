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

type Extractor interface {
	VisitToken(z *html.Tokenizer, tt html.TokenType, w io.Writer, depth *int)
}

type extractor struct {
	*html.Tokenizer
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func NewExtractor() extractor {
	return extractor{}
}

////////////////////////////////////////////////////////////////////////////
// Method definitions

func Walk(z *html.Tokenizer, e Extractor, w io.Writer) error {
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
		case html.StartTagToken, html.EndTagToken:
			if tt == html.StartTagToken {
				e.VisitToken(z, tt, w, &depth)
				depth++
			} else {
				depth--
			}
		default:
			e.VisitToken(z, tt, w, &depth)
		}
	}
}

func (e extractor) VisitToken(z *html.Tokenizer, tt html.TokenType, w io.Writer, depth *int) {
	verbose(2, ">: %d (%v)", *depth, tt)
	switch tt {
	case html.TextToken:
		if *depth > 0 {
			// emitBytes should copy the []byte it receives,
			// if it doesn't process it immediately.
			// emitBytes(z.Text())

			text := strings.TrimSpace(string(z.Text()))
			if text != "" {
				e.PrintElmt(w, *depth, text)
			}
		}
	case html.StartTagToken, html.SelfClosingTagToken:
		tn, _ := z.TagName()
		tag := string(tn)
		verbose(2, " T: %s", tag)
		if tag == "body" {
			*depth = 0
		}
		e.PrintElmt(w, *depth, tag)
	}
	verbose(2, "<: %d", *depth)
}

func (e extractor) PrintElmt(w io.Writer, depth int, s string) {
	fmt.Fprintf(w, "%*s%s\n", depth*2, "", s)
}
