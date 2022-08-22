// Code generated from Nginx.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nginx

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type NginxParser struct {
	*antlr.BaseParser
}

var nginxParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func nginxParserInit() {
	staticData := &nginxParserStaticData
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
		"config", "statement", "genericStatement", "regexHeaderStatement", "luaBlock",
		"luaStatement", "block", "genericBlockHeader", "ifStatement", "ifBody",
		"regexp", "locationBlockHeader", "rewriteStatement",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 206, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 4, 0, 30, 8, 0, 11, 0,
		12, 0, 31, 1, 1, 3, 1, 35, 8, 1, 1, 1, 1, 1, 1, 1, 3, 1, 40, 8, 1, 1, 1,
		1, 1, 3, 1, 44, 8, 1, 1, 2, 1, 2, 1, 2, 5, 2, 49, 8, 2, 10, 2, 12, 2, 52,
		9, 2, 1, 3, 1, 3, 1, 3, 1, 4, 3, 4, 58, 8, 4, 1, 4, 1, 4, 3, 4, 62, 8,
		4, 1, 4, 1, 4, 3, 4, 66, 8, 4, 1, 4, 5, 4, 69, 8, 4, 10, 4, 12, 4, 72,
		9, 4, 1, 4, 1, 4, 3, 4, 76, 8, 4, 1, 5, 4, 5, 79, 8, 5, 11, 5, 12, 5, 80,
		1, 5, 3, 5, 84, 8, 5, 1, 5, 3, 5, 87, 8, 5, 1, 5, 3, 5, 90, 8, 5, 1, 6,
		3, 6, 93, 8, 6, 1, 6, 1, 6, 3, 6, 97, 8, 6, 1, 6, 3, 6, 100, 8, 6, 1, 6,
		1, 6, 3, 6, 104, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 110, 8, 6, 1, 6, 3,
		6, 113, 8, 6, 5, 6, 115, 8, 6, 10, 6, 12, 6, 118, 9, 6, 1, 6, 1, 6, 3,
		6, 122, 8, 6, 1, 7, 1, 7, 1, 7, 5, 7, 127, 8, 7, 10, 7, 12, 7, 130, 9,
		7, 1, 8, 1, 8, 3, 8, 134, 8, 8, 1, 8, 1, 8, 3, 8, 138, 8, 8, 1, 8, 1, 8,
		3, 8, 142, 8, 8, 1, 8, 1, 8, 3, 8, 146, 8, 8, 5, 8, 148, 8, 8, 10, 8, 12,
		8, 151, 9, 8, 1, 8, 1, 8, 3, 8, 155, 8, 8, 1, 9, 1, 9, 3, 9, 159, 8, 9,
		1, 9, 4, 9, 162, 8, 9, 11, 9, 12, 9, 163, 1, 9, 3, 9, 167, 8, 9, 1, 9,
		1, 9, 3, 9, 171, 8, 9, 1, 9, 3, 9, 174, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10, 185, 8, 10, 11, 10, 12, 10, 186,
		1, 11, 1, 11, 3, 11, 191, 8, 11, 1, 11, 1, 11, 3, 11, 195, 8, 11, 1, 12,
		1, 12, 1, 12, 3, 12, 200, 8, 12, 1, 12, 1, 12, 3, 12, 204, 8, 12, 1, 12,
		0, 0, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 2, 2, 0, 4,
		6, 16, 16, 1, 0, 11, 14, 242, 0, 29, 1, 0, 0, 0, 2, 34, 1, 0, 0, 0, 4,
		45, 1, 0, 0, 0, 6, 53, 1, 0, 0, 0, 8, 57, 1, 0, 0, 0, 10, 86, 1, 0, 0,
		0, 12, 92, 1, 0, 0, 0, 14, 123, 1, 0, 0, 0, 16, 131, 1, 0, 0, 0, 18, 156,
		1, 0, 0, 0, 20, 184, 1, 0, 0, 0, 22, 188, 1, 0, 0, 0, 24, 196, 1, 0, 0,
		0, 26, 30, 3, 2, 1, 0, 27, 30, 3, 12, 6, 0, 28, 30, 3, 8, 4, 0, 29, 26,
		1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0,
		31, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 1, 1, 0, 0, 0, 33, 35, 5, 23,
		0, 0, 34, 33, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 39, 1, 0, 0, 0, 36, 40,
		3, 24, 12, 0, 37, 40, 3, 4, 2, 0, 38, 40, 3, 6, 3, 0, 39, 36, 1, 0, 0,
		0, 39, 37, 1, 0, 0, 0, 39, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 43,
		5, 1, 0, 0, 42, 44, 5, 23, 0, 0, 43, 42, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0,
		44, 3, 1, 0, 0, 0, 45, 50, 5, 16, 0, 0, 46, 49, 5, 16, 0, 0, 47, 49, 3,
		20, 10, 0, 48, 46, 1, 0, 0, 0, 48, 47, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0,
		50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 5, 1, 0, 0, 0, 52, 50, 1, 0,
		0, 0, 53, 54, 5, 19, 0, 0, 54, 55, 5, 16, 0, 0, 55, 7, 1, 0, 0, 0, 56,
		58, 5, 23, 0, 0, 57, 56, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 1, 0,
		0, 0, 59, 61, 5, 15, 0, 0, 60, 62, 5, 23, 0, 0, 61, 60, 1, 0, 0, 0, 61,
		62, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 65, 5, 2, 0, 0, 64, 66, 5, 23,
		0, 0, 65, 64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 70, 1, 0, 0, 0, 67, 69,
		3, 10, 5, 0, 68, 67, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0,
		70, 71, 1, 0, 0, 0, 71, 73, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 75, 5,
		3, 0, 0, 74, 76, 5, 23, 0, 0, 75, 74, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76,
		9, 1, 0, 0, 0, 77, 79, 7, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0,
		0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 83, 1, 0, 0, 0, 82, 84,
		5, 1, 0, 0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0,
		85, 87, 5, 1, 0, 0, 86, 78, 1, 0, 0, 0, 86, 85, 1, 0, 0, 0, 87, 89, 1,
		0, 0, 0, 88, 90, 5, 23, 0, 0, 89, 88, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90,
		11, 1, 0, 0, 0, 91, 93, 5, 23, 0, 0, 92, 91, 1, 0, 0, 0, 92, 93, 1, 0,
		0, 0, 93, 96, 1, 0, 0, 0, 94, 97, 3, 22, 11, 0, 95, 97, 3, 14, 7, 0, 96,
		94, 1, 0, 0, 0, 96, 95, 1, 0, 0, 0, 97, 99, 1, 0, 0, 0, 98, 100, 5, 23,
		0, 0, 99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101,
		103, 5, 2, 0, 0, 102, 104, 5, 23, 0, 0, 103, 102, 1, 0, 0, 0, 103, 104,
		1, 0, 0, 0, 104, 116, 1, 0, 0, 0, 105, 110, 3, 2, 1, 0, 106, 110, 3, 12,
		6, 0, 107, 110, 3, 16, 8, 0, 108, 110, 3, 8, 4, 0, 109, 105, 1, 0, 0, 0,
		109, 106, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 108, 1, 0, 0, 0, 110,
		112, 1, 0, 0, 0, 111, 113, 5, 23, 0, 0, 112, 111, 1, 0, 0, 0, 112, 113,
		1, 0, 0, 0, 113, 115, 1, 0, 0, 0, 114, 109, 1, 0, 0, 0, 115, 118, 1, 0,
		0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 119, 1, 0, 0, 0,
		118, 116, 1, 0, 0, 0, 119, 121, 5, 3, 0, 0, 120, 122, 5, 23, 0, 0, 121,
		120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 13, 1, 0, 0, 0, 123, 128, 5,
		16, 0, 0, 124, 127, 5, 16, 0, 0, 125, 127, 3, 20, 10, 0, 126, 124, 1, 0,
		0, 0, 126, 125, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0,
		128, 129, 1, 0, 0, 0, 129, 15, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 133,
		5, 4, 0, 0, 132, 134, 5, 23, 0, 0, 133, 132, 1, 0, 0, 0, 133, 134, 1, 0,
		0, 0, 134, 135, 1, 0, 0, 0, 135, 137, 3, 18, 9, 0, 136, 138, 5, 23, 0,
		0, 137, 136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		141, 5, 2, 0, 0, 140, 142, 5, 23, 0, 0, 141, 140, 1, 0, 0, 0, 141, 142,
		1, 0, 0, 0, 142, 149, 1, 0, 0, 0, 143, 145, 3, 2, 1, 0, 144, 146, 5, 23,
		0, 0, 145, 144, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 148, 1, 0, 0, 0,
		147, 143, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149,
		150, 1, 0, 0, 0, 150, 152, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 154,
		5, 3, 0, 0, 153, 155, 5, 23, 0, 0, 154, 153, 1, 0, 0, 0, 154, 155, 1, 0,
		0, 0, 155, 17, 1, 0, 0, 0, 156, 158, 5, 6, 0, 0, 157, 159, 5, 23, 0, 0,
		158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 161, 1, 0, 0, 0, 160,
		162, 5, 16, 0, 0, 161, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 161,
		1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 166, 1, 0, 0, 0, 165, 167, 5, 23,
		0, 0, 166, 165, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 170, 1, 0, 0, 0,
		168, 171, 5, 16, 0, 0, 169, 171, 3, 20, 10, 0, 170, 168, 1, 0, 0, 0, 170,
		169, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 173, 1, 0, 0, 0, 172, 174,
		5, 23, 0, 0, 173, 172, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 1, 0,
		0, 0, 175, 176, 5, 5, 0, 0, 176, 19, 1, 0, 0, 0, 177, 185, 5, 7, 0, 0,
		178, 185, 5, 8, 0, 0, 179, 185, 5, 16, 0, 0, 180, 181, 5, 6, 0, 0, 181,
		182, 3, 20, 10, 0, 182, 183, 5, 5, 0, 0, 183, 185, 1, 0, 0, 0, 184, 177,
		1, 0, 0, 0, 184, 178, 1, 0, 0, 0, 184, 179, 1, 0, 0, 0, 184, 180, 1, 0,
		0, 0, 185, 186, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0,
		187, 21, 1, 0, 0, 0, 188, 190, 5, 9, 0, 0, 189, 191, 5, 16, 0, 0, 190,
		189, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 194, 1, 0, 0, 0, 192, 195,
		5, 16, 0, 0, 193, 195, 3, 20, 10, 0, 194, 192, 1, 0, 0, 0, 194, 193, 1,
		0, 0, 0, 195, 23, 1, 0, 0, 0, 196, 199, 5, 10, 0, 0, 197, 200, 5, 16, 0,
		0, 198, 200, 3, 20, 10, 0, 199, 197, 1, 0, 0, 0, 199, 198, 1, 0, 0, 0,
		200, 201, 1, 0, 0, 0, 201, 203, 5, 16, 0, 0, 202, 204, 7, 1, 0, 0, 203,
		202, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 25, 1, 0, 0, 0, 43, 29, 31,
		34, 39, 43, 48, 50, 57, 61, 65, 70, 75, 80, 83, 86, 89, 92, 96, 99, 103,
		109, 112, 116, 121, 126, 128, 133, 137, 141, 145, 149, 154, 158, 163, 166,
		170, 173, 184, 186, 190, 194, 199, 203,
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

// NginxParserInit initializes any static state used to implement NginxParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNginxParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NginxParserInit() {
	staticData := &nginxParserStaticData
	staticData.once.Do(nginxParserInit)
}

// NewNginxParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNginxParser(input antlr.TokenStream) *NginxParser {
	NginxParserInit()
	this := new(NginxParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &nginxParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Nginx.g4"

	return this
}

// NginxParser tokens.
const (
	NginxParserEOF             = antlr.TokenEOF
	NginxParserT__0            = 1
	NginxParserT__1            = 2
	NginxParserT__2            = 3
	NginxParserT__3            = 4
	NginxParserT__4            = 5
	NginxParserT__5            = 6
	NginxParserT__6            = 7
	NginxParserT__7            = 8
	NginxParserT__8            = 9
	NginxParserT__9            = 10
	NginxParserT__10           = 11
	NginxParserT__11           = 12
	NginxParserT__12           = 13
	NginxParserT__13           = 14
	NginxParserLUA_HEADER      = 15
	NginxParserValue           = 16
	NginxParserSTR_EXT         = 17
	NginxParserLINE_COMMENT    = 18
	NginxParserREGEXP_PREFIXED = 19
	NginxParserQUOTED_STRING   = 20
	NginxParserSINGLE_QUOTED   = 21
	NginxParserWS              = 22
	NginxParserNEWLINE         = 23
)

// NginxParser rules.
const (
	NginxParserRULE_config               = 0
	NginxParserRULE_statement            = 1
	NginxParserRULE_genericStatement     = 2
	NginxParserRULE_regexHeaderStatement = 3
	NginxParserRULE_luaBlock             = 4
	NginxParserRULE_luaStatement         = 5
	NginxParserRULE_block                = 6
	NginxParserRULE_genericBlockHeader   = 7
	NginxParserRULE_ifStatement          = 8
	NginxParserRULE_ifBody               = 9
	NginxParserRULE_regexp               = 10
	NginxParserRULE_locationBlockHeader  = 11
	NginxParserRULE_rewriteStatement     = 12
)

