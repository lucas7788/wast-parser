
package ast

import "github.com/ontio/wast-parser/parser"


type Unreachable struct{ 
	
}

func (self *Unreachable) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *Unreachable) String() string {
	return "unreachable"
}

type Nop struct{ 
	
}

func (self *Nop) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *Nop) String() string {
	return "nop"
}

type Br struct{ 
	Index Index
}

func (self *Br) parseInstrBody(ps *parser.ParserBuffer) error { 
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *Br) String() string {
	return "br"
}

type BrIf struct{ 
	Index Index
}

func (self *BrIf) parseInstrBody(ps *parser.ParserBuffer) error { 
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *BrIf) String() string {
	return "br_if"
}

type BrTable struct{ 
	Indices BrTableIndices
}

func (self *BrTable) parseInstrBody(ps *parser.ParserBuffer) error { 
	err := self.Indices.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *BrTable) String() string {
	return "br_table"
}

type Return struct{ 
	
}

func (self *Return) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *Return) String() string {
	return "return"
}

type Call struct{ 
	Index Index
}

func (self *Call) parseInstrBody(ps *parser.ParserBuffer) error { 
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *Call) String() string {
	return "call"
}

type CallIndirect struct{ 
	Impl CallIndirectInner
}

func (self *CallIndirect) parseInstrBody(ps *parser.ParserBuffer) error { 
	err := self.Impl.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *CallIndirect) String() string {
	return "call_indirect"
}

type ReturnCall struct{ 
	Index Index
}

func (self *ReturnCall) parseInstrBody(ps *parser.ParserBuffer) error { 
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *ReturnCall) String() string {
	return "return_call"
}

type ReturnCallIndirect struct{ 
	Impl CallIndirectInner
}

func (self *ReturnCallIndirect) parseInstrBody(ps *parser.ParserBuffer) error { 
	err := self.Impl.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *ReturnCallIndirect) String() string {
	return "return_call_indirect"
}

type Drop struct{ 
	
}

func (self *Drop) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *Drop) String() string {
	return "drop"
}

type Select struct{ 
	SelectTypes SelectTypes
}

func (self *Select) parseInstrBody(ps *parser.ParserBuffer) error { 
	err := self.SelectTypes.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *Select) String() string {
	return "select"
}

type TableGet struct{ 
	Index Index
}

func (self *TableGet) parseInstrBody(ps *parser.ParserBuffer) error { 
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *TableGet) String() string {
	return "table.get"
}

type TableSet struct{ 
	Index Index
}

func (self *TableSet) parseInstrBody(ps *parser.ParserBuffer) error { 
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *TableSet) String() string {
	return "table.set"
}

type RefNull struct{ 
	
}

func (self *RefNull) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *RefNull) String() string {
	return "ref.null"
}

type RefIsNull struct{ 
	
}

func (self *RefIsNull) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *RefIsNull) String() string {
	return "ref.is_null"
}

type RefHost struct{ 
	Val uint32
}

func (self *RefHost) parseInstrBody(ps *parser.ParserBuffer) error { 
	val, err := ps.ExpectUint32()
	if err != nil {
		return err
	}
	self.Val = val

	return nil
}

func (self *RefHost) String() string {
	return "ref.host"
}

type RefFunc struct{ 
	Index Index
}

func (self *RefFunc) parseInstrBody(ps *parser.ParserBuffer) error { 
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *RefFunc) String() string {
	return "ref.func"
}

type I32Const struct{ 
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
	return "i32.const"
}

type I64Const struct{ 
	Val int64
}

func (self *I64Const) parseInstrBody(ps *parser.ParserBuffer) error { 
	val, err := ps.ExpectInt64()
	if err != nil {
		return err
	}
	self.Val = val

	return nil
}

func (self *I64Const) String() string {
	return "i64.const"
}

type F32Const struct{ 
	Val Float32
}

func (self *F32Const) parseInstrBody(ps *parser.ParserBuffer) error { 
	err := self.Val.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *F32Const) String() string {
	return "f32.const"
}

type F64Const struct{ 
	Val Float64
}

func (self *F64Const) parseInstrBody(ps *parser.ParserBuffer) error { 
	err := self.Val.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *F64Const) String() string {
	return "f64.const"
}

type I32Clz struct{ 
	
}

func (self *I32Clz) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Clz) String() string {
	return "i32.clz"
}

type I32Ctz struct{ 
	
}

func (self *I32Ctz) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Ctz) String() string {
	return "i32.ctz"
}

type I32Pocnt struct{ 
	
}

func (self *I32Pocnt) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Pocnt) String() string {
	return "i32.popcnt"
}

type I32Ad struct{ 
	
}

