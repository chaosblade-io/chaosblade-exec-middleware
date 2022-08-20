// Code generated from Nginx.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type NginxLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var nginxlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func nginxlexerLexerInit() {
	staticData := &nginxlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "';'", "'{'", "'}'", "'if'", "'('", "')'", "'\\.'", "'^'", "'location'",
		"'rewrite'", "'last'", "'break'", "'redirect'", "'permanent'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Value",
		"STR_EXT", "LINE_COMMENT", "REGEXP_PREFIXED", "QUOTED_STRING", "SINGLE_QUOTED",
		"WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "Value", "STR_EXT", "LINE_COMMENT",
		"REGEXP_PREFIXED", "QUOTED_STRING", "RegexpPrefix", "StringCharacters",
		"NON_ASCII", "EscapeSequence", "SINGLE_QUOTED", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 21, 204, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 3, 14,
		120, 8, 14, 1, 15, 1, 15, 4, 15, 124, 8, 15, 11, 15, 12, 15, 125, 1, 16,
		1, 16, 1, 16, 1, 16, 3, 16, 132, 8, 16, 1, 16, 5, 16, 135, 8, 16, 10, 16,
		12, 16, 138, 9, 16, 1, 16, 3, 16, 141, 8, 16, 1, 16, 1, 16, 3, 16, 145,
		8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 151, 8, 16, 1, 16, 1, 16, 3,
		16, 155, 8, 16, 3, 16, 157, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 4, 17, 163,
		8, 17, 11, 17, 12, 17, 164, 1, 18, 1, 18, 3, 18, 169, 8, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 3, 19, 175, 8, 19, 1, 20, 1, 20, 4, 20, 179, 8, 20, 11,
		20, 12, 20, 180, 1, 21, 1, 21, 1, 22, 1, 22, 3, 22, 187, 8, 22, 1, 23,
		1, 23, 5, 23, 191, 8, 23, 10, 23, 12, 23, 194, 9, 23, 1, 23, 1, 23, 1,
		24, 4, 24, 199, 8, 24, 11, 24, 12, 24, 200, 1, 24, 1, 24, 0, 0, 25, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 0,
		41, 0, 43, 0, 45, 0, 47, 20, 49, 21, 1, 0, 9, 10, 0, 33, 33, 35, 36, 38,
		38, 42, 58, 61, 61, 63, 91, 93, 95, 97, 122, 124, 124, 126, 126, 2, 0,
		10, 10, 13, 13, 10, 0, 33, 33, 35, 36, 38, 38, 40, 58, 61, 61, 63, 91,
		93, 95, 97, 122, 124, 124, 126, 126, 1, 0, 126, 126, 1, 0, 42, 42, 2, 0,
		34, 34, 92, 92, 8, 0, 34, 34, 39, 39, 92, 92, 98, 98, 102, 102, 110, 110,
		114, 114, 116, 116, 2, 0, 39, 39, 92, 92, 3, 0, 9, 10, 13, 13, 32, 32,
		218, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 53, 1, 0, 0,
		0, 5, 55, 1, 0, 0, 0, 7, 57, 1, 0, 0, 0, 9, 60, 1, 0, 0, 0, 11, 62, 1,
		0, 0, 0, 13, 64, 1, 0, 0, 0, 15, 67, 1, 0, 0, 0, 17, 69, 1, 0, 0, 0, 19,
		78, 1, 0, 0, 0, 21, 86, 1, 0, 0, 0, 23, 91, 1, 0, 0, 0, 25, 97, 1, 0, 0,
		0, 27, 106, 1, 0, 0, 0, 29, 119, 1, 0, 0, 0, 31, 123, 1, 0, 0, 0, 33, 156,
		1, 0, 0, 0, 35, 160, 1, 0, 0, 0, 37, 166, 1, 0, 0, 0, 39, 172, 1, 0, 0,
		0, 41, 178, 1, 0, 0, 0, 43, 182, 1, 0, 0, 0, 45, 184, 1, 0, 0, 0, 47, 188,
		1, 0, 0, 0, 49, 198, 1, 0, 0, 0, 51, 52, 5, 59, 0, 0, 52, 2, 1, 0, 0, 0,
		53, 54, 5, 123, 0, 0, 54, 4, 1, 0, 0, 0, 55, 56, 5, 125, 0, 0, 56, 6, 1,
		0, 0, 0, 57, 58, 5, 105, 0, 0, 58, 59, 5, 102, 0, 0, 59, 8, 1, 0, 0, 0,
		60, 61, 5, 40, 0, 0, 61, 10, 1, 0, 0, 0, 62, 63, 5, 41, 0, 0, 63, 12, 1,
		0, 0, 0, 64, 65, 5, 92, 0, 0, 65, 66, 5, 46, 0, 0, 66, 14, 1, 0, 0, 0,
		67, 68, 5, 94, 0, 0, 68, 16, 1, 0, 0, 0, 69, 70, 5, 108, 0, 0, 70, 71,
		5, 111, 0, 0, 71, 72, 5, 99, 0, 0, 72, 73, 5, 97, 0, 0, 73, 74, 5, 116,
		0, 0, 74, 75, 5, 105, 0, 0, 75, 76, 5, 111, 0, 0, 76, 77, 5, 110, 0, 0,
		77, 18, 1, 0, 0, 0, 78, 79, 5, 114, 0, 0, 79, 80, 5, 101, 0, 0, 80, 81,
		5, 119, 0, 0, 81, 82, 5, 114, 0, 0, 82, 83, 5, 105, 0, 0, 83, 84, 5, 116,
		0, 0, 84, 85, 5, 101, 0, 0, 85, 20, 1, 0, 0, 0, 86, 87, 5, 108, 0, 0, 87,
		88, 5, 97, 0, 0, 88, 89, 5, 115, 0, 0, 89, 90, 5, 116, 0, 0, 90, 22, 1,
		0, 0, 0, 91, 92, 5, 98, 0, 0, 92, 93, 5, 114, 0, 0, 93, 94, 5, 101, 0,
		0, 94, 95, 5, 97, 0, 0, 95, 96, 5, 107, 0, 0, 96, 24, 1, 0, 0, 0, 97, 98,
		5, 114, 0, 0, 98, 99, 5, 101, 0, 0, 99, 100, 5, 100, 0, 0, 100, 101, 5,
		105, 0, 0, 101, 102, 5, 114, 0, 0, 102, 103, 5, 101, 0, 0, 103, 104, 5,
		99, 0, 0, 104, 105, 5, 116, 0, 0, 105, 26, 1, 0, 0, 0, 106, 107, 5, 112,
		0, 0, 107, 108, 5, 101, 0, 0, 108, 109, 5, 114, 0, 0, 109, 110, 5, 109,
		0, 0, 110, 111, 5, 97, 0, 0, 111, 112, 5, 110, 0, 0, 112, 113, 5, 101,
		0, 0, 113, 114, 5, 110, 0, 0, 114, 115, 5, 116, 0, 0, 115, 28, 1, 0, 0,
		0, 116, 120, 3, 31, 15, 0, 117, 120, 3, 37, 18, 0, 118, 120, 3, 47, 23,
		0, 119, 116, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 118, 1, 0, 0, 0, 120,
		30, 1, 0, 0, 0, 121, 124, 7, 0, 0, 0, 122, 124, 3, 43, 21, 0, 123, 121,
		1, 0, 0, 0, 123, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 123, 1, 0,
		0, 0, 125, 126, 1, 0, 0, 0, 126, 32, 1, 0, 0, 0, 127, 128, 5, 45, 0, 0,
		128, 129, 5, 45, 0, 0, 129, 132, 5, 32, 0, 0, 130, 132, 5, 35, 0, 0, 131,
		127, 1, 0, 0, 0, 131, 130, 1, 0, 0, 0, 132, 136, 1, 0, 0, 0, 133, 135,
		8, 1, 0, 0, 134, 133, 1, 0, 0, 0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0,
		0, 0, 136, 137, 1, 0, 0, 0, 137, 144, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0,
		139, 141, 5, 13, 0, 0, 140, 139, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141,
		142, 1, 0, 0, 0, 142, 145, 5, 10, 0, 0, 143, 145, 5, 0, 0, 1, 144, 140,
		1, 0, 0, 0, 144, 143, 1, 0, 0, 0, 145, 157, 1, 0, 0, 0, 146, 147, 5, 45,
		0, 0, 147, 148, 5, 45, 0, 0, 148, 154, 1, 0, 0, 0, 149, 151, 5, 13, 0,
		0, 150, 149, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152,
		155, 5, 10, 0, 0, 153, 155, 5, 0, 0, 1, 154, 150, 1, 0, 0, 0, 154, 153,
		1, 0, 0, 0, 155, 157, 1, 0, 0, 0, 156, 131, 1, 0, 0, 0, 156, 146, 1, 0,
		0, 0, 157, 158, 1, 0, 0, 0, 158, 159, 6, 16, 0, 0, 159, 34, 1, 0, 0, 0,
		160, 162, 3, 39, 19, 0, 161, 163, 7, 2, 0, 0, 162, 161, 1, 0, 0, 0, 163,
		164, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 36, 1,
		0, 0, 0, 166, 168, 5, 34, 0, 0, 167, 169, 3, 41, 20, 0, 168, 167, 1, 0,
		0, 0, 168, 169, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 171, 5, 34, 0, 0,
		171, 38, 1, 0, 0, 0, 172, 174, 7, 3, 0, 0, 173, 175, 7, 4, 0, 0, 174, 173,
		1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 40, 1, 0, 0, 0, 176, 179, 8, 5,
		0, 0, 177, 179, 3, 45, 22, 0, 178, 176, 1, 0, 0, 0, 178, 177, 1, 0, 0,
		0, 179, 180, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181,
		42, 1, 0, 0, 0, 182, 183, 2, 128, 65535, 0, 183, 44, 1, 0, 0, 0, 184, 186,
		5, 92, 0, 0, 185, 187, 7, 6, 0, 0, 186, 185, 1, 0, 0, 0, 186, 187, 1, 0,
		0, 0, 187, 46, 1, 0, 0, 0, 188, 192, 5, 39, 0, 0, 189, 191, 8, 7, 0, 0,
		190, 189, 1, 0, 0, 0, 191, 194, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192,
		193, 1, 0, 0, 0, 193, 195, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 195, 196,
		5, 39, 0, 0, 196, 48, 1, 0, 0, 0, 197, 199, 7, 8, 0, 0, 198, 197, 1, 0,
		0, 0, 199, 200, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0,
		201, 202, 1, 0, 0, 0, 202, 203, 6, 24, 0, 0, 203, 50, 1, 0, 0, 0, 19, 0,
		119, 123, 125, 131, 136, 140, 144, 150, 154, 156, 164, 168, 174, 178, 180,
		186, 192, 200, 1, 6, 0, 0,
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