// IConfigContext is an interface to support dynamic dispatch.
type IConfigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfigContext differentiates from other interfaces.
	IsConfigContext()
}

type ConfigContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigContext() *ConfigContext {
	var p = new(ConfigContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_config
	return p
}

func (*ConfigContext) IsConfigContext() {}

func NewConfigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigContext {
	var p = new(ConfigContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_config

	return p
}

func (s *ConfigContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ConfigContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *ConfigContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *ConfigContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *ConfigContext) AllLuaBlock() []ILuaBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILuaBlockContext); ok {
			len++
		}
	}

	tst := make([]ILuaBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILuaBlockContext); ok {
			tst[i] = t.(ILuaBlockContext)
			i++
		}
	}

	return tst
}

func (s *ConfigContext) LuaBlock(i int) ILuaBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILuaBlockContext); ok {
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

	return t.(ILuaBlockContext)
}

func (s *ConfigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitConfig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) Config() (localctx IConfigContext) {
	this := p
	_ = this

	localctx = NewConfigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NginxParserRULE_config)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__8)|(1<<NginxParserT__9)|(1<<NginxParserLUA_HEADER)|(1<<NginxParserValue)|(1<<NginxParserREGEXP_PREFIXED)|(1<<NginxParserNEWLINE))) != 0) {
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(26)
				p.Statement()
			}

		case 2:
			{
				p.SetState(27)
				p.Block()
			}

		case 3:
			{
				p.SetState(28)
				p.LuaBlock()
			}

		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) RewriteStatement() IRewriteStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRewriteStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRewriteStatementContext)
}

