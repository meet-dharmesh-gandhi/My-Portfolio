package ast

type TypeKind int

const (
	KindInt TypeKind = iota
	KindString
	KindBool
	KindArray
	KindLiteralArray
	KindMap
	KindLiteralMap
)

type TypeNode interface {
	Kind() TypeKind
}

type NumberRange struct {
	Min *float64
	Max *float64
}

type CompoundType struct {
	Options []TypeNode
}

type Int struct {
	Number *NumberRange
}

func (Int) Kind() TypeKind {
	return KindInt
}

type String struct {
	Length *NumberRange
}

func (String) Kind() TypeKind {
	return KindString
}

type Bool struct {
	BoolValue *bool
}

func (Bool) Kind() TypeKind {
	return KindBool
}

type Array struct {
	Length        NumberRange
	ElementType   CompoundType
	ExplicitTypes []ArrayOverrides
}

type ArrayOverrides struct {
	Index       NumberRange
	ElementType CompoundType
}

func (Array) Kind() TypeKind {
	return KindArray
}

type LiteralArray struct {
	ElementTypes []CompoundType
}

func (LiteralArray) Kind() TypeKind {
	return KindLiteralArray
}

type Map struct {
	Length        NumberRange
	KeyType       CompoundType
	ValueType     CompoundType
	ExplicitTypes []MapOverrides
}

type MapOverrides struct {
	ExplicitKey string
	ValueType   CompoundType
}

func (Map) Kind() TypeKind {
	return KindMap
}

type LiteralMap struct {
	Elements []MapEntry
}

type MapEntry struct {
	Key       string
	ValueType CompoundType
}

func (LiteralMap) Kind() TypeKind {
	return KindLiteralMap
}
