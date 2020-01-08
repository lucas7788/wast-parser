package ast

import (
	"errors"
	"fmt"
	"github.com/ontio/wast-parser/lexer"
	"strconv"

	"github.com/ontio/wast-parser/parser"
)

var I32 = ValType{ty: 0}
var I64 = ValType{ty: 1}
var F32 = ValType{ty: 2}
var F64 = ValType{ty: 3}
var Anyref = ValType{ty: 4}
var Funcref = ValType{ty: 5}
var V128 = ValType{ty: 6}

type ValType struct {
	ty byte
}

func (self *ValType) Parse(ps *parser.ParserBuffer) error {
	kw, err := ps.ExpectKeyword()
	if err != nil {
		return err
	}

	switch kw {
	case "i32":
		*self = I32
	case "i64":
		*self = I64
	case "f32":
		*self = F32
	case "f64":
		*self = F64
	case "anyref":
		*self = Anyref
	case "funcref":
		*self = Funcref
	case "anyfunc":
		*self = Funcref
	case "v128":
		*self = V128
	default:
		return fmt.Errorf("parse valtype error, unexpected keyword: %s", kw)
	}

	return nil
}

type GlobalType struct {
	Type    ValType
	Mutable bool
}

func (self *GlobalType) Parse(ps *parser.ParserBuffer) error {
	//todo: parse mutable case
	err := self.Type.Parse(ps)
	if err != nil {
		return err
	}
	self.Mutable = false

	return nil
}

type Limits struct {
	Min uint32
	//opional 0 if None
	Max uint32
}

func (self *Limits) Parse(ps *parser.ParserBuffer) error {
	val, err := ps.ExpectInteger()
	if err != nil {
		return err
	}
	min, err := val.ToUint(32)
	if err != nil {
		return err
	}
	self.Min = uint32(min)


	val, err = ps.ExpectInteger()
	if err == nil {
		max, err := val.ToUint(32)
		if err != nil {
			return err
		}
		self.Max = uint32(max)
	}

	return nil
}


type MemoryType struct {
	Limits Limits
	Shared bool
}

func (self *MemoryType) Parse(ps *parser.ParserBuffer) error {
	err := self.Limits.Parse(ps)
	if err != nil {
		return err
	}

	err = ps.ExpectKeywordMatch("shared")
	if err == nil {
		self.Shared = true
	}

	return nil
}

type TableElemType struct {
	ty byte

}

var FuncRef = TableElemType{ty: 0}
var AnyRef = TableElemType{ty: 1}

func (self *TableElemType) Parse(ps *parser.ParserBuffer) error {
	kw, err := ps.ExpectKeyword()
	if err != nil {
		return err
	}

	switch kw {
	case "anyfunc":
		*self = FuncRef
	default:
		//todo: check impl
		panic("unimplement")
	}

	return nil
}

type TableType struct {
	Limits Limits
	Elem TableElemType
}

func (self *TableType) Parse(ps *parser.ParserBuffer) error {
	err := self.Limits.Parse(ps)
	if err != nil {
		return err
	}

	return self.Elem.Parse(ps)
}

type FunctionType struct {
	Params []FuncParam
	Results []ValType
}

type FuncParam struct {
	Id OptionId
	Val ValType
}


func (self *FunctionType)Parse(ps *parser.ParserBuffer) error {
	err := ps.ExpectKeywordMatch("func")
	if err != nil {
		return err
	}
	for {
		token := ps.PeekToken()
		if token != nil && token.Type() == lexer.LParenType {
			err := ps.Parens(func (ps *parser.ParserBuffer) error {
				kw, err := ps.ExpectKeyword()
				if err != nil {
					return err
				}
				switch kw {
				case "param":
					if len(self.Results) > 0 {
						return errors.New("result before parameter")
					}


				}

			})

			if err != nil {
				return err
			}
		}
	}

}