package ast

import (
	"errors"
	"fmt"
	"github.com/ontio/wast-parser/parser"
	"strconv"
	"strings"
)

type Id struct {
	Name string
}

func (self *Id) Parse(ps *parser.ParserBuffer) error {
	id := ps.TryGetId()
	if len(id) == 0 {
		return errors.New("expect an identifier")
	}
	self.Name = id
	return nil
}

type OptionId struct {
	name Id
}

func NoneOptionId() OptionId {
	return OptionId{}
}

func (self *OptionId) IsSome() bool {
	return self.name.Name != ""
}

func (self OptionId) ToId() Id {
	if !self.IsSome() {
		panic("empty option id")
	}

	return self.name
}

func (self *OptionId) Parse(ps *parser.ParserBuffer) {
	_ = ps.TryParse(&self.name)
}

type Index struct {
	isnum bool
	Num   uint32
	Id    Id
}

func NewNumIndex(num uint32) Index {
	return Index{
		isnum: true,
		Num:   num,
	}
}

type OptionIndex struct {
	isSome bool
	index  Index
}

func (self *OptionIndex) Parse(ps *parser.ParserBuffer) {
	self.isSome = false
	if ps.TryParse(&self.index) == nil {
		self.isSome = true
	}
}

func (self *OptionIndex) IsSome() bool {
	return self.isSome
}

func NewOptionIndex(ind Index) OptionIndex {
	return OptionIndex{
		isSome: true,
		index:  ind,
	}
}

func NoneOptionIndex() OptionIndex {
	return OptionIndex{
		isSome: false,
	}
}

func (self OptionIndex) ToIndex() Index {
	if !self.isSome {
		panic("assert some index")
	}
	return self.index
}

func (self OptionIndex) ToIndexOr(ind Index) Index {
	if !self.isSome {
		return ind
	}
	return self.index
}

func (self *Index) Parse(ps *parser.ParserBuffer) error {
	id := ps.TryGetId()
	if len(id) != 0 {
		self.isnum = false
		self.Id = Id{Name: id}
		return nil
	}

	num, err := ps.ExpectUint32()
	if err != nil {
		return err
	}
	self.isnum = true
	self.Num = uint32(num)

	return nil
}

type Float32 struct {
	bits uint32
}

func (self *Float32) Parse(ps *parser.ParserBuffer) error {
	res, err := ps.Float()
	if err != nil {
		return err
	}
	f, err := res.ToFloat32()
	if err != nil {
		return err
	}
	self.bits = f
	return nil
}

type Float64 struct {
	bits uint64
	val  float64
}

func (self *Float64) Parse(ps *parser.ParserBuffer) error {
	res, err := ps.Float()
	if err != nil {
		return err
	}
	f, err := res.ToFloat()
	if err != nil {
		return err
	}
	self.val = f
	return nil
}

type BlockType struct {
	Label Id
	Ty    TypeUse
}

func (self *BlockType) Parse(ps *parser.ParserBuffer) error {
	var id Id
	err := id.Parse(ps)
	if err != nil {
		return err
	}
	self.Label = id
	ty := TypeUse{}
	err = ty.ParseNoNames(ps)
	if err != nil {
		return err
	}
	self.Ty = ty
	return nil
}

type MemArg struct {
	Align  uint32
	Offset uint32
}

func (self *MemArg) Parse(ps *parser.ParserBuffer, defaultAlign uint32) error {

	parse_field := func(name string, ps *parser.ParserBuffer) (uint32, error) {
		kw, err := ps.TryKeyword()
		if err != nil {
			return 0, err
		}
		fmt.Printf("kw: %s, name: %s \n", kw, name)
		if strings.HasPrefix(kw, name) == false {
			return ps.Curr(), nil
		}
		kw = kw[len(name):]
		if strings.HasPrefix(kw, "=") == false {
			return ps.Curr(), nil
		}
		kw = kw[1:]
		base := 10
		if strings.HasPrefix(kw, "0x") {
			base = 16
			kw = kw[2:]
		}
		val, err := strconv.ParseUint(kw, base, 32)
		if err != nil {
			return 0, err
		}
		return uint32(val), nil
	}

	offset, err := parse_field("offset", ps)
	if err != nil {
		return err
	}
	self.Offset = offset
	temp, err := parse_field("align", ps)
	if err != nil {
		self.Align = defaultAlign
		return nil
	}
	if !isTwoPower(temp) {
		fmt.Println("temp:", temp)
		return fmt.Errorf("alignment must be a power of two, %d", temp)
	}
	self.Align = temp
	return nil
}

func isTwoPower(num uint32) bool {
	return num&(num-1) == 0
}