func (self *I32Ad) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Ad) String() string {
	return "i32.add"
}

type I32Sub struct{ 
	
}

func (self *I32Sub) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Sub) String() string {
	return "i32.sub"
}

type I32Mul struct{ 
	
}

func (self *I32Mul) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Mul) String() string {
	return "i32.mul"
}

type I32DivS struct{ 
	
}

func (self *I32DivS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32DivS) String() string {
	return "i32.div_s"
}

type I32DivU struct{ 
	
}

func (self *I32DivU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32DivU) String() string {
	return "i32.div_u"
}

type I32RemS struct{ 
	
}

func (self *I32RemS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32RemS) String() string {
	return "i32.rem_s"
}

type I32RemU struct{ 
	
}

func (self *I32RemU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32RemU) String() string {
	return "i32.rem_u"
}

type I32And struct{ 
	
}

func (self *I32And) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32And) String() string {
	return "i32.and"
}

type I32Or struct{ 
	
}

func (self *I32Or) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Or) String() string {
	return "i32.or"
}

type I32Xor struct{ 
	
}

func (self *I32Xor) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Xor) String() string {
	return "i32.xor"
}

type I32Shl struct{ 
	
}

func (self *I32Shl) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Shl) String() string {
	return "i32.shl"
}

type I32ShrS struct{ 
	
}

func (self *I32ShrS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32ShrS) String() string {
	return "i32.shr_s"
}

type I32ShrU struct{ 
	
}

func (self *I32ShrU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32ShrU) String() string {
	return "i32.shr_u"
}

type I32Rotl struct{ 
	
}

func (self *I32Rotl) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Rotl) String() string {
	return "i32.rotl"
}

type I32Rotr struct{ 
	
}

func (self *I32Rotr) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Rotr) String() string {
	return "i32.rotr"
}

type I64Clz struct{ 
	
}

func (self *I64Clz) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Clz) String() string {
	return "i64.clz"
}

type I64Ctz struct{ 
	
}

func (self *I64Ctz) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Ctz) String() string {
	return "i64.ctz"
}

type I64Popcnt struct{ 
	
}

func (self *I64Popcnt) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Popcnt) String() string {
	return "i64.popcnt"
}

type I64Add struct{ 
	
}

func (self *I64Add) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Add) String() string {
	return "i64.add"
}

type I64Sub struct{ 
	
}

func (self *I64Sub) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Sub) String() string {
	return "i64.sub"
}

type I64Mul struct{ 
	
}

func (self *I64Mul) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Mul) String() string {
	return "i64.mul"
}

type I64DivS struct{ 
	
}

func (self *I64DivS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64DivS) String() string {
	return "i64.div_s"
}

type I64DivU struct{ 
	
}

func (self *I64DivU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64DivU) String() string {
	return "i64.div_u"
}

type I64RemS struct{ 
	
}

func (self *I64RemS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64RemS) String() string {
	return "i64.rem_s"
}

type I64RemU struct{ 
	
}

func (self *I64RemU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64RemU) String() string {
	return "i64.rem_u"
}

type I64And struct{ 
	
}

func (self *I64And) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64And) String() string {
	return "i64.and"
}

type I64Or struct{ 
	
}

func (self *I64Or) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Or) String() string {
	return "i64.or"
}

type I64Xor struct{ 
	
}

func (self *I64Xor) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Xor) String() string {
	return "i64.xor"
}

type I64Shl struct{ 
	
}

func (self *I64Shl) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Shl) String() string {
	return "i64.shl"
}

type I64ShrS struct{ 
	
}

func (self *I64ShrS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64ShrS) String() string {
	return "i64.shr_s"
}

type I64ShrU struct{ 
	
}

func (self *I64ShrU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64ShrU) String() string {
	return "i64.shr_u"
}

type I64Rotl struct{ 
	
}

func (self *I64Rotl) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Rotl) String() string {
	return "i64.rotl"
}

type I64Rotr struct{ 
	
}

func (self *I64Rotr) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Rotr) String() string {
	return "i64.rotr"
}

type F32Abs struct{ 
	
}

func (self *F32Abs) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Abs) String() string {
	return "f32.abs"
}

type F32Neg struct{ 
	
}

func (self *F32Neg) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Neg) String() string {
	return "f32.neg"
}

type F32Ceil struct{ 
	
}

func (self *F32Ceil) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Ceil) String() string {
	return "f32.ceil"
}

type F32Floor struct{ 
	
}

func (self *F32Floor) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Floor) String() string {
	return "f32.floor"
}

type F32Trunc struct{ 
	
}

func (self *F32Trunc) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Trunc) String() string {
	return "f32.trunc"
}

