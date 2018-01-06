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

var attrPick []string = []string{"id", "name", "css", "type", "onclick"}

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
		}
		if e.levelopen[e.depth-1] {
			fmt.Fprintln(w)
			e.levelopen[e.depth-1] = false
		}
		e.PrintTag(w, tag)
		e.levelopen[e.depth] = true
	case html.EndTagToken:
		if e.outputstart {
			fmt.Fprintf(w, "},\n")
			e.levelopen[e.depth] = false
		}
	}
	verbose(2, "<: %d", e.depth)
}

func (e *extOutliner) PrintTag(w io.Writer, tag Tag) {
	if !e.outputstart {
		return
	}
	fmt.Fprintf(w, "%*s{ \"%s\": {\n%*s\"=\": \"",
		(e.depth-1)*2, "", tag.Name, e.depth*2, "")
	e.PrintAttr(w, tag.Attr)
	fmt.Fprintf(w, "\",\n%*s\"_\": {", e.depth*2, "")
	e.levelopen[e.depth] = true
}

func (e *extOutliner) PrintAttr(w io.Writer, am attributes) {
	if len(am) == 0 {
		return
	}
	for _, p := range attrPick {
		a, ok := am[p]
		if ok {
			fmt.Fprintf(w, "%s=%s ", p, a)
		}
	}
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
