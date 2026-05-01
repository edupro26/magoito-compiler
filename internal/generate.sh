#!/bin/sh

set -e

antlr4='java -jar ./antlr-4.13.2-complete.jar'

$antlr4 -Dlanguage=Go -package lexer -o lexer MagoitoLexer.g4
$antlr4 -Dlanguage=Go -no-listener -visitor -package parser -lib lexer -o parser MagoitoParser.g4