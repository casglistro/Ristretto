// Code generated from .//Ristretto.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Ristretto

import "github.com/antlr4-go/antlr/v4"

type BaseRistrettoVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRistrettoVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRistrettoVisitor) VisitDoctype(ctx *DoctypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRistrettoVisitor) VisitClosedTag(ctx *ClosedTagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRistrettoVisitor) VisitOpenedTag(ctx *OpenedTagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRistrettoVisitor) VisitTagname(ctx *TagnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRistrettoVisitor) VisitContent(ctx *ContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRistrettoVisitor) VisitTagWithAttr(ctx *TagWithAttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRistrettoVisitor) VisitAttrs(ctx *AttrsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRistrettoVisitor) VisitAttr(ctx *AttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRistrettoVisitor) VisitAttrValue(ctx *AttrValueContext) interface{} {
	return v.VisitChildren(ctx)
}
