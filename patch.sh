#!/bin/sh

## Make all relative imports inside github package to absolute to allow easier importing into other packages

for f in `find github/*.proto`; do
	tmpfile=$(mktemp /tmp/ghp.XXXXXX)
	sed -e 's/import "\([a-z]*\)\.proto"/import "github.com\/jhaynie\/go-github-protobuf\/github\/\1.proto"/g' $f > $tmpfile
	cp $tmpfile $f
	rm -rf $tmpfile
done
