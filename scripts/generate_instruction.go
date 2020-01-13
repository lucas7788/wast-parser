package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ontio/wast-parser/lexer"
	"github.com/ontio/wast-parser/parser"
	"github.com/valyala/fasttemplate"
)

type Instruction struct {
	Name   string
	Id     []string
	Fields []Field
}

type Field struct {
	Name string
	Type string
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
	if ps.PeekToken().Type() == lexer.LParenType {
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			for !ps.Empty() {
				id, err := expectKeyword(ps)
				if err != nil {
					return err
				}
				self.Id = append(self.Id, id)
			}
			return nil
		})

		if err != nil {
			return err
		}
	} else {
		id, err := expectKeyword(ps)
		if err != nil {
			return err
		}
		self.Id = append(self.Id, id)
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
type [Name] struct{ 
	[Fields]
}

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
		"Fields":    self.generateFields(),
		"parseBody": self.generateParseBody(),
		"Id":        self.Id[0],
	})
}

func (self Instruction) generateFields() string {
	var fields []string
	for _, field := range self.Fields {
		fields = append(fields, fmt.Sprintf("%s %s", field.Name, field.Type))
	}

	return strings.Join(fields, "\n")
}

func (self Instruction) generateParseBody() string {
	body := ""
	for _, field := range self.Fields {
		switch field.Type {
		case "uint32":
			body += parseInt(field.Name, "Uint32")
		case "int64":
			body += parseInt(field.Name, "Int64")
		default:
			body += parseGeneral(field.Name)
		}
	}

	return body
}

func parseGeneral(name string) string {
	return generate(
		`err := self.[Name].Parse(ps)
	if err != nil {
		return err
	}
`, map[string]interface{}{"Name": name})
}

func parseInt(name string, ty string) string {
	return generate(`val, err := ps.Expect[Type]()
	if err != nil {
		return err
	}
	self.[Name] = val
`, map[string]interface{}{"Name": name, "Type": ty})
}

func mustParseInstrs(source string) []Instruction {
	ps, err := parser.NewParserBuffer(source)
	if err != nil {
		panic(err)
	}
	var instrs []Instruction
	for !ps.Empty() {
		var inst Instruction
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			return inst.Parse(ps)
		})
		if err != nil {
			panic(err)
		}
		instrs = append(instrs, inst)
	}

	return instrs
}

func generateParseInstrution(instrs []Instruction) string {
	var cases []string
	for _, instr := range instrs {
		cases = append(cases, generate(` case "[Id]":
		inst = &[Name]{}`, map[string]interface{}{"Id": strings.Join(instr.Id, "\", \""), "Name": instr.Name}))
	}

	return generate(`
func parseInstr(ps *parser.ParserBuffer) (Instruction, error) {
	var inst Instruction
	kw, err := ps.ExpectKeyword()
	if err != nil {
		return nil, err
	}
	switch kw {
	[cases]
	default:
		panic("todo")
	}
	err = inst.parseInstrBody(ps)
	if err != nil {
		return nil, err
	}
	return inst, nil
}
`, map[string]interface{}{"cases": strings.Join(cases, "\n")})
}

