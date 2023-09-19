// Code generated from .//Ristretto.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Ristretto

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by RistrettoParser.
type RistrettoVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RistrettoParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by RistrettoParser#doctype.
	VisitDoctype(ctx *DoctypeContext) interface{}

	// Visit a parse tree produced by RistrettoParser#closedTag.
	VisitClosedTag(ctx *ClosedTagContext) interface{}

	// Visit a parse tree produced by RistrettoParser#openedTag.
	VisitOpenedTag(ctx *OpenedTagContext) interface{}

	// Visit a parse tree produced by RistrettoParser#tagname.
	VisitTagname(ctx *TagnameContext) interface{}

	// Visit a parse tree produced by RistrettoParser#content.
	VisitContent(ctx *ContentContext) interface{}

	// Visit a parse tree produced by RistrettoParser#tagWithAttr.
	VisitTagWithAttr(ctx *TagWithAttrContext) interface{}

	// Visit a parse tree produced by RistrettoParser#attrs.
	VisitAttrs(ctx *AttrsContext) interface{}

	// Visit a parse tree produced by RistrettoParser#attr.
	VisitAttr(ctx *AttrContext) interface{}

	// Visit a parse tree produced by RistrettoParser#attrValue.
	VisitAttrValue(ctx *AttrValueContext) interface{}
}
