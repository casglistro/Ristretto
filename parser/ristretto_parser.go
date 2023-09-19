// Code generated from .//Ristretto.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Ristretto

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

type RistrettoParser struct {
	*antlr.BaseParser
}

var RistrettoParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ristrettoParserInit() {
	staticData := &RistrettoParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "','", "'='", "", "", "'![DOCTYPE]'", "'#'", "'['", "']'",
		"'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "IDENTITY", "STRING", "DOCTYPE", "NUMBERSIGN", "LSBRACKET",
		"RSBRACKET", "LPAREN", "RPAREN", "LCBRACKET", "RCBRACKET", "WS", "DIGIT",
	}
	staticData.RuleNames = []string{
		"start", "doctype", "tag", "tagname", "content", "tagWithAttr", "attrs",
		"attr", "attrValue",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 74, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 3, 0, 20, 8, 0,
		1, 0, 5, 0, 23, 8, 0, 10, 0, 12, 0, 26, 9, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 43,
		8, 2, 10, 2, 12, 2, 46, 9, 2, 1, 2, 1, 2, 3, 2, 50, 8, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 3, 5, 58, 8, 5, 1, 6, 1, 6, 1, 6, 5, 6, 63, 8, 6,
		10, 6, 12, 6, 66, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 0, 0,
		9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 1, 1, 0, 4, 5, 71, 0, 19, 1, 0, 0,
		0, 2, 27, 1, 0, 0, 0, 4, 49, 1, 0, 0, 0, 6, 51, 1, 0, 0, 0, 8, 53, 1, 0,
		0, 0, 10, 55, 1, 0, 0, 0, 12, 59, 1, 0, 0, 0, 14, 67, 1, 0, 0, 0, 16, 71,
		1, 0, 0, 0, 18, 20, 3, 2, 1, 0, 19, 18, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0,
		20, 24, 1, 0, 0, 0, 21, 23, 3, 4, 2, 0, 22, 21, 1, 0, 0, 0, 23, 26, 1,
		0, 0, 0, 24, 22, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 1, 1, 0, 0, 0, 26,
		24, 1, 0, 0, 0, 27, 28, 5, 6, 0, 0, 28, 3, 1, 0, 0, 0, 29, 30, 5, 7, 0,
		0, 30, 31, 5, 8, 0, 0, 31, 32, 3, 10, 5, 0, 32, 33, 5, 9, 0, 0, 33, 34,
		5, 1, 0, 0, 34, 50, 1, 0, 0, 0, 35, 36, 5, 7, 0, 0, 36, 37, 5, 8, 0, 0,
		37, 38, 3, 10, 5, 0, 38, 39, 5, 9, 0, 0, 39, 44, 5, 12, 0, 0, 40, 43, 3,
		4, 2, 0, 41, 43, 3, 8, 4, 0, 42, 40, 1, 0, 0, 0, 42, 41, 1, 0, 0, 0, 43,
		46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 47, 1, 0, 0,
		0, 46, 44, 1, 0, 0, 0, 47, 48, 5, 13, 0, 0, 48, 50, 1, 0, 0, 0, 49, 29,
		1, 0, 0, 0, 49, 35, 1, 0, 0, 0, 50, 5, 1, 0, 0, 0, 51, 52, 5, 4, 0, 0,
		52, 7, 1, 0, 0, 0, 53, 54, 5, 5, 0, 0, 54, 9, 1, 0, 0, 0, 55, 57, 3, 6,
		3, 0, 56, 58, 3, 12, 6, 0, 57, 56, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58,
		11, 1, 0, 0, 0, 59, 64, 3, 14, 7, 0, 60, 61, 5, 2, 0, 0, 61, 63, 3, 14,
		7, 0, 62, 60, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65,
		1, 0, 0, 0, 65, 13, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 68, 5, 4, 0, 0,
		68, 69, 5, 3, 0, 0, 69, 70, 3, 16, 8, 0, 70, 15, 1, 0, 0, 0, 71, 72, 7,
		0, 0, 0, 72, 17, 1, 0, 0, 0, 7, 19, 24, 42, 44, 49, 57, 64,
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

// RistrettoParserInit initializes any static state used to implement RistrettoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRistrettoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RistrettoParserInit() {
	staticData := &RistrettoParserStaticData
	staticData.once.Do(ristrettoParserInit)
}

// NewRistrettoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRistrettoParser(input antlr.TokenStream) *RistrettoParser {
	RistrettoParserInit()
	this := new(RistrettoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &RistrettoParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Ristretto.g4"

	return this
}

// RistrettoParser tokens.
const (
	RistrettoParserEOF        = antlr.TokenEOF
	RistrettoParserT__0       = 1
	RistrettoParserT__1       = 2
	RistrettoParserT__2       = 3
	RistrettoParserIDENTITY   = 4
	RistrettoParserSTRING     = 5
	RistrettoParserDOCTYPE    = 6
	RistrettoParserNUMBERSIGN = 7
	RistrettoParserLSBRACKET  = 8
	RistrettoParserRSBRACKET  = 9
	RistrettoParserLPAREN     = 10
	RistrettoParserRPAREN     = 11
	RistrettoParserLCBRACKET  = 12
	RistrettoParserRCBRACKET  = 13
	RistrettoParserWS         = 14
	RistrettoParserDIGIT      = 15
)

// RistrettoParser rules.
const (
	RistrettoParserRULE_start       = 0
	RistrettoParserRULE_doctype     = 1
	RistrettoParserRULE_tag         = 2
	RistrettoParserRULE_tagname     = 3
	RistrettoParserRULE_content     = 4
	RistrettoParserRULE_tagWithAttr = 5
	RistrettoParserRULE_attrs       = 6
	RistrettoParserRULE_attr        = 7
	RistrettoParserRULE_attrValue   = 8
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Doctype() IDoctypeContext
	AllTag() []ITagContext
	Tag(i int) ITagContext

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RistrettoParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Doctype() IDoctypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoctypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoctypeContext)
}

func (s *StartContext) AllTag() []ITagContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITagContext); ok {
			len++
		}
	}

	tst := make([]ITagContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITagContext); ok {
			tst[i] = t.(ITagContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) Tag(i int) ITagContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
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

	return t.(ITagContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RistrettoVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RistrettoParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RistrettoParserRULE_start)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RistrettoParserDOCTYPE {
		{
			p.SetState(18)
			p.Doctype()
		}

	}
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RistrettoParserNUMBERSIGN {
		{
			p.SetState(21)
			p.Tag()
		}

		p.SetState(26)
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

// IDoctypeContext is an interface to support dynamic dispatch.
type IDoctypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOCTYPE() antlr.TerminalNode

	// IsDoctypeContext differentiates from other interfaces.
	IsDoctypeContext()
}

type DoctypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoctypeContext() *DoctypeContext {
	var p = new(DoctypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_doctype
	return p
}

func InitEmptyDoctypeContext(p *DoctypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_doctype
}

func (*DoctypeContext) IsDoctypeContext() {}

func NewDoctypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoctypeContext {
	var p = new(DoctypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RistrettoParserRULE_doctype

	return p
}

func (s *DoctypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DoctypeContext) DOCTYPE() antlr.TerminalNode {
	return s.GetToken(RistrettoParserDOCTYPE, 0)
}

func (s *DoctypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoctypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoctypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.EnterDoctype(s)
	}
}

func (s *DoctypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.ExitDoctype(s)
	}
}

func (s *DoctypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RistrettoVisitor:
		return t.VisitDoctype(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RistrettoParser) Doctype() (localctx IDoctypeContext) {
	localctx = NewDoctypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RistrettoParserRULE_doctype)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(RistrettoParserDOCTYPE)
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

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_tag
	return p
}

func InitEmptyTagContext(p *TagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_tag
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RistrettoParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) CopyAll(ctx *TagContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OpenedTagContext struct {
	TagContext
}

func NewOpenedTagContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpenedTagContext {
	var p = new(OpenedTagContext)

	InitEmptyTagContext(&p.TagContext)
	p.parser = parser
	p.CopyAll(ctx.(*TagContext))

	return p
}

func (s *OpenedTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenedTagContext) NUMBERSIGN() antlr.TerminalNode {
	return s.GetToken(RistrettoParserNUMBERSIGN, 0)
}

func (s *OpenedTagContext) LSBRACKET() antlr.TerminalNode {
	return s.GetToken(RistrettoParserLSBRACKET, 0)
}

func (s *OpenedTagContext) TagWithAttr() ITagWithAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagWithAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagWithAttrContext)
}

