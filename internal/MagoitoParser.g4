parser grammar MagoitoParser;

options {
    tokenVocab = MagoitoLexer;
}

program: declaration+ EOF;

declaration
    : CONST binder COLON typeExpr ASSIGN expr
    | FUN IDENTIFIER LPAREN binder (COMMA binder)* RPAREN COLON typeExpr ASSIGN expr
    ;

binder
    : IDENTIFIER
    | WILDCARD
    ;

typeExpr
    : nonTupleType (ARROW typeExpr)?   #arrowType
    | tupleType ARROW typeExpr         #tupleArrowType
    ;

nonTupleType
    : basicType
    | recordType
    | LPAREN typeExpr RPAREN
    ;

tupleType
    : LPAREN typeExpr COMMA typeExpr (COMMA typeExpr)* RPAREN
    ;

basicType
    : INT_TYPE
    | BOOL_TYPE
    | STRING_TYPE
    | UNIT_TYPE
    ;

recordType
    : LBRACE (recordTypeField (COMMA recordTypeField)*)? RBRACE
    ;

recordTypeField
    : IDENTIFIER COLON typeExpr
    ;

expr
    : seqExpr
    ;

seqExpr
    : controlExpr (SEMICOLON seqExpr)?
    ;

controlExpr
    : varDeclExpr
    | whileExpr
    | ifExpr
    | assignExpr
    | orExpr
    ;

varDeclExpr
    : VAR binder COLON typeExpr ASSIGN controlExpr
    ;

whileExpr
    : WHILE orExpr DO controlExpr
    ;

ifExpr
    : IF orExpr THEN controlExpr (ELSE controlExpr)?
    ;

assignExpr
    : IDENTIFIER ASSIGN controlExpr
    ;

orExpr
    : andExpr (OR andExpr)*
    ;

andExpr
    : equalityExpr (AND equalityExpr)*
    ;

equalityExpr
    : comparisonExpr ((EQ | NEQ) comparisonExpr)*
    ;

comparisonExpr
    : additiveExpr ((LT | LTE | GT | GTE) additiveExpr)*
    ;

additiveExpr
    : multiplicativeExpr ((PLUS | MINUS) multiplicativeExpr)*
    ;

multiplicativeExpr
    : powerExpr ((STAR | SLASH | PERCENT) powerExpr)*
    ;

powerExpr
    : unaryExpr (POWER powerExpr)?
    ;

unaryExpr
    : (MINUS | NOT) unaryExpr
    | projectionExpr
    ;

projectionExpr
    : primaryExpr (DOT IDENTIFIER)*
    ;

primaryExpr
    : INT_LITERAL
    | STRING_LITERAL
    | TRUE
    | FALSE
    | UNIT_VALUE
    | callExpr
    | IDENTIFIER
    | LPAREN expr RPAREN
    | recordExpr
    ;

callExpr
    : callee LPAREN expr (COMMA expr)* RPAREN
    ;

callee
    : IDENTIFIER
    | PRINT
    ;

recordExpr
    : LBRACE recordFieldList? RBRACE
    ;

recordFieldList
    : recordField (COMMA recordField)*
    ;

recordField
    : IDENTIFIER ASSIGN expr
    ;
