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
	*extractor
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func NewExtOutliner(i io.Reader) (*html.Tokenizer, *extOutliner) {
	z, e := NewExtractor(i)
	return z, &extOutliner{e}
}

////////////////////////////////////////////////////////////////////////////
// Method definitions

func (e *extOutliner) VisitToken(z *html.Tokenizer, tt html.TokenType, w io.Writer, depth *int) {
	verbose(2, ">: %d (%v)", *depth, tt)
	switch tt {
	case html.StartTagToken, html.SelfClosingTagToken:
		tag := TagParse(z)
		if tag.Name == "body" {
			*depth = 1
			e.outputstart = true
			fmt.Fprintln(w, `{`)
		}
		e.PrintTag(w, *depth, tag)
	case html.EndTagToken:
		fmt.Fprintf(w, "},\n")
	}
	verbose(2, "<: %d", *depth)
}

func (e *extOutliner) PrintTag(w io.Writer, depth int, tag Tag) {
	if !e.outputstart {
		return
	}
	fmt.Fprintf(w, "%*s{ \"%s\": {\n%*s  \"A\": \"",
		depth*2, "", tag.Name, depth*2, "")
	e.PrintAttr(w, tag.Attr)
	depth++
	fmt.Fprintf(w, "\",\n%*s\"C\": {", depth*2, "")
}

func (e *extOutliner) PrintAttr(w io.Writer, a attributes) {
	if len(a) == 0 {
		return
	}
	fmt.Fprintf(w, "%#v", a)
}

/*

JSON layout:

{
  "body": {
   "A": "",
   "C": {
    "h1": {
     "A": "",
     "C": {}
    },
    "h2": {
     "A": "",
     "C": {}
    },
    "ul": {
     "A": "",
     "C": {
      "li": {
        "A": "",
        "C": {}
      },
      "li2": {
        "A": "",
        "C": {}
      }
     }
    }
    "span": {
     "A": "",
     "C": {
      "img": {
        "A": "",
        "C": {}
      },
      "a": {
        "A": "",
        "C": {}
      }
     }
   }
  }
}

Note the "li2" can't be "li" again? Otherwise, will get
"Duplicate key" error

*/
