// Code generated from MagoitoParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MagoitoParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MagoitoParser struct {
	*antlr.BaseParser
}

var MagoitoParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func magoitoparserParserInit() {
	staticData := &MagoitoParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'const'", "'fun'", "'var'", "'if'", "'then'", "'else'", "'while'",
		"'do'", "'print'", "'true'", "'false'", "'unit'", "'Int'", "'Bool'",
		"'String'", "'Unit'", "'->'", "'=='", "'!='", "'<='", "'>='", "'||'",
		"'&&'", "'='", "'<'", "'>'", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'",
		"'!'", "'.'", "';'", "':'", "','", "'('", "')'", "'{'", "'}'", "'_'",
	}
	staticData.SymbolicNames = []string{
		"", "CONST", "FUN", "VAR", "IF", "THEN", "ELSE", "WHILE", "DO", "PRINT",
		"TRUE", "FALSE", "UNIT_VALUE", "INT_TYPE", "BOOL_TYPE", "STRING_TYPE",
		"UNIT_TYPE", "ARROW", "EQ", "NEQ", "LTE", "GTE", "OR", "AND", "ASSIGN",
		"LT", "GT", "PLUS", "MINUS", "STAR", "SLASH", "PERCENT", "POWER", "NOT",
		"DOT", "SEMICOLON", "COLON", "COMMA", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "WILDCARD", "IDENTIFIER", "INT_LITERAL", "STRING_LITERAL",
		"LINE_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "declaration", "binder", "typeExpr", "nonTupleType", "tupleType",
		"basicType", "recordType", "recordTypeField", "expr", "seqExpr", "controlExpr",
		"varDeclExpr", "whileExpr", "ifExpr", "assignExpr", "orExpr", "andExpr",
		"equalityExpr", "comparisonExpr", "additiveExpr", "multiplicativeExpr",
		"powerExpr", "unaryExpr", "projectionExpr", "primaryExpr", "callExpr",
		"callee", "recordExpr", "recordFieldList", "recordField",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 47, 299, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 4,
		0, 64, 8, 0, 11, 0, 12, 0, 65, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 83, 8, 1, 10, 1,
		12, 1, 86, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 94, 8, 1, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 101, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3,
		3, 107, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 115, 8, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 123, 8, 5, 10, 5, 12, 5, 126, 9, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 136, 8, 7, 10, 7,
		12, 7, 139, 9, 7, 3, 7, 141, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10, 154, 8, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 3, 11, 161, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 3, 14, 181, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 5, 16, 190, 8, 16, 10, 16, 12, 16, 193, 9, 16, 1, 17, 1, 17,
		1, 17, 5, 17, 198, 8, 17, 10, 17, 12, 17, 201, 9, 17, 1, 18, 1, 18, 1,
		18, 5, 18, 206, 8, 18, 10, 18, 12, 18, 209, 9, 18, 1, 19, 1, 19, 1, 19,
		5, 19, 214, 8, 19, 10, 19, 12, 19, 217, 9, 19, 1, 20, 1, 20, 1, 20, 5,
		20, 222, 8, 20, 10, 20, 12, 20, 225, 9, 20, 1, 21, 1, 21, 1, 21, 5, 21,
		230, 8, 21, 10, 21, 12, 21, 233, 9, 21, 1, 22, 1, 22, 1, 22, 3, 22, 238,
		8, 22, 1, 23, 1, 23, 1, 23, 3, 23, 243, 8, 23, 1, 24, 1, 24, 1, 24, 5,
		24, 248, 8, 24, 10, 24, 12, 24, 251, 9, 24, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 265, 8,
		25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 272, 8, 26, 10, 26, 12, 26,
		275, 9, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 3, 28, 283, 8, 28,
		1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 5, 29, 290, 8, 29, 10, 29, 12, 29, 293,
		9, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 0, 0, 31, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 56, 58, 60, 0, 8, 1, 0, 42, 43, 1, 0, 13, 16, 1, 0, 18,
		19, 2, 0, 20, 21, 25, 26, 1, 0, 27, 28, 1, 0, 29, 31, 2, 0, 28, 28, 33,
		33, 2, 0, 9, 9, 43, 43, 303, 0, 63, 1, 0, 0, 0, 2, 93, 1, 0, 0, 0, 4, 95,
		1, 0, 0, 0, 6, 106, 1, 0, 0, 0, 8, 114, 1, 0, 0, 0, 10, 116, 1, 0, 0, 0,
		12, 129, 1, 0, 0, 0, 14, 131, 1, 0, 0, 0, 16, 144, 1, 0, 0, 0, 18, 148,
		1, 0, 0, 0, 20, 150, 1, 0, 0, 0, 22, 160, 1, 0, 0, 0, 24, 162, 1, 0, 0,
		0, 26, 169, 1, 0, 0, 0, 28, 174, 1, 0, 0, 0, 30, 182, 1, 0, 0, 0, 32, 186,
		1, 0, 0, 0, 34, 194, 1, 0, 0, 0, 36, 202, 1, 0, 0, 0, 38, 210, 1, 0, 0,
		0, 40, 218, 1, 0, 0, 0, 42, 226, 1, 0, 0, 0, 44, 234, 1, 0, 0, 0, 46, 242,
		1, 0, 0, 0, 48, 244, 1, 0, 0, 0, 50, 264, 1, 0, 0, 0, 52, 266, 1, 0, 0,
		0, 54, 278, 1, 0, 0, 0, 56, 280, 1, 0, 0, 0, 58, 286, 1, 0, 0, 0, 60, 294,
		1, 0, 0, 0, 62, 64, 3, 2, 1, 0, 63, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0,
		65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 5,
		0, 0, 1, 68, 1, 1, 0, 0, 0, 69, 70, 5, 1, 0, 0, 70, 71, 3, 4, 2, 0, 71,
		72, 5, 36, 0, 0, 72, 73, 3, 6, 3, 0, 73, 74, 5, 24, 0, 0, 74, 75, 3, 18,
		9, 0, 75, 94, 1, 0, 0, 0, 76, 77, 5, 2, 0, 0, 77, 78, 5, 43, 0, 0, 78,
		79, 5, 38, 0, 0, 79, 84, 3, 4, 2, 0, 80, 81, 5, 37, 0, 0, 81, 83, 3, 4,
		2, 0, 82, 80, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85,
		1, 0, 0, 0, 85, 87, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 88, 5, 39, 0, 0,
		88, 89, 5, 36, 0, 0, 89, 90, 3, 6, 3, 0, 90, 91, 5, 24, 0, 0, 91, 92, 3,
		18, 9, 0, 92, 94, 1, 0, 0, 0, 93, 69, 1, 0, 0, 0, 93, 76, 1, 0, 0, 0, 94,
		3, 1, 0, 0, 0, 95, 96, 7, 0, 0, 0, 96, 5, 1, 0, 0, 0, 97, 100, 3, 8, 4,
		0, 98, 99, 5, 17, 0, 0, 99, 101, 3, 6, 3, 0, 100, 98, 1, 0, 0, 0, 100,
		101, 1, 0, 0, 0, 101, 107, 1, 0, 0, 0, 102, 103, 3, 10, 5, 0, 103, 104,
		5, 17, 0, 0, 104, 105, 3, 6, 3, 0, 105, 107, 1, 0, 0, 0, 106, 97, 1, 0,
		0, 0, 106, 102, 1, 0, 0, 0, 107, 7, 1, 0, 0, 0, 108, 115, 3, 12, 6, 0,
		109, 115, 3, 14, 7, 0, 110, 111, 5, 38, 0, 0, 111, 112, 3, 6, 3, 0, 112,
		113, 5, 39, 0, 0, 113, 115, 1, 0, 0, 0, 114, 108, 1, 0, 0, 0, 114, 109,
		1, 0, 0, 0, 114, 110, 1, 0, 0, 0, 115, 9, 1, 0, 0, 0, 116, 117, 5, 38,
		0, 0, 117, 118, 3, 6, 3, 0, 118, 119, 5, 37, 0, 0, 119, 124, 3, 6, 3, 0,
		120, 121, 5, 37, 0, 0, 121, 123, 3, 6, 3, 0, 122, 120, 1, 0, 0, 0, 123,
		126, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 127,
		1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 128, 5, 39, 0, 0, 128, 11, 1, 0,
		0, 0, 129, 130, 7, 1, 0, 0, 130, 13, 1, 0, 0, 0, 131, 140, 5, 40, 0, 0,
		132, 137, 3, 16, 8, 0, 133, 134, 5, 37, 0, 0, 134, 136, 3, 16, 8, 0, 135,
		133, 1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138,
		1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 132, 1, 0,
		0, 0, 140, 141, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 5, 41, 0, 0,
		143, 15, 1, 0, 0, 0, 144, 145, 5, 43, 0, 0, 145, 146, 5, 36, 0, 0, 146,
		147, 3, 6, 3, 0, 147, 17, 1, 0, 0, 0, 148, 149, 3, 20, 10, 0, 149, 19,
		1, 0, 0, 0, 150, 153, 3, 22, 11, 0, 151, 152, 5, 35, 0, 0, 152, 154, 3,
		20, 10, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 21, 1, 0, 0,
		0, 155, 161, 3, 24, 12, 0, 156, 161, 3, 26, 13, 0, 157, 161, 3, 28, 14,
		0, 158, 161, 3, 30, 15, 0, 159, 161, 3, 32, 16, 0, 160, 155, 1, 0, 0, 0,
		160, 156, 1, 0, 0, 0, 160, 157, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160,
		159, 1, 0, 0, 0, 161, 23, 1, 0, 0, 0, 162, 163, 5, 3, 0, 0, 163, 164, 3,
		4, 2, 0, 164, 165, 5, 36, 0, 0, 165, 166, 3, 6, 3, 0, 166, 167, 5, 24,
		0, 0, 167, 168, 3, 22, 11, 0, 168, 25, 1, 0, 0, 0, 169, 170, 5, 7, 0, 0,
		170, 171, 3, 32, 16, 0, 171, 172, 5, 8, 0, 0, 172, 173, 3, 22, 11, 0, 173,
		27, 1, 0, 0, 0, 174, 175, 5, 4, 0, 0, 175, 176, 3, 32, 16, 0, 176, 177,
		5, 5, 0, 0, 177, 180, 3, 22, 11, 0, 178, 179, 5, 6, 0, 0, 179, 181, 3,
		22, 11, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 29, 1, 0, 0,
		0, 182, 183, 5, 43, 0, 0, 183, 184, 5, 24, 0, 0, 184, 185, 3, 22, 11, 0,
		185, 31, 1, 0, 0, 0, 186, 191, 3, 34, 17, 0, 187, 188, 5, 22, 0, 0, 188,
		190, 3, 34, 17, 0, 189, 187, 1, 0, 0, 0, 190, 193, 1, 0, 0, 0, 191, 189,
		1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 33, 1, 0, 0, 0, 193, 191, 1, 0,
		0, 0, 194, 199, 3, 36, 18, 0, 195, 196, 5, 23, 0, 0, 196, 198, 3, 36, 18,
		0, 197, 195, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199,
		200, 1, 0, 0, 0, 200, 35, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 207, 3,
		38, 19, 0, 203, 204, 7, 2, 0, 0, 204, 206, 3, 38, 19, 0, 205, 203, 1, 0,
		0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0,
		208, 37, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 210, 215, 3, 40, 20, 0, 211,
		212, 7, 3, 0, 0, 212, 214, 3, 40, 20, 0, 213, 211, 1, 0, 0, 0, 214, 217,
		1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 39, 1, 0,
		0, 0, 217, 215, 1, 0, 0, 0, 218, 223, 3, 42, 21, 0, 219, 220, 7, 4, 0,
		0, 220, 222, 3, 42, 21, 0, 221, 219, 1, 0, 0, 0, 222, 225, 1, 0, 0, 0,
		223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 41, 1, 0, 0, 0, 225, 223,
		1, 0, 0, 0, 226, 231, 3, 44, 22, 0, 227, 228, 7, 5, 0, 0, 228, 230, 3,
		44, 22, 0, 229, 227, 1, 0, 0, 0, 230, 233, 1, 0, 0, 0, 231, 229, 1, 0,
		0, 0, 231, 232, 1, 0, 0, 0, 232, 43, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0,
		234, 237, 3, 46, 23, 0, 235, 236, 5, 32, 0, 0, 236, 238, 3, 44, 22, 0,
		237, 235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 45, 1, 0, 0, 0, 239, 240,
		7, 6, 0, 0, 240, 243, 3, 46, 23, 0, 241, 243, 3, 48, 24, 0, 242, 239, 1,
		0, 0, 0, 242, 241, 1, 0, 0, 0, 243, 47, 1, 0, 0, 0, 244, 249, 3, 50, 25,
		0, 245, 246, 5, 34, 0, 0, 246, 248, 5, 43, 0, 0, 247, 245, 1, 0, 0, 0,
		248, 251, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250,
		49, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 252, 265, 5, 44, 0, 0, 253, 265,
		5, 45, 0, 0, 254, 265, 5, 10, 0, 0, 255, 265, 5, 11, 0, 0, 256, 265, 5,
		12, 0, 0, 257, 265, 3, 52, 26, 0, 258, 265, 5, 43, 0, 0, 259, 260, 5, 38,
		0, 0, 260, 261, 3, 18, 9, 0, 261, 262, 5, 39, 0, 0, 262, 265, 1, 0, 0,
		0, 263, 265, 3, 56, 28, 0, 264, 252, 1, 0, 0, 0, 264, 253, 1, 0, 0, 0,
		264, 254, 1, 0, 0, 0, 264, 255, 1, 0, 0, 0, 264, 256, 1, 0, 0, 0, 264,
		257, 1, 0, 0, 0, 264, 258, 1, 0, 0, 0, 264, 259, 1, 0, 0, 0, 264, 263,
		1, 0, 0, 0, 265, 51, 1, 0, 0, 0, 266, 267, 3, 54, 27, 0, 267, 268, 5, 38,
		0, 0, 268, 273, 3, 18, 9, 0, 269, 270, 5, 37, 0, 0, 270, 272, 3, 18, 9,
		0, 271, 269, 1, 0, 0, 0, 272, 275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273,
		274, 1, 0, 0, 0, 274, 276, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 276, 277,
		5, 39, 0, 0, 277, 53, 1, 0, 0, 0, 278, 279, 7, 7, 0, 0, 279, 55, 1, 0,
		0, 0, 280, 282, 5, 40, 0, 0, 281, 283, 3, 58, 29, 0, 282, 281, 1, 0, 0,
		0, 282, 283, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 285, 5, 41, 0, 0, 285,
		57, 1, 0, 0, 0, 286, 291, 3, 60, 30, 0, 287, 288, 5, 37, 0, 0, 288, 290,
		3, 60, 30, 0, 289, 287, 1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 289, 1,
		0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 59, 1, 0, 0, 0, 293, 291, 1, 0, 0,
		0, 294, 295, 5, 43, 0, 0, 295, 296, 5, 24, 0, 0, 296, 297, 3, 18, 9, 0,
		297, 61, 1, 0, 0, 0, 25, 65, 84, 93, 100, 106, 114, 124, 137, 140, 153,
		160, 180, 191, 199, 207, 215, 223, 231, 237, 242, 249, 264, 273, 282, 291,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MagoitoParserInit initializes any static state used to implement MagoitoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMagoitoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MagoitoParserInit() {
	staticData := &MagoitoParserParserStaticData
	staticData.once.Do(magoitoparserParserInit)
}

// NewMagoitoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMagoitoParser(input antlr.TokenStream) *MagoitoParser {
	MagoitoParserInit()
	this := new(MagoitoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MagoitoParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MagoitoParser.g4"

	return this
}

// MagoitoParser tokens.
const (
	MagoitoParserEOF            = antlr.TokenEOF
	MagoitoParserCONST          = 1
	MagoitoParserFUN            = 2
	MagoitoParserVAR            = 3
	MagoitoParserIF             = 4
	MagoitoParserTHEN           = 5
	MagoitoParserELSE           = 6
	MagoitoParserWHILE          = 7
	MagoitoParserDO             = 8
	MagoitoParserPRINT          = 9
	MagoitoParserTRUE           = 10
	MagoitoParserFALSE          = 11
	MagoitoParserUNIT_VALUE     = 12
	MagoitoParserINT_TYPE       = 13
	MagoitoParserBOOL_TYPE      = 14
	MagoitoParserSTRING_TYPE    = 15
	MagoitoParserUNIT_TYPE      = 16
	MagoitoParserARROW          = 17
	MagoitoParserEQ             = 18
	MagoitoParserNEQ            = 19
	MagoitoParserLTE            = 20
	MagoitoParserGTE            = 21
	MagoitoParserOR             = 22
	MagoitoParserAND            = 23
	MagoitoParserASSIGN         = 24
	MagoitoParserLT             = 25
	MagoitoParserGT             = 26
	MagoitoParserPLUS           = 27
	MagoitoParserMINUS          = 28
	MagoitoParserSTAR           = 29
	MagoitoParserSLASH          = 30
	MagoitoParserPERCENT        = 31
	MagoitoParserPOWER          = 32
	MagoitoParserNOT            = 33
	MagoitoParserDOT            = 34
	MagoitoParserSEMICOLON      = 35
	MagoitoParserCOLON          = 36
	MagoitoParserCOMMA          = 37
	MagoitoParserLPAREN         = 38
	MagoitoParserRPAREN         = 39
	MagoitoParserLBRACE         = 40
	MagoitoParserRBRACE         = 41
	MagoitoParserWILDCARD       = 42
	MagoitoParserIDENTIFIER     = 43
	MagoitoParserINT_LITERAL    = 44
	MagoitoParserSTRING_LITERAL = 45
	MagoitoParserLINE_COMMENT   = 46
	MagoitoParserWS             = 47
)

// MagoitoParser rules.
const (
	MagoitoParserRULE_program            = 0
	MagoitoParserRULE_declaration        = 1
	MagoitoParserRULE_binder             = 2
	MagoitoParserRULE_typeExpr           = 3
	MagoitoParserRULE_nonTupleType       = 4
	MagoitoParserRULE_tupleType          = 5
	MagoitoParserRULE_basicType          = 6
	MagoitoParserRULE_recordType         = 7
	MagoitoParserRULE_recordTypeField    = 8
	MagoitoParserRULE_expr               = 9
	MagoitoParserRULE_seqExpr            = 10
	MagoitoParserRULE_controlExpr        = 11
	MagoitoParserRULE_varDeclExpr        = 12
	MagoitoParserRULE_whileExpr          = 13
	MagoitoParserRULE_ifExpr             = 14
	MagoitoParserRULE_assignExpr         = 15
	MagoitoParserRULE_orExpr             = 16
	MagoitoParserRULE_andExpr            = 17
	MagoitoParserRULE_equalityExpr       = 18
	MagoitoParserRULE_comparisonExpr     = 19
	MagoitoParserRULE_additiveExpr       = 20
	MagoitoParserRULE_multiplicativeExpr = 21
	MagoitoParserRULE_powerExpr          = 22
	MagoitoParserRULE_unaryExpr          = 23
	MagoitoParserRULE_projectionExpr     = 24
	MagoitoParserRULE_primaryExpr        = 25
	MagoitoParserRULE_callExpr           = 26
	MagoitoParserRULE_callee             = 27
	MagoitoParserRULE_recordExpr         = 28
	MagoitoParserRULE_recordFieldList    = 29
	MagoitoParserRULE_recordField        = 30
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(MagoitoParserEOF, 0)
}

func (s *ProgramContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MagoitoParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MagoitoParserCONST || _la == MagoitoParserFUN {
		{
			p.SetState(62)
			p.Declaration()
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(67)
		p.Match(MagoitoParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	AllBinder() []IBinderContext
	Binder(i int) IBinderContext
	COLON() antlr.TerminalNode
	TypeExpr() ITypeExprContext
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext
	FUN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) CONST() antlr.TerminalNode {
	return s.GetToken(MagoitoParserCONST, 0)
}

func (s *DeclarationContext) AllBinder() []IBinderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBinderContext); ok {
			len++
		}
	}

	tst := make([]IBinderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBinderContext); ok {
			tst[i] = t.(IBinderContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationContext) Binder(i int) IBinderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinderContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinderContext)
}

func (s *DeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(MagoitoParserCOLON, 0)
}

func (s *DeclarationContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *DeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserASSIGN, 0)
}

