package ast

import (
	"errors"
	"fmt"
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

func (self *OptionId) Parse(ps *parser.ParserBuffer) error {
	_ =  ps.Step(func(cursor *parser.Cursor) error {
		id := cursor.Id()
		if len(id) == 0 {
			return errors.New("expect an identifier")
		}
		self.name = id
		return nil
	})

	return nil
}

type Index struct {
	isnum bool
	Num   uint32
	Id    Id
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

type Module struct {
}

type Import struct {
	Module string
	Field  string

	// optional identifier
	Id   OptionId
	Item ImportItem
}

type ImportItem interface {
	ImportType() string
}

type ImportFunc struct {
}

type ImportGlobal struct {
	Global GlobalType
}

func (self ImportGlobal) ImportType() string {
	return "global"
}

type ImportMemory struct {
	Mem MemoryType
}

func (self ImportMemory) ImportType() string {
	return "memory"
}

type ImportTable struct {
	Table TableElemType
}

func (self ImportTable)ImportType() string {
	return "table"
}

func (self *Import) Parse(ps *parser.ParserBuffer) error {
	err := ps.ExpectKeywordMatch("import")
	if err != nil {
		return err
	}

	self.Module, err = ps.ExpectString()
	if err != nil {
		return err
	}
	self.Field, err = ps.ExpectString()
	if err != nil {
		return err
	}

	err = ps.Parens(func(ps *parser.ParserBuffer) error {
		_ = self.Id.Parse(ps)
		kw, err := ps.ExpectKeyword()
		if err != nil {
			return err
		}
		switch kw {
		case "func", "table":
			panic("unimplement")
		case "memory":
			var memory MemoryType
			err := memory.Parse(ps)
			if err != nil {
				return err
			}

			self.Item = ImportMemory{Mem: memory}
		case "global":
			var global GlobalType
			err := global.Parse(ps)
			if err != nil {
				return err
			}
			self.Item = ImportGlobal{Global: global}
		default:
			return fmt.Errorf("parse import item error: unexpected keyword: %s", kw)
		}

		return nil
	})

	return err
}
