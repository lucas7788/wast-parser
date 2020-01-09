package ast

import (
	"github.com/ontio/wast-parser/lexer"
	"github.com/ontio/wast-parser/parser"
)

type Data struct {
	implModuleField
	Name OptionId
	Kind DataKind
	Val  [][]byte
}

func (self *Data) Parse(ps *parser.ParserBuffer) error {
	err := ps.ExpectKeywordMatch("data")
	if err != nil {
		return err
	}

	self.Name.Parse(ps)
	if matchKeyword(ps.PeekToken(), "passive") {
		_ = ps.ExpectKeywordMatch("passive")
		self.Kind = DataKindPassive{}
	} else if ps.PeekToken().Type() == lexer.StringType {
		self.Kind = DataKindPassive{}
	} else {
		var memory OptionIndex
		if matchKeyword(ps.Peek2Token(), "memory") {
			err := ps.Parens(func(ps *parser.ParserBuffer) error {
				var index Index
				_ = ps.ExpectKeywordMatch("memory")
				err := index.Parse(ps)
				if err != nil {
					return err
				}
				memory = NewOptionIndex(index)
				return nil
			})
			if err != nil {
				return err
			}
		} else {
			memory.Parse(ps)
		}
		var expr Expression
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			if matchKeyword(ps.PeekToken(), "offset") {
				_ = ps.ExpectKeywordMatch("offset")
			}

			return expr.Parse(ps)
		})
		if err != nil {
			return err
		}
		self.Kind = DataKindActive{
			Memory: memory.ToIndexOr(NewNumIndex(0)),
			Offset: expr,
		}
	}

	for !ps.Empty() {
		str, err := ps.ExpectString()
		if err != nil {
			return err
		}
		self.Val = append(self.Val, []byte(str))
	}

	return nil
}

type DataKind interface {
	dataKind()
}

type implDataKind struct{}

func (self implDataKind) dataKind() {}

type DataKindPassive struct {
	implDataKind
}

type DataKindActive struct {
	implDataKind
	Memory Index
	Offset Expression
}
