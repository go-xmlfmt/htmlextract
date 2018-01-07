#!/bin/sh

set -e

# test first
htmlextract outline -i sample0.html -o /tmp/htmlextract.sample0.json.ref
echo $?
htmlextract outline -i sample1.html -o /tmp/htmlextract.sample1.json.ref
echo $?
htmlextract outline -i sample2.html -o /tmp/htmlextract.sample2.json.ref
echo $?

htmlextract outline -i sample0.html -o /tmp/htmlextract.sample0_2.json.ref -a dojotype -a style
echo $?

# check next
ls sample*.json | xargs -t -i sh -c 'diff -U1 {} /tmp/htmlextract.{}.ref'
ret=$?
echo $ret
exit $ret