func (s *OpenedTagContext) RSBRACKET() antlr.TerminalNode {
	return s.GetToken(RistrettoParserRSBRACKET, 0)
}

func (s *OpenedTagContext) LCBRACKET() antlr.TerminalNode {
	return s.GetToken(RistrettoParserLCBRACKET, 0)
}

func (s *OpenedTagContext) RCBRACKET() antlr.TerminalNode {
	return s.GetToken(RistrettoParserRCBRACKET, 0)
}

func (s *OpenedTagContext) AllTag() []ITagContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITagContext); ok {
			len++
		}
	}

	tst := make([]ITagContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITagContext); ok {
			tst[i] = t.(ITagContext)
			i++
		}
	}

	return tst
}

func (s *OpenedTagContext) Tag(i int) ITagContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
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

	return t.(ITagContext)
}

func (s *OpenedTagContext) AllContent() []IContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IContentContext); ok {
			len++
		}
	}

	tst := make([]IContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IContentContext); ok {
			tst[i] = t.(IContentContext)
			i++
		}
	}

	return tst
}

func (s *OpenedTagContext) Content(i int) IContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContentContext); ok {
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

	return t.(IContentContext)
}

func (s *OpenedTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.EnterOpenedTag(s)
	}
}

func (s *OpenedTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.ExitOpenedTag(s)
	}
}

func (s *OpenedTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RistrettoVisitor:
		return t.VisitOpenedTag(s)

	default:
		return t.VisitChildren(s)
	}
}

type ClosedTagContext struct {
	TagContext
}

func NewClosedTagContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ClosedTagContext {
	var p = new(ClosedTagContext)

	InitEmptyTagContext(&p.TagContext)
	p.parser = parser
	p.CopyAll(ctx.(*TagContext))

	return p
}

func (s *ClosedTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClosedTagContext) NUMBERSIGN() antlr.TerminalNode {
	return s.GetToken(RistrettoParserNUMBERSIGN, 0)
}

func (s *ClosedTagContext) LSBRACKET() antlr.TerminalNode {
	return s.GetToken(RistrettoParserLSBRACKET, 0)
}

func (s *ClosedTagContext) TagWithAttr() ITagWithAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagWithAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagWithAttrContext)
}

func (s *ClosedTagContext) RSBRACKET() antlr.TerminalNode {
	return s.GetToken(RistrettoParserRSBRACKET, 0)
}

func (s *ClosedTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.EnterClosedTag(s)
	}
}

func (s *ClosedTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.ExitClosedTag(s)
	}
}

