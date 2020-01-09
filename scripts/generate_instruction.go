package main

import (
	"errors"
	"fmt"

	"github.com/ontio/wast-parser/lexer"
	"github.com/ontio/wast-parser/parser"
	"github.com/valyala/fasttemplate"
)

type Instruction struct {
	Name   string  `json:"name"`
	Id     string  `json:"id"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func expectKeyword(ps *parser.ParserBuffer) (string, error) {
	token := ps.ReadToken()
	if token == nil {
		return "", errors.New("expect keyword")
	}
	switch val := token.(type) {
	case lexer.Keyword:
		return val.Val, nil
	case lexer.Reserved:
		return val.Val, nil
	default:
		return "", errors.New("expect keyword")
	}
}

func (self *Instruction) Parse(ps *parser.ParserBuffer) error {
	kw, err := expectKeyword(ps)
	if err != nil {
		return err
	}
	self.Name = kw
	self.Id, err = expectKeyword(ps)
	if err != nil {
		return err
	}

	for !ps.Empty() {
		err = ps.Parens(func(ps *parser.ParserBuffer) error {
			field, err := expectKeyword(ps)
			if err != nil {
				return err
			}
			ty, err := expectKeyword(ps)
			if err != nil {
				return err
			}

			self.Fields = append(self.Fields, Field{Name: field, Type: ty})
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func generate(template string, m map[string]interface{}) string {
	t := fasttemplate.New(template, "[", "]")
	return t.ExecuteString(m)
}

func (self Instruction) Generate() string {
	template := `
type [Name] struct{}

func (self *[Name]) parseInstrBody(ps *parser.ParserBuffer) error {
	[parseBody]
	return nil
}

func (self *[Name]) String() string {
	return "[Id]"
}
`
	return generate(template, map[string]interface{}{
		"Name":      self.Name,
		"parseBody": self.generateParseBody(),
		"Id":        self.Id,
	})
}

func (self Instruction) generateParseBody() string {
	body := ""
	for _, field := range self.Fields {
		switch field.Type {
		case "uint32":
			body += parseUint32(field.Name)
		default:
			body += parseGeneral(field.Name)
		}
	}

	return body
}

func parseGeneral(name string) string {
	return generate(`
	err := self.[Id].Parse(ps)
	if err != nil {
		return err
	}
`, map[string]interface{}{"Id": name})
}

func parseUint32(name string) string {
	return generate(`val, err := ps.ExpectUint32()
	if err != nil {
		return err
	}
	self.[Id] = val
`, map[string]interface{}{"Id": name})
}

func mustParseInstrs(source string) []Instruction {
	ps, err := parser.NewParserBuffer(source)
	if err != nil {
		panic(err)
	}
	var instrs []Instruction
	for !ps.Empty() {
		var inst Instruction
		err := ps.Parens(func (ps *parser.ParserBuffer)error {
			return inst.Parse(ps)
		})
		if err != nil {
			panic(err)
		}
		instrs = append(instrs, inst)
	}

	return instrs
}

func main() {
	instrs := `
(I32Const i32.const (Val uint32))
(Unreachable unreachable)
(Nop nop)
(Br br (Index Index))
(BrIf br_if (Index Index))
(BrTable br_table (Indices BrTableIndices)) 
(Return return)
`

	allInstrs := mustParseInstrs(instrs)
	all := ""
	for _, ins := range allInstrs {
		all += ins.Generate()
	}

	fmt.Println(all)
}
