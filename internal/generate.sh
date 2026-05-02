#!/bin/sh

set -e

antlr4='java -jar ../tools/antlr-4.13.2-complete.jar'

$antlr4 -Dlanguage=Go -package lexer lexer/*.g4 -o .
$antlr4 -Dlanguage=Go -no-listener -visitor -package parser -lib lexer parser/*.g4 -o .