type F32Nearest struct{ 
	
}

func (self *F32Nearest) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Nearest) String() string {
	return "f32.nearest"
}

type F32Sqrt struct{ 
	
}

func (self *F32Sqrt) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Sqrt) String() string {
	return "f32.sqrt"
}

type F32Add struct{ 
	
}

func (self *F32Add) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Add) String() string {
	return "f32.add"
}

type F32Sub struct{ 
	
}

func (self *F32Sub) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Sub) String() string {
	return "f32.sub"
}

type F32Mul struct{ 
	
}

func (self *F32Mul) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Mul) String() string {
	return "f32.mul"
}

type F32Div struct{ 
	
}

func (self *F32Div) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Div) String() string {
	return "f32.div"
}

type F32Min struct{ 
	
}

func (self *F32Min) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Min) String() string {
	return "f32.min"
}

type F32Max struct{ 
	
}

func (self *F32Max) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Max) String() string {
	return "f32.max"
}

type F32Copysign struct{ 
	
}

func (self *F32Copysign) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Copysign) String() string {
	return "f32.copysign"
}

type F64Abs struct{ 
	
}

func (self *F64Abs) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Abs) String() string {
	return "f64.abs"
}

type F64Neg struct{ 
	
}

func (self *F64Neg) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Neg) String() string {
	return "f64.neg"
}

type F64Ceil struct{ 
	
}

func (self *F64Ceil) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Ceil) String() string {
	return "f64.ceil"
}

type F64Floor struct{ 
	
}

func (self *F64Floor) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Floor) String() string {
	return "f64.floor"
}

type F64Trunc struct{ 
	
}

func (self *F64Trunc) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Trunc) String() string {
	return "f64.trunc"
}

type F64Nearest struct{ 
	
}

func (self *F64Nearest) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Nearest) String() string {
	return "f64.nearest"
}

type F64Sqrt struct{ 
	
}

func (self *F64Sqrt) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Sqrt) String() string {
	return "f64.sqrt"
}

type F64Add struct{ 
	
}

func (self *F64Add) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Add) String() string {
	return "f64.add"
}

type F64Sub struct{ 
	
}

func (self *F64Sub) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Sub) String() string {
	return "f64.sub"
}

type F64Mul struct{ 
	
}

func (self *F64Mul) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Mul) String() string {
	return "f64.mul"
}

type F64Div struct{ 
	
}

func (self *F64Div) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Div) String() string {
	return "f64.div"
}

type F64Min struct{ 
	
}

func (self *F64Min) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Min) String() string {
	return "f64.min"
}

type F64Max struct{ 
	
}

func (self *F64Max) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Max) String() string {
	return "f64.max"
}

type F64Copysign struct{ 
	
}

func (self *F64Copysign) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Copysign) String() string {
	return "f64.copysign"
}

type I32Eqz struct{ 
	
}

func (self *I32Eqz) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Eqz) String() string {
	return "i32.eqz"
}

type I32Eq struct{ 
	
}

func (self *I32Eq) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Eq) String() string {
	return "i32.eq"
}

type I32Ne struct{ 
	
}

func (self *I32Ne) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Ne) String() string {
	return "i32.ne"
}

type I32LtS struct{ 
	
}

func (self *I32LtS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32LtS) String() string {
	return "i32.lt_s"
}

type I32LtU struct{ 
	
}

func (self *I32LtU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32LtU) String() string {
	return "i32.lt_u"
}

type I32GtS struct{ 
	
}

func (self *I32GtS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32GtS) String() string {
	return "i32.gt_s"
}

type I32GtU struct{ 
	
}

func (self *I32GtU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32GtU) String() string {
	return "i32.gt_u"
}

type I32LeS struct{ 
	
}

func (self *I32LeS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32LeS) String() string {
	return "i32.le_s"
}

type I32LeU struct{ 
	
}

func (self *I32LeU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32LeU) String() string {
	return "i32.le_u"
}

type I32GeS struct{ 
	
}

func (self *I32GeS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32GeS) String() string {
	return "i32.ge_s"
}

type I32GeU struct{ 
	
}

func (self *I32GeU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32GeU) String() string {
	return "i32.ge_u"
}

type I64Eqz struct{ 
	
}

func (self *I64Eqz) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Eqz) String() string {
	return "i64.eqz"
}

type I64Eq struct{ 
	
}

func (self *I64Eq) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Eq) String() string {
	return "i64.eq"
}

type I64Ne struct{ 
	
}

func (self *I64Ne) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Ne) String() string {
	return "i64.ne"
}

type I64LtS struct{ 
	
}