func (s *DeclarationContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclarationContext) FUN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserFUN, 0)
}

func (s *DeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MagoitoParserIDENTIFIER, 0)
}

func (s *DeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserLPAREN, 0)
}

func (s *DeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserRPAREN, 0)
}

func (s *DeclarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserCOMMA)
}

func (s *DeclarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserCOMMA, i)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MagoitoParserRULE_declaration)
	var _la int

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MagoitoParserCONST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Match(MagoitoParserCONST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Binder()
		}
		{
			p.SetState(71)
			p.Match(MagoitoParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.TypeExpr()
		}
		{
			p.SetState(73)
			p.Match(MagoitoParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Expr()
		}

	case MagoitoParserFUN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.Match(MagoitoParserFUN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(MagoitoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Match(MagoitoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Binder()
		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MagoitoParserCOMMA {
			{
				p.SetState(80)
				p.Match(MagoitoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(81)
				p.Binder()
			}

			p.SetState(86)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(87)
			p.Match(MagoitoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Match(MagoitoParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.TypeExpr()
		}
		{
			p.SetState(90)
			p.Match(MagoitoParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Expr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinderContext is an interface to support dynamic dispatch.
type IBinderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	WILDCARD() antlr.TerminalNode

	// IsBinderContext differentiates from other interfaces.
	IsBinderContext()
}

type BinderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinderContext() *BinderContext {
	var p = new(BinderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_binder
	return p
}

func InitEmptyBinderContext(p *BinderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_binder
}

func (*BinderContext) IsBinderContext() {}

func NewBinderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinderContext {
	var p = new(BinderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_binder

	return p
}

func (s *BinderContext) GetParser() antlr.Parser { return s.parser }

func (s *BinderContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MagoitoParserIDENTIFIER, 0)
}

func (s *BinderContext) WILDCARD() antlr.TerminalNode {
	return s.GetToken(MagoitoParserWILDCARD, 0)
}

func (s *BinderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterBinder(s)
	}
}

func (s *BinderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitBinder(s)
	}
}

func (s *BinderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitBinder(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) Binder() (localctx IBinderContext) {
	localctx = NewBinderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MagoitoParserRULE_binder)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MagoitoParserWILDCARD || _la == MagoitoParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeExprContext is an interface to support dynamic dispatch.
type ITypeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeExprContext differentiates from other interfaces.
	IsTypeExprContext()
}

type TypeExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprContext() *TypeExprContext {
	var p = new(TypeExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_typeExpr
	return p
}

func InitEmptyTypeExprContext(p *TypeExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_typeExpr
}

func (*TypeExprContext) IsTypeExprContext() {}

func NewTypeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprContext {
	var p = new(TypeExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_typeExpr

	return p
}

func (s *TypeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeExprContext) CopyAll(ctx *TypeExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TupleArrowTypeContext struct {
	TypeExprContext
}

func NewTupleArrowTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleArrowTypeContext {
	var p = new(TupleArrowTypeContext)

	InitEmptyTypeExprContext(&p.TypeExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeExprContext))

	return p
}

func (s *TupleArrowTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleArrowTypeContext) TupleType() ITupleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleTypeContext)
}

func (s *TupleArrowTypeContext) ARROW() antlr.TerminalNode {
	return s.GetToken(MagoitoParserARROW, 0)
}

func (s *TupleArrowTypeContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *TupleArrowTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterTupleArrowType(s)
	}
}

func (s *TupleArrowTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitTupleArrowType(s)
	}
}

func (s *TupleArrowTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitTupleArrowType(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrowTypeContext struct {
	TypeExprContext
}

func NewArrowTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrowTypeContext {
	var p = new(ArrowTypeContext)

	InitEmptyTypeExprContext(&p.TypeExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeExprContext))

	return p
}

func (s *ArrowTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrowTypeContext) NonTupleType() INonTupleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INonTupleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INonTupleTypeContext)
}

func (s *ArrowTypeContext) ARROW() antlr.TerminalNode {
	return s.GetToken(MagoitoParserARROW, 0)
}

func (s *ArrowTypeContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *ArrowTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterArrowType(s)
	}
}

func (s *ArrowTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitArrowType(s)
	}
}

func (s *ArrowTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitArrowType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) TypeExpr() (localctx ITypeExprContext) {
	localctx = NewTypeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MagoitoParserRULE_typeExpr)
	var _la int

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArrowTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.NonTupleType()
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MagoitoParserARROW {
			{
				p.SetState(98)
				p.Match(MagoitoParserARROW)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(99)
				p.TypeExpr()
			}

		}

	case 2:
		localctx = NewTupleArrowTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.TupleType()
		}
		{
			p.SetState(103)
			p.Match(MagoitoParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.TypeExpr()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INonTupleTypeContext is an interface to support dynamic dispatch.
type INonTupleTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BasicType() IBasicTypeContext
	RecordType() IRecordTypeContext
	LPAREN() antlr.TerminalNode
	TypeExpr() ITypeExprContext
	RPAREN() antlr.TerminalNode

	// IsNonTupleTypeContext differentiates from other interfaces.
	IsNonTupleTypeContext()
}

type NonTupleTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonTupleTypeContext() *NonTupleTypeContext {
	var p = new(NonTupleTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_nonTupleType
	return p
}

func InitEmptyNonTupleTypeContext(p *NonTupleTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_nonTupleType
}

func (*NonTupleTypeContext) IsNonTupleTypeContext() {}

func NewNonTupleTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonTupleTypeContext {
	var p = new(NonTupleTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_nonTupleType

	return p
}

func (s *NonTupleTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *NonTupleTypeContext) BasicType() IBasicTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasicTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasicTypeContext)
}

func (s *NonTupleTypeContext) RecordType() IRecordTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecordTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecordTypeContext)
}

func (s *NonTupleTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserLPAREN, 0)
}