func (s *StatementContext) GenericStatement() IGenericStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericStatementContext)
}

func (s *StatementContext) RegexHeaderStatement() IRegexHeaderStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexHeaderStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexHeaderStatementContext)
}

func (s *StatementContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NginxParserNEWLINE)
}

func (s *StatementContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserNEWLINE, i)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NginxParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(33)
			p.Match(NginxParserNEWLINE)
		}

	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NginxParserT__9:
		{
			p.SetState(36)
			p.RewriteStatement()
		}

	case NginxParserValue:
		{
			p.SetState(37)
			p.GenericStatement()
		}

	case NginxParserREGEXP_PREFIXED:
		{
			p.SetState(38)
			p.RegexHeaderStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(41)
		p.Match(NginxParserT__0)
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(42)
			p.Match(NginxParserNEWLINE)
		}

	}

	return localctx
}

// IGenericStatementContext is an interface to support dynamic dispatch.
type IGenericStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenericStatementContext differentiates from other interfaces.
	IsGenericStatementContext()
}

type GenericStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericStatementContext() *GenericStatementContext {
	var p = new(GenericStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_genericStatement
	return p
}

func (*GenericStatementContext) IsGenericStatementContext() {}

func NewGenericStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericStatementContext {
	var p = new(GenericStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_genericStatement

	return p
}

func (s *GenericStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericStatementContext) AllValue() []antlr.TerminalNode {
	return s.GetTokens(NginxParserValue)
}

func (s *GenericStatementContext) Value(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserValue, i)
}

func (s *GenericStatementContext) AllRegexp() []IRegexpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegexpContext); ok {
			len++
		}
	}

	tst := make([]IRegexpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegexpContext); ok {
			tst[i] = t.(IRegexpContext)
			i++
		}
	}

	return tst
}

func (s *GenericStatementContext) Regexp(i int) IRegexpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexpContext); ok {
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

	return t.(IRegexpContext)
}

