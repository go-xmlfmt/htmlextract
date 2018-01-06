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

type attributes map[string]string

type Tag struct {
	Name string
	Attr attributes
}

type TagParser interface {
	TagName() (name []byte, hasAttr bool)
	TagAttr() (key, val []byte, moreAttr bool)
}

type Extractor interface {
	GetBase() *extractor
	VisitToken(z *html.Tokenizer, tt html.TokenType, w io.Writer)
}

type extractor struct {
	z     *html.Tokenizer
	depth int
	// for extOutliner
	outputstart bool
	levelopen   map[int]bool
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func TagParse(parser TagParser) Tag {

	name, hasAttr := parser.TagName()
	attr := make(attributes)

	if hasAttr {
		for {
			key, value, more := parser.TagAttr()
			attr[string(key)] = string(value)
			if !more {
				break
			}
		}
	}

	return Tag{Name: string(name), Attr: attr}
}

func NewExtractor(i io.Reader) *extractor {
	e := extractor{}
	e.z = html.NewTokenizer(i)
	e.levelopen = make(map[int]bool)
	return &e
}

func Walk(e Extractor, w io.Writer) error {
	b := e.GetBase()
	// https://godoc.org/golang.org/x/net/html
	for {
		tt := b.z.Next()
		switch tt {
		case html.ErrorToken:
			err := b.z.Err()
			if err == io.EOF {
				return nil // finished reading
			}
			return err
		case html.StartTagToken, html.EndTagToken:
			if tt == html.StartTagToken {
				b.depth++
			} else {
				b.depth--
			}
			e.VisitToken(b.z, tt, w)
		default:
			e.VisitToken(b.z, tt, w)
		}
	}
}

////////////////////////////////////////////////////////////////////////////
// Method definitions

func (e *extractor) GetBase() *extractor { return e }

func (e *extractor) VisitToken(z *html.Tokenizer, tt html.TokenType, w io.Writer) {
	verbose(2, ">: %d (%v)", e.depth, tt)
	switch tt {
	case html.TextToken:
		if e.depth > 0 {
			// emitBytes should copy the []byte it receives,
			// if it doesn't process it immediately.
			// emitBytes(z.Text())

			text := strings.TrimSpace(string(z.Text()))
			if text != "" {
				e.PrintElmt(w, text)
			}
		}
	case html.StartTagToken, html.SelfClosingTagToken:
		tag := TagParse(z)
		verbose(2, " T: %#v", tag)
		if tag.Name == "body" {
			e.depth = 0
		}
		e.PrintElmt(w, tag.Name)
	}
	verbose(2, "<: %d", e.depth)
}

func (e *extractor) PrintElmt(w io.Writer, s string) {
	fmt.Fprintf(w, "%*s%s\n", e.depth*2, "", s)
}