func (s *NonTupleTypeContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *NonTupleTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserRPAREN, 0)
}

func (s *NonTupleTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonTupleTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonTupleTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterNonTupleType(s)
	}
}

func (s *NonTupleTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitNonTupleType(s)
	}
}

func (s *NonTupleTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitNonTupleType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) NonTupleType() (localctx INonTupleTypeContext) {
	localctx = NewNonTupleTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MagoitoParserRULE_nonTupleType)
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MagoitoParserINT_TYPE, MagoitoParserBOOL_TYPE, MagoitoParserSTRING_TYPE, MagoitoParserUNIT_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.BasicType()
		}

	case MagoitoParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.RecordType()
		}

	case MagoitoParserLPAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.Match(MagoitoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.TypeExpr()
		}
		{
			p.SetState(112)
			p.Match(MagoitoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITupleTypeContext is an interface to support dynamic dispatch.
type ITupleTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllTypeExpr() []ITypeExprContext
	TypeExpr(i int) ITypeExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsTupleTypeContext differentiates from other interfaces.
	IsTupleTypeContext()
}

type TupleTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTupleTypeContext() *TupleTypeContext {
	var p = new(TupleTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_tupleType
	return p
}

func InitEmptyTupleTypeContext(p *TupleTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_tupleType
}

func (*TupleTypeContext) IsTupleTypeContext() {}

func NewTupleTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TupleTypeContext {
	var p = new(TupleTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_tupleType

	return p
}

func (s *TupleTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TupleTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserLPAREN, 0)
}

func (s *TupleTypeContext) AllTypeExpr() []ITypeExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeExprContext); ok {
			len++
		}
	}

	tst := make([]ITypeExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeExprContext); ok {
			tst[i] = t.(ITypeExprContext)
			i++
		}
	}

	return tst
}