func (s *GenericStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitGenericStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) GenericStatement() (localctx IGenericStatementContext) {
	this := p
	_ = this

	localctx = NewGenericStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NginxParserRULE_genericStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(NginxParserValue)
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__5)|(1<<NginxParserT__6)|(1<<NginxParserT__7)|(1<<NginxParserValue))) != 0 {
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(46)
				p.Match(NginxParserValue)
			}

		case 2:
			{
				p.SetState(47)
				p.Regexp()
			}

		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRegexHeaderStatementContext is an interface to support dynamic dispatch.
type IRegexHeaderStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegexHeaderStatementContext differentiates from other interfaces.
	IsRegexHeaderStatementContext()
}

type RegexHeaderStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexHeaderStatementContext() *RegexHeaderStatementContext {
	var p = new(RegexHeaderStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_regexHeaderStatement
	return p
}

func (*RegexHeaderStatementContext) IsRegexHeaderStatementContext() {}

func NewRegexHeaderStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexHeaderStatementContext {
	var p = new(RegexHeaderStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_regexHeaderStatement

	return p
}

func (s *RegexHeaderStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexHeaderStatementContext) REGEXP_PREFIXED() antlr.TerminalNode {
	return s.GetToken(NginxParserREGEXP_PREFIXED, 0)
}

func (s *RegexHeaderStatementContext) Value() antlr.TerminalNode {
	return s.GetToken(NginxParserValue, 0)
}

func (s *RegexHeaderStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexHeaderStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexHeaderStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitRegexHeaderStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) RegexHeaderStatement() (localctx IRegexHeaderStatementContext) {
	this := p
	_ = this

	localctx = NewRegexHeaderStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NginxParserRULE_regexHeaderStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(NginxParserREGEXP_PREFIXED)
	}
	{
		p.SetState(54)
		p.Match(NginxParserValue)
	}

	return localctx
}

// ILuaBlockContext is an interface to support dynamic dispatch.
type ILuaBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLuaBlockContext differentiates from other interfaces.
	IsLuaBlockContext()
}

type LuaBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLuaBlockContext() *LuaBlockContext {
	var p = new(LuaBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_luaBlock
	return p
}

func (*LuaBlockContext) IsLuaBlockContext() {}

func NewLuaBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LuaBlockContext {
	var p = new(LuaBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_luaBlock

	return p
}

func (s *LuaBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *LuaBlockContext) LUA_HEADER() antlr.TerminalNode {
	return s.GetToken(NginxParserLUA_HEADER, 0)
}

func (s *LuaBlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NginxParserNEWLINE)
}

func (s *LuaBlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserNEWLINE, i)
}

func (s *LuaBlockContext) AllLuaStatement() []ILuaStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILuaStatementContext); ok {
			len++
		}
	}

	tst := make([]ILuaStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILuaStatementContext); ok {
			tst[i] = t.(ILuaStatementContext)
			i++
		}
	}

	return tst
}

func (s *LuaBlockContext) LuaStatement(i int) ILuaStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILuaStatementContext); ok {
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

	return t.(ILuaStatementContext)
}

func (s *LuaBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LuaBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LuaBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitLuaBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) LuaBlock() (localctx ILuaBlockContext) {
	this := p
	_ = this

	localctx = NewLuaBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NginxParserRULE_luaBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(56)
			p.Match(NginxParserNEWLINE)
		}

	}
	{
		p.SetState(59)
		p.Match(NginxParserLUA_HEADER)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(60)
			p.Match(NginxParserNEWLINE)
		}

	}
	{
		p.SetState(63)
		p.Match(NginxParserT__1)
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(64)
			p.Match(NginxParserNEWLINE)
		}

	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__0)|(1<<NginxParserT__3)|(1<<NginxParserT__4)|(1<<NginxParserT__5)|(1<<NginxParserValue))) != 0 {
		{
			p.SetState(67)
			p.LuaStatement()
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(73)
		p.Match(NginxParserT__2)
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(74)
			p.Match(NginxParserNEWLINE)
		}

	}

	return localctx
}

// ILuaStatementContext is an interface to support dynamic dispatch.
type ILuaStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLuaStatementContext differentiates from other interfaces.
	IsLuaStatementContext()
}

type LuaStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLuaStatementContext() *LuaStatementContext {
	var p = new(LuaStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_luaStatement
	return p
}

func (*LuaStatementContext) IsLuaStatementContext() {}

func NewLuaStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LuaStatementContext {
	var p = new(LuaStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_luaStatement

	return p
}

func (s *LuaStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LuaStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(NginxParserNEWLINE, 0)
}

func (s *LuaStatementContext) AllValue() []antlr.TerminalNode {
	return s.GetTokens(NginxParserValue)
}

func (s *LuaStatementContext) Value(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserValue, i)
}

func (s *LuaStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LuaStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LuaStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitLuaStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) LuaStatement() (localctx ILuaStatementContext) {
	this := p
	_ = this

	localctx = NewLuaStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NginxParserRULE_luaStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NginxParserT__3, NginxParserT__4, NginxParserT__5, NginxParserValue:
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(77)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__3)|(1<<NginxParserT__4)|(1<<NginxParserT__5)|(1<<NginxParserValue))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(82)
				p.Match(NginxParserT__0)
			}

		}

	case NginxParserT__0:
		{
			p.SetState(85)
			p.Match(NginxParserT__0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(88)
			p.Match(NginxParserNEWLINE)
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LocationBlockHeader() ILocationBlockHeaderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocationBlockHeaderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocationBlockHeaderContext)
}

func (s *BlockContext) GenericBlockHeader() IGenericBlockHeaderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericBlockHeaderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericBlockHeaderContext)
}

func (s *BlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NginxParserNEWLINE)
}

func (s *BlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserNEWLINE, i)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *BlockContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *BlockContext) AllIfStatement() []IIfStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfStatementContext); ok {
			len++
		}
	}

	tst := make([]IIfStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfStatementContext); ok {
			tst[i] = t.(IIfStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) IfStatement(i int) IIfStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
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

	return t.(IIfStatementContext)
}

func (s *BlockContext) AllLuaBlock() []ILuaBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILuaBlockContext); ok {
			len++
		}
	}

	tst := make([]ILuaBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILuaBlockContext); ok {
			tst[i] = t.(ILuaBlockContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) LuaBlock(i int) ILuaBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILuaBlockContext); ok {
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

	return t.(ILuaBlockContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NginxParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(91)
			p.Match(NginxParserNEWLINE)
		}

	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NginxParserT__8:
		{
			p.SetState(94)
			p.LocationBlockHeader()
		}

	case NginxParserValue:
		{
			p.SetState(95)
			p.GenericBlockHeader()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(98)
			p.Match(NginxParserNEWLINE)
		}

	}
	{
		p.SetState(101)
		p.Match(NginxParserT__1)
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(102)
			p.Match(NginxParserNEWLINE)
		}

	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__3)|(1<<NginxParserT__8)|(1<<NginxParserT__9)|(1<<NginxParserLUA_HEADER)|(1<<NginxParserValue)|(1<<NginxParserREGEXP_PREFIXED)|(1<<NginxParserNEWLINE))) != 0 {
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(105)
				p.Statement()
			}

		case 2:
			{
				p.SetState(106)
				p.Block()
			}

		case 3:
			{
				p.SetState(107)
				p.IfStatement()
			}

		case 4:
			{
				p.SetState(108)
				p.LuaBlock()
			}

		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(111)
				p.Match(NginxParserNEWLINE)
			}

		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(119)
		p.Match(NginxParserT__2)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(120)
			p.Match(NginxParserNEWLINE)
		}

	}

	return localctx
}

// IGenericBlockHeaderContext is an interface to support dynamic dispatch.
type IGenericBlockHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenericBlockHeaderContext differentiates from other interfaces.
	IsGenericBlockHeaderContext()
}

type GenericBlockHeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericBlockHeaderContext() *GenericBlockHeaderContext {
	var p = new(GenericBlockHeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_genericBlockHeader
	return p
}

func (*GenericBlockHeaderContext) IsGenericBlockHeaderContext() {}

func NewGenericBlockHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericBlockHeaderContext {
	var p = new(GenericBlockHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_genericBlockHeader

	return p
}

func (s *GenericBlockHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericBlockHeaderContext) AllValue() []antlr.TerminalNode {
	return s.GetTokens(NginxParserValue)
}

func (s *GenericBlockHeaderContext) Value(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserValue, i)
}

func (s *GenericBlockHeaderContext) AllRegexp() []IRegexpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegexpContext); ok {
			len++
		}
	}

	tst := make([]IRegexpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegexpContext); ok {
			tst[i] = t.(IRegexpContext)
			i++
		}
	}

	return tst
}

