package ast

import (
	"fmt"
	"strings"
)

func PrintType(node TypeNode, level int) {
	switch n := node.(type) {
	case Int:
		fmt.Printf("%sInt\n", indent(level))
		if n.Number != nil {
			printRange(n.Number, level+1)
		}
	case String:
		fmt.Printf("%sString\n", indent(level))
		if n.Length != nil {
			fmt.Printf("%sLength:\n", indent(level+1))
			printRange(n.Length, level+2)
		}
	case Bool:
		fmt.Printf("%sBool\n", indent(level))
		if n.BoolValue != nil {
			fmt.Printf("%sValue: %v\n", indent(level+1), *n.BoolValue)
		} else {
			fmt.Printf("%sValue is nil\n", indent(level+1))
		}
	case *Array:
		fmt.Printf("%sArray\n", indent(level))
		fmt.Printf("%sLength:\n", indent(level))
		printRange(&n.Length, level+1)
		fmt.Printf("%sElement type(s):\n", indent(level))
		printCompoundType(n.ElementType, level+1)
		fmt.Printf("%sOverrides:\n", indent(level))
		for _, e := range n.ExplicitTypes {
			fmt.Printf("%sIndex:\n", indent(level+1))
			printRange(&e.Index, level+2)
			fmt.Printf("%sType:\n", indent(level+1))
			printCompoundType(e.ElementType, level+2)
		}
	case LiteralArray:
		fmt.Printf("%sLiteral Array\n", indent(level))
		for ind, e := range n.ElementTypes {
			fmt.Printf("%s[%d]:\n", indent(level+1), ind)
			printCompoundType(e, level+2)
		}
	case Map:
		fmt.Printf("%sMap\n", indent(level))
		fmt.Printf("%sLength\n", indent(level))
		printRange(&n.Length, level+1)
		fmt.Printf("%sKey type(s):\n", indent(level))
		printCompoundType(n.KeyType, level+1)
		fmt.Printf("%sValue type(s):\n", indent(level))
		printCompoundType(n.ValueType, level+1)
		fmt.Printf("%sOverrides:\n", indent(level))
		for _, e := range n.ExplicitTypes {
			fmt.Printf("%sKey:\n", indent(level+1))
			fmt.Printf("%s%s\n", indent(level+2), e.ExplicitKey)
			fmt.Printf("%sType:\n", indent(level+1))
			printCompoundType(e.ValueType, level+2)
		}
	case LiteralMap:
		fmt.Printf("%sLiteral Map\n", indent(level))
		for _, e := range n.Elements {
			fmt.Printf("%sKey:\n", indent(level+1))
			fmt.Printf("%s%s\n", indent(level+2), e.Key)
			printCompoundType(e.ValueType, level+1)
		}
	default:
		fmt.Printf("%sUnhandled case: %T\n", indent(level), node)
	}
}

func indent(level int) string {
	return strings.Repeat("| ", level)
}

func printRange(number *NumberRange, level int) {
	fmt.Printf("%sRange:\n", indent(level))
	if number.Min != nil {
		fmt.Printf("%sMin: %v\n", indent(level+1), *number.Min)
	}
	if number.Max != nil {
		fmt.Printf("%sMax: %v\n", indent(level+1), *number.Max)
	}
	if number.Max == nil && number.Min == nil {
		fmt.Printf("%sRange skipped entirely\n", indent(level))
	}
}

func printCompoundType(compoundType CompoundType, level int) {
	if len(compoundType.Options) == 1 {
		PrintType(compoundType.Options[0], level)
		return
	}

	fmt.Printf("%sCompound\n", indent(level))
	for _, singularType := range compoundType.Options {
		PrintType(singularType, level+1)
	}
}
