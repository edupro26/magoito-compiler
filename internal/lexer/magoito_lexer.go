// Code generated from MagoitoLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package lexer

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MagoitoLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MagoitoLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func magoitolexerLexerInit() {
	staticData := &MagoitoLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"CONST", "FUN", "VAR", "IF", "THEN", "ELSE", "WHILE", "DO", "PRINT",
		"TRUE", "FALSE", "UNIT_VALUE", "INT_TYPE", "BOOL_TYPE", "STRING_TYPE",
		"UNIT_TYPE", "ARROW", "EQ", "NEQ", "LTE", "GTE", "OR", "AND", "ASSIGN",
		"LT", "GT", "PLUS", "MINUS", "STAR", "SLASH", "PERCENT", "POWER", "NOT",
		"DOT", "SEMICOLON", "COLON", "COMMA", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "WILDCARD", "IDENTIFIER", "INT_LITERAL", "STRING_LITERAL",
		"LINE_COMMENT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 47, 277, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28,
		1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1,
		33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38,
		1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 5, 42, 236, 8,
		42, 10, 42, 12, 42, 239, 9, 42, 1, 43, 1, 43, 1, 43, 5, 43, 244, 8, 43,
		10, 43, 12, 43, 247, 9, 43, 3, 43, 249, 8, 43, 1, 44, 1, 44, 5, 44, 253,
		8, 44, 10, 44, 12, 44, 256, 9, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1,
		45, 5, 45, 264, 8, 45, 10, 45, 12, 45, 267, 9, 45, 1, 45, 1, 45, 1, 46,
		4, 46, 272, 8, 46, 11, 46, 12, 46, 273, 1, 46, 1, 46, 0, 0, 47, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30,
		61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39,
		79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 1, 0, 7,
		2, 0, 65, 90, 97, 122, 5, 0, 39, 39, 48, 57, 65, 90, 95, 95, 97, 122, 1,
		0, 49, 57, 1, 0, 48, 57, 3, 0, 10, 10, 13, 13, 34, 34, 2, 0, 10, 10, 13,
		13, 3, 0, 9, 10, 13, 13, 32, 32, 282, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0,
		0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1,
		0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81,
		1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0,
		89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 1, 95, 1, 0, 0, 0,
		3, 101, 1, 0, 0, 0, 5, 105, 1, 0, 0, 0, 7, 109, 1, 0, 0, 0, 9, 112, 1,
		0, 0, 0, 11, 117, 1, 0, 0, 0, 13, 122, 1, 0, 0, 0, 15, 128, 1, 0, 0, 0,
		17, 131, 1, 0, 0, 0, 19, 137, 1, 0, 0, 0, 21, 142, 1, 0, 0, 0, 23, 148,
		1, 0, 0, 0, 25, 153, 1, 0, 0, 0, 27, 157, 1, 0, 0, 0, 29, 162, 1, 0, 0,
		0, 31, 169, 1, 0, 0, 0, 33, 174, 1, 0, 0, 0, 35, 177, 1, 0, 0, 0, 37, 180,
		1, 0, 0, 0, 39, 183, 1, 0, 0, 0, 41, 186, 1, 0, 0, 0, 43, 189, 1, 0, 0,
		0, 45, 192, 1, 0, 0, 0, 47, 195, 1, 0, 0, 0, 49, 197, 1, 0, 0, 0, 51, 199,
		1, 0, 0, 0, 53, 201, 1, 0, 0, 0, 55, 203, 1, 0, 0, 0, 57, 205, 1, 0, 0,
		0, 59, 207, 1, 0, 0, 0, 61, 209, 1, 0, 0, 0, 63, 211, 1, 0, 0, 0, 65, 213,
		1, 0, 0, 0, 67, 215, 1, 0, 0, 0, 69, 217, 1, 0, 0, 0, 71, 219, 1, 0, 0,
		0, 73, 221, 1, 0, 0, 0, 75, 223, 1, 0, 0, 0, 77, 225, 1, 0, 0, 0, 79, 227,
		1, 0, 0, 0, 81, 229, 1, 0, 0, 0, 83, 231, 1, 0, 0, 0, 85, 233, 1, 0, 0,
		0, 87, 248, 1, 0, 0, 0, 89, 250, 1, 0, 0, 0, 91, 259, 1, 0, 0, 0, 93, 271,
		1, 0, 0, 0, 95, 96, 5, 99, 0, 0, 96, 97, 5, 111, 0, 0, 97, 98, 5, 110,
		0, 0, 98, 99, 5, 115, 0, 0, 99, 100, 5, 116, 0, 0, 100, 2, 1, 0, 0, 0,
		101, 102, 5, 102, 0, 0, 102, 103, 5, 117, 0, 0, 103, 104, 5, 110, 0, 0,
		104, 4, 1, 0, 0, 0, 105, 106, 5, 118, 0, 0, 106, 107, 5, 97, 0, 0, 107,
		108, 5, 114, 0, 0, 108, 6, 1, 0, 0, 0, 109, 110, 5, 105, 0, 0, 110, 111,
		5, 102, 0, 0, 111, 8, 1, 0, 0, 0, 112, 113, 5, 116, 0, 0, 113, 114, 5,
		104, 0, 0, 114, 115, 5, 101, 0, 0, 115, 116, 5, 110, 0, 0, 116, 10, 1,
		0, 0, 0, 117, 118, 5, 101, 0, 0, 118, 119, 5, 108, 0, 0, 119, 120, 5, 115,
		0, 0, 120, 121, 5, 101, 0, 0, 121, 12, 1, 0, 0, 0, 122, 123, 5, 119, 0,
		0, 123, 124, 5, 104, 0, 0, 124, 125, 5, 105, 0, 0, 125, 126, 5, 108, 0,
		0, 126, 127, 5, 101, 0, 0, 127, 14, 1, 0, 0, 0, 128, 129, 5, 100, 0, 0,
		129, 130, 5, 111, 0, 0, 130, 16, 1, 0, 0, 0, 131, 132, 5, 112, 0, 0, 132,
		133, 5, 114, 0, 0, 133, 134, 5, 105, 0, 0, 134, 135, 5, 110, 0, 0, 135,
		136, 5, 116, 0, 0, 136, 18, 1, 0, 0, 0, 137, 138, 5, 116, 0, 0, 138, 139,
		5, 114, 0, 0, 139, 140, 5, 117, 0, 0, 140, 141, 5, 101, 0, 0, 141, 20,
		1, 0, 0, 0, 142, 143, 5, 102, 0, 0, 143, 144, 5, 97, 0, 0, 144, 145, 5,
		108, 0, 0, 145, 146, 5, 115, 0, 0, 146, 147, 5, 101, 0, 0, 147, 22, 1,
		0, 0, 0, 148, 149, 5, 117, 0, 0, 149, 150, 5, 110, 0, 0, 150, 151, 5, 105,
		0, 0, 151, 152, 5, 116, 0, 0, 152, 24, 1, 0, 0, 0, 153, 154, 5, 73, 0,
		0, 154, 155, 5, 110, 0, 0, 155, 156, 5, 116, 0, 0, 156, 26, 1, 0, 0, 0,
		157, 158, 5, 66, 0, 0, 158, 159, 5, 111, 0, 0, 159, 160, 5, 111, 0, 0,
		160, 161, 5, 108, 0, 0, 161, 28, 1, 0, 0, 0, 162, 163, 5, 83, 0, 0, 163,
		164, 5, 116, 0, 0, 164, 165, 5, 114, 0, 0, 165, 166, 5, 105, 0, 0, 166,
		167, 5, 110, 0, 0, 167, 168, 5, 103, 0, 0, 168, 30, 1, 0, 0, 0, 169, 170,
		5, 85, 0, 0, 170, 171, 5, 110, 0, 0, 171, 172, 5, 105, 0, 0, 172, 173,
		5, 116, 0, 0, 173, 32, 1, 0, 0, 0, 174, 175, 5, 45, 0, 0, 175, 176, 5,
		62, 0, 0, 176, 34, 1, 0, 0, 0, 177, 178, 5, 61, 0, 0, 178, 179, 5, 61,
		0, 0, 179, 36, 1, 0, 0, 0, 180, 181, 5, 33, 0, 0, 181, 182, 5, 61, 0, 0,
		182, 38, 1, 0, 0, 0, 183, 184, 5, 60, 0, 0, 184, 185, 5, 61, 0, 0, 185,
		40, 1, 0, 0, 0, 186, 187, 5, 62, 0, 0, 187, 188, 5, 61, 0, 0, 188, 42,
		1, 0, 0, 0, 189, 190, 5, 124, 0, 0, 190, 191, 5, 124, 0, 0, 191, 44, 1,
		0, 0, 0, 192, 193, 5, 38, 0, 0, 193, 194, 5, 38, 0, 0, 194, 46, 1, 0, 0,
		0, 195, 196, 5, 61, 0, 0, 196, 48, 1, 0, 0, 0, 197, 198, 5, 60, 0, 0, 198,
		50, 1, 0, 0, 0, 199, 200, 5, 62, 0, 0, 200, 52, 1, 0, 0, 0, 201, 202, 5,
		43, 0, 0, 202, 54, 1, 0, 0, 0, 203, 204, 5, 45, 0, 0, 204, 56, 1, 0, 0,
		0, 205, 206, 5, 42, 0, 0, 206, 58, 1, 0, 0, 0, 207, 208, 5, 47, 0, 0, 208,
		60, 1, 0, 0, 0, 209, 210, 5, 37, 0, 0, 210, 62, 1, 0, 0, 0, 211, 212, 5,
		94, 0, 0, 212, 64, 1, 0, 0, 0, 213, 214, 5, 33, 0, 0, 214, 66, 1, 0, 0,
		0, 215, 216, 5, 46, 0, 0, 216, 68, 1, 0, 0, 0, 217, 218, 5, 59, 0, 0, 218,
		70, 1, 0, 0, 0, 219, 220, 5, 58, 0, 0, 220, 72, 1, 0, 0, 0, 221, 222, 5,
		44, 0, 0, 222, 74, 1, 0, 0, 0, 223, 224, 5, 40, 0, 0, 224, 76, 1, 0, 0,
		0, 225, 226, 5, 41, 0, 0, 226, 78, 1, 0, 0, 0, 227, 228, 5, 123, 0, 0,
		228, 80, 1, 0, 0, 0, 229, 230, 5, 125, 0, 0, 230, 82, 1, 0, 0, 0, 231,
		232, 5, 95, 0, 0, 232, 84, 1, 0, 0, 0, 233, 237, 7, 0, 0, 0, 234, 236,
		7, 1, 0, 0, 235, 234, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237, 235, 1, 0,
		0, 0, 237, 238, 1, 0, 0, 0, 238, 86, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0,
		240, 249, 5, 48, 0, 0, 241, 245, 7, 2, 0, 0, 242, 244, 7, 3, 0, 0, 243,
		242, 1, 0, 0, 0, 244, 247, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245, 246,
		1, 0, 0, 0, 246, 249, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 248, 240, 1, 0,
		0, 0, 248, 241, 1, 0, 0, 0, 249, 88, 1, 0, 0, 0, 250, 254, 5, 34, 0, 0,
		251, 253, 8, 4, 0, 0, 252, 251, 1, 0, 0, 0, 253, 256, 1, 0, 0, 0, 254,
		252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 257, 1, 0, 0, 0, 256, 254,
		1, 0, 0, 0, 257, 258, 5, 34, 0, 0, 258, 90, 1, 0, 0, 0, 259, 260, 5, 45,
		0, 0, 260, 261, 5, 45, 0, 0, 261, 265, 1, 0, 0, 0, 262, 264, 8, 5, 0, 0,
		263, 262, 1, 0, 0, 0, 264, 267, 1, 0, 0, 0, 265, 263, 1, 0, 0, 0, 265,
		266, 1, 0, 0, 0, 266, 268, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 268, 269,
		6, 45, 0, 0, 269, 92, 1, 0, 0, 0, 270, 272, 7, 6, 0, 0, 271, 270, 1, 0,
		0, 0, 272, 273, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0,
		274, 275, 1, 0, 0, 0, 275, 276, 6, 46, 0, 0, 276, 94, 1, 0, 0, 0, 7, 0,
		237, 245, 248, 254, 265, 273, 1, 6, 0, 0,
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

// MagoitoLexerInit initializes any static state used to implement MagoitoLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMagoitoLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MagoitoLexerInit() {
	staticData := &MagoitoLexerLexerStaticData
	staticData.once.Do(magoitolexerLexerInit)
}

// NewMagoitoLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMagoitoLexer(input antlr.CharStream) *MagoitoLexer {
	MagoitoLexerInit()
	l := new(MagoitoLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MagoitoLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "MagoitoLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MagoitoLexer tokens.
const (
	MagoitoLexerCONST          = 1
	MagoitoLexerFUN            = 2
	MagoitoLexerVAR            = 3
	MagoitoLexerIF             = 4
	MagoitoLexerTHEN           = 5
	MagoitoLexerELSE           = 6
	MagoitoLexerWHILE          = 7
	MagoitoLexerDO             = 8
	MagoitoLexerPRINT          = 9
	MagoitoLexerTRUE           = 10
	MagoitoLexerFALSE          = 11
	MagoitoLexerUNIT_VALUE     = 12
	MagoitoLexerINT_TYPE       = 13
	MagoitoLexerBOOL_TYPE      = 14
	MagoitoLexerSTRING_TYPE    = 15
	MagoitoLexerUNIT_TYPE      = 16
	MagoitoLexerARROW          = 17
	MagoitoLexerEQ             = 18
	MagoitoLexerNEQ            = 19
	MagoitoLexerLTE            = 20
	MagoitoLexerGTE            = 21
	MagoitoLexerOR             = 22
	MagoitoLexerAND            = 23
	MagoitoLexerASSIGN         = 24
	MagoitoLexerLT             = 25
	MagoitoLexerGT             = 26
	MagoitoLexerPLUS           = 27
	MagoitoLexerMINUS          = 28
	MagoitoLexerSTAR           = 29
	MagoitoLexerSLASH          = 30
	MagoitoLexerPERCENT        = 31
	MagoitoLexerPOWER          = 32
	MagoitoLexerNOT            = 33
	MagoitoLexerDOT            = 34
	MagoitoLexerSEMICOLON      = 35
	MagoitoLexerCOLON          = 36
	MagoitoLexerCOMMA          = 37
	MagoitoLexerLPAREN         = 38
	MagoitoLexerRPAREN         = 39
	MagoitoLexerLBRACE         = 40
	MagoitoLexerRBRACE         = 41
	MagoitoLexerWILDCARD       = 42
	MagoitoLexerIDENTIFIER     = 43
	MagoitoLexerINT_LITERAL    = 44
	MagoitoLexerSTRING_LITERAL = 45
	MagoitoLexerLINE_COMMENT   = 46
	MagoitoLexerWS             = 47
)