func (s *GenericBlockHeaderContext) Regexp(i int) IRegexpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexpContext); ok {
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

	return t.(IRegexpContext)
}

func (s *GenericBlockHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericBlockHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericBlockHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitGenericBlockHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) GenericBlockHeader() (localctx IGenericBlockHeaderContext) {
	this := p
	_ = this

	localctx = NewGenericBlockHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NginxParserRULE_genericBlockHeader)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(NginxParserValue)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__5)|(1<<NginxParserT__6)|(1<<NginxParserT__7)|(1<<NginxParserValue))) != 0 {
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(124)
				p.Match(NginxParserValue)
			}

		case 2:
			{
				p.SetState(125)
				p.Regexp()
			}

		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IfBody() IIfBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBodyContext)
}

func (s *IfStatementContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NginxParserNEWLINE)
}

func (s *IfStatementContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserNEWLINE, i)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) IfStatement() (localctx IIfStatementContext) {
	this := p
	_ = this

	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NginxParserRULE_ifStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(NginxParserT__3)
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(132)
			p.Match(NginxParserNEWLINE)
		}

	}
	{
		p.SetState(135)
		p.IfBody()
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(136)
			p.Match(NginxParserNEWLINE)
		}

	}
	{
		p.SetState(139)
		p.Match(NginxParserT__1)
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(140)
			p.Match(NginxParserNEWLINE)
		}

	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__9)|(1<<NginxParserValue)|(1<<NginxParserREGEXP_PREFIXED)|(1<<NginxParserNEWLINE))) != 0 {
		{
			p.SetState(143)
			p.Statement()
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(144)
				p.Match(NginxParserNEWLINE)
			}

		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(152)
		p.Match(NginxParserT__2)
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(153)
			p.Match(NginxParserNEWLINE)
		}

	}

	return localctx
}

// IIfBodyContext is an interface to support dynamic dispatch.
type IIfBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfBodyContext differentiates from other interfaces.
	IsIfBodyContext()
}

type IfBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBodyContext() *IfBodyContext {
	var p = new(IfBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_ifBody
	return p
}

func (*IfBodyContext) IsIfBodyContext() {}

func NewIfBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBodyContext {
	var p = new(IfBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_ifBody

	return p
}

func (s *IfBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBodyContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NginxParserNEWLINE)
}

func (s *IfBodyContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserNEWLINE, i)
}

func (s *IfBodyContext) AllValue() []antlr.TerminalNode {
	return s.GetTokens(NginxParserValue)
}

func (s *IfBodyContext) Value(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserValue, i)
}

func (s *IfBodyContext) Regexp() IRegexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexpContext)
}

func (s *IfBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitIfBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) IfBody() (localctx IIfBodyContext) {
	this := p
	_ = this

	localctx = NewIfBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NginxParserRULE_ifBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(NginxParserT__5)
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(157)
			p.Match(NginxParserNEWLINE)
		}

	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(160)
				p.Match(NginxParserValue)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(165)
			p.Match(NginxParserNEWLINE)
		}

	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(168)
			p.Match(NginxParserValue)
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(169)
			p.Regexp()
		}

	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(172)
			p.Match(NginxParserNEWLINE)
		}

	}
	{
		p.SetState(175)
		p.Match(NginxParserT__4)
	}

	return localctx
}

// IRegexpContext is an interface to support dynamic dispatch.
type IRegexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegexpContext differentiates from other interfaces.
	IsRegexpContext()
}

type RegexpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexpContext() *RegexpContext {
	var p = new(RegexpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_regexp
	return p
}

func (*RegexpContext) IsRegexpContext() {}

func NewRegexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexpContext {
	var p = new(RegexpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_regexp

	return p
}

func (s *RegexpContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexpContext) AllValue() []antlr.TerminalNode {
	return s.GetTokens(NginxParserValue)
}

func (s *RegexpContext) Value(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserValue, i)
}

func (s *RegexpContext) AllRegexp() []IRegexpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegexpContext); ok {
			len++
		}
	}

	tst := make([]IRegexpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegexpContext); ok {
			tst[i] = t.(IRegexpContext)
			i++
		}
	}

	return tst
}

