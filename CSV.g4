grammar CSV;

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
    | 'cook'
    | 'eat'
    | 'read'
    ;
OBJ : [a-zA-Z]+ ;
EMOTION
    : 'happy'
    | 'sad'
    | 'peace'
    ;
WEATHER:
    : 'good'
    | 'bad'
    ;