func (self *I64LtS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64LtS) String() string {
	return "i64.lt_s"
}

type I64LtU struct{ 
	
}

func (self *I64LtU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64LtU) String() string {
	return "i64.lt_u"
}

type I64GtS struct{ 
	
}

func (self *I64GtS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64GtS) String() string {
	return "i64.gt_s"
}

type I64GtU struct{ 
	
}

func (self *I64GtU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64GtU) String() string {
	return "i64.gt_u"
}

type I64LeS struct{ 
	
}

func (self *I64LeS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64LeS) String() string {
	return "i64.le_s"
}

type I64LeU struct{ 
	
}

func (self *I64LeU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64LeU) String() string {
	return "i64.le_u"
}

type I64GeS struct{ 
	
}

func (self *I64GeS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64GeS) String() string {
	return "i64.ge_s"
}

type I64GeU struct{ 
	
}

func (self *I64GeU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64GeU) String() string {
	return "i64.ge_u"
}

type F32Eq struct{ 
	
}

func (self *F32Eq) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Eq) String() string {
	return "f32.eq"
}

type F32Ne struct{ 
	
}

func (self *F32Ne) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Ne) String() string {
	return "f32.ne"
}

type F32Lt struct{ 
	
}

func (self *F32Lt) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Lt) String() string {
	return "f32.lt"
}

type F32Gt struct{ 
	
}

func (self *F32Gt) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Gt) String() string {
	return "f32.gt"
}

type F32Le struct{ 
	
}

func (self *F32Le) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Le) String() string {
	return "f32.le"
}

type F32Ge struct{ 
	
}

func (self *F32Ge) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32Ge) String() string {
	return "f32.ge"
}

type F64Eq struct{ 
	
}

func (self *F64Eq) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Eq) String() string {
	return "f64.eq"
}

type F64Ne struct{ 
	
}

func (self *F64Ne) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Ne) String() string {
	return "f64.ne"
}

type F64Lt struct{ 
	
}

func (self *F64Lt) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Lt) String() string {
	return "f64.lt"
}

type F64Gt struct{ 
	
}

func (self *F64Gt) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Gt) String() string {
	return "f64.gt"
}

type F64Le struct{ 
	
}

func (self *F64Le) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Le) String() string {
	return "f64.le"
}

type F64Ge struct{ 
	
}

func (self *F64Ge) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64Ge) String() string {
	return "f64.ge"
}

type I32Extend8S struct{ 
	
}

func (self *I32Extend8S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Extend8S) String() string {
	return "i32.extend8_s"
}

type I32Extend16S struct{ 
	
}

func (self *I32Extend16S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32Extend16S) String() string {
	return "i32.extend16_s"
}

type I64Extend8S struct{ 
	
}

func (self *I64Extend8S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Extend8S) String() string {
	return "i64.extend8_s"
}

type I64Extend16S struct{ 
	
}

func (self *I64Extend16S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Extend16S) String() string {
	return "i64.extend16_s"
}

type I64Extend32S struct{ 
	
}

func (self *I64Extend32S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64Extend32S) String() string {
	return "i64.extend32_s"
}

type I8x16Eq struct{ 
	
}

func (self *I8x16Eq) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16Eq) String() string {
	return "i8x16.eq"
}

type I8x16Ne struct{ 
	
}

func (self *I8x16Ne) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16Ne) String() string {
	return "i8x16.ne"
}

type I8x16LtS struct{ 
	
}

func (self *I8x16LtS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16LtS) String() string {
	return "i8x16.lt_s"
}

type I8x16LtU struct{ 
	
}

func (self *I8x16LtU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16LtU) String() string {
	return "i8x16.lt_u"
}

type I8x16GtS struct{ 
	
}

func (self *I8x16GtS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16GtS) String() string {
	return "i8x16.gt_s"
}

type I8x16GtU struct{ 
	
}

func (self *I8x16GtU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16GtU) String() string {
	return "i8x16.gt_u"
}

type I8x16LeS struct{ 
	
}

func (self *I8x16LeS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16LeS) String() string {
	return "i8x16.le_s"
}

type I8x16LeU struct{ 
	
}

func (self *I8x16LeU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16LeU) String() string {
	return "i8x16.le_u"
}

type I8x16GeS struct{ 
	
}

func (self *I8x16GeS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16GeS) String() string {
	return "i8x16.ge_s"
}

type I8x16GeU struct{ 
	
}

func (self *I8x16GeU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16GeU) String() string {
	return "i8x16.ge_u"
}

type I16x8Eq struct{ 
	
}

func (self *I16x8Eq) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8Eq) String() string {
	return "i16x8.eq"
}

type I16x8Ne struct{ 
	
}