func (s *ClosedTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RistrettoVisitor:
		return t.VisitClosedTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RistrettoParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RistrettoParserRULE_tag)
	var _la int

	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewClosedTagContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Match(RistrettoParserNUMBERSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.Match(RistrettoParserLSBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)
			p.TagWithAttr()
		}
		{
			p.SetState(32)
			p.Match(RistrettoParserRSBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(33)
			p.Match(RistrettoParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewOpenedTagContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Match(RistrettoParserNUMBERSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Match(RistrettoParserLSBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.TagWithAttr()
		}
		{
			p.SetState(38)
			p.Match(RistrettoParserRSBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(RistrettoParserLCBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RistrettoParserSTRING || _la == RistrettoParserNUMBERSIGN {
			p.SetState(42)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case RistrettoParserNUMBERSIGN:
				{
					p.SetState(40)
					p.Tag()
				}

			case RistrettoParserSTRING:
				{
					p.SetState(41)
					p.Content()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(47)
			p.Match(RistrettoParserRCBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// ITagnameContext is an interface to support dynamic dispatch.
type ITagnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTITY() antlr.TerminalNode

	// IsTagnameContext differentiates from other interfaces.
	IsTagnameContext()
}

type TagnameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagnameContext() *TagnameContext {
	var p = new(TagnameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_tagname
	return p
}

func InitEmptyTagnameContext(p *TagnameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_tagname
}

func (*TagnameContext) IsTagnameContext() {}

func NewTagnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagnameContext {
	var p = new(TagnameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RistrettoParserRULE_tagname

	return p
}

func (s *TagnameContext) GetParser() antlr.Parser { return s.parser }

func (s *TagnameContext) IDENTITY() antlr.TerminalNode {
	return s.GetToken(RistrettoParserIDENTITY, 0)
}

func (s *TagnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.EnterTagname(s)
	}
}

func (s *TagnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.ExitTagname(s)
	}
}

func (s *TagnameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RistrettoVisitor:
		return t.VisitTagname(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RistrettoParser) Tagname() (localctx ITagnameContext) {
	localctx = NewTagnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RistrettoParserRULE_tagname)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(RistrettoParserIDENTITY)
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

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsContentContext differentiates from other interfaces.
	IsContentContext()
}

type ContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentContext() *ContentContext {
	var p = new(ContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_content
	return p
}

func InitEmptyContentContext(p *ContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_content
}

func (*ContentContext) IsContentContext() {}

func NewContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentContext {
	var p = new(ContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RistrettoParserRULE_content

	return p
}

func (s *ContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentContext) STRING() antlr.TerminalNode {
	return s.GetToken(RistrettoParserSTRING, 0)
}

func (s *ContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.EnterContent(s)
	}
}

func (s *ContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.ExitContent(s)
	}
}

func (s *ContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RistrettoVisitor:
		return t.VisitContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RistrettoParser) Content() (localctx IContentContext) {
	localctx = NewContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RistrettoParserRULE_content)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(RistrettoParserSTRING)
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

// ITagWithAttrContext is an interface to support dynamic dispatch.
type ITagWithAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tagname() ITagnameContext
	Attrs() IAttrsContext

	// IsTagWithAttrContext differentiates from other interfaces.
	IsTagWithAttrContext()
}

type TagWithAttrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagWithAttrContext() *TagWithAttrContext {
	var p = new(TagWithAttrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_tagWithAttr
	return p
}

func InitEmptyTagWithAttrContext(p *TagWithAttrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_tagWithAttr
}

func (*TagWithAttrContext) IsTagWithAttrContext() {}

func NewTagWithAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagWithAttrContext {
	var p = new(TagWithAttrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RistrettoParserRULE_tagWithAttr

	return p
}

func (s *TagWithAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *TagWithAttrContext) Tagname() ITagnameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagnameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagnameContext)
}

func (s *TagWithAttrContext) Attrs() IAttrsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrsContext)
}

func (s *TagWithAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagWithAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagWithAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.EnterTagWithAttr(s)
	}
}

func (s *TagWithAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.ExitTagWithAttr(s)
	}
}

func (s *TagWithAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RistrettoVisitor:
		return t.VisitTagWithAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RistrettoParser) TagWithAttr() (localctx ITagWithAttrContext) {
	localctx = NewTagWithAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RistrettoParserRULE_tagWithAttr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Tagname()
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == RistrettoParserIDENTITY {
		{
			p.SetState(56)
			p.Attrs()
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

// IAttrsContext is an interface to support dynamic dispatch.
type IAttrsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAttr() []IAttrContext
	Attr(i int) IAttrContext

	// IsAttrsContext differentiates from other interfaces.
	IsAttrsContext()
}

type AttrsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrsContext() *AttrsContext {
	var p = new(AttrsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_attrs
	return p
}

func InitEmptyAttrsContext(p *AttrsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_attrs
}

func (*AttrsContext) IsAttrsContext() {}

func NewAttrsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrsContext {
	var p = new(AttrsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RistrettoParserRULE_attrs

	return p
}

func (s *AttrsContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrsContext) AllAttr() []IAttrContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttrContext); ok {
			len++
		}
	}

	tst := make([]IAttrContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttrContext); ok {
			tst[i] = t.(IAttrContext)
			i++
		}
	}

	return tst
}

func (s *AttrsContext) Attr(i int) IAttrContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrContext); ok {
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

	return t.(IAttrContext)
}

func (s *AttrsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.EnterAttrs(s)
	}
}

func (s *AttrsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.ExitAttrs(s)
	}
}

func (s *AttrsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RistrettoVisitor:
		return t.VisitAttrs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RistrettoParser) Attrs() (localctx IAttrsContext) {
	localctx = NewAttrsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RistrettoParserRULE_attrs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Attr()
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RistrettoParserT__1 {
		{
			p.SetState(60)
			p.Match(RistrettoParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.Attr()
		}

		p.SetState(66)
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

// IAttrContext is an interface to support dynamic dispatch.
type IAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTITY() antlr.TerminalNode
	AttrValue() IAttrValueContext

	// IsAttrContext differentiates from other interfaces.
	IsAttrContext()
}

type AttrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrContext() *AttrContext {
	var p = new(AttrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_attr
	return p
}

func InitEmptyAttrContext(p *AttrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_attr
}

func (*AttrContext) IsAttrContext() {}

func NewAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrContext {
	var p = new(AttrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RistrettoParserRULE_attr

	return p
}

func (s *AttrContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrContext) IDENTITY() antlr.TerminalNode {
	return s.GetToken(RistrettoParserIDENTITY, 0)
}

func (s *AttrContext) AttrValue() IAttrValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrValueContext)
}

func (s *AttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.EnterAttr(s)
	}
}

func (s *AttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.ExitAttr(s)
	}
}

func (s *AttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RistrettoVisitor:
		return t.VisitAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RistrettoParser) Attr() (localctx IAttrContext) {
	localctx = NewAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RistrettoParserRULE_attr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(RistrettoParserIDENTITY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(RistrettoParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.AttrValue()
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

// IAttrValueContext is an interface to support dynamic dispatch.
type IAttrValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTITY() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsAttrValueContext differentiates from other interfaces.
	IsAttrValueContext()
}

type AttrValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrValueContext() *AttrValueContext {
	var p = new(AttrValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_attrValue
	return p
}

func InitEmptyAttrValueContext(p *AttrValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RistrettoParserRULE_attrValue
}

func (*AttrValueContext) IsAttrValueContext() {}

func NewAttrValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrValueContext {
	var p = new(AttrValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RistrettoParserRULE_attrValue

	return p
}

func (s *AttrValueContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrValueContext) IDENTITY() antlr.TerminalNode {
	return s.GetToken(RistrettoParserIDENTITY, 0)
}

func (s *AttrValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(RistrettoParserSTRING, 0)
}

func (s *AttrValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.EnterAttrValue(s)
	}
}

func (s *AttrValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RistrettoListener); ok {
		listenerT.ExitAttrValue(s)
	}
}

func (s *AttrValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RistrettoVisitor:
		return t.VisitAttrValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RistrettoParser) AttrValue() (localctx IAttrValueContext) {
	localctx = NewAttrValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RistrettoParserRULE_attrValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RistrettoParserIDENTITY || _la == RistrettoParserSTRING) {
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
