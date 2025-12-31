package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func runSafe(inp string) (res any, panicked bool, panicVal any) {
    defer func() {
        if r := recover(); r != nil {
            panicked = true
            panicVal = r
        }
    }()
    res = parseData(inp)
    return
}

func TestParseDataDiagnostics(t *testing.T) {
    cases := []struct {
        name       string
        in         string
        wantPanic  bool
        wantKind   string // "string", "bool", "int", "float", "array", "map"
        wantExact  any
        wantInRange bool // when true, expect numeric random in [0,150)
    }{
        {"empty", "", false, "string", "", false},
        {"plain string", "hello world", false, "string", "hello world", false},
        {"bool true", "bool|true", false, "bool", true, false},
        {"bool false mixed case", "bool|False", false, "bool", false, false},
        {"bool random", "bool|", false, "bool", nil, false},
        {"int valid", "int|123", false, "int", 123, false},
        {"int negative", "int|-45", false, "int", -45, false},
        {"int malformed", "int|12a3", false, "int", nil, true},
        {"float valid", "float|12.34", false, "float", 12.34, false},
        {"float negative", "float|-0.56", false, "float", -0.56, false},
        {"float malformed", "float|12.a4", false, "float", nil, true},
        {"empty array", "[]", false, "array", nil, false},
        {"simple array", "[int|1, int|2, string]", false, "array", nil, false},
        {"empty map", "{}", false, "map", nil, false},
        {"simple map", "{'a':int|123}", false, "map", nil, false},
        {"nested structures", "{\"k\":[int|1,{'x':float|1.2}]}", false, "map", nil, false},
    }

    for _, c := range cases {
        res, pan, pv := runSafe(c.in)
        kind := "panic"
        if pan {
            if !c.wantPanic {
                t.Errorf("%s: unexpected panic for input %q: %v", c.name, c.in, pv)
            } else {
                t.Logf("%s: panicked as expected: %v", c.name, pv)
            }
            continue
        }

        if c.wantPanic {
            t.Errorf("%s: expected panic for input %q but got %+v", c.name, c.in, res)
            continue
        }

        if res != nil {
            rtype := reflect.TypeOf(res)
            switch rtype.Kind() {
            case reflect.Bool:
                kind = "bool"
            case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
                kind = "int"
            case reflect.Float32, reflect.Float64:
                kind = "float"
            case reflect.Slice:
                kind = "array"
            case reflect.Map:
                kind = "map"
            case reflect.String:
                kind = "string"
            default:
                kind = rtype.Kind().String()
            }
        } else {
            kind = "nil"
        }

        // Basic kind check
        if c.wantKind != "" && kind != c.wantKind {
            t.Errorf("%s: input=%q -> want kind %q but got kind %q (value=%#v)", c.name, c.in, c.wantKind, kind, res)
            continue
        }

        // Exact checks for deterministic cases
        if c.wantExact != nil {
            if !reflect.DeepEqual(res, c.wantExact) {
                t.Errorf("%s: input=%q -> want exact %#v but got %#v", c.name, c.in, c.wantExact, res)
            } else {
                t.Logf("%s: input=%q -> matched exact value %#v", c.name, c.in, res)
            }
            continue
        }

        // Range check for numeric fallbacks
        if c.wantInRange {
            switch v := res.(type) {
            case int:
                if v < 0 || v >= 150 {
                    t.Errorf("%s: input=%q -> expected random int in [0,150) but got %d", c.name, c.in, v)
                } else {
                    t.Logf("%s: input=%q -> got random int %d (in range)", c.name, c.in, v)
                }
            case float64:
                if v < 0 || v >= 150 {
                    t.Errorf("%s: input=%q -> expected random float in [0,150) but got %v", c.name, c.in, v)
                } else {
                    t.Logf("%s: input=%q -> got random float %v (in range)", c.name, c.in, v)
                }
            default:
                t.Errorf("%s: input=%q -> expected numeric random but got type %T value %#v", c.name, c.in, res, res)
            }
            continue
        }

        // For arrays and maps, just log the concrete value for manual inspection
        if kind == "array" || kind == "map" {
            t.Logf("%s: input=%q -> got %s value: %#v", c.name, c.in, kind, res)
            continue
        }

        t.Logf("%s: input=%q -> got kind=%s value=%#v", c.name, c.in, kind, res)
    }

    // Quick summary print (visible with -v)
    fmt.Println("Diagnostic run complete — inspect test output for mismatches and panics.")
}
