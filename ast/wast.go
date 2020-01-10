package ast

import "github.com/ontio/wast-parser/parser"

type Wat struct {
	Module Module
}

func (self *Wat) Parse(ps *parser.ParserBuffer) error {
	return ps.Parens(func(ps *parser.ParserBuffer) error {
		return self.Module.Parse(ps)
	})
}

type Wast struct {
	Directives []WastDirective
}

type WastDirective interface {
	wastDirective()
}

type implWastDirective struct {}
func (self implWastDirective)  wastDirective() {}

type AssertInvalidDirective struct {
	implWastDirective
	Module Module
	Msg string
}

type AssertMalformedDirective struct {
	implWastDirective
	Module QuoteModule
	Msg string
}

type RegisterDirective struct {
	implWastDirective
	Name string
	Module OptionId
}

type WastInvokeDirective struct {
	implWastDirective
}

type  AssertTrapDirective struct {
	implWastDirective
	Exec WastExecute
	msg string
}

type AssertReturnDirective struct {
	implWastDirective
	Exec WastExecute
	Results []Expression
}

type AssertReturnCanonicalNanDirective struct {
	implWastDirective
	Invoke WastInvokeDirective
}

type WastExecute interface {
	wastExecute()
}
type implWastExecute struct {}
func (self implWastExecute) wastExecute() {}

type WastInvoke struct {
	implWastExecute
	Module OptionId
	Name string
	Args []Expression
}

func (self *WastInvoke)Parse(ps *parser.ParserBuffer ) error {
	err := ps.ExpectKeywordMatch("invoke")
	if err != nil {
		return err
	}

	self.Module.Parse(ps)
	self.Name, err = ps.ExpectString()
	if err != nil {
		return err
	}

	for !ps.Empty() {
		err := ps.Parens(func (ps *parser.ParserBuffer)error {
			var expr Expression
			err := expr.Parse(ps)
			if err != nil {
				return err
			}

			self.Args = append(self.Args, expr)
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}

type QuoteModule interface {
	quoteModule()
}

type implQuoteModule struct {}
func (self implQuoteModule) quoteModule() {}

type Quote struct {
	implQuoteModule
	Data []string
}

func parseQuoteModule(ps *parser.ParserBuffer) (QuoteModule, error) {
	if matchKeyword(ps.Peek2Token(), "quote") {
		err := ps.ExpectKeywordMatch("module")
		if err != nil {
			return nil, err
		}
		_ = ps.ExpectKeywordMatch("quote")
		var quote Quote
		for !ps.Empty() {
			str, err:= ps.ExpectString()
			if err != nil {
				return nil, err
			}
			quote.Data = append(quote.Data, str)
		}
		return quote, nil
	}

	var module Module
	err := module.Parse(ps)

	return module, err
}

func parseWastDirective(ps *parser.ParserBuffer) (WastDirective, error) {
	kw , err := ps.PeekKeyword()
	if err != nil {
		return nil, err
	}
	switch kw {
	case "module":
		var module Module
		err := module.Parse(ps)
		return module, err
	case "assert_malformed":
		_, _ = ps.ExpectKeyword()
		var result AssertMalformedDirective
		err = ps.Parens(func (ps *parser.ParserBuffer)error {
			result.Module, err = parseQuoteModule(ps)
			return err
		})
		if err != nil {
			return nil, err
		}
		result.Msg, err = ps.ExpectString()
		if err != nil {
			return nil, err
		}
		return result, nil
	case "assert_invalid":
		_, _ = ps.ExpectKeyword()
		var result AssertInvalidDirective
		err = ps.Parens(result.Module.Parse)
		if err != nil {
			return nil, err
		}
		result.Msg, err = ps.ExpectString()
		if err != nil {
			return nil, err
		}
		return result, nil
	case "register":
		_, _ = ps.ExpectKeyword()
		var result RegisterDirective
		result.Name, err = ps.ExpectString()
		if err != nil {
			return nil, err
		}
		result.Module.Parse(ps)
		return result, nil
	case "invoke":

	}

}

