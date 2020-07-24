#!/bin/bash

curl $1 | grep -o "https://gist[^>]*\.js" | xargs -I{} ./get_gist_content.sh {}
