# nlp_ie_01.py
# reference: POS (https://universaldependencies.org/docs/u/pos/)
# import spacy
from random import randint
from IPython.display import Image, display
import visualise_spacy_tree
from spacy.matcher import Matcher
from spacy import displacy
import spacy

# load english language model
nlp = spacy.load('en_core_web_sm', disable=['ner', 'textcat'])

text = "This is a sample sentence."

# create spacy, and check the low level details for doc: tokens, token attributes
doc = nlp(text)

for token in doc:
    print(token.text, '->', token.pos_)

# nlp_ie_02.py

for token in doc:
    # check token pos

    if token.pos_ == 'NOUN':
        # print token
        print(token.text)

# nlp_ie_03.py
text = "The children love cream biscuits"

# create spacy
doc = nlp(text)

for token in doc:
    print(token.text, '->', token.pos_)

# nlp_ie_04.py
displacy.render(doc, style='dep', jupyter=True)

# nlp_ie_05.py

for token in doc:
    # extract subject

    if (token.dep_ == 'nsubj'):  # nominal subject
        print(token.text)
    # extract object
    elif (token.dep_ == 'dobj'):  # direct object
        print(token.text)

# nlp_ie_df.py
# Folder path
folders = glob.glob('./UNGD/UNGDC 1970-2018/Converted sessions/Session*')

# Dataframe
df = pd.DataFrame(columns={'Country', 'Speech', 'Session', 'Year'})

# Read speeches by India
i = 0

for folder in folders:

    speech = glob.glob(folder+'/IND*.txt')

    # only process one speed for one session
    with open(speech[0], encoding='utf8') as f:
        # Speech
        df.loc[i, 'Speech'] = f.read()
        # Year
        df.loc[i, 'Year'] = speech[0].split('_')[-1].split('.')[0]
        # Session
        df.loc[i, 'Session'] = speech[0].split('_')[-2]
        # Country
        df.loc[i, 'Country'] = speech[0].split('_')[0].split("\\")[-1]
        # Increment counter
        i += 1

df.head()

# nlp_ie_06.py
# function to preprocess speech


def clean(text):
    """Usual routine, clean the data"""

    # removing paragraph numbers
    text = re.sub('[0-9]+.\t', '', str(text))
    # removing new line characters
    text = re.sub('\n ', ' ', str(text))
    text = re.sub('\n', ' ', str(text))
    # removing apostrophes
    text = re.sub("'s", '', str(text))
    # removing hyphens
    text = re.sub("-", '', str(text))
    text = re.sub("— ", '', str(text))
    # removing quotation marks
    text = re.sub('\"', '', str(text))
    # removing salutations
    text = re.sub("Mr\.", 'Mr', str(text))
    text = re.sub("Mrs\.", 'Mrs', str(text))
    # removing any reference to outside text
    text = re.sub("[\(\[].*?[\)\]]", "", str(text))

    return text


# preprocessing speeches
df['Speech_clean'] = df['Speech'].apply(clean)

# nlp_ie_07.py
# split sentences


def sentences(text):
    """sentences.

    Args:
        text:
    """
    # split sentences and questions. split with '.' and '?' ...
    text = re.split('[.?!]', text)
    clean_sent = []

    for sent in text:
        clean_sent.append(sent)

    return clean_sent


# sentences
df['sent'] = df['Speech_clean'].apply(sentences)

# Sum: clean, split sentence (as the basic unit)
# nlp_ie_08.py


# load english language model
nlp = spacy.load('en_core_web_sm', disable=['ner', 'textcat'])

# nlp_ie_09.py
# function to find sentences containing PMs of India


def find_names(text):
    """find_names of Prime minister of India.

    Args:
        text: to be searched.
    """

    names = []

    # spacy doc
    doc = nlp(text)

    # pattern
    pattern = [{'LOWER': 'prime'},
               {'LOWER': 'minister'},
               {'POS': 'ADP', 'OP': '?'},
               {'POS': 'PROPN'}]

    # Matcher class object
    matcher = Matcher(nlp.vocab)
    matcher.add("names", None, pattern)

    matches = matcher(doc)

    # finding patterns in the text

    for i in range(0, len(matches)):

        # match: id, start, end
        token = doc[matches[i][1]:matches[i][2]]
        # append token to list
        names.append(str(token))

    # Only keep sentences containing Indian PMs

    for name in names:
        if (name.split()[2] == 'of') and (name.split()[3] != "India"):
                names.remove(name)

    return names


# apply function
df2['PM_Names'] = df2['Speech_clean'].apply(find_names)