func (self *I16x8Ne) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8Ne) String() string {
	return "i16x8.ne"
}

type I16x8LtS struct{ 
	
}

func (self *I16x8LtS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8LtS) String() string {
	return "i16x8.lt_s"
}

type I16x8LtU struct{ 
	
}

func (self *I16x8LtU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8LtU) String() string {
	return "i16x8.lt_u"
}

type I16x8GtS struct{ 
	
}

func (self *I16x8GtS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8GtS) String() string {
	return "i16x8.gt_s"
}

type I16x8GtU struct{ 
	
}

func (self *I16x8GtU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8GtU) String() string {
	return "i16x8.gt_u"
}

type I16x8LeS struct{ 
	
}

func (self *I16x8LeS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8LeS) String() string {
	return "i16x8.le_s"
}

type I16x8LeU struct{ 
	
}

func (self *I16x8LeU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8LeU) String() string {
	return "i16x8.le_u"
}

type I16x8GeS struct{ 
	
}

func (self *I16x8GeS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8GeS) String() string {
	return "i16x8.ge_s"
}

type I16x8GeU struct{ 
	
}

func (self *I16x8GeU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8GeU) String() string {
	return "i16x8.ge_u"
}

type I32x4Eq struct{ 
	
}

func (self *I32x4Eq) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4Eq) String() string {
	return "i32x4.eq"
}

type I32x4Ne struct{ 
	
}

func (self *I32x4Ne) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4Ne) String() string {
	return "i32x4.ne"
}

type I32x4LtS struct{ 
	
}

func (self *I32x4LtS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4LtS) String() string {
	return "i32x4.lt_s"
}

type I32x4LtU struct{ 
	
}

func (self *I32x4LtU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4LtU) String() string {
	return "i32x4.lt_u"
}

type I32x4GtS struct{ 
	
}

func (self *I32x4GtS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4GtS) String() string {
	return "i32x4.gt_s"
}

type I32x4GtU struct{ 
	
}

func (self *I32x4GtU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4GtU) String() string {
	return "i32x4.gt_u"
}

type I32x4LeS struct{ 
	
}

func (self *I32x4LeS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4LeS) String() string {
	return "i32x4.le_s"
}

type I32x4LeU struct{ 
	
}

func (self *I32x4LeU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4LeU) String() string {
	return "i32x4.le_u"
}

type I32x4GeS struct{ 
	
}

func (self *I32x4GeS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4GeS) String() string {
	return "i32x4.ge_s"
}

type I32x4GeU struct{ 
	
}

func (self *I32x4GeU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4GeU) String() string {
	return "i32x4.ge_u"
}

type F32x4Eq struct{ 
	
}

func (self *F32x4Eq) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Eq) String() string {
	return "f32x4.eq"
}

type F32x4Ne struct{ 
	
}

func (self *F32x4Ne) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Ne) String() string {
	return "f32x4.ne"
}

type F32x4Lt struct{ 
	
}

func (self *F32x4Lt) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Lt) String() string {
	return "f32x4.lt"
}

type F32x4Gt struct{ 
	
}

func (self *F32x4Gt) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Gt) String() string {
	return "f32x4.gt"
}

type F32x4Le struct{ 
	
}

func (self *F32x4Le) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Le) String() string {
	return "f32x4.le"
}

type F32x4Ge struct{ 
	
}

func (self *F32x4Ge) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Ge) String() string {
	return "f32x4.ge"
}

type F64x2Eq struct{ 
	
}

func (self *F64x2Eq) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Eq) String() string {
	return "f64x2.eq"
}

type F64x2Ne struct{ 
	
}

func (self *F64x2Ne) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Ne) String() string {
	return "f64x2.ne"
}

type F64x2Lt struct{ 
	
}

func (self *F64x2Lt) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Lt) String() string {
	return "f64x2.lt"
}

type F64x2Gt struct{ 
	
}

func (self *F64x2Gt) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Gt) String() string {
	return "f64x2.gt"
}

type F64x2Le struct{ 
	
}

func (self *F64x2Le) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Le) String() string {
	return "f64x2.le"
}

type F64x2Ge struct{ 
	
}

func (self *F64x2Ge) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Ge) String() string {
	return "f64x2.ge"
}

type V128Not struct{ 
	
}

func (self *V128Not) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *V128Not) String() string {
	return "v128.not"
}

type V128And struct{ 
	
}

func (self *V128And) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *V128And) String() string {
	return "v128.and"
}

type V128Or struct{ 
	
}

func (self *V128Or) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *V128Or) String() string {
	return "v128.or"
}

type V128Xor struct{ 
	
}