func (s *TupleTypeContext) TypeExpr(i int) ITypeExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *TupleTypeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserCOMMA)
}

func (s *TupleTypeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserCOMMA, i)
}

func (s *TupleTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserRPAREN, 0)
}

func (s *TupleTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TupleTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterTupleType(s)
	}
}

func (s *TupleTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitTupleType(s)
	}
}

func (s *TupleTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitTupleType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) TupleType() (localctx ITupleTypeContext) {
	localctx = NewTupleTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MagoitoParserRULE_tupleType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(MagoitoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.TypeExpr()
	}
	{
		p.SetState(118)
		p.Match(MagoitoParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.TypeExpr()
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserCOMMA {
		{
			p.SetState(120)
			p.Match(MagoitoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.TypeExpr()
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(127)
		p.Match(MagoitoParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBasicTypeContext is an interface to support dynamic dispatch.
type IBasicTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_TYPE() antlr.TerminalNode
	BOOL_TYPE() antlr.TerminalNode
	STRING_TYPE() antlr.TerminalNode
	UNIT_TYPE() antlr.TerminalNode

	// IsBasicTypeContext differentiates from other interfaces.
	IsBasicTypeContext()
}

type BasicTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicTypeContext() *BasicTypeContext {
	var p = new(BasicTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_basicType
	return p
}

func InitEmptyBasicTypeContext(p *BasicTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_basicType
}

func (*BasicTypeContext) IsBasicTypeContext() {}

func NewBasicTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicTypeContext {
	var p = new(BasicTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_basicType

	return p
}

func (s *BasicTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicTypeContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserINT_TYPE, 0)
}

func (s *BasicTypeContext) BOOL_TYPE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserBOOL_TYPE, 0)
}

func (s *BasicTypeContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserSTRING_TYPE, 0)
}

func (s *BasicTypeContext) UNIT_TYPE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserUNIT_TYPE, 0)
}

func (s *BasicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterBasicType(s)
	}
}

func (s *BasicTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitBasicType(s)
	}
}

func (s *BasicTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitBasicType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) BasicType() (localctx IBasicTypeContext) {
	localctx = NewBasicTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MagoitoParserRULE_basicType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&122880) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRecordTypeContext is an interface to support dynamic dispatch.
type IRecordTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllRecordTypeField() []IRecordTypeFieldContext
	RecordTypeField(i int) IRecordTypeFieldContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsRecordTypeContext differentiates from other interfaces.
	IsRecordTypeContext()
}

type RecordTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecordTypeContext() *RecordTypeContext {
	var p = new(RecordTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_recordType
	return p
}

func InitEmptyRecordTypeContext(p *RecordTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_recordType
}

func (*RecordTypeContext) IsRecordTypeContext() {}

func NewRecordTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordTypeContext {
	var p = new(RecordTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_recordType

	return p
}

func (s *RecordTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordTypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserLBRACE, 0)
}

func (s *RecordTypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserRBRACE, 0)
}

func (s *RecordTypeContext) AllRecordTypeField() []IRecordTypeFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRecordTypeFieldContext); ok {
			len++
		}
	}

	tst := make([]IRecordTypeFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRecordTypeFieldContext); ok {
			tst[i] = t.(IRecordTypeFieldContext)
			i++
		}
	}

	return tst
}

func (s *RecordTypeContext) RecordTypeField(i int) IRecordTypeFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecordTypeFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecordTypeFieldContext)
}

func (s *RecordTypeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserCOMMA)
}

func (s *RecordTypeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserCOMMA, i)
}

func (s *RecordTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterRecordType(s)
	}
}

func (s *RecordTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitRecordType(s)
	}
}

