Important notes:
What is needed? needed in life


# landing pages
- Added thoughts
- todo
- Daily/weekly review. show statics. analyze. or KANBAN system. or we use another thing 

API , different parts of life.

# Project todo


## research paper
how to enjoy life. So it can be better. (for the book in kindle Jin Jing 精进)

## data structure define
Ans: just input as lines of text describing what we have done today, it will be put to structured relational database, output graph data(a presentation)

## client restful API (external)
it could do thing for us (retrieve data for yout)

## we will need a datastorage thing. we just use mariadb. For this we just do pure sql, as I am learing this.

- simple example.
- pipline needs do work first.
- script for test local (faster loop)
- let CI run normal tests
- add document generation
- image generation part( some message queue can be use, to view real time data) (knowledge graph)
- grading part
- IDE for input(later), we just use NLP first, should be easier
- learn from bill board grammar, then just do it with antlr. For basic dataset operation, just a wrapper for SQL. Or we just add new items automatically.
- after we parsed our input, we put it to related dataset, then data analysis.

# Internal Implementation
- UML
- code generation (use case and test codes)

## go and python
[data dog page introducing about the inter-operation](https://www.datadoghq.com/blog/engineering/cgo-and-python/)
Just python pool and save to database [YOU PROBABLY DON’T NEED A MESSAGE QUEUE by Bozho](https://techblog.bozho.net/you-probably-dont-need-a-message-queue/)

> You put a row with a processed=false flag in the database. A scheduled job
> runs, picks all unprocessed ones and processes them asynchronously. Then,
> when processing is finished, set the flag to true. I have used this approach
> a number of times, including large production systems, and it works pretty
> well.

So answer: just run scheduled tasks. Not go calling python... Not Over-enginerring
over-complicate it. Not easy to maintain.

## learn from other application
### things
**todo**
You main thinking tool
- Today
- Upcoming
- Anytime
- Someday ( some time in the future)
**quick find**
nice to have

**goal tracking**
- progress split
- milestone
- notes taking (automatically review / reminder)(GNU Cash like, accounting)

## technically
Put to restful API, easier for exchange
### generate test cases

## Implementation ideals
### parse by python, and then send it back to our go API?
DO DO DO

### Message queue or not
Queue: let python handle inputed sentences, and put it to handled ones. The un-handled 
ones can be leave there for analysis.
Just like UDP.

Anyway,  we need parse out our input, and put it to dataset. Or read it out from 
dataset, then back to dataset.

### for display
We use Javascript highcharts thing to show graph (html tools is more available, interactive graphs)

## Components
client: flutter for input of sentences (https better), output of the graph/statistics
server: highcharts export server (only for UI) + nlp handling (nlp use python scapy) + (go or python? save to ) dataset
- spaCy server https://spacy.io/universe/project/spacy-server

Spacy Server:
- docker pull neelkamath/spacy-server
- https://hub.docker.com/r/neelkamath/spacy-server
- docker pull neelkamath/spacy-server:2-en_core_web_sm-sense2vec

# Ideal
- plan your day, and tracking the progress

    [review for Reflectly](https://medium.com/@bigdchang/reflectly-product-analysis-cd584a2aa98a)
    and the suggestion from this guy:
Find out your user's end goals/wishes/ambitions. Ask questions that help them
reflect about their progress and life. Create a best friend/father/mother
figure in your user's life that he or she never had.

Now, with journal entries and reflections that are relevant to a user's life
and goals, you have a few new avenues to encourage habitual usage.

- One, see progress. Mood tracking, but forget about the arbitrary smiley/blue faces. **Monitor my journey towards my goals**.
- Two, memories. Bring up my thoughts from a week, a month ago. Show me how far I've progressed.
- Three, blatantly ask what they could have done today to make their day even better. Then, ask them to do that tomorrow. And check the next day and ask about their progress again. Celebrate with them when they take a step forward. Raise them up when they fail.

