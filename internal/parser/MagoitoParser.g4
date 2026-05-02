parser grammar MagoitoParser;

options {
    tokenVocab = MagoitoLexer;
}

program: declaration+ EOF;

declaration
    : CONST binder COLON typeExpr ASSIGN expr #constDeclaration
    | FUN IDENTIFIER LPAREN binder (COMMA binder)* RPAREN COLON typeExpr ASSIGN expr #funDeclaration
    ;

binder
    : IDENTIFIER #idBinder
    | WILDCARD   #wildcardBinder
    ;

typeExpr
    : nonTupleType (ARROW typeExpr)?   #arrowType
    | tupleType ARROW typeExpr         #tupleArrowType
    ;

nonTupleType
    : basicType            #basicNonTupleType
    | recordType           #recordNonTupleType
    | LPAREN typeExpr RPAREN #parenNonTupleType
    ;

tupleType
    : LPAREN typeExpr COMMA typeExpr (COMMA typeExpr)* RPAREN
    ;

basicType
    : INT_TYPE    #intBasicType
    | BOOL_TYPE   #boolBasicType
    | STRING_TYPE #stringBasicType
    | UNIT_TYPE   #unitBasicType
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
    : varDeclExpr #varDeclControl
    | whileExpr   #whileControl
    | ifExpr      #ifControl
    | assignExpr  #assignControl
    | orExpr      #orControl
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
    : INT_LITERAL        #intLiteralPrimary
    | STRING_LITERAL     #stringLiteralPrimary
    | TRUE               #truePrimary
    | FALSE              #falsePrimary
    | UNIT_VALUE         #unitPrimary
    | callExpr           #callPrimary
    | IDENTIFIER         #identifierPrimary
    | LPAREN expr RPAREN #parenPrimary
    | recordExpr         #recordPrimary
    ;

callExpr
    : callee LPAREN expr (COMMA expr)* RPAREN
    ;

callee
    : IDENTIFIER #identCallee
    | PRINT      #printCallee
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