# nlp_ie_10.py
# to check if keyswords like 'programs','schemes', etc. present in sentences


def prog_sent(text):
    """prog_sent get sentences which contain programs. Search with case insens-
    itivity.

    Args:
        text:
    """

    patterns = [r'\b(?i)'+'plan'+r'\b',
                r'\b(?i)'+'programme'+r'\b',
                r'\b(?i)'+'scheme'+r'\b',
                r'\b(?i)'+'campaign'+r'\b',
                r'\b(?i)'+'initiative'+r'\b',
                r'\b(?i)'+'conference'+r'\b',
                r'\b(?i)'+'agreement'+r'\b',
                r'\b(?i)'+'alliance'+r'\b']

    output = []
    flag = 0

    for pat in patterns:
        if re.search(pat, text) != None:
            flag = 1

            break

    return flag


# apply function
df2['Check_Schemes'] = df2['Sent'].apply(prog_sent)

# nlp_ie_11.py
# to extract initiatives using pattern matching


def all_schemes(text, check):
    """all_schemes.

    Args:
        text:
        check:
    """

    schemes = []

    doc = nlp(text)

    # initiatives
    prog_list = ['programme', 'scheme',
                 'initiative', 'campaign',
                 'agreement', 'conference',
                 'alliance', 'plan']

    # pattern to match initiatives names
    # e.g. The Paul Bill plan
    pattern = [{'POS': 'DET'},  # e.g. the
               {'POS': 'PROPN', 'DEP': 'compound'},  # proper noun
               {'POS': 'PROPN', 'DEP': 'compound'},
               {'POS': 'PROPN', 'OP': '?'},
               {'POS': 'PROPN', 'OP': '?'},
               {'POS': 'PROPN', 'OP': '?'},
               {'LOWER': {'IN': prog_list}, 'OP': '+'}
               ]

    if check == 0:
        # return blank list

        return schemes

    # Matcher class object
    matcher = Matcher(nlp.vocab)
    matcher.add("matching", None, pattern)
    matches = matcher(doc)

    for i in range(0, len(matches)):

        # match: id, start, end
        start, end = matches[i][1], matches[i][2]

        if doc[start].pos_ == 'DET':
            start = start+1

        # matched string
        span = str(doc[start:end])

        if (len(schemes) != 0) and (schemes[-1] in span):
            schemes[-1] = span
        else:
            schemes.append(span)

    return schemes


# apply function
df2['Schemes1'] = df2.apply(
    lambda x: all_schemes(x.Sent, x.Check_Schemes), axis=1)