func main() {
	instrs := `
(Block block (BlockType BlockType))
(If if (BlockType BlockType))
(Else else (Id OptionId))
;;Loop(BlockType<'a>) : [0x03] : "loop",
;;End(Option<ast::Id<'a>>) : [0x0b] : "end",

(Unreachable unreachable)
(Nop nop)
(Br br (Index Index))
(BrIf br_if (Index Index))
(BrTable br_table (Indices BrTableIndices)) 
(Return return)
(Call call (Index Index))
(CallIndirect call_indirect (Impl CallIndirectInner))
(ReturnCall return_call (Index Index))
(ReturnCallIndirect return_call_indirect (Impl CallIndirectInner))
(Drop drop)
(Select select (SelectTypes SelectTypes))
(LocalGet (local.get get_local) (Index Index)) 
(LocalSet (local.set set_local) (Index Index)) 
(LocalTee (local.tee tee_local) (Index Index)) 
(GlobalGet (global.get get_global) (Index Index)) 
(GlobalSet (global.set set_global) (Index Index)) 

(TableGet table.get (Index Index))
(TableSet table.set (Index Index))

;;(I32Load(MemArg<4>) i32.load)
;;(I64Load(MemArg<8>) i64.load)
;;(F32Load(MemArg<4>) f32.load)
;;(F64Load(MemArg<8>) f64.load)
;;(I32Load8s(MemArg<1>) i32.load8_s)
;;(I32Load8u(MemArg<1>) i32.load8_u)
;;(I32Load16s(MemArg<2>) i32.load16_s)
;;(I32Load16u(MemArg<2>) i32.load16_u)
;;(I64Load8s(MemArg<1>) i64.load8_s)
;;(I64Load8u(MemArg<1>) i64.load8_u)
;;(I64Load16s(MemArg<2>) i64.load16_s)
;;(I64Load16u(MemArg<2>) i64.load16_u)
;;(I64Load32s(MemArg<4>) i64.load32_s)
;;(I64Load32u(MemArg<4>) i64.load32_u)
;;(I32Store(MemArg<4>) i32.store)
;;(I64Store(MemArg<8>) i64.store)
;;(F32Store(MemArg<4>) f32.store)
;;(F64Store(MemArg<8>) f64.store)
;;(I32Store8(MemArg<1>) i32.store8)
;;(I32Store16(MemArg<2>) i32.store16)
;;(I64Store8(MemArg<1>) i64.store8)
;;(I64Store16(MemArg<2>) i64.store16)
;;(I64Store32(MemArg<4>) i64.store32)

;; Lots of bulk memory proposal here as well
(MemorySize (memory.size current_memory))
(MemoryGrow (memory.grow grow_memory))
;;MemoryInit(MemoryInit<'a>) : [0xfc, 0x08] : "memory.init",
(MemoryCopy memory.copy)
(MemoryFill memory.fill)
(DataDrop data.drop (Index Index))
(ElemDrop elem.drop (Index Index))
;;(TableInit table.init (Impl TableInit))
(TableCopy table.copy)
(TableFill table.fill (Index Index))
(TableSize table.size (Index Index))
(TableGrow table.grow (Index Index))

(RefNull ref.null)
(RefIsNull ref.is_null)
(RefHost ref.host (Val uint32))
(RefFunc ref.func (Index Index))

(I32Const i32.const (Val uint32))
(I64Const i64.const (Val int64))
(F32Const f32.const (Val Float32))
(F64Const f64.const (Val Float64))

(I32Clz i32.clz)
(I32Ctz i32.ctz)
(I32Pocnt i32.popcnt)
(I32Add i32.add)
(I32Sub i32.sub)
(I32Mul i32.mul)
(I32DivS i32.div_s)
(I32DivU i32.div_u)
(I32RemS i32.rem_s)
(I32RemU i32.rem_u)
(I32And i32.and)
(I32Or i32.or)
(I32Xor i32.xor)
(I32Shl i32.shl)
(I32ShrS i32.shr_s)
(I32ShrU i32.shr_u)
(I32Rotl i32.rotl)
(I32Rotr i32.rotr)

(I64Clz i64.clz)
(I64Ctz i64.ctz)
(I64Popcnt i64.popcnt)
(I64Add i64.add)
(I64Sub i64.sub)
(I64Mul i64.mul)
(I64DivS i64.div_s)
(I64DivU i64.div_u)
(I64RemS i64.rem_s)
(I64RemU i64.rem_u)
(I64And i64.and)
(I64Or i64.or)
(I64Xor i64.xor)
(I64Shl i64.shl)
(I64ShrS i64.shr_s)
(I64ShrU i64.shr_u)
(I64Rotl i64.rotl)
(I64Rotr i64.rotr)

(F32Abs f32.abs)
(F32Neg f32.neg)
(F32Ceil f32.ceil)
(F32Floor f32.floor)
(F32Trunc f32.trunc)
(F32Nearest f32.nearest)
(F32Sqrt f32.sqrt)
(F32Add f32.add)
(F32Sub f32.sub)
(F32Mul f32.mul)
(F32Div f32.div)
(F32Min f32.min)
(F32Max f32.max)
(F32Copysign f32.copysign)

(F64Abs f64.abs)
(F64Neg f64.neg)
(F64Ceil f64.ceil)
(F64Floor f64.floor)
(F64Trunc f64.trunc)
(F64Nearest f64.nearest)
(F64Sqrt f64.sqrt)
(F64Add f64.add)
(F64Sub f64.sub)
(F64Mul f64.mul)
(F64Div f64.div)
(F64Min f64.min)
(F64Max f64.max)
(F64Copysign f64.copysign)

(I32Eqz i32.eqz)
(I32Eq i32.eq)
(I32Ne i32.ne)
(I32LtS i32.lt_s)
(I32LtU i32.lt_u)
(I32GtS i32.gt_s)
(I32GtU i32.gt_u)
(I32LeS i32.le_s)
(I32LeU i32.le_u)
(I32GeS i32.ge_s)
(I32GeU i32.ge_u)

(I64Eqz i64.eqz)
(I64Eq i64.eq)
(I64Ne i64.ne)
(I64LtS i64.lt_s)
(I64LtU i64.lt_u)
(I64GtS i64.gt_s)
(I64GtU i64.gt_u)
(I64LeS i64.le_s)
(I64LeU i64.le_u)
(I64GeS i64.ge_s)
(I64GeU i64.ge_u)

(F32Eq f32.eq)
(F32Ne f32.ne)
(F32Lt f32.lt)
(F32Gt f32.gt)
(F32Le f32.le)
(F32Ge f32.ge)

(F64Eq f64.eq)
(F64Ne f64.ne)
(F64Lt f64.lt)
(F64Gt f64.gt)
(F64Le f64.le)
(F64Ge f64.ge)

(I32WrapI64 (i32.wrap_i64 i32.wrap/i64))
(I32TruncF32S (i32.trunc_f32_s i32.trunc_s/f32))
(I32TruncF32U (i32.trunc_f32_u i32.trunc_u/f32))
(I32TruncF64S (i32.trunc_f64_s i32.trunc_s/f64))
(I32TruncF64U (i32.trunc_f64_u i32.trunc_u/f64))
(I64ExtendI32S (i64.extend_i32_s i64.extend_s/i32))
(I64ExtendI32U (i64.extend_i32_u i64.extend_u/i32))
(I64TruncF32S (i64.trunc_f32_s i64.trunc_s/f32))
(I64TruncF32U (i64.trunc_f32_u i64.trunc_u/f32))
(I64TruncF64S (i64.trunc_f64_s i64.trunc_s/f64))
(I64TruncF64U (i64.trunc_f64_u i64.trunc_u/f64))
(F32ConvertI32S (f32.convert_i32_s f32.convert_s/i32))
(F32ConvertI32U (f32.convert_i32_u f32.convert_u/i32))
(F32ConvertI64S (f32.convert_i64.s f32.convert_s/i64))
(F32ConvertI64U (f32.convert_i64.u f32.convert_u/i64))
(F32DemoteF64 (f32.demote_f64 f32.demote/f64))
(F64ConvertI32S (f64.convert_i32_s f64.convert_s/i32))
(F64ConvertI32U (f64.convert_i32_u f64.convert_u/i32))
(F64ConvertI64S (f64.convert_i64.s f64.convert_s/i64))
(F64ConvertI64U (f64.convert_i64.u f64.convert_u/i64))
(F64PromoteF32 (f64.promote_f32 f64.promote/f32))
(I32ReinterpretF32 (i32.reinterpret_f32 i32.reinterpret/f32))
(I64ReinterpretF64 (i64.reinterpret_f64 i64.reinterpret/f64))
(F32ReinterpretI32 (f32.reinterpret_i32 f32.reinterpret/i32))
(F64ReinterpretI64 (f64.reinterpret_i64 f64.reinterpret/i64))

;;(// non-trapping float to int
(I32TruncSatF32S (i32.trunc_sat_f32_s i32.trunc_s:sat/f32))
(I32TruncSatF32U (i32.trunc_sat_f32_u i32.trunc_u:sat/f32))
(I32TruncSatF64S (i32.trunc_sat_f64_s i32.trunc_s:sat/f64))
(I32TruncSatF64U (i32.trunc_sat_f64_u i32.trunc_u:sat/f64))
(I64TruncSatF32S (i64.trunc_sat_f32_s i64.trunc_s:sat/f32))
(I64TruncSatF32U (i64.trunc_sat_f32_u i64.trunc_u:sat/f32))
(I64TruncSatF64S (i64.trunc_sat_f64_s i64.trunc_s:sat/f64))
(I64TruncSatF64U (i64.trunc_sat_f64_u i64.trunc_u:sat/f64))

;; sign extension proposal
(I32Extend8S i32.extend8_s)
(I32Extend16S i32.extend16_s)
(I64Extend8S i64.extend8_s)
(I64Extend16S i64.extend16_s)
(I64Extend32S i64.extend32_s)

;; atomics proposal
;;AtomicNotify(MemArg<4>) : [0xfe, 0x00] : "atomic.notify",
;;I32AtomicWait(MemArg<4>) : [0xfe, 0x01] : "i32.atomic.wait",
;;I64AtomicWait(MemArg<8>) : [0xfe, 0x02] : "i64.atomic.wait",
;;AtomicFence : [0xfe, 0x03] : "atomic.fence",
;;
;;I32AtomicLoad(MemArg<4>) : [0xfe, 0x10] : "i32.atomic.load",
;;I64AtomicLoad(MemArg<8>) : [0xfe, 0x11] : "i64.atomic.load",
;;I32AtomicLoad8u(MemArg<1>) : [0xfe, 0x12] : "i32.atomic.load8_u",
;;I32AtomicLoad16u(MemArg<2>) : [0xfe, 0x13] : "i32.atomic.load16_u",
;;I64AtomicLoad8u(MemArg<1>) : [0xfe, 0x14] : "i64.atomic.load8_u",
;;I64AtomicLoad16u(MemArg<2>) : [0xfe, 0x15] : "i64.atomic.load16_u",
;;I64AtomicLoad32u(MemArg<4>) : [0xfe, 0x16] : "i64.atomic.load32_u",
;;I32AtomicStore(MemArg<4>) : [0xfe, 0x17] : "i32.atomic.store",
;;I64AtomicStore(MemArg<8>) : [0xfe, 0x18] : "i64.atomic.store",
;;I32AtomicStore8(MemArg<1>) : [0xfe, 0x19] : "i32.atomic.store8",
;;I32AtomicStore16(MemArg<2>) : [0xfe, 0x1a] : "i32.atomic.store16",
;;I64AtomicStore8(MemArg<1>) : [0xfe, 0x1b] : "i64.atomic.store8",
;;I64AtomicStore16(MemArg<2>) : [0xfe, 0x1c] : "i64.atomic.store16",
;;I64AtomicStore32(MemArg<4>) : [0xfe, 0x1d] : "i64.atomic.store32",
;;
;;I32AtomicRmwAdd(MemArg<4>) : [0xfe, 0x1e] : "i32.atomic.rmw.add",
;;I64AtomicRmwAdd(MemArg<8>) : [0xfe, 0x1f] : "i64.atomic.rmw.add",
;;I32AtomicRmw8AddU(MemArg<1>) : [0xfe, 0x20] : "i32.atomic.rmw8.add_u",
;;I32AtomicRmw16AddU(MemArg<2>) : [0xfe, 0x21] : "i32.atomic.rmw16.add_u",
;;I64AtomicRmw8AddU(MemArg<1>) : [0xfe, 0x22] : "i64.atomic.rmw8.add_u",
;;I64AtomicRmw16AddU(MemArg<2>) : [0xfe, 0x23] : "i64.atomic.rmw16.add_u",
;;I64AtomicRmw32AddU(MemArg<4>) : [0xfe, 0x24] : "i64.atomic.rmw32.add_u",
;;
;;I32AtomicRmwSub(MemArg<4>) : [0xfe, 0x25] : "i32.atomic.rmw.sub",
;;I64AtomicRmwSub(MemArg<8>) : [0xfe, 0x26] : "i64.atomic.rmw.sub",
;;I32AtomicRmw8SubU(MemArg<1>) : [0xfe, 0x27] : "i32.atomic.rmw8.sub_u",
;;I32AtomicRmw16SubU(MemArg<2>) : [0xfe, 0x28] : "i32.atomic.rmw16.sub_u",
;;I64AtomicRmw8SubU(MemArg<1>) : [0xfe, 0x29] : "i64.atomic.rmw8.sub_u",
;;I64AtomicRmw16SubU(MemArg<2>) : [0xfe, 0x2a] : "i64.atomic.rmw16.sub_u",
;;I64AtomicRmw32SubU(MemArg<4>) : [0xfe, 0x2b] : "i64.atomic.rmw32.sub_u",
;;
;;I32AtomicRmwAnd(MemArg<4>) : [0xfe, 0x2c] : "i32.atomic.rmw.and",
;;I64AtomicRmwAnd(MemArg<8>) : [0xfe, 0x2d] : "i64.atomic.rmw.and",
;;I32AtomicRmw8AndU(MemArg<1>) : [0xfe, 0x2e] : "i32.atomic.rmw8.and_u",
;;I32AtomicRmw16AndU(MemArg<2>) : [0xfe, 0x2f] : "i32.atomic.rmw16.and_u",
;;I64AtomicRmw8AndU(MemArg<1>) : [0xfe, 0x30] : "i64.atomic.rmw8.and_u",
;;I64AtomicRmw16AndU(MemArg<2>) : [0xfe, 0x31] : "i64.atomic.rmw16.and_u",
;;I64AtomicRmw32AndU(MemArg<4>) : [0xfe, 0x32] : "i64.atomic.rmw32.and_u",
;;
;;I32AtomicRmwOr(MemArg<4>) : [0xfe, 0x33] : "i32.atomic.rmw.or",
;;I64AtomicRmwOr(MemArg<8>) : [0xfe, 0x34] : "i64.atomic.rmw.or",
;;I32AtomicRmw8OrU(MemArg<1>) : [0xfe, 0x35] : "i32.atomic.rmw8.or_u",
;;I32AtomicRmw16OrU(MemArg<2>) : [0xfe, 0x36] : "i32.atomic.rmw16.or_u",
;;I64AtomicRmw8OrU(MemArg<1>) : [0xfe, 0x37] : "i64.atomic.rmw8.or_u",
;;I64AtomicRmw16OrU(MemArg<2>) : [0xfe, 0x38] : "i64.atomic.rmw16.or_u",
;;I64AtomicRmw32OrU(MemArg<4>) : [0xfe, 0x39] : "i64.atomic.rmw32.or_u",
;;
;;I32AtomicRmwXor(MemArg<4>) : [0xfe, 0x3a] : "i32.atomic.rmw.xor",
;;I64AtomicRmwXor(MemArg<8>) : [0xfe, 0x3b] : "i64.atomic.rmw.xor",
;;I32AtomicRmw8XorU(MemArg<1>) : [0xfe, 0x3c] : "i32.atomic.rmw8.xor_u",
;;I32AtomicRmw16XorU(MemArg<2>) : [0xfe, 0x3d] : "i32.atomic.rmw16.xor_u",
;;I64AtomicRmw8XorU(MemArg<1>) : [0xfe, 0x3e] : "i64.atomic.rmw8.xor_u",
;;I64AtomicRmw16XorU(MemArg<2>) : [0xfe, 0x3f] : "i64.atomic.rmw16.xor_u",
;;I64AtomicRmw32XorU(MemArg<4>) : [0xfe, 0x40] : "i64.atomic.rmw32.xor_u",
;;
;;I32AtomicRmwXchg(MemArg<4>) : [0xfe, 0x41] : "i32.atomic.rmw.xchg",
;;I64AtomicRmwXchg(MemArg<8>) : [0xfe, 0x42] : "i64.atomic.rmw.xchg",
;;I32AtomicRmw8XchgU(MemArg<1>) : [0xfe, 0x43] : "i32.atomic.rmw8.xchg_u",
;;I32AtomicRmw16XchgU(MemArg<2>) : [0xfe, 0x44] : "i32.atomic.rmw16.xchg_u",
;;I64AtomicRmw8XchgU(MemArg<1>) : [0xfe, 0x45] : "i64.atomic.rmw8.xchg_u",
;;I64AtomicRmw16XchgU(MemArg<2>) : [0xfe, 0x46] : "i64.atomic.rmw16.xchg_u",
;;I64AtomicRmw32XchgU(MemArg<4>) : [0xfe, 0x47] : "i64.atomic.rmw32.xchg_u",
;;
;;I32AtomicRmwCmpxchg(MemArg<4>) : [0xfe, 0x48] : "i32.atomic.rmw.cmpxchg",
;;I64AtomicRmwCmpxchg(MemArg<8>) : [0xfe, 0x49] : "i64.atomic.rmw.cmpxchg",
;;I32AtomicRmw8CmpxchgU(MemArg<1>) : [0xfe, 0x4a] : "i32.atomic.rmw8.cmpxchg_u",
;;I32AtomicRmw16CmpxchgU(MemArg<2>) : [0xfe, 0x4b] : "i32.atomic.rmw16.cmpxchg_u",
;;I64AtomicRmw8CmpxchgU(MemArg<1>) : [0xfe, 0x4c] : "i64.atomic.rmw8.cmpxchg_u",
;;I64AtomicRmw16CmpxchgU(MemArg<2>) : [0xfe, 0x4d] : "i64.atomic.rmw16.cmpxchg_u",
;;I64AtomicRmw32CmpxchgU(MemArg<4>) : [0xfe, 0x4e] : "i64.atomic.rmw32.cmpxchg_u",

;;
;;V128Load(MemArg<16>) : [0xfd, 0x00] : "v128.load",
;;V128Store(MemArg<16>) : [0xfd, 0x01] : "v128.store",
;;V128Const(V128Const) : [0xfd, 0x02] : "v128.const",
;;
;;I8x16Splat : [0xfd, 0x04] : "i8x16.splat",
;;I8x16ExtractLaneS(i32) : [0xfd, 0x05] : "i8x16.extract_lane_s",
;;I8x16ExtractLaneU(i32) : [0xfd, 0x06] : "i8x16.extract_lane_u",
;;I8x16ReplaceLane(i32) : [0xfd, 0x07] : "i8x16.replace_lane",
;;I16x8Splat : [0xfd, 0x08] : "i16x8.splat",
;;I16x8ExtractLaneS(i32) : [0xfd, 0x09] : "i16x8.extract_lane_s",
;;I16x8ExtractLaneU(i32) : [0xfd, 0x0a] : "i16x8.extract_lane_u",
;;I16x8ReplaceLane(i32) : [0xfd, 0x0b] : "i16x8.replace_lane",
;;I32x4Splat : [0xfd, 0x0c] : "i32x4.splat",
;;I32x4ExtractLane(i32) : [0xfd, 0x0d] : "i32x4.extract_lane",
;;I32x4ReplaceLane(i32) : [0xfd, 0x0e] : "i32x4.replace_lane",
;;I64x2Splat : [0xfd, 0x0f] : "i64x2.splat",
;;I64x2ExtractLane(i32) : [0xfd, 0x10] : "i64x2.extract_lane",
;;I64x2ReplaceLane(i32) : [0xfd, 0x11] : "i64x2.replace_lane",
;;F32x4Splat : [0xfd, 0x12] : "f32x4.splat",
;;F32x4ExtractLane(i32) : [0xfd, 0x13] : "f32x4.extract_lane",
;;F32x4ReplaceLane(i32) : [0xfd, 0x14] : "f32x4.replace_lane",
;;F64x2Splat : [0xfd, 0x15] : "f64x2.splat",
;;F64x2ExtractLane(i32) : [0xfd, 0x16] : "f64x2.extract_lane",
;;F64x2ReplaceLane(i32) : [0xfd, 0x17] : "f64x2.replace_lane",

(I8x16Eq i8x16.eq)
(I8x16Ne i8x16.ne)
(I8x16LtS i8x16.lt_s)
(I8x16LtU i8x16.lt_u)
(I8x16GtS i8x16.gt_s)
(I8x16GtU i8x16.gt_u)
(I8x16LeS i8x16.le_s)
(I8x16LeU i8x16.le_u)
(I8x16GeS i8x16.ge_s)
(I8x16GeU i8x16.ge_u)
(I16x8Eq i16x8.eq)
(I16x8Ne i16x8.ne)
(I16x8LtS i16x8.lt_s)
(I16x8LtU i16x8.lt_u)
(I16x8GtS i16x8.gt_s)
(I16x8GtU i16x8.gt_u)
(I16x8LeS i16x8.le_s)
(I16x8LeU i16x8.le_u)
(I16x8GeS i16x8.ge_s)
(I16x8GeU i16x8.ge_u)
(I32x4Eq i32x4.eq)
(I32x4Ne i32x4.ne)
(I32x4LtS i32x4.lt_s)
(I32x4LtU i32x4.lt_u)
(I32x4GtS i32x4.gt_s)
(I32x4GtU i32x4.gt_u)
(I32x4LeS i32x4.le_s)
(I32x4LeU i32x4.le_u)
(I32x4GeS i32x4.ge_s)
(I32x4GeU i32x4.ge_u)

(F32x4Eq f32x4.eq)
(F32x4Ne f32x4.ne)
(F32x4Lt f32x4.lt)
(F32x4Gt f32x4.gt)
(F32x4Le f32x4.le)
(F32x4Ge f32x4.ge)
(F64x2Eq f64x2.eq)
(F64x2Ne f64x2.ne)
(F64x2Lt f64x2.lt)
(F64x2Gt f64x2.gt)
(F64x2Le f64x2.le)
(F64x2Ge f64x2.ge)

(V128Not v128.not)
(V128And v128.and)
(V128Or v128.or)
(V128Xor v128.xor)
(V128Bitselect v128.bitselect)

(I8x16Neg i8x16.neg)
(I8x16AnyTrue i8x16.any_true)
(I8x16AllTrue i8x16.all_true)
(I8x16Shl i8x16.shl)
(I8x16ShrS i8x16.shr_s)
(I8x16ShrU i8x16.shr_u)
(I8x16Add i8x16.add)
(I8x16AddSaturateS i8x16.add_saturate_s)
(I8x16AddSaturateU i8x16.add_saturate_u)
(I8x16Sub i8x16.sub)
(I8x16SubSaturateS i8x16.sub_saturate_s)
(I8x16SubSaturateU i8x16.sub_saturate_u)
(I8x16Mul i8x16.mul)

(I16x8Neg i16x8.neg)
(I16x8AnyTrue i16x8.any_true)
(I16x8AllTrue i16x8.all_true)
(I16x8Shl i16x8.shl)
(I16x8ShrS i16x8.shr_s)
(I16x8ShrU i16x8.shr_u)
(I16x8Add i16x8.add)
(I16x8AddSaturateS i16x8.add_saturate_s)
(I16x8AddSaturateU i16x8.add_saturate_u)
(I16x8Sub i16x8.sub)
(I16x8SubSaturateS i16x8.sub_saturate_s)
(I16x8SubSaturateU i16x8.sub_saturate_u)
(I16x8Mul i16x8.mul)

(I32x4Neg i32x4.neg)
(I32x4AnyTrue i32x4.any_true)
(I32x4AllTrue i32x4.all_true)
(I32x4Shl i32x4.shl)
(I32x4ShrS i32x4.shr_s)
(I32x4ShrU i32x4.shr_u)
(I32x4Add i32x4.add)
(I32x4Sub i32x4.sub)
(I32x4Mul i32x4.mul)

(I64x2Neg i64x2.neg)
(I64x2AnyTrue i64x2.any_true)
(I64x2AllTrue i64x2.all_true)
(I64x2Shl i64x2.shl)
(I64x2ShrS i64x2.shr_s)
(I64x2ShrU i64x2.shr_u)
(I64x2Add i64x2.add)
(I64x2Sub i64x2.sub)
(I64x2Mul i64x2.mul)

(F32x4Abs f32x4.abs)
(F32x4Neg f32x4.neg)
(F32x4Sqrt f32x4.sqrt)
(F32x4Add f32x4.add)
(F32x4Sub f32x4.sub)
(F32x4Mul f32x4.mul)
(F32x4Div f32x4.div)
(F32x4Min f32x4.min)
(F32x4Max f32x4.max)

(F64x2Abs f64x2.abs)
(F64x2Neg f64x2.neg)
(F64x2Sqrt f64x2.sqrt)
(F64x2Add f64x2.add)
(F64x2Sub f64x2.sub)
(F64x2Mul f64x2.mul)
(F64x2Div f64x2.div)
(F64x2Min f64x2.min)
(F64x2Max f64x2.max)

(I32x4TruncSatF32x4S i32x4.trunc_sat_f32x4_s)
(I32x4TruncSatF32x4U i32x4.trunc_sat_f32x4_u)
(I64x2TruncSatF64x2S i64x2.trunc_sat_f64x2_s)
(I64x2TruncSatF64x2U i64x2.trunc_sat_f64x2_u)
(F32x4ConvertI32x4S f32x4.convert_i32x4_s)
(F32x4ConvertI32x4U f32x4.convert_i32x4_u)
(F64x2ConvertI64x2S f64x2.convert_i64x2_s)
(F64x2ConvertI64x2U f64x2.convert_i64x2_u)
(V8x16Swizzle v8x16.swizzle)

;;V8x16Shuffle(V8x16Shuffle) : [0xfd, 0xc1] : "v8x16.shuffle",
;;V8x16LoadSplat(MemArg<1>) : [0xfd, 0xc2] : "v8x16.load_splat",
;;V16x8LoadSplat(MemArg<2>) : [0xfd, 0xc3] : "v16x8.load_splat",
;;V32x4LoadSplat(MemArg<4>) : [0xfd, 0xc4] : "v32x4.load_splat",
;;V64x2LoadSplat(MemArg<8>) : [0xfd, 0xc5] : "v64x2.load_splat",

(I8x16NarrowI16x8S i8x16.narrow_i16x8_s)
(I8x16NarrowI16x8U i8x16.narrow_i16x8_u)
(I16x8NarrowI32x4S i16x8.narrow_i32x4_s)
(I16x8NarrowI32x4U i16x8.narrow_i32x4_u)

(I16x8WidenLowI8x16S i16x8.widen_low_i8x16_s)
(I16x8WidenHighI8x16S i16x8.widen_high_i8x16_s)
(I16x8WidenLowI8x16U i16x8.widen_low_i8x16_u)
(I16x8WidenHighI8x16u i16x8.widen_high_i8x16_u)
(I32x4WidenLowI16x8S i32x4.widen_low_i16x8_s)
(I32x4WidenHighI16x8S i32x4.widen_high_i16x8_s)
(I32x4WidenLowI16x8U i32x4.widen_low_i16x8_u)
(I32x4WidenHighI16x8u i32x4.widen_high_i16x8_u)

;;I16x8Load8x8S(MemArg<1>) : [0xfd, 0xd2] : "i16x8.load8x8_s",
;;I16x8Load8x8U(MemArg<1>) : [0xfd, 0xd3] : "i16x8.load8x8_u",
;;I32x4Load16x4S(MemArg<2>) : [0xfd, 0xd4] : "i32x4.load16x4_s",
;;I32x4Load16x4U(MemArg<2>) : [0xfd, 0xd5] : "i32x4.load16x4_u",
;;I64x2Load32x2S(MemArg<4>) : [0xfd, 0xd6] : "i64x2.load32x2_s",
;;I64x2Load32x2U(MemArg<4>) : [0xfd, 0xd7] : "i64x2.load32x2_u",
(V128Andnot v128.andnot)

`

	allInstrs := mustParseInstrs(instrs)
	all := ""
	for _, ins := range allInstrs {
		all += ins.Generate()
	}

	parseInstr := generateParseInstrution(allInstrs)

	goFile := generate(`
package ast

import "github.com/ontio/wast-parser/parser"

[Instrs]
[parseInstr]
`, map[string]interface{}{"Instrs": all, "parseInstr": parseInstr})

	err := ioutil.WriteFile("../ast/instruction.go", []byte(goFile), 0666)
	if err != nil {
		fmt.Printf("write file error: %s", err)
	}
}