func (s *RecordTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitRecordType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) RecordType() (localctx IRecordTypeContext) {
	localctx = NewRecordTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MagoitoParserRULE_recordType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(MagoitoParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MagoitoParserIDENTIFIER {
		{
			p.SetState(132)
			p.RecordTypeField()
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MagoitoParserCOMMA {
			{
				p.SetState(133)
				p.Match(MagoitoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(134)
				p.RecordTypeField()
			}

			p.SetState(139)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(142)
		p.Match(MagoitoParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRecordTypeFieldContext is an interface to support dynamic dispatch.
type IRecordTypeFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	TypeExpr() ITypeExprContext

	// IsRecordTypeFieldContext differentiates from other interfaces.
	IsRecordTypeFieldContext()
}

type RecordTypeFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecordTypeFieldContext() *RecordTypeFieldContext {
	var p = new(RecordTypeFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_recordTypeField
	return p
}

func InitEmptyRecordTypeFieldContext(p *RecordTypeFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_recordTypeField
}

func (*RecordTypeFieldContext) IsRecordTypeFieldContext() {}

func NewRecordTypeFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordTypeFieldContext {
	var p = new(RecordTypeFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_recordTypeField

	return p
}

func (s *RecordTypeFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordTypeFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MagoitoParserIDENTIFIER, 0)
}

func (s *RecordTypeFieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(MagoitoParserCOLON, 0)
}

func (s *RecordTypeFieldContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *RecordTypeFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordTypeFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordTypeFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterRecordTypeField(s)
	}
}

func (s *RecordTypeFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitRecordTypeField(s)
	}
}

func (s *RecordTypeFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitRecordTypeField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) RecordTypeField() (localctx IRecordTypeFieldContext) {
	localctx = NewRecordTypeFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MagoitoParserRULE_recordTypeField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(MagoitoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(MagoitoParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.TypeExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SeqExpr() ISeqExprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) SeqExpr() ISeqExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISeqExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISeqExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MagoitoParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.SeqExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISeqExprContext is an interface to support dynamic dispatch.
type ISeqExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ControlExpr() IControlExprContext
	SEMICOLON() antlr.TerminalNode
	SeqExpr() ISeqExprContext

	// IsSeqExprContext differentiates from other interfaces.
	IsSeqExprContext()
}

type SeqExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeqExprContext() *SeqExprContext {
	var p = new(SeqExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_seqExpr
	return p
}

func InitEmptySeqExprContext(p *SeqExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_seqExpr
}

func (*SeqExprContext) IsSeqExprContext() {}

func NewSeqExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeqExprContext {
	var p = new(SeqExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_seqExpr

	return p
}

func (s *SeqExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SeqExprContext) ControlExpr() IControlExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IControlExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IControlExprContext)
}

func (s *SeqExprContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(MagoitoParserSEMICOLON, 0)
}

func (s *SeqExprContext) SeqExpr() ISeqExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISeqExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISeqExprContext)
}

func (s *SeqExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeqExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeqExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterSeqExpr(s)
	}
}

func (s *SeqExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitSeqExpr(s)
	}
}

func (s *SeqExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitSeqExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) SeqExpr() (localctx ISeqExprContext) {
	localctx = NewSeqExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MagoitoParserRULE_seqExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.ControlExpr()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MagoitoParserSEMICOLON {
		{
			p.SetState(151)
			p.Match(MagoitoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.SeqExpr()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IControlExprContext is an interface to support dynamic dispatch.
type IControlExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDeclExpr() IVarDeclExprContext
	WhileExpr() IWhileExprContext
	IfExpr() IIfExprContext
	AssignExpr() IAssignExprContext
	OrExpr() IOrExprContext

	// IsControlExprContext differentiates from other interfaces.
	IsControlExprContext()
}

type ControlExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControlExprContext() *ControlExprContext {
	var p = new(ControlExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_controlExpr
	return p
}

func InitEmptyControlExprContext(p *ControlExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_controlExpr
}

func (*ControlExprContext) IsControlExprContext() {}

func NewControlExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlExprContext {
	var p = new(ControlExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_controlExpr

	return p
}

func (s *ControlExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ControlExprContext) VarDeclExpr() IVarDeclExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclExprContext)
}

func (s *ControlExprContext) WhileExpr() IWhileExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileExprContext)
}

func (s *ControlExprContext) IfExpr() IIfExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfExprContext)
}

func (s *ControlExprContext) AssignExpr() IAssignExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignExprContext)
}

func (s *ControlExprContext) OrExpr() IOrExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrExprContext)
}

func (s *ControlExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ControlExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterControlExpr(s)
	}
}

func (s *ControlExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitControlExpr(s)
	}
}

func (s *ControlExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitControlExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) ControlExpr() (localctx IControlExprContext) {
	localctx = NewControlExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MagoitoParserRULE_controlExpr)
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.VarDeclExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.WhileExpr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)
			p.IfExpr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(158)
			p.AssignExpr()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(159)
			p.OrExpr()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclExprContext is an interface to support dynamic dispatch.
type IVarDeclExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	Binder() IBinderContext
	COLON() antlr.TerminalNode
	TypeExpr() ITypeExprContext
	ASSIGN() antlr.TerminalNode
	ControlExpr() IControlExprContext

	// IsVarDeclExprContext differentiates from other interfaces.
	IsVarDeclExprContext()
}

type VarDeclExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclExprContext() *VarDeclExprContext {
	var p = new(VarDeclExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_varDeclExpr
	return p
}

func InitEmptyVarDeclExprContext(p *VarDeclExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_varDeclExpr
}

func (*VarDeclExprContext) IsVarDeclExprContext() {}

func NewVarDeclExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclExprContext {
	var p = new(VarDeclExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_varDeclExpr

	return p
}

func (s *VarDeclExprContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclExprContext) VAR() antlr.TerminalNode {
	return s.GetToken(MagoitoParserVAR, 0)
}

func (s *VarDeclExprContext) Binder() IBinderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinderContext)
}

func (s *VarDeclExprContext) COLON() antlr.TerminalNode {
	return s.GetToken(MagoitoParserCOLON, 0)
}

func (s *VarDeclExprContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *VarDeclExprContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserASSIGN, 0)
}

func (s *VarDeclExprContext) ControlExpr() IControlExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IControlExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IControlExprContext)
}

func (s *VarDeclExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterVarDeclExpr(s)
	}
}

func (s *VarDeclExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitVarDeclExpr(s)
	}
}

func (s *VarDeclExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitVarDeclExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) VarDeclExpr() (localctx IVarDeclExprContext) {
	localctx = NewVarDeclExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MagoitoParserRULE_varDeclExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(MagoitoParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Binder()
	}
	{
		p.SetState(164)
		p.Match(MagoitoParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.TypeExpr()
	}
	{
		p.SetState(166)
		p.Match(MagoitoParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.ControlExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileExprContext is an interface to support dynamic dispatch.
type IWhileExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	OrExpr() IOrExprContext
	DO() antlr.TerminalNode
	ControlExpr() IControlExprContext

	// IsWhileExprContext differentiates from other interfaces.
	IsWhileExprContext()
}

type WhileExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileExprContext() *WhileExprContext {
	var p = new(WhileExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_whileExpr
	return p
}

func InitEmptyWhileExprContext(p *WhileExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_whileExpr
}

func (*WhileExprContext) IsWhileExprContext() {}

func NewWhileExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileExprContext {
	var p = new(WhileExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_whileExpr

	return p
}

func (s *WhileExprContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileExprContext) WHILE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserWHILE, 0)
}

func (s *WhileExprContext) OrExpr() IOrExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrExprContext)
}

func (s *WhileExprContext) DO() antlr.TerminalNode {
	return s.GetToken(MagoitoParserDO, 0)
}

func (s *WhileExprContext) ControlExpr() IControlExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IControlExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IControlExprContext)
}

func (s *WhileExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterWhileExpr(s)
	}
}

func (s *WhileExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitWhileExpr(s)
	}
}

func (s *WhileExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitWhileExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) WhileExpr() (localctx IWhileExprContext) {
	localctx = NewWhileExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MagoitoParserRULE_whileExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(MagoitoParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.OrExpr()
	}
	{
		p.SetState(171)
		p.Match(MagoitoParserDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.ControlExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfExprContext is an interface to support dynamic dispatch.
type IIfExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	OrExpr() IOrExprContext
	THEN() antlr.TerminalNode
	AllControlExpr() []IControlExprContext
	ControlExpr(i int) IControlExprContext
	ELSE() antlr.TerminalNode

	// IsIfExprContext differentiates from other interfaces.
	IsIfExprContext()
}

type IfExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfExprContext() *IfExprContext {
	var p = new(IfExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_ifExpr
	return p
}

func InitEmptyIfExprContext(p *IfExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_ifExpr
}

func (*IfExprContext) IsIfExprContext() {}

func NewIfExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfExprContext {
	var p = new(IfExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_ifExpr

	return p
}

func (s *IfExprContext) GetParser() antlr.Parser { return s.parser }

func (s *IfExprContext) IF() antlr.TerminalNode {
	return s.GetToken(MagoitoParserIF, 0)
}

func (s *IfExprContext) OrExpr() IOrExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrExprContext)
}

func (s *IfExprContext) THEN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserTHEN, 0)
}

func (s *IfExprContext) AllControlExpr() []IControlExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IControlExprContext); ok {
			len++
		}
	}

	tst := make([]IControlExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IControlExprContext); ok {
			tst[i] = t.(IControlExprContext)
			i++
		}
	}

	return tst
}