# nlp_ie_12.py
# To understand the structure of the sentence.
doc = nlp(' Last year, I spoke about the Ujjwala programme, through which,
          I am happy to report, 50 million free liquid-gas connections have
          been provided so far')
png = visualise_spacy_tree.create_png(doc)
display(Image(png))

# nlp_ie_13.py
# rule to extract initiative name


def sent_subtree(text):
    """sent_subtree, In tree view to get the proper nouns in the subtree for
    initiative name.

    Args:
        text:
    """

    # pattern match for schemes or initiatives
    patterns = [r'\b(?i)'+'plan'+r'\b',
                r'\b(?i)'+'programme'+r'\b',
                r'\b(?i)'+'scheme'+r'\b',
                r'\b(?i)'+'campaign'+r'\b',
                r'\b(?i)'+'initiative'+r'\b',
                r'\b(?i)'+'conference'+r'\b',
                r'\b(?i)'+'agreement'+r'\b',
                r'\b(?i)'+'alliance'+r'\b']

    schemes = []
    doc = nlp(text)
    flag = 0
    # if no initiative present in sentence

    for pat in patterns:

        if re.search(pat, text) != None:
            flag = 1

            break

    if flag == 0:
        return schemes

    # iterating over sentence tokens

    for token in doc:

        for pat in patterns:

            # if we get a pattern match

            if re.search(pat, token.text) != None:

                word = ''
                # iterating over token subtree

                for node in token.subtree:
                    # only extract the proper nouns

                    if (node.pos_ == 'PROPN'):
                        word += node.text+' '

                if len(word) != 0:
                    schemes.append(word)

    return schemes


# derive initiatives
df2['Schemes2'] = df2['Sent'].apply(sent_subtree)

# nlp_ie_14.py
row_list = []
# df2 contains all sentences from all speeches

for i in range(len(df2)):
    sent = df2.loc[i, 'Sent']

    if (',' not in sent) and (len(sent.split()) <= 15):

        year = df2.loc[i, 'Year']
        length = len(sent.split())

        dict1 = {'Year': year, 'Sent': sent, 'Len': length}
        row_list.append(dict1)

# df with shorter sentences
df3 = pd.DataFrame(columns=['Year', 'Sent', "Len"])
df3 = pd.DataFrame(row_list)

# nlp_ie_15.py


def rand_sent(df):
    """rand_sent.

    Args:
        df:
    """

    index = randint(0, len(df))
    print('Index = ', index)
    doc = nlp(df.loc[index, 'Sent'][1:])
    displacy.render(doc, style='dep', jupyter=True)

    return index

# nlp_ie_16.py
# function to check output percentage for a rule


def output_per(df, out_col):
    """output_per is a function to check output percentage for a rule
    (percentage of non-trivial in a column).

    Args:
        df:
        out_col:
    """

    result = 0

    for out in df[out_col]:
        if len(out) != 0:
            result += 1

    per = result/len(df)
    per *= 100

    return per

# nlp_ie_18.py
# function for rule 1: noun(subject), verb, noun(object)


def rule1(text):
    """rule1: noun(subject), verb, noun(object)

    Args:
        text: to extract SVO info
    """

    doc = nlp(text)
    sent = []

    for token in doc:

        # if the token is a verb

        if (token.pos_ == 'VERB'):

            phrase = ''

            # only extract noun or pronoun subjects

            for sub_tok in token.lefts:

                if (sub_tok.dep_ in ['nsubj', 'nsubjpass']) and (sub_tok.pos_ in ['NOUN', 'PROPN', 'PRON']):

                    # add subject to the phrase
                    phrase += sub_tok.text

                    # save the root of the verb (our main token) in phrase
                    phrase += ' '+token.lemma_

                    # check for noun or pronoun direct objects

                    for sub_tok in token.rights:

                        # save the object in the phrase

                        if (sub_tok.dep_ in ['dobj']) and (sub_tok.pos_ in ['NOUN', 'PROPN']):

                            phrase += ' '+sub_tok.text
                            sent.append(phrase)

    return sent


# nlp_ie_19.py
# create a df containing sentence and its output for rule 1
row_list = []

for i in range(len(df3)):

    sent = df3.loc[i, 'Sent']
    year = df3.loc[i, 'Year']
    output = rule1(sent)
    dict1 = {'Year': year, 'Sent': sent, 'Output': output}
    row_list.append(dict1)

df_rule1 = pd.DataFrame(row_list)

# rule 1 achieves 20% result on simple sentences
output_per(df_rule1, 'Output')

# nlp_ie_20.py
# create a df containing sentence and its output for rule 1
row_list = []

# df2 contains all the sentences from all the speeches

for i in range(len(df2)):

    sent = df2.loc[i, 'Sent']
    year = df2.loc[i, 'Year']
    output = rule1(sent)
    dict1 = {'Year': year, 'Sent': sent, 'Output': output}
    row_list.append(dict1)

df_rule1_all = pd.DataFrame(row_list)

# check rule1 output on complete speeches
output_per(df_rule1_all, 'Output')

# nlp_ie_21.py
# Information Extraction #4 – Rule on Adjective Noun Structure
# function for rule 2


def rule2(text):
    """rule2, noun.

    Args:
        text:
    """

    doc = nlp(text)

    pat = []

    # iterate over tokens

    for token in doc:
        phrase = ''
        # if the word is a subject noun or an object noun

        if (token.pos_ == 'NOUN')\
                and (token.dep_ in ['dobj', 'pobj', 'nsubj', 'nsubjpass']):

            # iterate over the children nodes

            for subtoken in token.children:
                # if word is an adjective or has a compound dependency

                if (subtoken.pos_ == 'ADJ') or (subtoken.dep_ == 'compound'):
                    phrase += subtoken.text + ' '

            if len(phrase) != 0:
                phrase += token.text

        if len(phrase) != 0:
            pat.append(phrase)

    return pat

# nlp_ie_22.py


def rule2_mod(text, index):
    """rule2_mod.

    Args:
        text:
        index:
    """

    doc = nlp(text)

    phrase = ''

    for token in doc:

        if token.i == index:

            for subtoken in token.children:
                if (subtoken.pos_ == 'ADJ'):
                    phrase += ' '+subtoken.text

            break

    return phrase

# nlp_ie_23.py
# rule 1 modified function


def rule1_mod(text):
    """rule1_mod.

    Args:
        text:
    """

    doc = nlp(text)

    sent = []

    for token in doc:
        # root word

        if (token.pos_ == 'VERB'):

            phrase = ''

            # only extract noun or pronoun subjects

            for sub_tok in token.lefts:

                if (sub_tok.dep_ in ['nsubj', 'nsubjpass']) and (sub_tok.pos_ in ['NOUN', 'PROPN', 'PRON']):

                    # look for subject modifier
                    adj = rule2_mod(text, sub_tok.i)

                    phrase += adj + ' ' + sub_tok.text

                    # save the root word of the word
                    phrase += ' '+token.lemma_

                    # check for noun or pronoun direct objects

                    for sub_tok in token.rights:

                        if (sub_tok.dep_ in ['dobj']) and (sub_tok.pos_ in ['NOUN', 'PROPN']):

                            # look for object modifier
                            adj = rule2_mod(text, sub_tok.i)

                            phrase += adj+' '+sub_tok.text
                            sent.append(phrase)

    return sent

# nlp_ie_24.py
# rule 3 function


def rule3(text):
    """rule3.

    Args:
        text:
    """

    doc = nlp(text)

    sent = []

    for token in doc:

        # look for prepositions

        if token.pos_ == 'ADP':

            phrase = ''

            # if its head word is a noun

            if token.head.pos_ == 'NOUN':

                # append noun and preposition to phrase
                phrase += token.head.text
                phrase += ' '+token.text

                # check the nodes to the right of the preposition

                for right_tok in token.rights:
                    # append if it is a noun or proper noun

                    if (right_tok.pos_ in ['NOUN', 'PROPN']):
                        phrase += ' '+right_tok.text

                if len(phrase) > 2:
                    sent.append(phrase)

    return sent

# nlp_ie_25.py
# rule 0


def rule0(text, index):
    """rule0.

    Args:
        text:
        index:
    """

    doc = nlp(text)

    token = doc[index]

    entity = ''

    for sub_tok in token.children:
        if (sub_tok.dep_ in ['compound', 'amod']):
            entity += sub_tok.text+' '

    entity += token.text

    return entity

# nlp_ie_26.py
# rule 3 function


def rule3_mod(text):
    """rule3_mod.

    Args:
        text:
    """

    doc = nlp(text)

    sent = []

    for token in doc:

        if token.pos_ == 'ADP':

            phrase = ''

            if token.head.pos_ == 'NOUN':

                # appended rule
                append = rule0(text, token.head.i)

                if len(append) != 0:
                    phrase += append
                else:
                    phrase += token.head.text
                phrase += ' '+token.text

                for right_tok in token.rights:
                    if (right_tok.pos_ in ['NOUN', 'PROPN']):

                        right_phrase = ''
                        # appended rule
                        append = rule0(text, right_tok.i)

                        if len(append) != 0:
                            right_phrase += ' '+append
                        else:
                            right_phrase += ' '+right_tok.text

                        phrase += right_phrase

                if len(phrase) > 2:
                    sent.append(phrase)

    return sent


# nlp_ie_27.py
# create a dataframe containing sentences
df2 = pd.DataFrame(columns=['Sent', 'Year', 'Len'])

row_list = []

for i in range(len(df)):
    for sent in df.loc[i, 'sent']:

        wordcount = len(sent.split())
        year = df.loc[i, 'Year']

        dict1 = {'Year': year, 'Sent': sent, 'Len': wordcount}
        row_list.append(dict1)

df2 = pd.DataFrame(row_list)

# nlp_ie_28.py
# selecting non-empty output rows
df_show = pd.DataFrame(columns=df_rule1_all.columns)

for row in range(len(df_rule1_all)):

    if len(df_rule1_all.loc[row, 'Output']) != 0:
        df_show = df_show.append(df_rule1_all.loc[row, :])

# reset the index
df_show.reset_index(inplace=True)
df_show.drop('index', axis=1, inplace=True)

# nlp_ie_29.py
# separate subject, verb and object

verb_dict = dict()
dis_dict = dict()
dis_list = []

# iterating over all the sentences

for i in range(len(df_show)):

    # sentence containing the output
    sentence = df_show.loc[i, 'Sent']
    # year of the sentence
    year = df_show.loc[i, 'Year']
    # output of the sentence
    output = df_show.loc[i, 'Output']

    # iterating over all the outputs from the sentence

    for sent in output:

        # separate subject, verb and object
        n1, v, n2 = sent.split()[:1], sent.split()[1], sent.split()[2:]

        # append to list, along with the sentence
        dis_dict = {'Sent': sentence, 'Year': year,
                    'Noun1': n1, 'Verb': v, 'Noun2': n2}
        dis_list.append(dis_dict)

        # counting the number of sentences containing the verb
        verb = sent.split()[1]

        if verb in verb_dict:
            verb_dict[verb] += 1
        else:
            verb_dict[verb] = 1

df_sep = pd.DataFrame(dis_list)

# nlp_ie_30.py
# create a df containing sentence and its output for rule 2
row_list = []

for i in range(len(df3)):

    sent = df3.loc[i, 'Sent']
    year = df3.loc[i, 'Year']
    # rule
    output = rule2(sent)

    dict1 = {'Year': year, 'Sent': sent, 'Output': output}
    row_list.append(dict1)

df_rule2 = pd.DataFrame(row_list)

# nlp_ie_31.py
# create a df containing sentence and its output for rule 2
row_list = []

# df2 contains all the sentences from all the speeches

for i in range(len(df2)):

    sent = df2.loc[i, 'Sent']
    year = df2.loc[i, 'Year']
    output = rule2(sent)
    dict1 = {'Year': year, 'Sent': sent, 'Output': output}
    row_list.append(dict1)

df_rule2_all = pd.DataFrame(row_list)

# check rule output on complete speeches
output_per(df_rule2_all, 'Output')

# nlp_ie_32.py
# selecting non-empty outputs
df_show2 = pd.DataFrame(columns=df_rule2_all.columns)

for row in range(len(df_rule2_all)):

    if len(df_rule2_all.loc[row, 'Output']) != 0:
        df_show2 = df_show2.append(df_rule2_all.loc[row, :])

# reset the index
df_show2.reset_index(inplace=True)
df_show2.drop('index', axis=1, inplace=True)

# nlp_ie_33.py
# create a df containing sentence and its output for modified rule 1
row_list = []

# df2 contains all the sentences from all the speeches

for i in range(len(df2)):

    sent = df2.loc[i, 'Sent']
    year = df2.loc[i, 'Year']
    output = rule1_mod(sent)
    dict1 = {'Year': year, 'Sent': sent, 'Output': output}
    row_list.append(dict1)

df_rule1_mod_all = pd.DataFrame(row_list)
# check rule1 output on complete speeches
output_per(df_rule1_mod_all, 'Output')

# nlp_ie_34.py
# create a df containing sentence and its output for rule 3
row_list = []

for i in range(len(df3)):

    sent = df3.loc[i, 'Sent']
    year = df3.loc[i, 'Year']

    # rule
    output = rule3(sent)

    dict1 = {'Year': year, 'Sent': sent, 'Output': output}
    row_list.append(dict1)

df_rule3 = pd.DataFrame(row_list)
# output percentage for rule 3
output_per(df_rule3, 'Output')

# nlp_ie_35.py
# create a df containing sentence and its output for rule 3
row_list = []

# df2 contains all the sentences from all the speeches

for i in range(len(df2)):

    sent = df2.loc[i, 'Sent']
    year = df2.loc[i, 'Year']
    output = rule3(sent)
    dict1 = {'Year': year, 'Sent': sent, 'Output': output}
    row_list.append(dict1)

df_rule3_all = pd.DataFrame(row_list)
# check rule3 output on complete speeches
output_per(df_rule3_all, 'Output')

# nlp_ie_36.py
# select non-empty outputs
df_show3 = pd.DataFrame(columns=df_rule3_all.columns)

for row in range(len(df_rule3_all)):

    if len(df_rule3_all.loc[row, 'Output']) != 0:
        df_show3 = df_show3.append(df_rule3_all.loc[row, :])

# reset the index
df_show3.reset_index(inplace=True)
df_show3.drop('index', axis=1, inplace=True)

# nlp_ie_37.py
# separate noun, preposition and noun

prep_dict = dict()
dis_dict = dict()
dis_list = []

# iterating over all the sentences

for i in range(len(df_show3)):

    # sentence containing the output
    sentence = df_show3.loc[i, 'Sent']
    # year of the sentence
    year = df_show3.loc[i, 'Year']
    # output of the sentence
    output = df_show3.loc[i, 'Output']

    # iterating over all the outputs from the sentence

    for sent in output:

        # separate subject, verb and object
        n1, p, n2 = sent.split()[0], sent.split()[1], sent.split()[2:]

        # append to list, along with the sentence
        dis_dict = {'Sent': sentence, 'Year': year,
                    'Noun1': n1, 'Preposition': p, 'Noun2': n2}
        dis_list.append(dis_dict)

        # counting the number of sentences containing the verb
        prep = sent.split()[1]

        if prep in prep_dict:
            prep_dict[prep] += 1
        else:
            prep_dict[prep] = 1

df_sep3 = pd.DataFrame(dis_list)


# nlp_ie_prep.py
sort = sorted(prep_dict.items(), key=lambda d: (d[1], d[0]), reverse=True)
sort[:10]
