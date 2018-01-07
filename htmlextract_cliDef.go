////////////////////////////////////////////////////////////////////////////
// Program: htmlextract
// Purpose: HTML Extraction Tool
// Authors: Antonio Sun (c) 2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/go-easygen/cli"
	clix "github.com/go-easygen/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

//==========================================================================
// htmlextract

type rootT struct {
	cli.Helper
	Case    int         `cli:"c,case" usage:"1: lowercase tags, 2: uppercase, 0: no-change."`
	Verbose cli.Counter `cli:"v,verbose" usage:"Verbose mode (Multiple -v options increase the verbosity.)"`
}

var root = &cli.Command{
	Name:   "htmlextract",
	Desc:   "HTML Extraction Tool\nVersion " + version + " built on " + date,
	Text:   "Tool to extract info from HTML",
	Global: true,
	Argv:   func() interface{} { return new(rootT) },
	Fn:     htmlextract,

	NumArg: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
//  type OptsT struct {
//  	Case	int
//  	Verbose	cli.Counter
//  	Verbose int
//  }

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "htmlextract"
//          version   = "0.1.0"
//          date = "2018-01-07"

//  	rootArgv *rootT
//  	// Opts store all the configurable options
//  	Opts OptsT
//  )

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
//  	//NOTE: You can set any writer implements io.Writer
//  	// default writer is os.Stdout
//  	if err := cli.Root(root,
//  		cli.Tree(outlineDef),
//  		cli.Tree(cleanDef),
//  		cli.Tree(h2mdDef)).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Main dispatcher

//  func htmlextract(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here

////////////////////////////////////////////////////////////////////////////
// outline

//  func outlineCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*outlineT)
//  	fmt.Printf("[outline]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
//  	Opts.Case, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Case, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	return nil
//  }

type outlineT struct {
	Filei      *clix.Reader `cli:"*i,input" usage:"The file/url to extract from (mandatory)"`
	Fileo      *clix.Writer `cli:"o,output" usage:"The output outline file (default: input.json)"`
	Attributes []string     `cli:"a,attributes" usage:"extra attributes to include (may be more than one)"`
}

var outlineDef = &cli.Command{
	Name: "outline",
	Desc: "Extract HTML structure as outline",
	Text: "Usage:\n  htmlextract outline -i /tmp/f.html",
	Argv: func() interface{} { return new(outlineT) },
	Fn:   outlineCLI,

	NumOption: cli.AtLeast(1),
}

////////////////////////////////////////////////////////////////////////////
// clean

//  func cleanCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*cleanT)
//  	fmt.Printf("[clean]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
//  	Opts.Case, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Case, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	return nil
//  }

type cleanT struct {
	Filei *clix.Reader `cli:"*i,input" usage:"The file/url to extract from (mandatory)"`
	Fileo *clix.Writer `cli:"*o,output" usage:"The output html file (mandatory)"`
}

var cleanDef = &cli.Command{
	Name: "clean",
	Desc: "Clean up HTML tags & attributes as much as possible",
	Text: "Usage:\n  htmlextract clean -i /tmp/fi.html -o /tmp/fo.html",
	Argv: func() interface{} { return new(cleanT) },
	Fn:   cleanCLI,

	NumArg:      cli.AtLeast(1),
	CanSubRoute: true,
}

////////////////////////////////////////////////////////////////////////////
// h2md

//  func h2mdCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*h2mdT)
//  	fmt.Printf("[h2md]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
//  	Opts.Case, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Case, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	return nil
//  }

type h2mdT struct {
	Filei *clix.Reader `cli:"*i,input" usage:"The file/url to extract from (mandatory)"`
	Fileo *clix.Writer `cli:"o,output" usage:"The output .md file (default: input.md)"`
}

var h2mdDef = &cli.Command{
	Name: "h2md",
	Desc: "Convert HTML to .md file (on top of above clean up)",
	Text: "Usage:\n  htmlextract h2md -i /tmp/f.html",
	Argv: func() interface{} { return new(h2mdT) },
	Fn:   h2mdCLI,

	NumArg:      cli.AtLeast(1),
	CanSubRoute: true,
}
