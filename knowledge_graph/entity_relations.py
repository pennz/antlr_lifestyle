#!/usr/bin/env python
# coding: utf8
"""A simple example of extracting relations between phrases and entities using
spaCy's named entity recognizer and the dependency parse. Here, we extract
money and currency values (entities labelled as MONEY) and then check the
dependency tree to find the noun phrase they are referring to – for example:
$9.4 million --> Net income.

Compatible with: spaCy v2.0.0+
Last tested with: v2.2.1
"""
from __future__ import unicode_literals, print_function

import plac
import spacy


TEXTS = [
    "I have made a sandwitch today.",
    "I have happily eaten a sandwitch today at Supermarket because I am hungry."
    " It made me feel good.",
    "I solve the problem by following the method found on Internet.",
]


@plac.annotations(
    model=("Model to load (needs parser and NER)", "positional", None, str)
)
def main(model="en_core_web_sm"):
    nlp = spacy.load(model)
    print("Loaded model '%s'" % model)
    print("Processing %d texts" % len(TEXTS))

    for text in TEXTS:
        print("Handling:", text)
        doc = nlp(text)
        relations = extract_currency_relations(doc)

        for r1, r2 in relations:
            print("{:<10}\t{}\t{}".format(r1.text, r2.ent_type_, r2.text))


def filter_spans(spans):
    # Filter a sequence of spans so they don't contain overlaps
    # For spaCy 2.1.4+: this function is available as spacy.util.filter_spans()
    get_sort_key = lambda span: (span.end - span.start, -span.start)
    sorted_spans = sorted(spans, key=get_sort_key, reverse=True)
    result = []
    seen_tokens = set()

    for span in sorted_spans:
        # Check for end - 1 here because boundaries are inclusive

        if span.start not in seen_tokens and span.end - 1 not in seen_tokens:
            result.append(span)
        seen_tokens.update(range(span.start, span.end))
    result = sorted(result, key=lambda span: span.start)

    return result


def extract_currency_relations(doc):
    # Merge entities and noun chunks into one token
    spans = list(doc.ents) + list(doc.noun_chunks)
    spans = filter_spans(spans)
    with doc.retokenize() as retokenizer:
        for span in spans:
            retokenizer.merge(span)

    relations = []

    for money in filter(lambda w: w.ent_type_ == "MONEY", doc):
        # example: "Net income was $9.4 million compared to the prior year of
        # $2.7 million.",
        # Here, we have net income = $9.4 million, as dobj

        # https://spacy.io/api/annotation#dependency-parsing-english
        # - dobj: direct oject
        # - nsubj: nominal subject
        # - pobj: object of preposition
        # - prep: prepositional modifier

        if money.dep_ in ("attr", "dobj"):
            # https://spacy.io/usage/linguistic-features#navigating
            # spaCy uses the terms head and child to describe the words
            # connected by a single arc in the dependency tree.
            subject = [w for w in money.head.lefts if w.dep_ == "nsubj"]

            if subject:
                subject = subject[0]
                relations.append((subject, money))
        elif money.dep_ == "pobj" and money.head.dep_ == "prep":
            relations.append((money.head.head, money))

    return relations


if __name__ == "__main__":
    plac.call(main)

    # Expected output:
    # Net income      MONEY   $9.4 million
    # the prior year  MONEY   $2.7 million
    # Revenue         MONEY   twelve billion dollars
    # a loss          MONEY   1b
