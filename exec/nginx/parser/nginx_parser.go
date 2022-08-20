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
		"", "';'", "'{'", "'}'", "'if'", "'('", "')'", "'\\.'", "'^'", "'location'",
		"'rewrite'", "'last'", "'break'", "'redirect'", "'permanent'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Value",
		"STR_EXT", "LINE_COMMENT", "REGEXP_PREFIXED", "QUOTED_STRING", "SINGLE_QUOTED",
		"WS",
	}
	staticData.ruleNames = []string{
		"config", "statement", "genericStatement", "regexHeaderStatement", "block",
		"genericBlockHeader", "ifStatement", "ifBody", "regexp", "locationBlockHeader",
		"rewriteStatement",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 120, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 4, 0, 25, 8, 0, 11, 0, 12, 0, 26, 1, 1, 1, 1, 1, 1, 3,
		1, 32, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 39, 8, 2, 10, 2, 12, 2,
		42, 9, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 49, 8, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 5, 4, 55, 8, 4, 10, 4, 12, 4, 58, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 5, 5, 65, 8, 5, 10, 5, 12, 5, 68, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5,
		6, 74, 8, 6, 10, 6, 12, 6, 77, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7,
		84, 8, 7, 1, 7, 1, 7, 3, 7, 88, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 4, 8, 99, 8, 8, 11, 8, 12, 8, 100, 1, 9, 1, 9, 3,
		9, 105, 8, 9, 1, 9, 1, 9, 3, 9, 109, 8, 9, 1, 10, 1, 10, 1, 10, 3, 10,
		114, 8, 10, 1, 10, 1, 10, 3, 10, 118, 8, 10, 1, 10, 0, 0, 11, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 0, 1, 1, 0, 11, 14, 132, 0, 24, 1, 0, 0,
		0, 2, 31, 1, 0, 0, 0, 4, 35, 1, 0, 0, 0, 6, 43, 1, 0, 0, 0, 8, 48, 1, 0,
		0, 0, 10, 61, 1, 0, 0, 0, 12, 69, 1, 0, 0, 0, 14, 80, 1, 0, 0, 0, 16, 98,
		1, 0, 0, 0, 18, 102, 1, 0, 0, 0, 20, 110, 1, 0, 0, 0, 22, 25, 3, 2, 1,
		0, 23, 25, 3, 8, 4, 0, 24, 22, 1, 0, 0, 0, 24, 23, 1, 0, 0, 0, 25, 26,
		1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 1, 1, 0, 0, 0,
		28, 32, 3, 20, 10, 0, 29, 32, 3, 4, 2, 0, 30, 32, 3, 6, 3, 0, 31, 28, 1,
		0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33,
		34, 5, 1, 0, 0, 34, 3, 1, 0, 0, 0, 35, 40, 5, 15, 0, 0, 36, 39, 5, 15,
		0, 0, 37, 39, 3, 16, 8, 0, 38, 36, 1, 0, 0, 0, 38, 37, 1, 0, 0, 0, 39,
		42, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 5, 1, 0, 0,
		0, 42, 40, 1, 0, 0, 0, 43, 44, 5, 18, 0, 0, 44, 45, 5, 15, 0, 0, 45, 7,
		1, 0, 0, 0, 46, 49, 3, 18, 9, 0, 47, 49, 3, 10, 5, 0, 48, 46, 1, 0, 0,
		0, 48, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 56, 5, 2, 0, 0, 51, 55,
		3, 2, 1, 0, 52, 55, 3, 8, 4, 0, 53, 55, 3, 12, 6, 0, 54, 51, 1, 0, 0, 0,
		54, 52, 1, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1,
		0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 59, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59,
		60, 5, 3, 0, 0, 60, 9, 1, 0, 0, 0, 61, 66, 5, 15, 0, 0, 62, 65, 5, 15,
		0, 0, 63, 65, 3, 16, 8, 0, 64, 62, 1, 0, 0, 0, 64, 63, 1, 0, 0, 0, 65,
		68, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 11, 1, 0, 0,
		0, 68, 66, 1, 0, 0, 0, 69, 70, 5, 4, 0, 0, 70, 71, 3, 14, 7, 0, 71, 75,
		5, 2, 0, 0, 72, 74, 3, 2, 1, 0, 73, 72, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0,
		75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 78, 1, 0, 0, 0, 77, 75, 1,
		0, 0, 0, 78, 79, 5, 3, 0, 0, 79, 13, 1, 0, 0, 0, 80, 81, 5, 5, 0, 0, 81,
		83, 5, 15, 0, 0, 82, 84, 5, 15, 0, 0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0,
		0, 0, 84, 87, 1, 0, 0, 0, 85, 88, 5, 15, 0, 0, 86, 88, 3, 16, 8, 0, 87,
		85, 1, 0, 0, 0, 87, 86, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 1, 0, 0,
		0, 89, 90, 5, 6, 0, 0, 90, 15, 1, 0, 0, 0, 91, 99, 5, 7, 0, 0, 92, 99,
		5, 8, 0, 0, 93, 99, 5, 15, 0, 0, 94, 95, 5, 5, 0, 0, 95, 96, 3, 16, 8,
		0, 96, 97, 5, 6, 0, 0, 97, 99, 1, 0, 0, 0, 98, 91, 1, 0, 0, 0, 98, 92,
		1, 0, 0, 0, 98, 93, 1, 0, 0, 0, 98, 94, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0,
		100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 17, 1, 0, 0, 0, 102, 104,
		5, 9, 0, 0, 103, 105, 5, 15, 0, 0, 104, 103, 1, 0, 0, 0, 104, 105, 1, 0,
		0, 0, 105, 108, 1, 0, 0, 0, 106, 109, 5, 15, 0, 0, 107, 109, 3, 16, 8,
		0, 108, 106, 1, 0, 0, 0, 108, 107, 1, 0, 0, 0, 109, 19, 1, 0, 0, 0, 110,
		113, 5, 10, 0, 0, 111, 114, 5, 15, 0, 0, 112, 114, 3, 16, 8, 0, 113, 111,
		1, 0, 0, 0, 113, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 117, 5, 15,
		0, 0, 116, 118, 7, 0, 0, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0,
		118, 21, 1, 0, 0, 0, 19, 24, 26, 31, 38, 40, 48, 54, 56, 64, 66, 75, 83,
		87, 98, 100, 104, 108, 113, 117,
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
	NginxParserValue           = 15
	NginxParserSTR_EXT         = 16
	NginxParserLINE_COMMENT    = 17
	NginxParserREGEXP_PREFIXED = 18
	NginxParserQUOTED_STRING   = 19
	NginxParserSINGLE_QUOTED   = 20
	NginxParserWS              = 21
)

