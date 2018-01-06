////////////////////////////////////////////////////////////////////////////
// Program: extractor.go
// Purpose: HTML Extractor
// Authors: Antonio Sun (c) 2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type extOutliner struct {
	extractor
	output bool
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func NewExtOutliner(i io.Reader) (*html.Tokenizer, *extOutliner) {
	return html.NewTokenizer(i), &extOutliner{}
}

////////////////////////////////////////////////////////////////////////////
// Method definitions

func (e *extOutliner) VisitToken(z *html.Tokenizer, tt html.TokenType, w io.Writer, depth *int) {
	verbose(2, ">: %d (%v)", *depth, tt)
	switch tt {
	case html.StartTagToken, html.SelfClosingTagToken:
		tag := TagParse(z)
		if tag.Name == "body" {
			*depth = 0
			e.output = true
		}
		e.PrintTag(w, *depth, tag)
	}
	verbose(2, "<: %d", *depth)
}

func (e *extOutliner) PrintTag(w io.Writer, depth int, tag Tag) {
	if !e.output {
		return
	}
	fmt.Fprintf(w, "%*s{ \"T:\", \"%s\"\n", depth*2, "", tag.Name)
}
