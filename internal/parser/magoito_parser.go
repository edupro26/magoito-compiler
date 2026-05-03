// Code generated from parser/MagoitoParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
		4, 1, 47, 307, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 4,
		0, 64, 8, 0, 11, 0, 12, 0, 65, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 83, 8, 1, 10, 1,
		12, 1, 86, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 94, 8, 1, 1,
		2, 1, 2, 3, 2, 98, 8, 2, 1, 3, 1, 3, 1, 3, 3, 3, 103, 8, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 109, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4,
		117, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 125, 8, 5, 10, 5,
		12, 5, 128, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 136, 8, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 142, 8, 7, 10, 7, 12, 7, 145, 9, 7, 3, 7,
		147, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 3, 10, 160, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 167,
		8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 187,
		8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 5, 16, 196, 8,
		16, 10, 16, 12, 16, 199, 9, 16, 1, 17, 1, 17, 1, 17, 5, 17, 204, 8, 17,
		10, 17, 12, 17, 207, 9, 17, 1, 18, 1, 18, 1, 18, 5, 18, 212, 8, 18, 10,
		18, 12, 18, 215, 9, 18, 1, 19, 1, 19, 1, 19, 5, 19, 220, 8, 19, 10, 19,
		12, 19, 223, 9, 19, 1, 20, 1, 20, 1, 20, 5, 20, 228, 8, 20, 10, 20, 12,
		20, 231, 9, 20, 1, 21, 1, 21, 1, 21, 5, 21, 236, 8, 21, 10, 21, 12, 21,
		239, 9, 21, 1, 22, 1, 22, 1, 22, 3, 22, 244, 8, 22, 1, 23, 1, 23, 1, 23,
		3, 23, 249, 8, 23, 1, 24, 1, 24, 1, 24, 5, 24, 254, 8, 24, 10, 24, 12,
		24, 257, 9, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 271, 8, 25, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 5, 26, 278, 8, 26, 10, 26, 12, 26, 281, 9, 26, 1, 26, 1, 26,
		1, 27, 1, 27, 3, 27, 287, 8, 27, 1, 28, 1, 28, 3, 28, 291, 8, 28, 1, 28,
		1, 28, 1, 29, 1, 29, 1, 29, 5, 29, 298, 8, 29, 10, 29, 12, 29, 301, 9,
		29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 0, 0, 31, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		50, 52, 54, 56, 58, 60, 0, 5, 1, 0, 18, 19, 2, 0, 20, 21, 25, 26, 1, 0,
		27, 28, 1, 0, 29, 31, 2, 0, 28, 28, 33, 33, 316, 0, 63, 1, 0, 0, 0, 2,
		93, 1, 0, 0, 0, 4, 97, 1, 0, 0, 0, 6, 108, 1, 0, 0, 0, 8, 116, 1, 0, 0,
		0, 10, 118, 1, 0, 0, 0, 12, 135, 1, 0, 0, 0, 14, 137, 1, 0, 0, 0, 16, 150,
		1, 0, 0, 0, 18, 154, 1, 0, 0, 0, 20, 156, 1, 0, 0, 0, 22, 166, 1, 0, 0,
		0, 24, 168, 1, 0, 0, 0, 26, 175, 1, 0, 0, 0, 28, 180, 1, 0, 0, 0, 30, 188,
		1, 0, 0, 0, 32, 192, 1, 0, 0, 0, 34, 200, 1, 0, 0, 0, 36, 208, 1, 0, 0,
		0, 38, 216, 1, 0, 0, 0, 40, 224, 1, 0, 0, 0, 42, 232, 1, 0, 0, 0, 44, 240,
		1, 0, 0, 0, 46, 248, 1, 0, 0, 0, 48, 250, 1, 0, 0, 0, 50, 270, 1, 0, 0,
		0, 52, 272, 1, 0, 0, 0, 54, 286, 1, 0, 0, 0, 56, 288, 1, 0, 0, 0, 58, 294,
		1, 0, 0, 0, 60, 302, 1, 0, 0, 0, 62, 64, 3, 2, 1, 0, 63, 62, 1, 0, 0, 0,
		64, 65, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 1,
		0, 0, 0, 67, 68, 5, 0, 0, 1, 68, 1, 1, 0, 0, 0, 69, 70, 5, 1, 0, 0, 70,
		71, 3, 4, 2, 0, 71, 72, 5, 36, 0, 0, 72, 73, 3, 6, 3, 0, 73, 74, 5, 24,
		0, 0, 74, 75, 3, 18, 9, 0, 75, 94, 1, 0, 0, 0, 76, 77, 5, 2, 0, 0, 77,
		78, 5, 43, 0, 0, 78, 79, 5, 38, 0, 0, 79, 84, 3, 4, 2, 0, 80, 81, 5, 37,
		0, 0, 81, 83, 3, 4, 2, 0, 82, 80, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82,
		1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0,
		87, 88, 5, 39, 0, 0, 88, 89, 5, 36, 0, 0, 89, 90, 3, 6, 3, 0, 90, 91, 5,
		24, 0, 0, 91, 92, 3, 18, 9, 0, 92, 94, 1, 0, 0, 0, 93, 69, 1, 0, 0, 0,
		93, 76, 1, 0, 0, 0, 94, 3, 1, 0, 0, 0, 95, 98, 5, 43, 0, 0, 96, 98, 5,
		42, 0, 0, 97, 95, 1, 0, 0, 0, 97, 96, 1, 0, 0, 0, 98, 5, 1, 0, 0, 0, 99,
		102, 3, 8, 4, 0, 100, 101, 5, 17, 0, 0, 101, 103, 3, 6, 3, 0, 102, 100,
		1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 109, 1, 0, 0, 0, 104, 105, 3, 10,
		5, 0, 105, 106, 5, 17, 0, 0, 106, 107, 3, 6, 3, 0, 107, 109, 1, 0, 0, 0,
		108, 99, 1, 0, 0, 0, 108, 104, 1, 0, 0, 0, 109, 7, 1, 0, 0, 0, 110, 117,
		3, 12, 6, 0, 111, 117, 3, 14, 7, 0, 112, 113, 5, 38, 0, 0, 113, 114, 3,
		6, 3, 0, 114, 115, 5, 39, 0, 0, 115, 117, 1, 0, 0, 0, 116, 110, 1, 0, 0,
		0, 116, 111, 1, 0, 0, 0, 116, 112, 1, 0, 0, 0, 117, 9, 1, 0, 0, 0, 118,
		119, 5, 38, 0, 0, 119, 120, 3, 6, 3, 0, 120, 121, 5, 37, 0, 0, 121, 126,
		3, 6, 3, 0, 122, 123, 5, 37, 0, 0, 123, 125, 3, 6, 3, 0, 124, 122, 1, 0,
		0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0,
		127, 129, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 130, 5, 39, 0, 0, 130,
		11, 1, 0, 0, 0, 131, 136, 5, 13, 0, 0, 132, 136, 5, 14, 0, 0, 133, 136,
		5, 15, 0, 0, 134, 136, 5, 16, 0, 0, 135, 131, 1, 0, 0, 0, 135, 132, 1,
		0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 134, 1, 0, 0, 0, 136, 13, 1, 0, 0,
		0, 137, 146, 5, 40, 0, 0, 138, 143, 3, 16, 8, 0, 139, 140, 5, 37, 0, 0,
		140, 142, 3, 16, 8, 0, 141, 139, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143,
		141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 147, 1, 0, 0, 0, 145, 143,
		1, 0, 0, 0, 146, 138, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 1, 0,
		0, 0, 148, 149, 5, 41, 0, 0, 149, 15, 1, 0, 0, 0, 150, 151, 5, 43, 0, 0,
		151, 152, 5, 36, 0, 0, 152, 153, 3, 6, 3, 0, 153, 17, 1, 0, 0, 0, 154,
		155, 3, 20, 10, 0, 155, 19, 1, 0, 0, 0, 156, 159, 3, 22, 11, 0, 157, 158,
		5, 35, 0, 0, 158, 160, 3, 20, 10, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1,
		0, 0, 0, 160, 21, 1, 0, 0, 0, 161, 167, 3, 24, 12, 0, 162, 167, 3, 26,
		13, 0, 163, 167, 3, 28, 14, 0, 164, 167, 3, 30, 15, 0, 165, 167, 3, 32,
		16, 0, 166, 161, 1, 0, 0, 0, 166, 162, 1, 0, 0, 0, 166, 163, 1, 0, 0, 0,
		166, 164, 1, 0, 0, 0, 166, 165, 1, 0, 0, 0, 167, 23, 1, 0, 0, 0, 168, 169,
		5, 3, 0, 0, 169, 170, 3, 4, 2, 0, 170, 171, 5, 36, 0, 0, 171, 172, 3, 6,
		3, 0, 172, 173, 5, 24, 0, 0, 173, 174, 3, 22, 11, 0, 174, 25, 1, 0, 0,
		0, 175, 176, 5, 7, 0, 0, 176, 177, 3, 32, 16, 0, 177, 178, 5, 8, 0, 0,
		178, 179, 3, 22, 11, 0, 179, 27, 1, 0, 0, 0, 180, 181, 5, 4, 0, 0, 181,
		182, 3, 32, 16, 0, 182, 183, 5, 5, 0, 0, 183, 186, 3, 22, 11, 0, 184, 185,
		5, 6, 0, 0, 185, 187, 3, 22, 11, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1,
		0, 0, 0, 187, 29, 1, 0, 0, 0, 188, 189, 5, 43, 0, 0, 189, 190, 5, 24, 0,
		0, 190, 191, 3, 22, 11, 0, 191, 31, 1, 0, 0, 0, 192, 197, 3, 34, 17, 0,
		193, 194, 5, 22, 0, 0, 194, 196, 3, 34, 17, 0, 195, 193, 1, 0, 0, 0, 196,
		199, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 33, 1,
		0, 0, 0, 199, 197, 1, 0, 0, 0, 200, 205, 3, 36, 18, 0, 201, 202, 5, 23,
		0, 0, 202, 204, 3, 36, 18, 0, 203, 201, 1, 0, 0, 0, 204, 207, 1, 0, 0,
		0, 205, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 35, 1, 0, 0, 0, 207,
		205, 1, 0, 0, 0, 208, 213, 3, 38, 19, 0, 209, 210, 7, 0, 0, 0, 210, 212,
		3, 38, 19, 0, 211, 209, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1,
		0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 37, 1, 0, 0, 0, 215, 213, 1, 0, 0,
		0, 216, 221, 3, 40, 20, 0, 217, 218, 7, 1, 0, 0, 218, 220, 3, 40, 20, 0,
		219, 217, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221,
		222, 1, 0, 0, 0, 222, 39, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224, 229, 3,
		42, 21, 0, 225, 226, 7, 2, 0, 0, 226, 228, 3, 42, 21, 0, 227, 225, 1, 0,
		0, 0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0,
		230, 41, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232, 237, 3, 44, 22, 0, 233,
		234, 7, 3, 0, 0, 234, 236, 3, 44, 22, 0, 235, 233, 1, 0, 0, 0, 236, 239,
		1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 43, 1, 0,
		0, 0, 239, 237, 1, 0, 0, 0, 240, 243, 3, 46, 23, 0, 241, 242, 5, 32, 0,
		0, 242, 244, 3, 44, 22, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0,
		244, 45, 1, 0, 0, 0, 245, 246, 7, 4, 0, 0, 246, 249, 3, 46, 23, 0, 247,
		249, 3, 48, 24, 0, 248, 245, 1, 0, 0, 0, 248, 247, 1, 0, 0, 0, 249, 47,
		1, 0, 0, 0, 250, 255, 3, 50, 25, 0, 251, 252, 5, 34, 0, 0, 252, 254, 5,
		43, 0, 0, 253, 251, 1, 0, 0, 0, 254, 257, 1, 0, 0, 0, 255, 253, 1, 0, 0,
		0, 255, 256, 1, 0, 0, 0, 256, 49, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 258,
		271, 5, 44, 0, 0, 259, 271, 5, 45, 0, 0, 260, 271, 5, 10, 0, 0, 261, 271,
		5, 11, 0, 0, 262, 271, 5, 12, 0, 0, 263, 271, 3, 52, 26, 0, 264, 271, 5,
		43, 0, 0, 265, 266, 5, 38, 0, 0, 266, 267, 3, 18, 9, 0, 267, 268, 5, 39,
		0, 0, 268, 271, 1, 0, 0, 0, 269, 271, 3, 56, 28, 0, 270, 258, 1, 0, 0,
		0, 270, 259, 1, 0, 0, 0, 270, 260, 1, 0, 0, 0, 270, 261, 1, 0, 0, 0, 270,
		262, 1, 0, 0, 0, 270, 263, 1, 0, 0, 0, 270, 264, 1, 0, 0, 0, 270, 265,
		1, 0, 0, 0, 270, 269, 1, 0, 0, 0, 271, 51, 1, 0, 0, 0, 272, 273, 3, 54,
		27, 0, 273, 274, 5, 38, 0, 0, 274, 279, 3, 18, 9, 0, 275, 276, 5, 37, 0,
		0, 276, 278, 3, 18, 9, 0, 277, 275, 1, 0, 0, 0, 278, 281, 1, 0, 0, 0, 279,
		277, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 282, 1, 0, 0, 0, 281, 279,
		1, 0, 0, 0, 282, 283, 5, 39, 0, 0, 283, 53, 1, 0, 0, 0, 284, 287, 5, 43,
		0, 0, 285, 287, 5, 9, 0, 0, 286, 284, 1, 0, 0, 0, 286, 285, 1, 0, 0, 0,
		287, 55, 1, 0, 0, 0, 288, 290, 5, 40, 0, 0, 289, 291, 3, 58, 29, 0, 290,
		289, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 293,
		5, 41, 0, 0, 293, 57, 1, 0, 0, 0, 294, 299, 3, 60, 30, 0, 295, 296, 5,
		37, 0, 0, 296, 298, 3, 60, 30, 0, 297, 295, 1, 0, 0, 0, 298, 301, 1, 0,
		0, 0, 299, 297, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 59, 1, 0, 0, 0,
		301, 299, 1, 0, 0, 0, 302, 303, 5, 43, 0, 0, 303, 304, 5, 24, 0, 0, 304,
		305, 3, 18, 9, 0, 305, 61, 1, 0, 0, 0, 28, 65, 84, 93, 97, 102, 108, 116,
		126, 135, 143, 146, 159, 166, 186, 197, 205, 213, 221, 229, 237, 243, 248,
		255, 270, 279, 286, 290, 299,
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

