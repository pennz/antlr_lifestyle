grammar DO;

# just subjects, objects, and vert first.

file : hdr row+ ;
hdr : row ;

row : field (',' field)* '\r'? '\n' ;

field
    :   TEXT	# text
    |   STRING	# string
    |		# empty
    ;

TEXT : ~[,\n\r"]+ ;
STRING : '"' ('""'|~'"')* '"' ;

ACTION
    : 'work'
    | 'buy'
    | 'cook'
    | 'eat'
    | 'read'
    | 'talk'
    ;

OBJ : [a-zA-Z]+ ;

# https://en.wikipedia.org/wiki/Sense
EMOTION
    : 'happy'
    | 'sad'
    | 'peace'
    ;



WEATHER
    : 'good'
    | 'bad'
    ;
