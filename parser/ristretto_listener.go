// Code generated from .//Ristretto.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Ristretto

import "github.com/antlr4-go/antlr/v4"

// RistrettoListener is a complete listener for a parse tree produced by RistrettoParser.
type RistrettoListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterDoctype is called when entering the doctype production.
	EnterDoctype(c *DoctypeContext)

	// EnterClosedTag is called when entering the closedTag production.
	EnterClosedTag(c *ClosedTagContext)

	// EnterOpenedTag is called when entering the openedTag production.
	EnterOpenedTag(c *OpenedTagContext)

	// EnterTagname is called when entering the tagname production.
	EnterTagname(c *TagnameContext)

	// EnterContent is called when entering the content production.
	EnterContent(c *ContentContext)

	// EnterTagWithAttr is called when entering the tagWithAttr production.
	EnterTagWithAttr(c *TagWithAttrContext)

	// EnterAttrs is called when entering the attrs production.
	EnterAttrs(c *AttrsContext)

	// EnterAttr is called when entering the attr production.
	EnterAttr(c *AttrContext)

	// EnterAttrValue is called when entering the attrValue production.
	EnterAttrValue(c *AttrValueContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitDoctype is called when exiting the doctype production.
	ExitDoctype(c *DoctypeContext)

	// ExitClosedTag is called when exiting the closedTag production.
	ExitClosedTag(c *ClosedTagContext)

	// ExitOpenedTag is called when exiting the openedTag production.
	ExitOpenedTag(c *OpenedTagContext)

	// ExitTagname is called when exiting the tagname production.
	ExitTagname(c *TagnameContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)

	// ExitTagWithAttr is called when exiting the tagWithAttr production.
	ExitTagWithAttr(c *TagWithAttrContext)

	// ExitAttrs is called when exiting the attrs production.
	ExitAttrs(c *AttrsContext)

	// ExitAttr is called when exiting the attr production.
	ExitAttr(c *AttrContext)

	// ExitAttrValue is called when exiting the attrValue production.
	ExitAttrValue(c *AttrValueContext)
}
