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

func NewExtOutliner(i io.Reader) *extOutliner {
	return &extOutliner{NewExtractor(i)}
}

////////////////////////////////////////////////////////////////////////////
// Method definitions

func (e *extOutliner) VisitToken(z *html.Tokenizer, tt html.TokenType, w io.Writer) {
	verbose(2, ">: %d (%v)", e.depth, tt)
	switch tt {
	case html.StartTagToken, html.SelfClosingTagToken:
		tag := TagParse(z)
		if tag.Name == "body" {
			e.depth = 1
			e.outputstart = true
			fmt.Fprintln(w, `{`)
		}
		e.PrintTag(w, tag)
	case html.EndTagToken:
		fmt.Fprintf(w, "},\n")
	}
	verbose(2, "<: %d", e.depth)
}

func (e *extOutliner) PrintTag(w io.Writer, tag Tag) {
	if !e.outputstart {
		return
	}
	fmt.Fprintf(w, "%*s{ \"%s\": {\n%*s  \"A\": \"",
		(e.depth-1)*2, "", tag.Name, e.depth*2, "")
	e.PrintAttr(w, tag.Attr)
	fmt.Fprintf(w, "\",\n%*s\"C\": {", e.depth*2, "")
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