// NginxParser rules.
const (
	NginxParserRULE_config               = 0
	NginxParserRULE_statement            = 1
	NginxParserRULE_genericStatement     = 2
	NginxParserRULE_regexHeaderStatement = 3
	NginxParserRULE_block                = 4
	NginxParserRULE_genericBlockHeader   = 5
	NginxParserRULE_ifStatement          = 6
	NginxParserRULE_ifBody               = 7
	NginxParserRULE_regexp               = 8
	NginxParserRULE_locationBlockHeader  = 9
	NginxParserRULE_rewriteStatement     = 10
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
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__8)|(1<<NginxParserT__9)|(1<<NginxParserValue)|(1<<NginxParserREGEXP_PREFIXED))) != 0) {
		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(22)
				p.Statement()
			}

		case 2:
			{
				p.SetState(23)
				p.Block()
			}

		}

		p.SetState(26)
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
	p.SetState(31)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NginxParserT__9:
		{
			p.SetState(28)
			p.RewriteStatement()
		}

	case NginxParserValue:
		{
			p.SetState(29)
			p.GenericStatement()
		}

	case NginxParserREGEXP_PREFIXED:
		{
			p.SetState(30)
			p.RegexHeaderStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(33)
		p.Match(NginxParserT__0)
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
		p.SetState(35)
		p.Match(NginxParserValue)
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__4)|(1<<NginxParserT__6)|(1<<NginxParserT__7)|(1<<NginxParserValue))) != 0 {
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(36)
				p.Match(NginxParserValue)
			}

		case 2:
			{
				p.SetState(37)
				p.Regexp()
			}

		}

		p.SetState(42)
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
		p.SetState(43)
		p.Match(NginxParserREGEXP_PREFIXED)
	}
	{
		p.SetState(44)
		p.Match(NginxParserValue)
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
	p.EnterRule(localctx, 8, NginxParserRULE_block)
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
	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NginxParserT__8:
		{
			p.SetState(46)
			p.LocationBlockHeader()
		}

	case NginxParserValue:
		{
			p.SetState(47)
			p.GenericBlockHeader()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(50)
		p.Match(NginxParserT__1)
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__3)|(1<<NginxParserT__8)|(1<<NginxParserT__9)|(1<<NginxParserValue)|(1<<NginxParserREGEXP_PREFIXED))) != 0 {
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(51)
				p.Statement()
			}

		case 2:
			{
				p.SetState(52)
				p.Block()
			}

		case 3:
			{
				p.SetState(53)
				p.IfStatement()
			}

		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(59)
		p.Match(NginxParserT__2)
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
	p.EnterRule(localctx, 10, NginxParserRULE_genericBlockHeader)
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
		p.SetState(61)
		p.Match(NginxParserValue)
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__4)|(1<<NginxParserT__6)|(1<<NginxParserT__7)|(1<<NginxParserValue))) != 0 {
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(62)
				p.Match(NginxParserValue)
			}

		case 2:
			{
				p.SetState(63)
				p.Regexp()
			}

		}

		p.SetState(68)
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
	p.EnterRule(localctx, 12, NginxParserRULE_ifStatement)
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
		p.SetState(69)
		p.Match(NginxParserT__3)
	}
	{
		p.SetState(70)
		p.IfBody()
	}
	{
		p.SetState(71)
		p.Match(NginxParserT__1)
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__9)|(1<<NginxParserValue)|(1<<NginxParserREGEXP_PREFIXED))) != 0 {
		{
			p.SetState(72)
			p.Statement()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
		p.Match(NginxParserT__2)
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
	p.EnterRule(localctx, 14, NginxParserRULE_ifBody)

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
		p.SetState(80)
		p.Match(NginxParserT__4)
	}
	{
		p.SetState(81)
		p.Match(NginxParserValue)
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(82)
			p.Match(NginxParserValue)
		}

	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(85)
			p.Match(NginxParserValue)
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(86)
			p.Regexp()
		}

	}
	{
		p.SetState(89)
		p.Match(NginxParserT__5)
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
	p.EnterRule(localctx, 16, NginxParserRULE_regexp)

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
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(98)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case NginxParserT__6:
				{
					p.SetState(91)
					p.Match(NginxParserT__6)
				}

			case NginxParserT__7:
				{
					p.SetState(92)
					p.Match(NginxParserT__7)
				}

			case NginxParserValue:
				{
					p.SetState(93)
					p.Match(NginxParserValue)
				}

			case NginxParserT__4:
				{
					p.SetState(94)
					p.Match(NginxParserT__4)
				}
				{
					p.SetState(95)
					p.Regexp()
				}
				{
					p.SetState(96)
					p.Match(NginxParserT__5)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 18, NginxParserRULE_locationBlockHeader)

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
		p.SetState(102)
		p.Match(NginxParserT__8)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(103)
			p.Match(NginxParserValue)
		}

	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(106)
			p.Match(NginxParserValue)
		}

	case 2:
		{
			p.SetState(107)
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
	p.EnterRule(localctx, 20, NginxParserRULE_rewriteStatement)
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
		p.SetState(110)
		p.Match(NginxParserT__9)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(111)
			p.Match(NginxParserValue)
		}

	case 2:
		{
			p.SetState(112)
			p.Regexp()
		}

	}
	{
		p.SetState(115)
		p.Match(NginxParserValue)
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__10)|(1<<NginxParserT__11)|(1<<NginxParserT__12)|(1<<NginxParserT__13))) != 0 {
		{
			p.SetState(116)
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
