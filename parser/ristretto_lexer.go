// Code generated from .//Ristretto.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

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

type RistrettoLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var RistrettoLexerLexerStaticData struct {
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

func ristrettolexerLexerInit() {
	staticData := &RistrettoLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "';'", "','", "'='", "", "", "'![DOCTYPE]'", "'#'", "'['", "']'",
		"'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "IDENTITY", "STRING", "DOCTYPE", "NUMBERSIGN", "LSBRACKET",
		"RSBRACKET", "LPAREN", "RPAREN", "LCBRACKET", "RCBRACKET", "WS", "DIGIT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "IDENTITY", "STRING", "DOCTYPE", "NUMBERSIGN",
		"LSBRACKET", "RSBRACKET", "LPAREN", "RPAREN", "LCBRACKET", "RCBRACKET",
		"WS", "DIGIT", "LETTER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 90, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 4, 3, 42, 8, 3,
		11, 3, 12, 3, 43, 1, 4, 1, 4, 5, 4, 48, 8, 4, 10, 4, 12, 4, 51, 9, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 13, 4, 13, 81, 8, 13, 11, 13, 12, 13, 82, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 0, 0, 16, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 0, 1, 0, 4, 3, 0, 9, 10, 13, 13, 34, 34, 3, 0, 9, 10, 13, 13,
		32, 32, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 92, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 1, 33, 1, 0, 0, 0, 3, 35, 1, 0, 0,
		0, 5, 37, 1, 0, 0, 0, 7, 41, 1, 0, 0, 0, 9, 45, 1, 0, 0, 0, 11, 54, 1,
		0, 0, 0, 13, 65, 1, 0, 0, 0, 15, 67, 1, 0, 0, 0, 17, 69, 1, 0, 0, 0, 19,
		71, 1, 0, 0, 0, 21, 73, 1, 0, 0, 0, 23, 75, 1, 0, 0, 0, 25, 77, 1, 0, 0,
		0, 27, 80, 1, 0, 0, 0, 29, 86, 1, 0, 0, 0, 31, 88, 1, 0, 0, 0, 33, 34,
		5, 59, 0, 0, 34, 2, 1, 0, 0, 0, 35, 36, 5, 44, 0, 0, 36, 4, 1, 0, 0, 0,
		37, 38, 5, 61, 0, 0, 38, 6, 1, 0, 0, 0, 39, 42, 3, 29, 14, 0, 40, 42, 3,
		31, 15, 0, 41, 39, 1, 0, 0, 0, 41, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0,
		43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 8, 1, 0, 0, 0, 45, 49, 5, 34,
		0, 0, 46, 48, 8, 0, 0, 0, 47, 46, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47,
		1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0,
		52, 53, 5, 34, 0, 0, 53, 10, 1, 0, 0, 0, 54, 55, 5, 33, 0, 0, 55, 56, 5,
		91, 0, 0, 56, 57, 5, 68, 0, 0, 57, 58, 5, 79, 0, 0, 58, 59, 5, 67, 0, 0,
		59, 60, 5, 84, 0, 0, 60, 61, 5, 89, 0, 0, 61, 62, 5, 80, 0, 0, 62, 63,
		5, 69, 0, 0, 63, 64, 5, 93, 0, 0, 64, 12, 1, 0, 0, 0, 65, 66, 5, 35, 0,
		0, 66, 14, 1, 0, 0, 0, 67, 68, 5, 91, 0, 0, 68, 16, 1, 0, 0, 0, 69, 70,
		5, 93, 0, 0, 70, 18, 1, 0, 0, 0, 71, 72, 5, 40, 0, 0, 72, 20, 1, 0, 0,
		0, 73, 74, 5, 41, 0, 0, 74, 22, 1, 0, 0, 0, 75, 76, 5, 123, 0, 0, 76, 24,
		1, 0, 0, 0, 77, 78, 5, 125, 0, 0, 78, 26, 1, 0, 0, 0, 79, 81, 7, 1, 0,
		0, 80, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83,
		1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 6, 13, 0, 0, 85, 28, 1, 0, 0, 0,
		86, 87, 7, 2, 0, 0, 87, 30, 1, 0, 0, 0, 88, 89, 7, 3, 0, 0, 89, 32, 1,
		0, 0, 0, 5, 0, 41, 43, 49, 82, 1, 6, 0, 0,
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

// RistrettoLexerInit initializes any static state used to implement RistrettoLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRistrettoLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func RistrettoLexerInit() {
	staticData := &RistrettoLexerLexerStaticData
	staticData.once.Do(ristrettolexerLexerInit)
}

// NewRistrettoLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRistrettoLexer(input antlr.CharStream) *RistrettoLexer {
	RistrettoLexerInit()
	l := new(RistrettoLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &RistrettoLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Ristretto.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RistrettoLexer tokens.
const (
	RistrettoLexerT__0       = 1
	RistrettoLexerT__1       = 2
	RistrettoLexerT__2       = 3
	RistrettoLexerIDENTITY   = 4
	RistrettoLexerSTRING     = 5
	RistrettoLexerDOCTYPE    = 6
	RistrettoLexerNUMBERSIGN = 7
	RistrettoLexerLSBRACKET  = 8
	RistrettoLexerRSBRACKET  = 9
	RistrettoLexerLPAREN     = 10
	RistrettoLexerRPAREN     = 11
	RistrettoLexerLCBRACKET  = 12
	RistrettoLexerRCBRACKET  = 13
	RistrettoLexerWS         = 14
	RistrettoLexerDIGIT      = 15
)
