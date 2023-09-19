package syntax

import (
	"Ristretto/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/liyue201/gostl/ds/stack"
	"strings"
)

type RistrettoVisitor struct {
	*parser.BaseRistrettoVisitor
	S *stack.Stack[int]
}

func (v *RistrettoVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *RistrettoVisitor) VisitAttr(ctx *parser.AttrContext) interface{} {
	return fmt.Sprintf("%v=%v", ctx.IDENTITY().GetText(),
		ctx.AttrValue().GetText())
}

func (v *RistrettoVisitor) VisitAttrs(ctx *parser.AttrsContext) interface{} {
	attrs := make([]string, 0)

	for _, e := range ctx.AllAttr() {
		attr, _ := v.Visit(e).(string)
		attrs = append(attrs, attr)
	}

	return strings.Join(attrs, " ")
}

func (v *RistrettoVisitor) VisitTagWithAttr(ctx *parser.TagWithAttrContext) interface{} {
	tagname := ctx.Tagname().GetText()

	if ctx.Attrs() == nil {
		return fmt.Sprintf("%v", tagname)
	} else {
		attrs, _ := v.Visit(ctx.Attrs()).(string)

		return fmt.Sprintf("%v %v", tagname,
			attrs)
	}
}

func (v *RistrettoVisitor) VisitDoctype(_ *parser.DoctypeContext) interface{} {
	return "<!DOCTYPE html>"
}

func (v *RistrettoVisitor) VisitContent(ctx *parser.ContentContext) interface{} {
	level := v.S.Pop()

	return strings.Repeat("  ", level-1) + ctx.GetText()[1:len(ctx.GetText())-1]
}

func (v *RistrettoVisitor) VisitClosedTag(ctx *parser.ClosedTagContext) interface{} {
	level := v.S.Pop()

	tag, _ := v.Visit(ctx.TagWithAttr()).(string)

	return strings.Repeat("  ", level-1) + fmt.Sprintf("<%v />", tag)
}

func (v *RistrettoVisitor) VisitOpenedTag(ctx *parser.OpenedTagContext) interface{} {
	level := v.S.Pop()

	tagname := ctx.TagWithAttr().Tagname().GetText()

	tagWithAttr, _ := v.Visit(ctx.TagWithAttr()).(string)

	children := make([]string, 0)

	for _, e := range ctx.GetChildren() {
		childString := ""

		if child, ok := e.(parser.IContentContext); ok {
			v.S.Push(level + 1)
			childString, _ = v.Visit(child).(string)
		}

		if child, ok := e.(parser.ITagContext); ok {
			v.S.Push(level + 1)
			childString, _ = v.Visit(child).(string)
		}

		if childString != "" {
			children = append(children, childString)
		}
	}

	if len(children) == 0 {
		return strings.Repeat("  ", level-1) + fmt.Sprintf("<%v></%v>", tagWithAttr, tagname)
	} else {
		return strings.Repeat("  ", level-1) + fmt.Sprintf("<%v>\n%v\n", tagWithAttr,
			strings.Join(children, "\n")) + strings.Repeat("  ", level-1) + fmt.Sprintf("</%v>", tagname)
	}
}

func (v *RistrettoVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	doctype := ""

	if ctx.Doctype() != nil {
		doctype, _ = v.Visit(ctx.Doctype()).(string)
	}

	tags := make([]string, 0)

	for _, e := range ctx.AllTag() {
		v.S.Push(1)
		tag, _ := v.Visit(e).(string)
		tags = append(tags, tag)
	}

	return fmt.Sprintf("%v\n%v", doctype, strings.Join(tags, "\n"))
}
