/*
 * Copyright 2025 The ChaosBlade Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated from Nginx.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var (
	_ = fmt.Printf
	_ = sync.Once{}
	_ = unicode.IsLetter
)

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
		"", "';'", "'{'", "'}'", "'if'", "')'", "'('", "'\\.'", "'^'", "'location'",
		"'rewrite'", "'last'", "'break'", "'redirect'", "'permanent'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "LUA_HEADER",
		"Value", "STR_EXT", "LINE_COMMENT", "REGEXP_PREFIXED", "QUOTED_STRING",
		"SINGLE_QUOTED", "WS", "NEWLINE",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "LUA_HEADER", "Value", "STR_EXT",
		"LINE_COMMENT", "REGEXP_PREFIXED", "QUOTED_STRING", "RegexpPrefix",
		"NON_ASCII", "EscapeSequence", "SINGLE_QUOTED", "WS", "NEWLINE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 23, 230, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 4, 14, 120,
		8, 14, 11, 14, 12, 14, 121, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 15,
		140, 8, 15, 1, 16, 1, 16, 4, 16, 144, 8, 16, 11, 16, 12, 16, 145, 1, 17,
		1, 17, 1, 17, 1, 17, 3, 17, 152, 8, 17, 1, 17, 5, 17, 155, 8, 17, 10, 17,
		12, 17, 158, 9, 17, 1, 17, 3, 17, 161, 8, 17, 1, 17, 1, 17, 3, 17, 165,
		8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 171, 8, 17, 1, 17, 1, 17, 3,
		17, 175, 8, 17, 3, 17, 177, 8, 17, 1, 17, 3, 17, 180, 8, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 4, 18, 186, 8, 18, 11, 18, 12, 18, 187, 1, 19, 1, 19,
		1, 19, 5, 19, 193, 8, 19, 10, 19, 12, 19, 196, 9, 19, 1, 19, 1, 19, 1,
		20, 1, 20, 3, 20, 202, 8, 20, 1, 21, 1, 21, 1, 22, 1, 22, 3, 22, 208, 8,
		22, 1, 23, 1, 23, 5, 23, 212, 8, 23, 10, 23, 12, 23, 215, 9, 23, 1, 23,
		1, 23, 1, 24, 4, 24, 220, 8, 24, 11, 24, 12, 24, 221, 1, 24, 1, 24, 1,
		25, 4, 25, 227, 8, 25, 11, 25, 12, 25, 228, 0, 0, 26, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 0, 43, 0, 45,
		0, 47, 21, 49, 22, 51, 23, 1, 0, 9, 2, 0, 95, 95, 97, 122, 10, 0, 33, 33,
		35, 36, 38, 38, 40, 58, 61, 61, 63, 91, 93, 95, 97, 122, 124, 124, 126,
		126, 2, 0, 10, 10, 13, 13, 1, 0, 34, 34, 1, 0, 126, 126, 1, 0, 42, 42,
		8, 0, 34, 34, 39, 39, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116,
		116, 1, 0, 39, 39, 2, 0, 9, 9, 32, 32, 247, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 1, 53, 1, 0, 0, 0, 3, 55, 1, 0,
		0, 0, 5, 57, 1, 0, 0, 0, 7, 59, 1, 0, 0, 0, 9, 62, 1, 0, 0, 0, 11, 64,
		1, 0, 0, 0, 13, 66, 1, 0, 0, 0, 15, 69, 1, 0, 0, 0, 17, 71, 1, 0, 0, 0,
		19, 80, 1, 0, 0, 0, 21, 88, 1, 0, 0, 0, 23, 93, 1, 0, 0, 0, 25, 99, 1,
		0, 0, 0, 27, 108, 1, 0, 0, 0, 29, 119, 1, 0, 0, 0, 31, 139, 1, 0, 0, 0,
		33, 143, 1, 0, 0, 0, 35, 176, 1, 0, 0, 0, 37, 183, 1, 0, 0, 0, 39, 189,
		1, 0, 0, 0, 41, 199, 1, 0, 0, 0, 43, 203, 1, 0, 0, 0, 45, 205, 1, 0, 0,
		0, 47, 209, 1, 0, 0, 0, 49, 219, 1, 0, 0, 0, 51, 226, 1, 0, 0, 0, 53, 54,
		5, 59, 0, 0, 54, 2, 1, 0, 0, 0, 55, 56, 5, 123, 0, 0, 56, 4, 1, 0, 0, 0,
		57, 58, 5, 125, 0, 0, 58, 6, 1, 0, 0, 0, 59, 60, 5, 105, 0, 0, 60, 61,
		5, 102, 0, 0, 61, 8, 1, 0, 0, 0, 62, 63, 5, 41, 0, 0, 63, 10, 1, 0, 0,
		0, 64, 65, 5, 40, 0, 0, 65, 12, 1, 0, 0, 0, 66, 67, 5, 92, 0, 0, 67, 68,
		5, 46, 0, 0, 68, 14, 1, 0, 0, 0, 69, 70, 5, 94, 0, 0, 70, 16, 1, 0, 0,
		0, 71, 72, 5, 108, 0, 0, 72, 73, 5, 111, 0, 0, 73, 74, 5, 99, 0, 0, 74,
		75, 5, 97, 0, 0, 75, 76, 5, 116, 0, 0, 76, 77, 5, 105, 0, 0, 77, 78, 5,
		111, 0, 0, 78, 79, 5, 110, 0, 0, 79, 18, 1, 0, 0, 0, 80, 81, 5, 114, 0,
		0, 81, 82, 5, 101, 0, 0, 82, 83, 5, 119, 0, 0, 83, 84, 5, 114, 0, 0, 84,
		85, 5, 105, 0, 0, 85, 86, 5, 116, 0, 0, 86, 87, 5, 101, 0, 0, 87, 20, 1,
		0, 0, 0, 88, 89, 5, 108, 0, 0, 89, 90, 5, 97, 0, 0, 90, 91, 5, 115, 0,
		0, 91, 92, 5, 116, 0, 0, 92, 22, 1, 0, 0, 0, 93, 94, 5, 98, 0, 0, 94, 95,
		5, 114, 0, 0, 95, 96, 5, 101, 0, 0, 96, 97, 5, 97, 0, 0, 97, 98, 5, 107,
		0, 0, 98, 24, 1, 0, 0, 0, 99, 100, 5, 114, 0, 0, 100, 101, 5, 101, 0, 0,
		101, 102, 5, 100, 0, 0, 102, 103, 5, 105, 0, 0, 103, 104, 5, 114, 0, 0,
		104, 105, 5, 101, 0, 0, 105, 106, 5, 99, 0, 0, 106, 107, 5, 116, 0, 0,
		107, 26, 1, 0, 0, 0, 108, 109, 5, 112, 0, 0, 109, 110, 5, 101, 0, 0, 110,
		111, 5, 114, 0, 0, 111, 112, 5, 109, 0, 0, 112, 113, 5, 97, 0, 0, 113,
		114, 5, 110, 0, 0, 114, 115, 5, 101, 0, 0, 115, 116, 5, 110, 0, 0, 116,
		117, 5, 116, 0, 0, 117, 28, 1, 0, 0, 0, 118, 120, 7, 0, 0, 0, 119, 118,
		1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 121, 122, 1, 0,
		0, 0, 122, 123, 1, 0, 0, 0, 123, 124, 5, 98, 0, 0, 124, 125, 5, 121, 0,
		0, 125, 126, 5, 95, 0, 0, 126, 127, 5, 108, 0, 0, 127, 128, 5, 117, 0,
		0, 128, 129, 5, 97, 0, 0, 129, 130, 5, 95, 0, 0, 130, 131, 5, 98, 0, 0,
		131, 132, 5, 108, 0, 0, 132, 133, 5, 111, 0, 0, 133, 134, 5, 99, 0, 0,
		134, 135, 5, 107, 0, 0, 135, 30, 1, 0, 0, 0, 136, 140, 3, 33, 16, 0, 137,
		140, 3, 39, 19, 0, 138, 140, 3, 47, 23, 0, 139, 136, 1, 0, 0, 0, 139, 137,
		1, 0, 0, 0, 139, 138, 1, 0, 0, 0, 140, 32, 1, 0, 0, 0, 141, 144, 7, 1,
		0, 0, 142, 144, 3, 43, 21, 0, 143, 141, 1, 0, 0, 0, 143, 142, 1, 0, 0,
		0, 144, 145, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146,
		34, 1, 0, 0, 0, 147, 148, 5, 45, 0, 0, 148, 149, 5, 45, 0, 0, 149, 152,
		5, 32, 0, 0, 150, 152, 5, 35, 0, 0, 151, 147, 1, 0, 0, 0, 151, 150, 1,
		0, 0, 0, 152, 156, 1, 0, 0, 0, 153, 155, 8, 2, 0, 0, 154, 153, 1, 0, 0,
		0, 155, 158, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157,
		164, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 159, 161, 5, 13, 0, 0, 160, 159,
		1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 165, 5, 10,
		0, 0, 163, 165, 5, 0, 0, 1, 164, 160, 1, 0, 0, 0, 164, 163, 1, 0, 0, 0,
		165, 177, 1, 0, 0, 0, 166, 167, 5, 45, 0, 0, 167, 168, 5, 45, 0, 0, 168,
		174, 1, 0, 0, 0, 169, 171, 5, 13, 0, 0, 170, 169, 1, 0, 0, 0, 170, 171,
		1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 175, 5, 10, 0, 0, 173, 175, 5, 0,
		0, 1, 174, 170, 1, 0, 0, 0, 174, 173, 1, 0, 0, 0, 175, 177, 1, 0, 0, 0,
		176, 151, 1, 0, 0, 0, 176, 166, 1, 0, 0, 0, 177, 179, 1, 0, 0, 0, 178,
		180, 3, 51, 25, 0, 179, 178, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 181,
		1, 0, 0, 0, 181, 182, 6, 17, 0, 0, 182, 36, 1, 0, 0, 0, 183, 185, 3, 41,
		20, 0, 184, 186, 7, 1, 0, 0, 185, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0,
		187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 38, 1, 0, 0, 0, 189, 194,
		5, 34, 0, 0, 190, 193, 8, 3, 0, 0, 191, 193, 3, 45, 22, 0, 192, 190, 1,
		0, 0, 0, 192, 191, 1, 0, 0, 0, 193, 196, 1, 0, 0, 0, 194, 192, 1, 0, 0,
		0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 197,
		198, 5, 34, 0, 0, 198, 40, 1, 0, 0, 0, 199, 201, 7, 4, 0, 0, 200, 202,
		7, 5, 0, 0, 201, 200, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 42, 1, 0,
		0, 0, 203, 204, 2, 128, 65535, 0, 204, 44, 1, 0, 0, 0, 205, 207, 5, 92,
		0, 0, 206, 208, 7, 6, 0, 0, 207, 206, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0,
		208, 46, 1, 0, 0, 0, 209, 213, 5, 39, 0, 0, 210, 212, 8, 7, 0, 0, 211,
		210, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214,
		1, 0, 0, 0, 214, 216, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 217, 5, 39,
		0, 0, 217, 48, 1, 0, 0, 0, 218, 220, 7, 8, 0, 0, 219, 218, 1, 0, 0, 0,
		220, 221, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222,
		223, 1, 0, 0, 0, 223, 224, 6, 24, 0, 0, 224, 50, 1, 0, 0, 0, 225, 227,
		7, 2, 0, 0, 226, 225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 226, 1, 0,
		0, 0, 228, 229, 1, 0, 0, 0, 229, 52, 1, 0, 0, 0, 21, 0, 121, 139, 143,
		145, 151, 156, 160, 164, 170, 174, 176, 179, 187, 192, 194, 201, 207, 213,
		221, 228, 1, 6, 0, 0,
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
	NginxLexerLUA_HEADER      = 15
	NginxLexerValue           = 16
	NginxLexerSTR_EXT         = 17
	NginxLexerLINE_COMMENT    = 18
	NginxLexerREGEXP_PREFIXED = 19
	NginxLexerQUOTED_STRING   = 20
	NginxLexerSINGLE_QUOTED   = 21
	NginxLexerWS              = 22
	NginxLexerNEWLINE         = 23
)