func (s *DeclarationContext) CopyAll(ctx *DeclarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConstDeclarationContext struct {
	DeclarationContext
}

func NewConstDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstDeclarationContext {
	var p = new(ConstDeclarationContext)

	InitEmptyDeclarationContext(&p.DeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationContext))

	return p
}

func (s *ConstDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDeclarationContext) CONST() antlr.TerminalNode {
	return s.GetToken(MagoitoParserCONST, 0)
}

func (s *ConstDeclarationContext) Binder() IBinderContext {
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

func (s *ConstDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(MagoitoParserCOLON, 0)
}

func (s *ConstDeclarationContext) TypeExpr() ITypeExprContext {
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

func (s *ConstDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserASSIGN, 0)
}

func (s *ConstDeclarationContext) Expr() IExprContext {
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

func (s *ConstDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitConstDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunDeclarationContext struct {
	DeclarationContext
}

func NewFunDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunDeclarationContext {
	var p = new(FunDeclarationContext)

	InitEmptyDeclarationContext(&p.DeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclarationContext))

	return p
}

func (s *FunDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunDeclarationContext) FUN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserFUN, 0)
}

func (s *FunDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MagoitoParserIDENTIFIER, 0)
}

func (s *FunDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserLPAREN, 0)
}

func (s *FunDeclarationContext) AllBinder() []IBinderContext {
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

func (s *FunDeclarationContext) Binder(i int) IBinderContext {
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

func (s *FunDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserRPAREN, 0)
}

func (s *FunDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(MagoitoParserCOLON, 0)
}

func (s *FunDeclarationContext) TypeExpr() ITypeExprContext {
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

func (s *FunDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserASSIGN, 0)
}

func (s *FunDeclarationContext) Expr() IExprContext {
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

func (s *FunDeclarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MagoitoParserCOMMA)
}

func (s *FunDeclarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MagoitoParserCOMMA, i)
}

