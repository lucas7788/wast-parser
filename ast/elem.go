package ast

import (
	"github.com/ontio/wast-parser/lexer"
	"github.com/ontio/wast-parser/parser"
)

type Elem struct {
	implModuleField
	Name    OptionId
	Kind    ElemKind
	Payload ElemPayload

	forceNonZero bool
}

type ElemKind interface {
	elemKind()
}

type implElemKind struct{}

func (self implElemKind) elemKind() {}

type ElemKindPassive struct {
	implElemKind
}

type ElemKindActive struct {
	implElemKind
	Table  Index
	Offset Expression
}

func (self *Elem) Parse(ps *parser.ParserBuffer) error {
	err := ps.ExpectKeywordMatch("elem")
	if err != nil {
		return err
	}

	self.Name.Parse(ps)
	self.forceNonZero = false
	if ps.PeekToken().Type() == lexer.LParenType || ps.PeekUint32() {
		var table OptionIndex
		if matchKeyword(ps.Peek2Token(), "table") {
			self.forceNonZero = true
			err := ps.Parens(func(ps *parser.ParserBuffer) error {
				_ = ps.ExpectKeywordMatch("table")
				var index Index
				err := index.Parse(ps)
				if err != nil {
					return err
				}
				table = NewOptionIndex(index)
				return nil
			})
			if err != nil {
				return err
			}
		} else if ps.PeekUint32() {
			var index Index
			err := index.Parse(ps)
			if err != nil {
				return err
			}
			table = NewOptionIndex(index)
		}
		var expr Expression
		err = ps.Parens(func(ps *parser.ParserBuffer) error {
			if matchKeyword(ps.PeekToken(), "offset") {
				_ = ps.ExpectKeywordMatch("offset")
			}
			return expr.Parse(ps)
		})
		if err != nil {
			return err
		}
		self.Kind = ElemKindActive{
			Table:  table.ToIndexOr(NewNumIndex(0)),
			Offset: expr,
		}
	} else {
		self.Kind = ElemKindPassive{}
	}

	self.Payload, err = parseElemPayload(ps)
	return err
}
