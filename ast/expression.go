package ast

import (
	"fmt"

	"github.com/ontio/wast-parser/lexer"
	"github.com/ontio/wast-parser/parser"
)

type Expression struct {
	Instrs []Instruction
}

func (self *Expression) Parse(ps *parser.ParserBuffer) error {
	var instrs instructions
	err := instrs.parseFoldedInstrs(ps)
	if err != nil {
		return err
	}
	self.Instrs = instrs.Instrs

	return nil
}

func (self *Expression) String() string {
	str := ""
	for _, inst := range self.Instrs {
		str += inst.String() + " "
	}

	return str
}

//todo
type Instruction interface {
	parseInstrBody(ps *parser.ParserBuffer) error
	String() string
}

type Unreachable struct{}

func (self *Unreachable) parseInstrBody(ps *parser.ParserBuffer) error {
	return nil
}

func (self *Unreachable) String() string {
	return "unreachable"
}

type Nop struct{}

func (self *Nop) parseInstrBody(ps *parser.ParserBuffer) error {
	return nil
}

func (self *Nop) String() string {
	return "nop"
}

type End struct {
	Id OptionId
}

func (self *End) parseInstrBody(ps *parser.ParserBuffer) error {
	self.Id.Parse(ps)
	return nil
}

func (self *End) String() string {
	return "end"
}

type Block struct {
}

func (self *Block) parseInstrBody(ps *parser.ParserBuffer) error {
	return nil
}

func (self *Block) String() string {
	return "block"
}

type I32Const struct {
	Val uint32
}

func (self *I32Const) parseInstrBody(ps *parser.ParserBuffer) error {
	val, err := ps.ExpectUint32()
	if err != nil {
		return err
	}
	self.Val = val
	return nil
}

func (self *I32Const) String() string {
	return fmt.Sprintf("(i32.const %d)", self.Val)
}

func parseInstr(ps *parser.ParserBuffer) (Instruction, error) {
	var inst Instruction
	kw, err := ps.ExpectKeyword()
	if err != nil {
		return nil, err
	}
	switch kw {
	case "unreachable":
		inst = &Unreachable{}
	case "nop":
		inst = &Nop{}
	case "block":
		inst = &Block{}
	case "i32.const":
		inst = &I32Const{}
	default:
		panic("todo")
	}
	err = inst.parseInstrBody(ps)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

type instructions struct {
	Instrs []Instruction
}

func (self *instructions) parseFoldedInstrs(ps *parser.ParserBuffer) error {
	for !ps.Empty() {
		err := self.parseOneInstr(ps)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *instructions) parseOneInstr(ps *parser.ParserBuffer) error {
	if ps.PeekToken().Type() != lexer.LParenType {
		instr, err := parseInstr(ps)
		if err != nil {
			return err
		}
		self.Instrs = append(self.Instrs, instr)
		return nil
	}

	return ps.Parens(func(ps *parser.ParserBuffer) error {
		instr, err := parseInstr(ps)
		if err != nil {
			return err
		}
		switch val := instr.(type) {
		case *Block: //loop
			self.Instrs = append(self.Instrs, val)

			err := self.parseFoldedInstrs(ps)
			if err != nil {
				return err
			}
			self.Instrs = append(self.Instrs, &End{Id: NoneOptionId()})
		//case If : todo
		default:
			err := self.parseFoldedInstrs(ps)
			if err != nil {
				return err
			}

			self.Instrs = append(self.Instrs, val)
		}
		return nil
	})
}
