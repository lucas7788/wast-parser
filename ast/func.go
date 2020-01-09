package ast

import "github.com/ontio/wast-parser/parser"

type Func struct {
	implModuleField
	Name    OptionId
	Exports InlineExport
	Kind    FuncKind
	Type    TypeUse
}

type FuncKind interface {
	funcKind()
}

type implFuncKind struct{}

func (self implFuncKind) funcKind() {}

type FuncKindImport struct {
	implFuncKind
	Module string
	Name   string
}

type FuncKindInline struct {
	implFuncKind
	Locals []Local
	Expr   Expression
}

type Local struct {
	Id      OptionId
	ValType ValType
}

func (self *Func) Parse(ps *parser.ParserBuffer) error {
	err := ps.ExpectKeywordMatch("func")
	if err != nil {
		return err
	}
	self.Name.Parse(ps)
	err = self.Exports.Parse(ps)
	if err != nil {
		return err
	}

	if matchKeyword(ps.Peek2Token(), "import") {
		var module, name string
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			err := ps.ExpectKeywordMatch("import")
			if err != nil {
				panic(err)
			}

			module, err = ps.ExpectString()
			if err != nil {
				return err
			}

			name, err = ps.ExpectString()
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}

		err = self.Type.Parse(ps)
		if err != nil {
			return err
		}

		self.Kind = FuncKindImport{Module: module, Name: name}
		return nil
	}
	err = self.Type.Parse(ps)
	if err != nil {
		return err
	}
	var locals []Local
	for matchKeyword(ps.Peek2Token(), "local") {
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			err := ps.ExpectKeywordMatch("local")
			if err != nil {
				panic(err)
			}

			if ps.Empty() {
				return nil
			}

			var local Local
			local.Id.Parse(ps)
			err = local.ValType.Parse(ps)
			if err != nil {
				return err
			}
			more := !local.Id.IsSome()
			locals = append(locals, local)
			for more && !ps.Empty() {
				local := Local{Id: NoneOptionId()}
				err := local.ValType.Parse(ps)
				if err != nil {
					return err
				}
				locals = append(locals, local)
			}

			return nil
		})

		if err != nil {
			return err
		}
	}

	var expr Expression
	err = expr.Parse(ps)
	if err != nil {
		return err
	}

	self.Kind = FuncKindInline{Locals: locals, Expr: expr}
	return nil
}
