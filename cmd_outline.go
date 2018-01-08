////////////////////////////////////////////////////////////////////////////
// Program: htmlextract
// Purpose: HTML Extraction Tool
// Authors: Antonio Sun (c) 2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"net/url"
	"os"
	"regexp"

	"github.com/go-easygen/cli"
)

////////////////////////////////////////////////////////////////////////////
// outline

func outlineCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*outlineT)
	// fmt.Printf("[outline]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	Opts.Case, Opts.Verbose =
		rootArgv.Case, rootArgv.Verbose.Value()
	attrPick = append(attrPick, argv.Attributes...)

	fin := argv.Filei.Name()
	verbose(2, "Input file: '%s'", fin)
	if regexp.MustCompile(`(?i)^http`).MatchString(fin) {
		u, e := url.QueryUnescape(fin)
		abortOn("Input url error", e)
		// in case url contains ?param1=...&param2=...
		up, _ := url.Parse(u)
		verbose(2, "url: %#v", up)
		// in case up.Path is empty
		u = up.Host + up.Path
		// in case of ending '/'
		u = regexp.MustCompile(`^(.*)/$`).ReplaceAllString(u, "${1}.")
		// get the name from the last part, less extension
		fin = regexp.MustCompile(`^.*/(.*)\.[^.]*$`).ReplaceAllString(u, "${1}")
		fin += ".html"
		verbose(2, "Input file: '%s'", fin)
	}
	if !ctx.IsSet("--output") {
		fileo, err := os.Create(
			regexp.MustCompile(`(?i).html?$`).
				ReplaceAllLiteralString(fin, ".json"))
		abortOn("Creating output file", err)
		argv.Fileo.SetWriter(fileo)
	}
	fileo := argv.Fileo
	defer fileo.Close()

	//e := NewExtractor(argv.Filei)
	e := NewExtOutliner(argv.Filei)
	return Walk(e, fileo)
}
