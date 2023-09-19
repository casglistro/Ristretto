// Code generated from .//Ristretto.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Ristretto

import "github.com/antlr4-go/antlr/v4"

// BaseRistrettoListener is a complete listener for a parse tree produced by RistrettoParser.
type BaseRistrettoListener struct{}

var _ RistrettoListener = &BaseRistrettoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRistrettoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRistrettoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRistrettoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRistrettoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseRistrettoListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseRistrettoListener) ExitStart(ctx *StartContext) {}

// EnterDoctype is called when production doctype is entered.
func (s *BaseRistrettoListener) EnterDoctype(ctx *DoctypeContext) {}

// ExitDoctype is called when production doctype is exited.
func (s *BaseRistrettoListener) ExitDoctype(ctx *DoctypeContext) {}

// EnterClosedTag is called when production closedTag is entered.
func (s *BaseRistrettoListener) EnterClosedTag(ctx *ClosedTagContext) {}

// ExitClosedTag is called when production closedTag is exited.
func (s *BaseRistrettoListener) ExitClosedTag(ctx *ClosedTagContext) {}

// EnterOpenedTag is called when production openedTag is entered.
func (s *BaseRistrettoListener) EnterOpenedTag(ctx *OpenedTagContext) {}

// ExitOpenedTag is called when production openedTag is exited.
func (s *BaseRistrettoListener) ExitOpenedTag(ctx *OpenedTagContext) {}

// EnterTagname is called when production tagname is entered.
func (s *BaseRistrettoListener) EnterTagname(ctx *TagnameContext) {}

// ExitTagname is called when production tagname is exited.
func (s *BaseRistrettoListener) ExitTagname(ctx *TagnameContext) {}

// EnterContent is called when production content is entered.
func (s *BaseRistrettoListener) EnterContent(ctx *ContentContext) {}

// ExitContent is called when production content is exited.
func (s *BaseRistrettoListener) ExitContent(ctx *ContentContext) {}

// EnterTagWithAttr is called when production tagWithAttr is entered.
func (s *BaseRistrettoListener) EnterTagWithAttr(ctx *TagWithAttrContext) {}

// ExitTagWithAttr is called when production tagWithAttr is exited.
func (s *BaseRistrettoListener) ExitTagWithAttr(ctx *TagWithAttrContext) {}

// EnterAttrs is called when production attrs is entered.
func (s *BaseRistrettoListener) EnterAttrs(ctx *AttrsContext) {}

// ExitAttrs is called when production attrs is exited.
func (s *BaseRistrettoListener) ExitAttrs(ctx *AttrsContext) {}

// EnterAttr is called when production attr is entered.
func (s *BaseRistrettoListener) EnterAttr(ctx *AttrContext) {}

// ExitAttr is called when production attr is exited.
func (s *BaseRistrettoListener) ExitAttr(ctx *AttrContext) {}

// EnterAttrValue is called when production attrValue is entered.
func (s *BaseRistrettoListener) EnterAttrValue(ctx *AttrValueContext) {}

// ExitAttrValue is called when production attrValue is exited.
func (s *BaseRistrettoListener) ExitAttrValue(ctx *AttrValueContext) {}