func (self *V128Xor) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *V128Xor) String() string {
	return "v128.xor"
}

type V128Bitselect struct{ 
	
}

func (self *V128Bitselect) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *V128Bitselect) String() string {
	return "v128.bitselect"
}

type I8x16Neg struct{ 
	
}

func (self *I8x16Neg) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16Neg) String() string {
	return "i8x16.neg"
}

type I8x16AnyTrue struct{ 
	
}

func (self *I8x16AnyTrue) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16AnyTrue) String() string {
	return "i8x16.any_true"
}

type I8x16AllTrue struct{ 
	
}

func (self *I8x16AllTrue) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16AllTrue) String() string {
	return "i8x16.all_true"
}

type I8x16Shl struct{ 
	
}

func (self *I8x16Shl) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16Shl) String() string {
	return "i8x16.shl"
}

type I8x16ShrS struct{ 
	
}

func (self *I8x16ShrS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16ShrS) String() string {
	return "i8x16.shr_s"
}

type I8x16ShrU struct{ 
	
}

func (self *I8x16ShrU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16ShrU) String() string {
	return "i8x16.shr_u"
}

type I8x16Add struct{ 
	
}

func (self *I8x16Add) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16Add) String() string {
	return "i8x16.add"
}

type I8x16AddSaturateS struct{ 
	
}

func (self *I8x16AddSaturateS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16AddSaturateS) String() string {
	return "i8x16.add_saturate_s"
}

type I8x16AddSaturateU struct{ 
	
}

func (self *I8x16AddSaturateU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16AddSaturateU) String() string {
	return "i8x16.add_saturate_u"
}

type I8x16Sub struct{ 
	
}

func (self *I8x16Sub) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16Sub) String() string {
	return "i8x16.sub"
}

type I8x16SubSaturateS struct{ 
	
}

func (self *I8x16SubSaturateS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16SubSaturateS) String() string {
	return "i8x16.sub_saturate_s"
}

type I8x16SubSaturateU struct{ 
	
}

func (self *I8x16SubSaturateU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16SubSaturateU) String() string {
	return "i8x16.sub_saturate_u"
}

type I8x16Mul struct{ 
	
}

func (self *I8x16Mul) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16Mul) String() string {
	return "i8x16.mul"
}

type I16x8Neg struct{ 
	
}

func (self *I16x8Neg) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8Neg) String() string {
	return "i16x8.neg"
}

type I16x8AnyTrue struct{ 
	
}

func (self *I16x8AnyTrue) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8AnyTrue) String() string {
	return "i16x8.any_true"
}

type I16x8AllTrue struct{ 
	
}

func (self *I16x8AllTrue) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8AllTrue) String() string {
	return "i16x8.all_true"
}

type I16x8Shl struct{ 
	
}

func (self *I16x8Shl) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8Shl) String() string {
	return "i16x8.shl"
}

type I16x8ShrS struct{ 
	
}

func (self *I16x8ShrS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8ShrS) String() string {
	return "i16x8.shr_s"
}

type I16x8ShrU struct{ 
	
}

func (self *I16x8ShrU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8ShrU) String() string {
	return "i16x8.shr_u"
}

type I16x8Add struct{ 
	
}

func (self *I16x8Add) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8Add) String() string {
	return "i16x8.add"
}

type I16x8AddSaturateS struct{ 
	
}

func (self *I16x8AddSaturateS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8AddSaturateS) String() string {
	return "i16x8.add_saturate_s"
}

type I16x8AddSaturateU struct{ 
	
}

func (self *I16x8AddSaturateU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8AddSaturateU) String() string {
	return "i16x8.add_saturate_u"
}

type I16x8Sub struct{ 
	
}

func (self *I16x8Sub) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8Sub) String() string {
	return "i16x8.sub"
}

type I16x8SubSaturateS struct{ 
	
}

func (self *I16x8SubSaturateS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8SubSaturateS) String() string {
	return "i16x8.sub_saturate_s"
}

type I16x8SubSaturateU struct{ 
	
}

func (self *I16x8SubSaturateU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8SubSaturateU) String() string {
	return "i16x8.sub_saturate_u"
}

type I16x8Mul struct{ 
	
}

func (self *I16x8Mul) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8Mul) String() string {
	return "i16x8.mul"
}

type I32x4Neg struct{ 
	
}

func (self *I32x4Neg) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4Neg) String() string {
	return "i32x4.neg"
}

type I32x4AnyTrue struct{ 
	
}

func (self *I32x4AnyTrue) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4AnyTrue) String() string {
	return "i32x4.any_true"
}

type I32x4AllTrue struct{ 
	
}

