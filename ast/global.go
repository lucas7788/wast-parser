package ast

import "github.com/ontio/wast-parser/parser"

type Global struct {
	Name    OptionId
	Exports InlineExport
	ValType GlobalValType
	Kind    GlobalType
}

type GlobalType interface {
	globalType()
}

type implGobalType struct {}
func (self implGobalType) globalType() {}

type GlobalTypeImport struct {
	implGobalType
	Module string
	Field string
}

type GlobalTypeInline struct {
	Expr string // todo
}

func (self *Global) Parse(ps *parser.ParserBuffer) error {
	err := ps.ExpectKeywordMatch("global")
	if err != nil {
		return err
	}
	self.Name.Parse(ps)

	err = self.Exports.Parse(ps)
	if err != nil {
		return err
	}

	token := ps.Peek2Token()
	if matchKeyword(token, "import") {
		var imp GlobalTypeImport
		err = ps.Parens(func (ps *parser.ParserBuffer)error {
			err := ps.ExpectKeywordMatch("import")
			if err != nil {
				panic(err)
			}
			imp.Module, err = ps.ExpectString()
			if err != nil {
				return err
			}
			imp.Field, err = ps.ExpectString()
			return err
		})
		if err != nil {
			return err
		}

		err = self.ValType.Parse(ps)
		if err != nil {
			return err
		}

		self.Kind = imp
		return nil
	} else {
		//todo: parse inline export
		panic("unimplemented")
	}
}