func (s *RegexpContext) Regexp(i int) IRegexpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexpContext); ok {
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

	return t.(IRegexpContext)
}

func (s *RegexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitRegexp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) Regexp() (localctx IRegexpContext) {
	this := p
	_ = this

	localctx = NewRegexpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NginxParserRULE_regexp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(184)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case NginxParserT__6:
				{
					p.SetState(177)
					p.Match(NginxParserT__6)
				}

			case NginxParserT__7:
				{
					p.SetState(178)
					p.Match(NginxParserT__7)
				}

			case NginxParserValue:
				{
					p.SetState(179)
					p.Match(NginxParserValue)
				}

			case NginxParserT__5:
				{
					p.SetState(180)
					p.Match(NginxParserT__5)
				}
				{
					p.SetState(181)
					p.Regexp()
				}
				{
					p.SetState(182)
					p.Match(NginxParserT__4)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
	}

	return localctx
}

// ILocationBlockHeaderContext is an interface to support dynamic dispatch.
type ILocationBlockHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocationBlockHeaderContext differentiates from other interfaces.
	IsLocationBlockHeaderContext()
}

type LocationBlockHeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocationBlockHeaderContext() *LocationBlockHeaderContext {
	var p = new(LocationBlockHeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_locationBlockHeader
	return p
}

func (*LocationBlockHeaderContext) IsLocationBlockHeaderContext() {}

func NewLocationBlockHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocationBlockHeaderContext {
	var p = new(LocationBlockHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_locationBlockHeader

	return p
}

func (s *LocationBlockHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *LocationBlockHeaderContext) AllValue() []antlr.TerminalNode {
	return s.GetTokens(NginxParserValue)
}

func (s *LocationBlockHeaderContext) Value(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserValue, i)
}

func (s *LocationBlockHeaderContext) Regexp() IRegexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexpContext)
}

func (s *LocationBlockHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocationBlockHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocationBlockHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitLocationBlockHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) LocationBlockHeader() (localctx ILocationBlockHeaderContext) {
	this := p
	_ = this

	localctx = NewLocationBlockHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NginxParserRULE_locationBlockHeader)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(NginxParserT__8)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(189)
			p.Match(NginxParserValue)
		}

	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(192)
			p.Match(NginxParserValue)
		}

	case 2:
		{
			p.SetState(193)
			p.Regexp()
		}

	}

	return localctx
}

// IRewriteStatementContext is an interface to support dynamic dispatch.
type IRewriteStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRewriteStatementContext differentiates from other interfaces.
	IsRewriteStatementContext()
}

type RewriteStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRewriteStatementContext() *RewriteStatementContext {
	var p = new(RewriteStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_rewriteStatement
	return p
}

func (*RewriteStatementContext) IsRewriteStatementContext() {}

func NewRewriteStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RewriteStatementContext {
	var p = new(RewriteStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_rewriteStatement

	return p
}

func (s *RewriteStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *RewriteStatementContext) AllValue() []antlr.TerminalNode {
	return s.GetTokens(NginxParserValue)
}

func (s *RewriteStatementContext) Value(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserValue, i)
}

func (s *RewriteStatementContext) Regexp() IRegexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexpContext)
}

func (s *RewriteStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RewriteStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RewriteStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitRewriteStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) RewriteStatement() (localctx IRewriteStatementContext) {
	this := p
	_ = this

	localctx = NewRewriteStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NginxParserRULE_rewriteStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(NginxParserT__9)
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(197)
			p.Match(NginxParserValue)
		}

	case 2:
		{
			p.SetState(198)
			p.Regexp()
		}

	}
	{
		p.SetState(201)
		p.Match(NginxParserValue)
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__10)|(1<<NginxParserT__11)|(1<<NginxParserT__12)|(1<<NginxParserT__13))) != 0 {
		{
			p.SetState(202)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__10)|(1<<NginxParserT__11)|(1<<NginxParserT__12)|(1<<NginxParserT__13))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}
