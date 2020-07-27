Important notes:
What is needed? needed in life


# landing pages
- Added thoughts
- todo
- Daily/weekly review. show statics. analyze. or KANBAN system. or we use another thing 

API , different parts of life.

# Project todo

## research paper
how to enjoy life. So it can be better. (for the book in kindle Jin Jing)

## data structure define
Ans: just input as lines of text describing what we have done today, it will be put to structured relational database, output graph data(a presentation)

## client restful API (external)
it could do thing for us (retrieve data for yout)

## we will need a datastorage thing. we just use mariadb. For this we just do pure sql, as I am learing this

- simple example.
- pipline needs do work first.
- script for test local (faster loop)
- let CI run normal tests
- add document generation
- image generation part( some message queue can be use, to view real time data) (knowledge graph)
- grading part
- ide for input(later), we just use nlp first, should be easier
- learn from bill board grammar, then just do it with antlr. For basic dataset operation, just a wrapper for SQL. Or we just add new items automatically.
- after we parsed our input, we put it to related dataset, then data analysis.

# Internal Implementation
## go and python
[data dog page introducing about the inter-operation](https://www.datadoghq.com/blog/engineering/cgo-and-python/)
Just python pool and save to database [YOU PROBABLY DON’T NEED A MESSAGE QUEUE by Bozho](https://techblog.bozho.net/you-probably-dont-need-a-message-queue/)

> You put a row with a processed=false flag in the database. A scheduled job
> runs, picks all unprocessed ones and processes them asynchronously. Then,
> when processing is finished, set the flag to true. I have used this approach
> a number of times, including large production systems, and it works pretty
> well.

So answer: just run scheduled tasks. Not go calling python... Not
over-complicate it.

## Message queue or not
Queue: let python handle inputed sentences, and put it to handled ones. The un-handled 
ones can be leave there for analysis.

## for display
Follow datadog agent? go just call python module?
