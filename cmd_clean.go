////////////////////////////////////////////////////////////////////////////
// Program: htmlextract
// Purpose: HTML Extraction Tool
// Authors: Antonio Sun (c) 2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/go-easygen/cli"
)

////////////////////////////////////////////////////////////////////////////
// clean

func cleanCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*cleanT)
	fmt.Printf("[clean]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	Opts.Case, Opts.Verbose =
		rootArgv.Case, rootArgv.Verbose.Value()
	return nil
}
