#!/bin/bash

get_gist_content () {
	curl $1 | sed -n 's/.*gist-meta.*\(https[^\\]*\).*view raw.*/\1/p' | sed 's/github.com/githubusercontent.com/' | xargs -I{} wget {}
}

export -f get_gist_content
curl $1 | grep -o "https://gist[^>]*\.js" | xargs -I{} bash -c 'get_gist_content {}'
