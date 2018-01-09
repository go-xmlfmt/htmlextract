
# htmlextract

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-xmlfmt/htmlextract?status.svg)](http://godoc.org/github.com/go-xmlfmt/htmlextract)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-xmlfmt/htmlextract)](https://goreportcard.com/report/github.com/go-xmlfmt/htmlextract)
[![travis Status](https://travis-ci.org/go-xmlfmt/htmlextract.svg?branch=master)](https://travis-ci.org/go-xmlfmt/htmlextract)

## TOC
- [htmlextract - HTML Extraction Tool](#htmlextract---html-extraction-tool)
- [Usage](#usage)
  - [$ htmlextract](#-htmlextract)
  - [$ htmlextract outline](#-htmlextract-outline)
  - [$ htmlextract clean](#-htmlextract-clean)
  - [$ htmlextract h2md](#-htmlextract-h2md)
- [Examples](#examples)
  - [Outline](#outline)
    - [$ htmlextract outline -i test/sample0.html -o](#-htmlextract-outline--i-testsample0html--o)
    - [Advantages](#advantages)
    - [Usage](#usage-1)
      - [Specifying more attributes](#specifying-more-attributes)
      - [Work with URL directly](#work-with-url-directly)
- [Download binaries](#download-binaries)
- [Debian package](#debian-package)
- [Install Source](#install-source)
  - [Author(s) & Contributor(s)](#author(s)-&-contributor(s))

## htmlextract - HTML Extraction Tool

The `htmlextract` makes it easy to look at the HTML files from different aspects. 

- **`htmlextract outline`** will extract HTML structure as outline so as to focus more easily on the structure, not the details.
- **`htmlextract clean`** will clean up HTML tags & attributes as much as possible, so as to go back to the plain text version as easy as possible. 
- **`htmlextract h2md`** will convert HTML to .md file on top of above clean up.

# Usage

### $ htmlextract
```sh
HTML Extraction Tool
Version 0.2.0 built on 2018-01-08

Tool to extract info from HTML

Options:

  -h, --help      display help information
  -c, --case      1: lowercase tags, 2: uppercase, 0: no-change.
  -v, --verbose   Verbose mode (Multiple -v options increase the verbosity.)

Commands:

  outline   Extract HTML structure as outline
  clean     Clean up HTML tags & attributes as much as possible
  h2md      Convert HTML to .md file (on top of above clean up)
```

### $ htmlextract outline
```sh
Extract HTML structure as outline

Usage:
  htmlextract outline -i /tmp/f.html

Options:

  -h, --help         display help information
  -c, --case         1: lowercase tags, 2: uppercase, 0: no-change.
  -v, --verbose      Verbose mode (Multiple -v options increase the verbosity.)
  -i, --input       *The file/url to extract from (mandatory)
  -o, --output       The output outline file (default: input.json)
  -a, --attributes   extra attributes to include (may be more than one)
```

### $ htmlextract clean
```sh
Clean up HTML tags & attributes as much as possible

Usage:
  htmlextract clean -i /tmp/fi.html -o /tmp/fo.html

Options:

  -h, --help      display help information
  -c, --case      1: lowercase tags, 2: uppercase, 0: no-change.
  -v, --verbose   Verbose mode (Multiple -v options increase the verbosity.)
  -i, --input    *The file/url to extract from (mandatory)
  -o, --output   *The output html file (mandatory)
```

### $ htmlextract h2md
```sh
Convert HTML to .md file (on top of above clean up)

Usage:
  htmlextract h2md -i /tmp/f.html

Options:

  -h, --help      display help information
  -c, --case      1: lowercase tags, 2: uppercase, 0: no-change.
  -v, --verbose   Verbose mode (Multiple -v options increase the verbosity.)
  -i, --input    *The file/url to extract from (mandatory)
  -o, --output    The output .md file (default: input.md)
```


# Examples

## Outline

### $ htmlextract outline -i test/sample0.html -o
```json
{

"body": {
  "=": "css=.claro ",
  "_": {
  "h1": {
    "=": "css=.title ",
    "_": {}},
  "div": {
    "=": "id=ctrlBtns ",
    "_": {}},
  "br": {
    "=": "",
    "_": {}},
  "div": {
    "=": "id=pluginList ",
    "_": {}},
  "div": {
    "=": "id=gridContainer ",
    "_": {}},
  "br": {
    "=": "",
    "_": {}},
  "div": {
    "=": "id=ctrlBtns2 ",
    "_": {}},
  "br": {
    "=": "",
    "_": {}},
  "div": {
    "=": "id=menusSupport ",
    "_": {
    "div": {
      "=": "id=headerMenu ",
      "_": {
      "div": {
        "=": "",
        "_": {}},
      "div": {
        "=": "",
        "_": {}},
}},
    "div": {
      "=": "id=rowMenu ",
      "_": {
      "div": {
        "=": "",
        "_": {}},
}},
}},
  "div": {
    "=": "id=exporterSupport css=.supporter ",
    "_": {
    "input": {
      "=": "id=exportAllCSV type=button ",
      "_": {}},
    "textarea": {
      "=": "id=csvResults ",
      "_": {}},
}},
  "div": {
    "=": "id=printerSupport css=.supporter ",
    "_": {
    "input": {
      "=": "css=.print.all type=button onclick=printGrid() ",
      "_": {}},
    "input": {
      "=": "css=.print.sel type=button onclick=printSelected() ",
      "_": {}},
    "input": {
      "=": "css=.print.preview' type=button onclick=printPreview() ",
      "_": {}},
    "input": {
      "=": "id=print_title type=text ",
      "_": {}},
}},
  "div": {
    "=": "id=cellMergeSupport css=.supporter ",
    "_": {
    "table": {
      "=": "",
      "_": {
      "tr": {
        "=": "",
        "_": {
        "td": {
          "=": "",
          "_": {}},
        "td": {
          "=": "",
          "_": {
          "input": {
            "=": "id=inputRow type=text ",
            "_": {}},
}},
}},
      "tr": {
        "=": "",
        "_": {
        "td": {
          "=": "",
          "_": {}},
        "td": {
          "=": "",
          "_": {
          "input": {
            "=": "id=inputStart type=text ",
            "_": {}},
}},
}},
      "tr": {
        "=": "",
        "_": {
        "td": {
          "=": "",
          "_": {}},
        "td": {
          "=": "",
          "_": {
          "input": {
            "=": "id=inputEnd type=text ",
            "_": {}},
}},
}},
      "tr": {
        "=": "",
        "_": {
        "td": {
          "=": "",
          "_": {}},
        "td": {
          "=": "",
          "_": {
          "input": {
            "=": "id=inputMajor type=text ",
            "_": {}},
}},
}},
      "tr": {
        "=": "",
        "_": {
        "td": {
          "=": "",
          "_": {
          "button": {
            "=": "id=mergeCell onclick=mergeCells() ",
            "_": {}},
}},
}},
}},
}},
  "div": {
    "=": "id=paginationSupport css=.supporter ",
    "_": {
    "table": {
      "=": "",
      "_": {
      "tr": {
        "=": "",
        "_": {
        "td": {
          "=": "",
          "_": {
          "input": {
            "=": "id=inputScrollToRowIdx type=text ",
            "_": {}},
}},
        "td": {
          "=": "",
          "_": {
          "button": {
            "=": "id=scrollToRow ",
            "_": {}},
}},
}},
}},
}},
  "div": {
    "=": "css=.tutor.footer.howtouse ",
    "_": {
    "p": {
      "=": "",
      "_": {
      "h2": {
        "=": "",
        "_": {}},
      "ol": {
        "=": "",
        "_": {
        "li": {
          "=": "",
          "_": {}},
        "li": {
          "=": "",
          "_": {}},
}},
}},
}},
  "div": {
    "=": "id=repeatcounter ",
    "_": {}},
}},
}},
```

### Advantages

- By extracting HTML structure as outline, the `htmlextract outline` will make it easier to analyse the file structure, by eliminating all the glory details out of the way, which is most often needed when doing web scrapping or WebDriver code developing.
- The output is mindfully chosen as the JSON format so as to easily take advantage of the dynamic folding feature that the text editors provide. Or you can use the [jsonformatter.org](https://jsonformatter.org/) online as well, even without a text editor.

Here is a screenshot of viewing the result of `htmlextract outline -i test/sample0.html`:

![sample.png](sample.png "Sample screenshot")

### Usage

#### Specifying more attributes

If the predefined attribute selection is not enough, the it is easily to add your own by the `-a, --attributes` switch. Note that you can use the switch as many times as you wish, to provide as many attributes as you need:

```sh
$ htmlextract outline -a dojotype -a style -i test/sample0.html -o | grep -1 dojotype | head -3 
  "div": {
    "=": "id=pluginList dojotype=PluginTable style=float:right; ",
    "_": {}},
```

#### Work with URL directly

Starting with version `0.2.0`, `htmlextract` can extract from URL directly:

```sh
$ htmlextract outline -i http://demoaut.katalon.com/profile.php -o | head -35
...
"body": {
  "=": "",
  "_": {
  "a": {
    "=": "id=menu-toggle css=.btn.btn-dark.btn-lg.toggle ",
    "_": {
    "i": {
      "=": "css=.fa.fa-bars ",
      "_": {}},
}},
  "nav": {
    "=": "id=sidebar-wrapper ",
    "_": {
    "ul": {
      "=": "css=.sidebar-nav ",
      "_": {
      "a": {
        "=": "id=menu-close css=.btn.btn-light.btn-lg.pull-right.toggle ",
        "_": {
        "i": {
          "=": "css=.fa.fa-times ",
          "_": {}},
...
```


# Download binaries

- The latest binary executables are available under  
https://bintray.com/antoniosun/bin/htmlextract/latest, or directly under  
https://bintray.com/version/files/antoniosun/bin/htmlextract/latest  
as the result of the Continuous-Integration process.
- I.e., they are built during every git push, automatically by [travis-ci](https://travis-ci.org/), right from the source code, truely WYSIWYG.
- Pick & choose the binary executable that suits your OS and its architecture. E.g., for Linux, it would most probably be the `htmlextract-linux-amd64` file. If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- You may want to rename it to a shorter name instead, e.g., `htmlextract`, after downloading it. 


# Debian package

Available at https://bintray.com/antoniosun/deb/htmlextract,  
or directly at  https://dl.bintray.com/antoniosun/deb:

```
echo "deb [trusted=yes] https://dl.bintray.com/antoniosun/deb all main" | sudo tee /etc/apt/sources.list.d/antoniosun-debs.list
sudo apt-get update

sudo chmod 644 /etc/apt/sources.list.d/antoniosun-debs.list
apt-cache policy htmlextract

sudo apt-get install -y htmlextract
```



# Install Source

To install the source code instead:

```
go get github.com/go-xmlfmt/htmlextract
```


## Author(s) & Contributor(s)

- [Antonio SUN](https://github.com/AntonioSun)

_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe), the _one-stop wire-framing solution_ for Go cli based projects, from start to deploy.

All patches welcome. 
