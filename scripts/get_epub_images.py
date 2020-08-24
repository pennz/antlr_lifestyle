#!/usr/bin/env python3

import ebooklib
from ebooklib import epub
import ipdb

book = epub.read_epub('../JingJin.epub')

for image in book.get_items_of_type(ebooklib.ITEM_IMAGE):
    file_name = image.file_name
    file_name = file_name.replace("/","_")
    print(file_name)
    with open(file_name, 'wb') as f:
        f.write(image.content)
