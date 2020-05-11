grammar Chat;

/*
 * Parser Rules
 */

chat				: line+ EOF ;

line				: name command NEWLINE ;

message				: (emoticon | link | color | mention | IDENT )+ ;

name				: IDENT ;

command				: (SAYS | SHOUTS) ':' message
                    | verb 'to'? object? ('because' reason)? ('by' method)? ('for' time)?
                    ;

verb: IDENT ;
object: IDENT ;
reason: message ;
method: message ;
time: NUMBER TIME_UNIT ;
					 					
emoticon			: ':' '-'? ')'
					| ':' '-'? '('
					;

link				: TEXT TEXT ;

color				: '/' IDENT '/' message '/';

mention				: '@' IDENT ;


/*
 * Lexer Rules
 */

fragment A			: ('A'|'a') ;
fragment S			: ('S'|'s') ;
fragment Y          : ('Y'|'y') ;
fragment H          : ('H'|'h') ;
fragment O          : ('O'|'o') ;
fragment U          : ('U'|'u') ;
fragment T          : ('T'|'t') ;

fragment 
ALP: [a-zA-Z];
fragment 
DIGIT: [0-9] ;
TIME_UNIT: 'day' | 'days'
            | 'hours' 
            | 'hour' 
            | 'month'
            | 'months'
            | 'year'
            | 'years'
            | 'minute' 
            | 'minutes' 
            | 'second' 
            | 'seconds'
            ;

IDENT : ALP (ALP | DIGIT | '_')* ;

SAYS				: S A Y S ;

SHOUTS				: S H O U T S ;

NEWLINE				: ('\r'? '\n' | '\r')+ ;

TEXT				: ('['|'(') ~[\])]+ (']'|')');

NUMBER: [1-9] DIGIT* ;

WS : [ \t]+ -> skip;  // toss out whitespace