func (self *I32x4AllTrue) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4AllTrue) String() string {
	return "i32x4.all_true"
}

type I32x4Shl struct{ 
	
}

func (self *I32x4Shl) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4Shl) String() string {
	return "i32x4.shl"
}

type I32x4ShrS struct{ 
	
}

func (self *I32x4ShrS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4ShrS) String() string {
	return "i32x4.shr_s"
}

type I32x4ShrU struct{ 
	
}

func (self *I32x4ShrU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4ShrU) String() string {
	return "i32x4.shr_u"
}

type I32x4Add struct{ 
	
}

func (self *I32x4Add) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4Add) String() string {
	return "i32x4.add"
}

type I32x4Sub struct{ 
	
}

func (self *I32x4Sub) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4Sub) String() string {
	return "i32x4.sub"
}

type I32x4Mul struct{ 
	
}

func (self *I32x4Mul) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4Mul) String() string {
	return "i32x4.mul"
}

type I64x2Neg struct{ 
	
}

func (self *I64x2Neg) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64x2Neg) String() string {
	return "i64x2.neg"
}

type I64x2AnyTrue struct{ 
	
}

func (self *I64x2AnyTrue) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64x2AnyTrue) String() string {
	return "i64x2.any_true"
}

type I64x2AllTrue struct{ 
	
}

func (self *I64x2AllTrue) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64x2AllTrue) String() string {
	return "i64x2.all_true"
}

type I64x2Shl struct{ 
	
}

func (self *I64x2Shl) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64x2Shl) String() string {
	return "i64x2.shl"
}

type I64x2ShrS struct{ 
	
}

func (self *I64x2ShrS) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64x2ShrS) String() string {
	return "i64x2.shr_s"
}

type I64x2ShrU struct{ 
	
}

func (self *I64x2ShrU) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64x2ShrU) String() string {
	return "i64x2.shr_u"
}

type I64x2Add struct{ 
	
}

func (self *I64x2Add) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64x2Add) String() string {
	return "i64x2.add"
}

type I64x2Sub struct{ 
	
}

func (self *I64x2Sub) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64x2Sub) String() string {
	return "i64x2.sub"
}

type I64x2Mul struct{ 
	
}

func (self *I64x2Mul) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64x2Mul) String() string {
	return "i64x2.mul"
}

type F32x4Abs struct{ 
	
}

func (self *F32x4Abs) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Abs) String() string {
	return "f32x4.abs"
}

type F32x4Neg struct{ 
	
}

func (self *F32x4Neg) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Neg) String() string {
	return "f32x4.neg"
}

type F32x4Sqrt struct{ 
	
}

func (self *F32x4Sqrt) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Sqrt) String() string {
	return "f32x4.sqrt"
}

type F32x4Add struct{ 
	
}

func (self *F32x4Add) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Add) String() string {
	return "f32x4.add"
}

type F32x4Sub struct{ 
	
}

func (self *F32x4Sub) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Sub) String() string {
	return "f32x4.sub"
}

type F32x4Mul struct{ 
	
}

func (self *F32x4Mul) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Mul) String() string {
	return "f32x4.mul"
}

type F32x4Div struct{ 
	
}

func (self *F32x4Div) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Div) String() string {
	return "f32x4.div"
}

type F32x4Min struct{ 
	
}

func (self *F32x4Min) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Min) String() string {
	return "f32x4.min"
}

type F32x4Max struct{ 
	
}

func (self *F32x4Max) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4Max) String() string {
	return "f32x4.max"
}

type F64x2Abs struct{ 
	
}

func (self *F64x2Abs) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Abs) String() string {
	return "f64x2.abs"
}

type F64x2Neg struct{ 
	
}

func (self *F64x2Neg) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Neg) String() string {
	return "f64x2.neg"
}

type F64x2Sqrt struct{ 
	
}

func (self *F64x2Sqrt) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Sqrt) String() string {
	return "f64x2.sqrt"
}

type F64x2Add struct{ 
	
}

func (self *F64x2Add) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Add) String() string {
	return "f64x2.add"
}

type F64x2Sub struct{ 
	
}

func (self *F64x2Sub) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Sub) String() string {
	return "f64x2.sub"
}

type F64x2Mul struct{ 
	
}

func (self *F64x2Mul) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Mul) String() string {
	return "f64x2.mul"
}

type F64x2Div struct{ 
	
}

func (self *F64x2Div) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Div) String() string {
	return "f64x2.div"
}

type F64x2Min struct{ 
	
}

func (self *F64x2Min) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Min) String() string {
	return "f64x2.min"
}

type F64x2Max struct{ 
	
}

func (self *F64x2Max) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2Max) String() string {
	return "f64x2.max"
}

