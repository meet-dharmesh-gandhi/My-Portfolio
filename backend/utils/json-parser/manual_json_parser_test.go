package utils

import (
	"reflect"
	"testing"
)

func TestParseData(t *testing.T) {
	cases := []string{
		"string|101",
		"string|",
		"aabbddcc",
		"bool|",
		"bool|true",
		"bool|false",
		"int|",
		"int|0",
		"int|45643",
		"int|-45643",
		"float|",
		"float|-",
		"float|57432",
		"float|-57432",
		"float|34567.4354243",
		"float|-34567.4354243",
		"[string|4]",
		"['a', 'b', 'c', int|123, bool|true, float|12.34, string|4, [{'hardcoded string': v, 'generated string': string|4, 'int': int|, 'float': float|, 'array': ['looks like this should parse', bool|, int|, string|4], 'map': {'nested map key': 'nested map value', 'nested map int': int|}}, bool|false, string|4], float|-56.78, int|-987, 'lorem ipsum', string]",
		"{'key1': 'value1', 'key2': int|2345, 'key3': float|23.45, 'key4': bool|false, 'key5': string|4, 'key6': [int|1, int|2, int|3, string|4], 'key7': {'nestedKey1': 'nestedValue1', 'nestedKey2': float|-67.89, 'nestedKey3': bool|true}}",
		"array;5;string|",
	}

	for _, c := range cases {
		println("\n-----------------------------------------------------------------------------\n")
		t.Logf("Input: %s\n--------------------\n", c)
		result := parseData(c)
		t.Logf("Output: Type: %v, %#v\n\n\n\n", reflect.TypeOf(result), result)
	}
}