func (s *IfExprContext) ControlExpr(i int) IControlExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IControlExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IControlExprContext)
}

func (s *IfExprContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserELSE, 0)
}

func (s *IfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterIfExpr(s)
	}
}

func (s *IfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitIfExpr(s)
	}
}

func (s *IfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitIfExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) IfExpr() (localctx IIfExprContext) {
	localctx = NewIfExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MagoitoParserRULE_ifExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(MagoitoParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.OrExpr()
	}
	{
		p.SetState(176)
		p.Match(MagoitoParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.ControlExpr()
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(178)
			p.Match(MagoitoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.ControlExpr()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignExprContext is an interface to support dynamic dispatch.
type IAssignExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	ControlExpr() IControlExprContext

	// IsAssignExprContext differentiates from other interfaces.
	IsAssignExprContext()
}

type AssignExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignExprContext() *AssignExprContext {
	var p = new(AssignExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_assignExpr
	return p
}

func InitEmptyAssignExprContext(p *AssignExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_assignExpr
}

func (*AssignExprContext) IsAssignExprContext() {}

func NewAssignExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignExprContext {
	var p = new(AssignExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_assignExpr

	return p
}

func (s *AssignExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MagoitoParserIDENTIFIER, 0)
}

func (s *AssignExprContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserASSIGN, 0)
}

func (s *AssignExprContext) ControlExpr() IControlExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IControlExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IControlExprContext)
}

func (s *AssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterAssignExpr(s)
	}
}

func (s *AssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitAssignExpr(s)
	}
}

func (s *AssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) AssignExpr() (localctx IAssignExprContext) {
	localctx = NewAssignExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MagoitoParserRULE_assignExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(MagoitoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.Match(MagoitoParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.ControlExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrExprContext is an interface to support dynamic dispatch.
type IOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAndExpr() []IAndExprContext
	AndExpr(i int) IAndExprContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsOrExprContext differentiates from other interfaces.
	IsOrExprContext()
}

type OrExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExprContext() *OrExprContext {
	var p = new(OrExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_orExpr
	return p
}

func InitEmptyOrExprContext(p *OrExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_orExpr
}

func (*OrExprContext) IsOrExprContext() {}

func NewOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExprContext {
	var p = new(OrExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_orExpr

	return p
}

func (s *OrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExprContext) AllAndExpr() []IAndExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAndExprContext); ok {
			len++
		}
	}

	tst := make([]IAndExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAndExprContext); ok {
			tst[i] = t.(IAndExprContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) AndExpr(i int) IAndExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndExprContext)
}

func (s *OrExprContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserOR)
}

func (s *OrExprContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserOR, i)
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) OrExpr() (localctx IOrExprContext) {
	localctx = NewOrExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MagoitoParserRULE_orExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.AndExpr()
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserOR {
		{
			p.SetState(187)
			p.Match(MagoitoParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.AndExpr()
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAndExprContext is an interface to support dynamic dispatch.
type IAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEqualityExpr() []IEqualityExprContext
	EqualityExpr(i int) IEqualityExprContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsAndExprContext differentiates from other interfaces.
	IsAndExprContext()
}

type AndExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExprContext() *AndExprContext {
	var p = new(AndExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_andExpr
	return p
}

func InitEmptyAndExprContext(p *AndExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_andExpr
}

func (*AndExprContext) IsAndExprContext() {}

func NewAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExprContext {
	var p = new(AndExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_andExpr

	return p
}

func (s *AndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExprContext) AllEqualityExpr() []IEqualityExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualityExprContext); ok {
			len++
		}
	}

	tst := make([]IEqualityExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualityExprContext); ok {
			tst[i] = t.(IEqualityExprContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) EqualityExpr(i int) IEqualityExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExprContext)
}

func (s *AndExprContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserAND)
}

func (s *AndExprContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserAND, i)
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) AndExpr() (localctx IAndExprContext) {
	localctx = NewAndExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MagoitoParserRULE_andExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.EqualityExpr()
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserAND {
		{
			p.SetState(195)
			p.Match(MagoitoParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.EqualityExpr()
		}

		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualityExprContext is an interface to support dynamic dispatch.
type IEqualityExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComparisonExpr() []IComparisonExprContext
	ComparisonExpr(i int) IComparisonExprContext
	AllEQ() []antlr.TerminalNode
	EQ(i int) antlr.TerminalNode
	AllNEQ() []antlr.TerminalNode
	NEQ(i int) antlr.TerminalNode

	// IsEqualityExprContext differentiates from other interfaces.
	IsEqualityExprContext()
}

type EqualityExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExprContext() *EqualityExprContext {
	var p = new(EqualityExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_equalityExpr
	return p
}

func InitEmptyEqualityExprContext(p *EqualityExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_equalityExpr
}

func (*EqualityExprContext) IsEqualityExprContext() {}

func NewEqualityExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExprContext {
	var p = new(EqualityExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_equalityExpr

	return p
}

func (s *EqualityExprContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExprContext) AllComparisonExpr() []IComparisonExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparisonExprContext); ok {
			len++
		}
	}

	tst := make([]IComparisonExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparisonExprContext); ok {
			tst[i] = t.(IComparisonExprContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExprContext) ComparisonExpr(i int) IComparisonExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonExprContext)
}

func (s *EqualityExprContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserEQ)
}

func (s *EqualityExprContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserEQ, i)
}

func (s *EqualityExprContext) AllNEQ() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserNEQ)
}

func (s *EqualityExprContext) NEQ(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserNEQ, i)
}

func (s *EqualityExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterEqualityExpr(s)
	}
}

func (s *EqualityExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitEqualityExpr(s)
	}
}

func (s *EqualityExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitEqualityExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) EqualityExpr() (localctx IEqualityExprContext) {
	localctx = NewEqualityExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MagoitoParserRULE_equalityExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.ComparisonExpr()
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserEQ || _la == MagoitoParserNEQ {
		{
			p.SetState(203)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MagoitoParserEQ || _la == MagoitoParserNEQ) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(204)
			p.ComparisonExpr()
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparisonExprContext is an interface to support dynamic dispatch.
type IComparisonExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAdditiveExpr() []IAdditiveExprContext
	AdditiveExpr(i int) IAdditiveExprContext
	AllLT() []antlr.TerminalNode
	LT(i int) antlr.TerminalNode
	AllLTE() []antlr.TerminalNode
	LTE(i int) antlr.TerminalNode
	AllGT() []antlr.TerminalNode
	GT(i int) antlr.TerminalNode
	AllGTE() []antlr.TerminalNode
	GTE(i int) antlr.TerminalNode

	// IsComparisonExprContext differentiates from other interfaces.
	IsComparisonExprContext()
}

type ComparisonExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonExprContext() *ComparisonExprContext {
	var p = new(ComparisonExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_comparisonExpr
	return p
}

func InitEmptyComparisonExprContext(p *ComparisonExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_comparisonExpr
}

func (*ComparisonExprContext) IsComparisonExprContext() {}

func NewComparisonExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonExprContext {
	var p = new(ComparisonExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_comparisonExpr

	return p
}

func (s *ComparisonExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonExprContext) AllAdditiveExpr() []IAdditiveExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveExprContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveExprContext); ok {
			tst[i] = t.(IAdditiveExprContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonExprContext) AdditiveExpr(i int) IAdditiveExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExprContext)
}

func (s *ComparisonExprContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserLT)
}

func (s *ComparisonExprContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserLT, i)
}

func (s *ComparisonExprContext) AllLTE() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserLTE)
}

func (s *ComparisonExprContext) LTE(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserLTE, i)
}

func (s *ComparisonExprContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserGT)
}

func (s *ComparisonExprContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserGT, i)
}

func (s *ComparisonExprContext) AllGTE() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserGTE)
}

func (s *ComparisonExprContext) GTE(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserGTE, i)
}

func (s *ComparisonExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterComparisonExpr(s)
	}
}

func (s *ComparisonExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitComparisonExpr(s)
	}
}