type I32x4TruncSatF32x4S struct{ 
	
}

func (self *I32x4TruncSatF32x4S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4TruncSatF32x4S) String() string {
	return "i32x4.trunc_sat_f32x4_s"
}

type I32x4TruncSatF32x4U struct{ 
	
}

func (self *I32x4TruncSatF32x4U) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4TruncSatF32x4U) String() string {
	return "i32x4.trunc_sat_f32x4_u"
}

type I64x2TruncSatF64x2S struct{ 
	
}

func (self *I64x2TruncSatF64x2S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64x2TruncSatF64x2S) String() string {
	return "i64x2.trunc_sat_f64x2_s"
}

type I64x2TruncSatF64x2U struct{ 
	
}

func (self *I64x2TruncSatF64x2U) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I64x2TruncSatF64x2U) String() string {
	return "i64x2.trunc_sat_f64x2_u"
}

type F32x4ConvertI32x4S struct{ 
	
}

func (self *F32x4ConvertI32x4S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4ConvertI32x4S) String() string {
	return "f32x4.convert_i32x4_s"
}

type F32x4ConvertI32x4U struct{ 
	
}

func (self *F32x4ConvertI32x4U) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F32x4ConvertI32x4U) String() string {
	return "f32x4.convert_i32x4_u"
}

type F64x2ConvertI64x2S struct{ 
	
}

func (self *F64x2ConvertI64x2S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2ConvertI64x2S) String() string {
	return "f64x2.convert_i64x2_s"
}

type F64x2ConvertI64x2U struct{ 
	
}

func (self *F64x2ConvertI64x2U) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *F64x2ConvertI64x2U) String() string {
	return "f64x2.convert_i64x2_u"
}

type V8x16Swizzle struct{ 
	
}

func (self *V8x16Swizzle) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *V8x16Swizzle) String() string {
	return "v8x16.swizzle"
}

type I8x16NarrowI16x8S struct{ 
	
}

func (self *I8x16NarrowI16x8S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16NarrowI16x8S) String() string {
	return "i8x16.narrow_i16x8_s"
}

type I8x16NarrowI16x8U struct{ 
	
}

func (self *I8x16NarrowI16x8U) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I8x16NarrowI16x8U) String() string {
	return "i8x16.narrow_i16x8_u"
}

type I16x8NarrowI32x4S struct{ 
	
}

func (self *I16x8NarrowI32x4S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8NarrowI32x4S) String() string {
	return "i16x8.narrow_i32x4_s"
}

type I16x8NarrowI32x4U struct{ 
	
}

func (self *I16x8NarrowI32x4U) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8NarrowI32x4U) String() string {
	return "i16x8.narrow_i32x4_u"
}

type I16x8WidenLowI8x16S struct{ 
	
}

func (self *I16x8WidenLowI8x16S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8WidenLowI8x16S) String() string {
	return "i16x8.widen_low_i8x16_s"
}

type I16x8WidenHighI8x16S struct{ 
	
}

func (self *I16x8WidenHighI8x16S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8WidenHighI8x16S) String() string {
	return "i16x8.widen_high_i8x16_s"
}

type I16x8WidenLowI8x16U struct{ 
	
}

func (self *I16x8WidenLowI8x16U) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8WidenLowI8x16U) String() string {
	return "i16x8.widen_low_i8x16_u"
}

type I16x8WidenHighI8x16u struct{ 
	
}

func (self *I16x8WidenHighI8x16u) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I16x8WidenHighI8x16u) String() string {
	return "i16x8.widen_high_i8x16_u"
}

type I32x4WidenLowI16x8S struct{ 
	
}

func (self *I32x4WidenLowI16x8S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4WidenLowI16x8S) String() string {
	return "i32x4.widen_low_i16x8_s"
}

type I32x4WidenHighI16x8S struct{ 
	
}

func (self *I32x4WidenHighI16x8S) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4WidenHighI16x8S) String() string {
	return "i32x4.widen_high_i16x8_s"
}

type I32x4WidenLowI16x8U struct{ 
	
}

func (self *I32x4WidenLowI16x8U) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4WidenLowI16x8U) String() string {
	return "i32x4.widen_low_i16x8_u"
}

type I32x4WidenHighI16x8u struct{ 
	
}

func (self *I32x4WidenHighI16x8u) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *I32x4WidenHighI16x8u) String() string {
	return "i32x4.widen_high_i16x8_u"
}

type V128Andnot struct{ 
	
}

func (self *V128Andnot) parseInstrBody(ps *parser.ParserBuffer) error { 
	
	return nil
}

func (self *V128Andnot) String() string {
	return "v128.andnot"
}

