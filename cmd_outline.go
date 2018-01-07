////////////////////////////////////////////////////////////////////////////
// Program: htmlextract
// Purpose: HTML Extraction Tool
// Authors: Antonio Sun (c) 2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"
	"regexp"

	"github.com/mkideal/cli"
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

	verbose(2, "Input file: '%s'", argv.Filei.Name())
	if !ctx.IsSet("--output") {
		fileo, err := os.Create(
			regexp.MustCompile(`(?i)html?`).
				ReplaceAllLiteralString(argv.Filei.Name(), "json"))
		abortOn("Creating output file", err)
		argv.Fileo.SetWriter(fileo)
	}
	fileo := argv.Fileo
	defer fileo.Close()

	//e := NewExtractor(argv.Filei)
	e := NewExtOutliner(argv.Filei)
	return Walk(e, fileo)
}