func (s *ComparisonExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitComparisonExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) ComparisonExpr() (localctx IComparisonExprContext) {
	localctx = NewComparisonExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MagoitoParserRULE_comparisonExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.AdditiveExpr()
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&103809024) != 0 {
		{
			p.SetState(211)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&103809024) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(212)
			p.AdditiveExpr()
		}

		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdditiveExprContext is an interface to support dynamic dispatch.
type IAdditiveExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMultiplicativeExpr() []IMultiplicativeExprContext
	MultiplicativeExpr(i int) IMultiplicativeExprContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsAdditiveExprContext differentiates from other interfaces.
	IsAdditiveExprContext()
}

type AdditiveExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExprContext() *AdditiveExprContext {
	var p = new(AdditiveExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_additiveExpr
	return p
}

func InitEmptyAdditiveExprContext(p *AdditiveExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_additiveExpr
}

func (*AdditiveExprContext) IsAdditiveExprContext() {}

func NewAdditiveExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExprContext {
	var p = new(AdditiveExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_additiveExpr

	return p
}

func (s *AdditiveExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExprContext) AllMultiplicativeExpr() []IMultiplicativeExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeExprContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeExprContext); ok {
			tst[i] = t.(IMultiplicativeExprContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExprContext) MultiplicativeExpr(i int) IMultiplicativeExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExprContext)
}

func (s *AdditiveExprContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserPLUS)
}

func (s *AdditiveExprContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserPLUS, i)
}

func (s *AdditiveExprContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserMINUS)
}

func (s *AdditiveExprContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserMINUS, i)
}

func (s *AdditiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitAdditiveExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) AdditiveExpr() (localctx IAdditiveExprContext) {
	localctx = NewAdditiveExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MagoitoParserRULE_additiveExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.MultiplicativeExpr()
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserPLUS || _la == MagoitoParserMINUS {
		{
			p.SetState(219)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MagoitoParserPLUS || _la == MagoitoParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(220)
			p.MultiplicativeExpr()
		}

		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplicativeExprContext is an interface to support dynamic dispatch.
type IMultiplicativeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPowerExpr() []IPowerExprContext
	PowerExpr(i int) IPowerExprContext
	AllSTAR() []antlr.TerminalNode
	STAR(i int) antlr.TerminalNode
	AllSLASH() []antlr.TerminalNode
	SLASH(i int) antlr.TerminalNode
	AllPERCENT() []antlr.TerminalNode
	PERCENT(i int) antlr.TerminalNode

	// IsMultiplicativeExprContext differentiates from other interfaces.
	IsMultiplicativeExprContext()
}

type MultiplicativeExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExprContext() *MultiplicativeExprContext {
	var p = new(MultiplicativeExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_multiplicativeExpr
	return p
}

func InitEmptyMultiplicativeExprContext(p *MultiplicativeExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_multiplicativeExpr
}

func (*MultiplicativeExprContext) IsMultiplicativeExprContext() {}

func NewMultiplicativeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExprContext {
	var p = new(MultiplicativeExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_multiplicativeExpr

	return p
}

func (s *MultiplicativeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExprContext) AllPowerExpr() []IPowerExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPowerExprContext); ok {
			len++
		}
	}

	tst := make([]IPowerExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPowerExprContext); ok {
			tst[i] = t.(IPowerExprContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExprContext) PowerExpr(i int) IPowerExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerExprContext)
}

func (s *MultiplicativeExprContext) AllSTAR() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserSTAR)
}

func (s *MultiplicativeExprContext) STAR(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserSTAR, i)
}

func (s *MultiplicativeExprContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserSLASH)
}

func (s *MultiplicativeExprContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserSLASH, i)
}

func (s *MultiplicativeExprContext) AllPERCENT() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserPERCENT)
}

func (s *MultiplicativeExprContext) PERCENT(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserPERCENT, i)
}

func (s *MultiplicativeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterMultiplicativeExpr(s)
	}
}

func (s *MultiplicativeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitMultiplicativeExpr(s)
	}
}

func (s *MultiplicativeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitMultiplicativeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) MultiplicativeExpr() (localctx IMultiplicativeExprContext) {
	localctx = NewMultiplicativeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MagoitoParserRULE_multiplicativeExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.PowerExpr()
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3758096384) != 0 {
		{
			p.SetState(227)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3758096384) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(228)
			p.PowerExpr()
		}

		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPowerExprContext is an interface to support dynamic dispatch.
type IPowerExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpr() IUnaryExprContext
	POWER() antlr.TerminalNode
	PowerExpr() IPowerExprContext

	// IsPowerExprContext differentiates from other interfaces.
	IsPowerExprContext()
}

type PowerExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPowerExprContext() *PowerExprContext {
	var p = new(PowerExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_powerExpr
	return p
}

func InitEmptyPowerExprContext(p *PowerExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_powerExpr
}

func (*PowerExprContext) IsPowerExprContext() {}

func NewPowerExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowerExprContext {
	var p = new(PowerExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_powerExpr

	return p
}

func (s *PowerExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PowerExprContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *PowerExprContext) POWER() antlr.TerminalNode {
	return s.GetToken(MagoitoParserPOWER, 0)
}

func (s *PowerExprContext) PowerExpr() IPowerExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerExprContext)
}

func (s *PowerExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowerExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterPowerExpr(s)
	}
}

func (s *PowerExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitPowerExpr(s)
	}
}

func (s *PowerExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitPowerExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) PowerExpr() (localctx IPowerExprContext) {
	localctx = NewPowerExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MagoitoParserRULE_powerExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.UnaryExpr()
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MagoitoParserPOWER {
		{
			p.SetState(235)
			p.Match(MagoitoParserPOWER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.PowerExpr()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExprContext is an interface to support dynamic dispatch.
type IUnaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpr() IUnaryExprContext
	MINUS() antlr.TerminalNode
	NOT() antlr.TerminalNode
	ProjectionExpr() IProjectionExprContext

	// IsUnaryExprContext differentiates from other interfaces.
	IsUnaryExprContext()
}

type UnaryExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExprContext() *UnaryExprContext {
	var p = new(UnaryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_unaryExpr
	return p
}

func InitEmptyUnaryExprContext(p *UnaryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_unaryExpr
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_unaryExpr

	return p
}

func (s *UnaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *UnaryExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MagoitoParserMINUS, 0)
}

func (s *UnaryExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(MagoitoParserNOT, 0)
}

func (s *UnaryExprContext) ProjectionExpr() IProjectionExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectionExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectionExprContext)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) UnaryExpr() (localctx IUnaryExprContext) {
	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MagoitoParserRULE_unaryExpr)
	var _la int

	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MagoitoParserMINUS, MagoitoParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(239)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MagoitoParserMINUS || _la == MagoitoParserNOT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(240)
			p.UnaryExpr()
		}

	case MagoitoParserPRINT, MagoitoParserTRUE, MagoitoParserFALSE, MagoitoParserUNIT_VALUE, MagoitoParserLPAREN, MagoitoParserLBRACE, MagoitoParserIDENTIFIER, MagoitoParserINT_LITERAL, MagoitoParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(241)
			p.ProjectionExpr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProjectionExprContext is an interface to support dynamic dispatch.
type IProjectionExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpr() IPrimaryExprContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsProjectionExprContext differentiates from other interfaces.
	IsProjectionExprContext()
}

type ProjectionExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionExprContext() *ProjectionExprContext {
	var p = new(ProjectionExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_projectionExpr
	return p
}

func InitEmptyProjectionExprContext(p *ProjectionExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_projectionExpr
}

func (*ProjectionExprContext) IsProjectionExprContext() {}

func NewProjectionExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionExprContext {
	var p = new(ProjectionExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_projectionExpr

	return p
}

func (s *ProjectionExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionExprContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ProjectionExprContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserDOT)
}

func (s *ProjectionExprContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserDOT, i)
}

func (s *ProjectionExprContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserIDENTIFIER)
}

func (s *ProjectionExprContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserIDENTIFIER, i)
}

func (s *ProjectionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterProjectionExpr(s)
	}
}

func (s *ProjectionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitProjectionExpr(s)
	}
}

