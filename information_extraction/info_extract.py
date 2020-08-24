"""
https://www.analyticsvidhya.com/blog/2019/09/introduction-information-extraction-python-spacy/

Goal:
	extract hypernym-hyponym pairs by using these patterns/rules
Knowledge:
	https://spacy.io/usage/linguistic-features
	lemma, POS, Tag(detailed), Dep(syntactic dep), Shape(Cap, punct, digits), is alpha, is stop(stop word)
	Most of the tags and labels look pretty abstract, and they vary between languages. spacy.explain will show you a short description – for example, spacy.explain("VBZ") returns “verb, 3rd person singular present”.

"""

# ie_load_libraries.py
import re
import string
import nltk
import spacy
import pandas as pd
import numpy as np
import math
from tqdm import tqdm

from spacy.matcher import Matcher
from spacy.tokens import Span
from spacy import displacy

pd.set_option('display.max_colwidth', 200)

# ie_load_spacy_model.py
# load spaCy model
nlp = spacy.load("en_core_web_sm")

# ie_hearst_pattern_1.py
# sample text
text = "GDP in developing countries such as Vietnam will continue growing at a high rate."

# create a spaCy object
doc = nlp(text)

# ie_dep_parse_1.py
# print token, dependency, POS tag

for tok in doc:
  print(tok.text, "-->", tok.dep_, "-->", tok.pos_)

# ie_define_pattern.py
#define the pattern
pattern = [{'POS': 'NOUN'},
           {'LOWER': 'such'},
           {'LOWER': 'as'},
           {'POS': 'PROPN'}  # proper noun
           ]

# ie_spacy_matcher_1.py
# Matcher class object
matcher = Matcher(nlp.vocab)
matcher.add("matching_1", None, pattern)

matches = matcher(doc)
span = doc[matches[0][1]:matches[0][2]]

print(span.text)

# ie_spacy_matcher_2.py
# Matcher class object
matcher = Matcher(nlp.vocab)

#define the pattern
pattern = [{'DEP': 'amod', 'OP': "?"},  # adjectival modifier
           {'POS': 'NOUN'},
           {'LOWER': 'such'},
           {'LOWER': 'as'},
           {'POS': 'PROPN'}]

matcher.add("matching_1", None, pattern)
matches = matcher(doc)

span = doc[matches[0][1]:matches[0][2]]
print(span.text)

# ie_dep_parse_2.py
doc = nlp("Here is how you can keep your car and other vehicles clean.")

# print dependency tags and POS tags

for tok in doc:
  print(tok.text, "-->", tok.dep_, "-->", tok.pos_)

# ie_spacy_matcher_3.py
# Matcher class object
matcher = Matcher(nlp.vocab)

#define the pattern
pattern = [{'DEP': 'amod', 'OP': "?"},
           {'POS': 'NOUN'},
           {'LOWER': 'and', 'OP': "?"},
           {'LOWER': 'or', 'OP': "?"},
           {'LOWER': 'other'},
           {'POS': 'NOUN'}]

matcher.add("matching_1", None, pattern)

matches = matcher(doc)
span = doc[matches[0][1]:matches[0][2]]
print(span.text)

# ie_dep_parse_3.py
doc = nlp("Eight people, including two children, were injured in the explosion")

for tok in doc:
  print(tok.text, "-->", tok.dep_, "-->", tok.pos_)

# ie_spacy_matcher_4.py
# Matcher class object
matcher = Matcher(nlp.vocab)

#define the pattern
pattern = [{'DEP': 'nummod', 'OP': "?"},  # numeric modifier
           {'DEP': 'amod', 'OP': "?"},  # adjectival modifier
           {'POS': 'NOUN'},
           {'IS_PUNCT': True},
           {'LOWER': 'including'},
           {'DEP': 'nummod', 'OP': "?"},
           {'DEP': 'amod', 'OP': "?"},
           {'POS': 'NOUN'}]

matcher.add("matching_1", None, pattern)

matches = matcher(doc)
span = doc[matches[0][1]:matches[0][2]]
print(span.text)

# ie_dep_parse_4.py
doc = nlp("A healthy eating pattern includes fruits, especially whole fruits.")

for tok in doc:
  print(tok.text, tok.dep_, tok.pos_)

# ie_spacy_matcher_5.py
# Matcher class object
matcher = Matcher(nlp.vocab)

#define the pattern
pattern = [{'DEP': 'nummod', 'OP': "?"},
           {'DEP': 'amod', 'OP': "?"},
           {'POS': 'NOUN'},
           {'IS_PUNCT': True},
           {'LOWER': 'especially'},
           {'DEP': 'nummod', 'OP': "?"},
           {'DEP': 'amod', 'OP': "?"},
           {'POS': 'NOUN'}]

matcher.add("matching_1", None, pattern)

matches = matcher(doc)
span = doc[matches[0][1]:matches[0][2]]
print(span.text)

# ie_dep_tree_plot.py
"""The simple rule-based methods work well for information extraction tasks. However, they have a few drawbacks and shortcomings.
- need to exhaust all rule for all possible patterns,
- not generalization
"""

text = "Tableau was recently acquired by Salesforce."

# Plot the dependency graph
doc = nlp(text)
displacy.render(doc, style='dep', jupyter=True)

# ie_dep_parse_5.py
text = "Tableau was recently acquired by Salesforce."
doc = nlp(text)

for tok in doc:
  print(tok.text, "-->", tok.dep_, "-->", tok.pos_)


def subtree_matcher(doc):
  subjpass = 0

  for i, tok in enumerate(doc):
    # find dependency tag that contains the text "subjpass"

    if tok.dep_.find("subjpass") == True:
      subjpass = 1

  x = ''
  y = ''

  # if subjpass == 1 then sentence is passive

  if subjpass == 1:
    for i, tok in enumerate(doc):
      if tok.dep_.find("subjpass") == True:
        y = tok.text

      if tok.dep_.endswith("obj") == True:
        x = tok.text

  # if subjpass == 0 then sentence is not passive
  else:
    for i, tok in enumerate(doc):
      if tok.dep_.endswith("subj") == True:
        x = tok.text

      if tok.dep_.endswith("obj") == True:
        y = tok.text

  return x, y


# ie_subtree_match_1.py
text_2 = "Careem, a ride hailing major in middle east, was acquired by Uber."

doc_2 = nlp(text_2)
subtree_matcher(doc_2)

# ie_subtree_match_2.py
text_3 = "Salesforce recently acquired Tableau."
doc_3 = nlp(text_3)
subtree_matcher(doc_3)

# ie_dep_parse_6.py

for tok in doc_3:
  print(tok.text, "-->", tok.dep_, "-->", tok.pos_)

# ie_subtree_match_3.py


print(subtree_matcher(doc_3))
print(subtree_matcher(nlp("Tableau was recently acquired by Salesforce.")))
