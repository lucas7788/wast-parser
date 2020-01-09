package ast

import (
	"errors"
	"github.com/ontio/wast-parser/parser"
)

type Id struct {
	Name string
}

func (self *Id) Parse(ps *parser.ParserBuffer) error {
	return ps.Step(func(cursor *parser.Cursor) error {
		id := cursor.Id()
		if len(id) == 0 {
			return errors.New("expect an identifier")
		}
		self.Name = id
		return nil
	})
}

type OptionId struct {
	name string
}

func (self *OptionId)IsSome() bool {
	return self.name != ""
}

func (self OptionId)ToId() Id {
	if !self.IsSome() {
		panic("empty option id")
	}

	return Id {Name:self.name}
}

func (self *OptionId) Parse(ps *parser.ParserBuffer)  {
	_ =  ps.Step(func(cursor *parser.Cursor) error {
		id := cursor.Id()
		if len(id) == 0 {
			return errors.New("expect an identifier")
		}
		self.name = id
		return nil
	})
}

type Index struct {
	isnum bool
	Num   uint32
	Id    Id
}

type OptionIndex struct {
	isSome bool
	index Index
}

func (self *OptionIndex)IsSome() bool {
	return self.isSome
}

func NewOptionIndex(ind Index) OptionIndex {
	return OptionIndex{
		isSome:true,
		index:ind,
	}
}

func NoneOptionIndex() OptionIndex {
	return OptionIndex{
		isSome:false,
	}
}

func (self OptionIndex)ToIndex() Index {
	if !self.isSome {
		panic("assert some index")
	}
	return self.index
}

func (self *Index) Parse(ps *parser.ParserBuffer) error {
	return ps.Step(func(cursor *parser.Cursor) error {
		c2 := cursor.Clone()
		id := c2.Id()
		if len(id) != 0 {
			self.isnum = false
			self.Id = Id{Name: id}
		}
		val, err := cursor.Integer()
		if err != nil {
			return err
		}
		num, err := val.ToUint(32)
		if err != nil {
			return err
		}
		self.isnum = true
		self.Num = uint32(num)

		return nil
	})
}