func (s *ProjectionExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitProjectionExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) ProjectionExpr() (localctx IProjectionExprContext) {
	localctx = NewProjectionExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MagoitoParserRULE_projectionExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.PrimaryExpr()
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserDOT {
		{
			p.SetState(245)
			p.Match(MagoitoParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Match(MagoitoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_LITERAL() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	UNIT_VALUE() antlr.TerminalNode
	CallExpr() ICallExprContext
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	RecordExpr() IRecordExprContext

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_primaryExpr
	return p
}

func InitEmptyPrimaryExprContext(p *PrimaryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_primaryExpr
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(MagoitoParserINT_LITERAL, 0)
}

func (s *PrimaryExprContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(MagoitoParserSTRING_LITERAL, 0)
}

func (s *PrimaryExprContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserTRUE, 0)
}

func (s *PrimaryExprContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserFALSE, 0)
}

func (s *PrimaryExprContext) UNIT_VALUE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserUNIT_VALUE, 0)
}

func (s *PrimaryExprContext) CallExpr() ICallExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallExprContext)
}

func (s *PrimaryExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MagoitoParserIDENTIFIER, 0)
}

func (s *PrimaryExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserLPAREN, 0)
}

func (s *PrimaryExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrimaryExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserRPAREN, 0)
}

func (s *PrimaryExprContext) RecordExpr() IRecordExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecordExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecordExprContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MagoitoParserRULE_primaryExpr)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(252)
			p.Match(MagoitoParserINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.Match(MagoitoParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(254)
			p.Match(MagoitoParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(255)
			p.Match(MagoitoParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(256)
			p.Match(MagoitoParserUNIT_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(257)
			p.CallExpr()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(258)
			p.Match(MagoitoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(259)
			p.Match(MagoitoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)
			p.Expr()
		}
		{
			p.SetState(261)
			p.Match(MagoitoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(263)
			p.RecordExpr()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICallExprContext is an interface to support dynamic dispatch.
type ICallExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Callee() ICalleeContext
	LPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCallExprContext differentiates from other interfaces.
	IsCallExprContext()
}

type CallExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallExprContext() *CallExprContext {
	var p = new(CallExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_callExpr
	return p
}

func InitEmptyCallExprContext(p *CallExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_callExpr
}

func (*CallExprContext) IsCallExprContext() {}

func NewCallExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallExprContext {
	var p = new(CallExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_callExpr

	return p
}

func (s *CallExprContext) GetParser() antlr.Parser { return s.parser }

func (s *CallExprContext) Callee() ICalleeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalleeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalleeContext)
}

func (s *CallExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserLPAREN, 0)
}

func (s *CallExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CallExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CallExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserRPAREN, 0)
}

func (s *CallExprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserCOMMA)
}

func (s *CallExprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserCOMMA, i)
}

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

func (s *CallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) CallExpr() (localctx ICallExprContext) {
	localctx = NewCallExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MagoitoParserRULE_callExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Callee()
	}
	{
		p.SetState(267)
		p.Match(MagoitoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Expr()
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserCOMMA {
		{
			p.SetState(269)
			p.Match(MagoitoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.Expr()
		}

		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(276)
		p.Match(MagoitoParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICalleeContext is an interface to support dynamic dispatch.
type ICalleeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	PRINT() antlr.TerminalNode

	// IsCalleeContext differentiates from other interfaces.
	IsCalleeContext()
}

type CalleeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalleeContext() *CalleeContext {
	var p = new(CalleeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_callee
	return p
}

func InitEmptyCalleeContext(p *CalleeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_callee
}

func (*CalleeContext) IsCalleeContext() {}

func NewCalleeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalleeContext {
	var p = new(CalleeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_callee

	return p
}

func (s *CalleeContext) GetParser() antlr.Parser { return s.parser }

func (s *CalleeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MagoitoParserIDENTIFIER, 0)
}

func (s *CalleeContext) PRINT() antlr.TerminalNode {
	return s.GetToken(MagoitoParserPRINT, 0)
}

func (s *CalleeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalleeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalleeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterCallee(s)
	}
}

func (s *CalleeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitCallee(s)
	}
}

func (s *CalleeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitCallee(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) Callee() (localctx ICalleeContext) {
	localctx = NewCalleeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MagoitoParserRULE_callee)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MagoitoParserPRINT || _la == MagoitoParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRecordExprContext is an interface to support dynamic dispatch.
type IRecordExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	RecordFieldList() IRecordFieldListContext

	// IsRecordExprContext differentiates from other interfaces.
	IsRecordExprContext()
}

type RecordExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecordExprContext() *RecordExprContext {
	var p = new(RecordExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_recordExpr
	return p
}

func InitEmptyRecordExprContext(p *RecordExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_recordExpr
}

func (*RecordExprContext) IsRecordExprContext() {}

func NewRecordExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordExprContext {
	var p = new(RecordExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_recordExpr

	return p
}

func (s *RecordExprContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordExprContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserLBRACE, 0)
}

func (s *RecordExprContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserRBRACE, 0)
}

func (s *RecordExprContext) RecordFieldList() IRecordFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecordFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecordFieldListContext)
}

func (s *RecordExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterRecordExpr(s)
	}
}

func (s *RecordExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitRecordExpr(s)
	}
}

func (s *RecordExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitRecordExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) RecordExpr() (localctx IRecordExprContext) {
	localctx = NewRecordExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, MagoitoParserRULE_recordExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Match(MagoitoParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MagoitoParserIDENTIFIER {
		{
			p.SetState(281)
			p.RecordFieldList()
		}

	}
	{
		p.SetState(284)
		p.Match(MagoitoParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRecordFieldListContext is an interface to support dynamic dispatch.
type IRecordFieldListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRecordField() []IRecordFieldContext
	RecordField(i int) IRecordFieldContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsRecordFieldListContext differentiates from other interfaces.
	IsRecordFieldListContext()
}

type RecordFieldListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecordFieldListContext() *RecordFieldListContext {
	var p = new(RecordFieldListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_recordFieldList
	return p
}

func InitEmptyRecordFieldListContext(p *RecordFieldListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_recordFieldList
}

func (*RecordFieldListContext) IsRecordFieldListContext() {}

func NewRecordFieldListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordFieldListContext {
	var p = new(RecordFieldListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_recordFieldList

	return p
}

func (s *RecordFieldListContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordFieldListContext) AllRecordField() []IRecordFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRecordFieldContext); ok {
			len++
		}
	}

	tst := make([]IRecordFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRecordFieldContext); ok {
			tst[i] = t.(IRecordFieldContext)
			i++
		}
	}

	return tst
}

func (s *RecordFieldListContext) RecordField(i int) IRecordFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecordFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecordFieldContext)
}

func (s *RecordFieldListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserCOMMA)
}

func (s *RecordFieldListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserCOMMA, i)
}

func (s *RecordFieldListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordFieldListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordFieldListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterRecordFieldList(s)
	}
}

func (s *RecordFieldListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitRecordFieldList(s)
	}
}

func (s *RecordFieldListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitRecordFieldList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) RecordFieldList() (localctx IRecordFieldListContext) {
	localctx = NewRecordFieldListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MagoitoParserRULE_recordFieldList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.RecordField()
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserCOMMA {
		{
			p.SetState(287)
			p.Match(MagoitoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)
			p.RecordField()
		}

		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRecordFieldContext is an interface to support dynamic dispatch.
type IRecordFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext

	// IsRecordFieldContext differentiates from other interfaces.
	IsRecordFieldContext()
}

type RecordFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecordFieldContext() *RecordFieldContext {
	var p = new(RecordFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_recordField
	return p
}

func InitEmptyRecordFieldContext(p *RecordFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MagoitoParserRULE_recordField
}

func (*RecordFieldContext) IsRecordFieldContext() {}

func NewRecordFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordFieldContext {
	var p = new(RecordFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagoitoParserRULE_recordField

	return p
}

func (s *RecordFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MagoitoParserIDENTIFIER, 0)
}

func (s *RecordFieldContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserASSIGN, 0)
}

func (s *RecordFieldContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RecordFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.EnterRecordField(s)
	}
}

func (s *RecordFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagoitoParserListener); ok {
		listenerT.ExitRecordField(s)
	}
}

func (s *RecordFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitRecordField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) RecordField() (localctx IRecordFieldContext) {
	localctx = NewRecordFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MagoitoParserRULE_recordField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Match(MagoitoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
		p.Match(MagoitoParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
		p.Expr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