// NginxLexerInit initializes any static state used to implement NginxLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewNginxLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func NginxLexerInit() {
	staticData := &nginxlexerLexerStaticData
	staticData.once.Do(nginxlexerLexerInit)
}

// NewNginxLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewNginxLexer(input antlr.CharStream) *NginxLexer {
	NginxLexerInit()
	l := new(NginxLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &nginxlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Nginx.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NginxLexer tokens.
const (
	NginxLexerT__0            = 1
	NginxLexerT__1            = 2
	NginxLexerT__2            = 3
	NginxLexerT__3            = 4
	NginxLexerT__4            = 5
	NginxLexerT__5            = 6
	NginxLexerT__6            = 7
	NginxLexerT__7            = 8
	NginxLexerT__8            = 9
	NginxLexerT__9            = 10
	NginxLexerT__10           = 11
	NginxLexerT__11           = 12
	NginxLexerT__12           = 13
	NginxLexerT__13           = 14
	NginxLexerValue           = 15
	NginxLexerSTR_EXT         = 16
	NginxLexerLINE_COMMENT    = 17
	NginxLexerREGEXP_PREFIXED = 18
	NginxLexerQUOTED_STRING   = 19
	NginxLexerSINGLE_QUOTED   = 20
	NginxLexerWS              = 21
)