func (s *FunDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitFunDeclaration(s)

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
		localctx = NewConstDeclarationContext(p, localctx)
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
		localctx = NewFunDeclarationContext(p, localctx)
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

func (s *BinderContext) CopyAll(ctx *BinderContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BinderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdBinderContext struct {
	BinderContext
}

func NewIdBinderContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdBinderContext {
	var p = new(IdBinderContext)

	InitEmptyBinderContext(&p.BinderContext)
	p.parser = parser
	p.CopyAll(ctx.(*BinderContext))

	return p
}

func (s *IdBinderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdBinderContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MagoitoParserIDENTIFIER, 0)
}

func (s *IdBinderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitIdBinder(s)

	default:
		return t.VisitChildren(s)
	}
}

type WildcardBinderContext struct {
	BinderContext
}

func NewWildcardBinderContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WildcardBinderContext {
	var p = new(WildcardBinderContext)

	InitEmptyBinderContext(&p.BinderContext)
	p.parser = parser
	p.CopyAll(ctx.(*BinderContext))

	return p
}

func (s *WildcardBinderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WildcardBinderContext) WILDCARD() antlr.TerminalNode {
	return s.GetToken(MagoitoParserWILDCARD, 0)
}

func (s *WildcardBinderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitWildcardBinder(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) Binder() (localctx IBinderContext) {
	localctx = NewBinderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MagoitoParserRULE_binder)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MagoitoParserIDENTIFIER:
		localctx = NewIdBinderContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.Match(MagoitoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MagoitoParserWILDCARD:
		localctx = NewWildcardBinderContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.Match(MagoitoParserWILDCARD)
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

	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArrowTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.NonTupleType()
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MagoitoParserARROW {
			{
				p.SetState(100)
				p.Match(MagoitoParserARROW)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(101)
				p.TypeExpr()
			}

		}

	case 2:
		localctx = NewTupleArrowTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.TupleType()
		}
		{
			p.SetState(105)
			p.Match(MagoitoParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
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

func (s *NonTupleTypeContext) CopyAll(ctx *NonTupleTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *NonTupleTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonTupleTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BasicNonTupleTypeContext struct {
	NonTupleTypeContext
}

func NewBasicNonTupleTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BasicNonTupleTypeContext {
	var p = new(BasicNonTupleTypeContext)

	InitEmptyNonTupleTypeContext(&p.NonTupleTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*NonTupleTypeContext))

	return p
}

func (s *BasicNonTupleTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicNonTupleTypeContext) BasicType() IBasicTypeContext {
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

func (s *BasicNonTupleTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitBasicNonTupleType(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenNonTupleTypeContext struct {
	NonTupleTypeContext
}

func NewParenNonTupleTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenNonTupleTypeContext {
	var p = new(ParenNonTupleTypeContext)

	InitEmptyNonTupleTypeContext(&p.NonTupleTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*NonTupleTypeContext))

	return p
}

func (s *ParenNonTupleTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenNonTupleTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserLPAREN, 0)
}

func (s *ParenNonTupleTypeContext) TypeExpr() ITypeExprContext {
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

func (s *ParenNonTupleTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserRPAREN, 0)
}

func (s *ParenNonTupleTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitParenNonTupleType(s)

	default:
		return t.VisitChildren(s)
	}
}

type RecordNonTupleTypeContext struct {
	NonTupleTypeContext
}

func NewRecordNonTupleTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecordNonTupleTypeContext {
	var p = new(RecordNonTupleTypeContext)

	InitEmptyNonTupleTypeContext(&p.NonTupleTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*NonTupleTypeContext))

	return p
}

func (s *RecordNonTupleTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordNonTupleTypeContext) RecordType() IRecordTypeContext {
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

func (s *RecordNonTupleTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitRecordNonTupleType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) NonTupleType() (localctx INonTupleTypeContext) {
	localctx = NewNonTupleTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MagoitoParserRULE_nonTupleType)
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MagoitoParserINT_TYPE, MagoitoParserBOOL_TYPE, MagoitoParserSTRING_TYPE, MagoitoParserUNIT_TYPE:
		localctx = NewBasicNonTupleTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.BasicType()
		}

	case MagoitoParserLBRACE:
		localctx = NewRecordNonTupleTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.RecordType()
		}

	case MagoitoParserLPAREN:
		localctx = NewParenNonTupleTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.Match(MagoitoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.TypeExpr()
		}
		{
			p.SetState(114)
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
		p.SetState(118)
		p.Match(MagoitoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.TypeExpr()
	}
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

	for _la == MagoitoParserCOMMA {
		{
			p.SetState(122)
			p.Match(MagoitoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.TypeExpr()
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(129)
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

func (s *BasicTypeContext) CopyAll(ctx *BasicTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BasicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolBasicTypeContext struct {
	BasicTypeContext
}

func NewBoolBasicTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolBasicTypeContext {
	var p = new(BoolBasicTypeContext)

	InitEmptyBasicTypeContext(&p.BasicTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*BasicTypeContext))

	return p
}

func (s *BoolBasicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolBasicTypeContext) BOOL_TYPE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserBOOL_TYPE, 0)
}

func (s *BoolBasicTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitBoolBasicType(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnitBasicTypeContext struct {
	BasicTypeContext
}

func NewUnitBasicTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnitBasicTypeContext {
	var p = new(UnitBasicTypeContext)

	InitEmptyBasicTypeContext(&p.BasicTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*BasicTypeContext))

	return p
}

func (s *UnitBasicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitBasicTypeContext) UNIT_TYPE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserUNIT_TYPE, 0)
}

func (s *UnitBasicTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitUnitBasicType(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringBasicTypeContext struct {
	BasicTypeContext
}

func NewStringBasicTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringBasicTypeContext {
	var p = new(StringBasicTypeContext)

	InitEmptyBasicTypeContext(&p.BasicTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*BasicTypeContext))

	return p
}

func (s *StringBasicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringBasicTypeContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserSTRING_TYPE, 0)
}

func (s *StringBasicTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitStringBasicType(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntBasicTypeContext struct {
	BasicTypeContext
}

func NewIntBasicTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntBasicTypeContext {
	var p = new(IntBasicTypeContext)

	InitEmptyBasicTypeContext(&p.BasicTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*BasicTypeContext))

	return p
}

func (s *IntBasicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntBasicTypeContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserINT_TYPE, 0)
}

func (s *IntBasicTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitIntBasicType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) BasicType() (localctx IBasicTypeContext) {
	localctx = NewBasicTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MagoitoParserRULE_basicType)
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MagoitoParserINT_TYPE:
		localctx = NewIntBasicTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.Match(MagoitoParserINT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MagoitoParserBOOL_TYPE:
		localctx = NewBoolBasicTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.Match(MagoitoParserBOOL_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MagoitoParserSTRING_TYPE:
		localctx = NewStringBasicTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)
			p.Match(MagoitoParserSTRING_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MagoitoParserUNIT_TYPE:
		localctx = NewUnitBasicTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(134)
			p.Match(MagoitoParserUNIT_TYPE)
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
		p.SetState(137)
		p.Match(MagoitoParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MagoitoParserIDENTIFIER {
		{
			p.SetState(138)
			p.RecordTypeField()
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MagoitoParserCOMMA {
			{
				p.SetState(139)
				p.Match(MagoitoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(140)
				p.RecordTypeField()
			}

			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(148)
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
		p.SetState(150)
		p.Match(MagoitoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Match(MagoitoParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
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
		p.SetState(154)
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
		p.SetState(156)
		p.ControlExpr()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MagoitoParserSEMICOLON {
		{
			p.SetState(157)
			p.Match(MagoitoParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
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

func (s *ControlExprContext) CopyAll(ctx *ControlExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ControlExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VarDeclControlContext struct {
	ControlExprContext
}

func NewVarDeclControlContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDeclControlContext {
	var p = new(VarDeclControlContext)

	InitEmptyControlExprContext(&p.ControlExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ControlExprContext))

	return p
}

func (s *VarDeclControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclControlContext) VarDeclExpr() IVarDeclExprContext {
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

func (s *VarDeclControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitVarDeclControl(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignControlContext struct {
	ControlExprContext
}

func NewAssignControlContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignControlContext {
	var p = new(AssignControlContext)

	InitEmptyControlExprContext(&p.ControlExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ControlExprContext))

	return p
}

func (s *AssignControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignControlContext) AssignExpr() IAssignExprContext {
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

func (s *AssignControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitAssignControl(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrControlContext struct {
	ControlExprContext
}

func NewOrControlContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrControlContext {
	var p = new(OrControlContext)

	InitEmptyControlExprContext(&p.ControlExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ControlExprContext))

	return p
}

func (s *OrControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrControlContext) OrExpr() IOrExprContext {
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

func (s *OrControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitOrControl(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfControlContext struct {
	ControlExprContext
}

func NewIfControlContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfControlContext {
	var p = new(IfControlContext)

	InitEmptyControlExprContext(&p.ControlExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ControlExprContext))

	return p
}

func (s *IfControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfControlContext) IfExpr() IIfExprContext {
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

func (s *IfControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitIfControl(s)

	default:
		return t.VisitChildren(s)
	}
}

type WhileControlContext struct {
	ControlExprContext
}

func NewWhileControlContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileControlContext {
	var p = new(WhileControlContext)

	InitEmptyControlExprContext(&p.ControlExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ControlExprContext))

	return p
}

func (s *WhileControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileControlContext) WhileExpr() IWhileExprContext {
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

func (s *WhileControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitWhileControl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) ControlExpr() (localctx IControlExprContext) {
	localctx = NewControlExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MagoitoParserRULE_controlExpr)
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVarDeclControlContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.VarDeclExpr()
		}

	case 2:
		localctx = NewWhileControlContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.WhileExpr()
		}

	case 3:
		localctx = NewIfControlContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(163)
			p.IfExpr()
		}

	case 4:
		localctx = NewAssignControlContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(164)
			p.AssignExpr()
		}

	case 5:
		localctx = NewOrControlContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(165)
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
		p.SetState(168)
		p.Match(MagoitoParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Binder()
	}
	{
		p.SetState(170)
		p.Match(MagoitoParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.TypeExpr()
	}
	{
		p.SetState(172)
		p.Match(MagoitoParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
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
		p.SetState(175)
		p.Match(MagoitoParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.OrExpr()
	}
	{
		p.SetState(177)
		p.Match(MagoitoParserDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
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
		p.SetState(180)
		p.Match(MagoitoParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.OrExpr()
	}
	{
		p.SetState(182)
		p.Match(MagoitoParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.ControlExpr()
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(184)
			p.Match(MagoitoParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
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
		p.SetState(188)
		p.Match(MagoitoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.Match(MagoitoParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
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
		p.SetState(192)
		p.AndExpr()
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserOR {
		{
			p.SetState(193)
			p.Match(MagoitoParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.AndExpr()
		}

		p.SetState(199)
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
		p.SetState(200)
		p.EqualityExpr()
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserAND {
		{
			p.SetState(201)
			p.Match(MagoitoParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.EqualityExpr()
		}

		p.SetState(207)
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
		p.SetState(208)
		p.ComparisonExpr()
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserEQ || _la == MagoitoParserNEQ {
		{
			p.SetState(209)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MagoitoParserEQ || _la == MagoitoParserNEQ) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(210)
			p.ComparisonExpr()
		}

		p.SetState(215)
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
		p.SetState(216)
		p.AdditiveExpr()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&103809024) != 0 {
		{
			p.SetState(217)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&103809024) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(218)
			p.AdditiveExpr()
		}

		p.SetState(223)
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
		p.SetState(224)
		p.MultiplicativeExpr()
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserPLUS || _la == MagoitoParserMINUS {
		{
			p.SetState(225)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MagoitoParserPLUS || _la == MagoitoParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(226)
			p.MultiplicativeExpr()
		}

		p.SetState(231)
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
		p.SetState(232)
		p.PowerExpr()
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3758096384) != 0 {
		{
			p.SetState(233)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3758096384) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(234)
			p.PowerExpr()
		}

		p.SetState(239)
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
		p.SetState(240)
		p.UnaryExpr()
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MagoitoParserPOWER {
		{
			p.SetState(241)
			p.Match(MagoitoParserPOWER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
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

	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MagoitoParserMINUS, MagoitoParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(245)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MagoitoParserMINUS || _la == MagoitoParserNOT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(246)
			p.UnaryExpr()
		}

	case MagoitoParserPRINT, MagoitoParserTRUE, MagoitoParserFALSE, MagoitoParserUNIT_VALUE, MagoitoParserLPAREN, MagoitoParserLBRACE, MagoitoParserIDENTIFIER, MagoitoParserINT_LITERAL, MagoitoParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(247)
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
		p.SetState(250)
		p.PrimaryExpr()
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserDOT {
		{
			p.SetState(251)
			p.Match(MagoitoParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(252)
			p.Match(MagoitoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(257)
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

func (s *PrimaryExprContext) CopyAll(ctx *PrimaryExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringLiteralPrimaryContext struct {
	PrimaryExprContext
}

func NewStringLiteralPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralPrimaryContext {
	var p = new(StringLiteralPrimaryContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *StringLiteralPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralPrimaryContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(MagoitoParserSTRING_LITERAL, 0)
}

func (s *StringLiteralPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitStringLiteralPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierPrimaryContext struct {
	PrimaryExprContext
}

func NewIdentifierPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierPrimaryContext {
	var p = new(IdentifierPrimaryContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *IdentifierPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierPrimaryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MagoitoParserIDENTIFIER, 0)
}

func (s *IdentifierPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitIdentifierPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenPrimaryContext struct {
	PrimaryExprContext
}

func NewParenPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenPrimaryContext {
	var p = new(ParenPrimaryContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *ParenPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenPrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserLPAREN, 0)
}

func (s *ParenPrimaryContext) Expr() IExprContext {
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

func (s *ParenPrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MagoitoParserRPAREN, 0)
}

func (s *ParenPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitParenPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntLiteralPrimaryContext struct {
	PrimaryExprContext
}

func NewIntLiteralPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralPrimaryContext {
	var p = new(IntLiteralPrimaryContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *IntLiteralPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralPrimaryContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(MagoitoParserINT_LITERAL, 0)
}

func (s *IntLiteralPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitIntLiteralPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type FalsePrimaryContext struct {
	PrimaryExprContext
}

func NewFalsePrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalsePrimaryContext {
	var p = new(FalsePrimaryContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *FalsePrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalsePrimaryContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserFALSE, 0)
}

func (s *FalsePrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitFalsePrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type TruePrimaryContext struct {
	PrimaryExprContext
}

func NewTruePrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TruePrimaryContext {
	var p = new(TruePrimaryContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *TruePrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TruePrimaryContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserTRUE, 0)
}

func (s *TruePrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitTruePrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallPrimaryContext struct {
	PrimaryExprContext
}

func NewCallPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallPrimaryContext {
	var p = new(CallPrimaryContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *CallPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallPrimaryContext) CallExpr() ICallExprContext {
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

func (s *CallPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitCallPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnitPrimaryContext struct {
	PrimaryExprContext
}

func NewUnitPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnitPrimaryContext {
	var p = new(UnitPrimaryContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *UnitPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitPrimaryContext) UNIT_VALUE() antlr.TerminalNode {
	return s.GetToken(MagoitoParserUNIT_VALUE, 0)
}

func (s *UnitPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitUnitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type RecordPrimaryContext struct {
	PrimaryExprContext
}

func NewRecordPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecordPrimaryContext {
	var p = new(RecordPrimaryContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *RecordPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordPrimaryContext) RecordExpr() IRecordExprContext {
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

func (s *RecordPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitRecordPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MagoitoParserRULE_primaryExpr)
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIntLiteralPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)
			p.Match(MagoitoParserINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewStringLiteralPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(259)
			p.Match(MagoitoParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewTruePrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(260)
			p.Match(MagoitoParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewFalsePrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(261)
			p.Match(MagoitoParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewUnitPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(262)
			p.Match(MagoitoParserUNIT_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewCallPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(263)
			p.CallExpr()
		}

	case 7:
		localctx = NewIdentifierPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(264)
			p.Match(MagoitoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewParenPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(265)
			p.Match(MagoitoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)
			p.Expr()
		}
		{
			p.SetState(267)
			p.Match(MagoitoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewRecordPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(269)
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
		p.SetState(272)
		p.Callee()
	}
	{
		p.SetState(273)
		p.Match(MagoitoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
		p.Expr()
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserCOMMA {
		{
			p.SetState(275)
			p.Match(MagoitoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Expr()
		}

		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(282)
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

func (s *CalleeContext) CopyAll(ctx *CalleeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CalleeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalleeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentCalleeContext struct {
	CalleeContext
}

func NewIdentCalleeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentCalleeContext {
	var p = new(IdentCalleeContext)

	InitEmptyCalleeContext(&p.CalleeContext)
	p.parser = parser
	p.CopyAll(ctx.(*CalleeContext))

	return p
}

func (s *IdentCalleeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentCalleeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MagoitoParserIDENTIFIER, 0)
}

func (s *IdentCalleeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitIdentCallee(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintCalleeContext struct {
	CalleeContext
}

func NewPrintCalleeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintCalleeContext {
	var p = new(PrintCalleeContext)

	InitEmptyCalleeContext(&p.CalleeContext)
	p.parser = parser
	p.CopyAll(ctx.(*CalleeContext))

	return p
}

func (s *PrintCalleeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintCalleeContext) PRINT() antlr.TerminalNode {
	return s.GetToken(MagoitoParserPRINT, 0)
}

func (s *PrintCalleeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MagoitoParserVisitor:
		return t.VisitPrintCallee(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MagoitoParser) Callee() (localctx ICalleeContext) {
	localctx = NewCalleeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MagoitoParserRULE_callee)
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MagoitoParserIDENTIFIER:
		localctx = NewIdentCalleeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(284)
			p.Match(MagoitoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MagoitoParserPRINT:
		localctx = NewPrintCalleeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)
			p.Match(MagoitoParserPRINT)
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
		p.SetState(288)
		p.Match(MagoitoParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MagoitoParserIDENTIFIER {
		{
			p.SetState(289)
			p.RecordFieldList()
		}

	}
	{
		p.SetState(292)
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
		p.SetState(294)
		p.RecordField()
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MagoitoParserCOMMA {
		{
			p.SetState(295)
			p.Match(MagoitoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)
			p.RecordField()
		}

		p.SetState(301)
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
		p.SetState(302)
		p.Match(MagoitoParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.Match(MagoitoParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
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
