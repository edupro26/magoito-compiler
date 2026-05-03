lexer grammar MagoitoLexer;

CONST: 'const';
FUN: 'fun';
VAR: 'var';
IF: 'if';
THEN: 'then';
ELSE: 'else';
WHILE: 'while';
DO: 'do';
PRINT: 'print';
TRUE: 'true';
FALSE: 'false';
UNIT_VALUE: 'unit';

INT_TYPE: 'Int';
BOOL_TYPE: 'Bool';
STRING_TYPE: 'String';
UNIT_TYPE: 'Unit';

ARROW: '->';
EQ: '==';
NEQ: '!=';
LTE: '<=';
GTE: '>=';
OR: '||';
AND: '&&';
ASSIGN: '=';
LT: '<';
GT: '>';
PLUS: '+';
MINUS: '-';
STAR: '*';
SLASH: '/';
PERCENT: '%';
POWER: '^';
NOT: '!';
DOT: '.';
SEMICOLON: ';';
COLON: ':';
COMMA: ',';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
WILDCARD: '_';

IDENTIFIER
    : [a-zA-Z] [a-zA-Z0-9_']*
    ;

INT_LITERAL
    : '0'
    | [1-9] [0-9]*
    ;

STRING_LITERAL
    : '"' ~["\r\n]* '"'
    ;

LINE_COMMENT
    : '--' ~[\r\n]* -> skip
    ;

WS
    : [ \t\r\n]+ -> skip
    ;
