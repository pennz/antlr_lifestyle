#!/bin/bash
curl $1 | sed -n 's/.*gist-meta.*\(https[^\\]*\).*view raw.*/\1/p' | sed 's/github.com/githubusercontent.com/' | xargs -I{} wget {}
