package utils

import (
	"math/rand"
	"strconv"
	"strings"
)

func generateRandomString(length int) string {
	var wordset = [...]string{"lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit",
		"a", "ac", "accumsan", "ad", "aenean", "aliquam", "aliquet", "ante",
		"aptent", "arcu", "at", "auctor", "augue", "bibendum", "blandit",
		"class", "commodo", "condimentum", "congue", "consequat", "conubia",
		"convallis", "cras", "cubilia", "curabitur", "curae", "cursus",
		"dapibus", "diam", "dictum", "dictumst", "dignissim", "dis", "donec",
		"dui", "duis", "efficitur", "egestas", "eget", "eleifend", "elementum",
		"enim", "erat", "eros", "est", "et", "etiam", "eu", "euismod", "ex",
		"facilisi", "facilisis", "fames", "faucibus", "felis", "fermentum",
		"feugiat", "finibus", "fringilla", "fusce", "gravida", "habitant",
		"habitasse", "hac", "hendrerit", "himenaeos", "iaculis", "id",
		"imperdiet", "in", "inceptos", "integer", "interdum", "justo",
		"lacinia", "lacus", "laoreet", "lectus", "leo", "libero", "ligula",
		"litora", "lobortis", "luctus", "maecenas", "magna", "magnis",
		"malesuada", "massa", "mattis", "mauris", "maximus", "metus", "mi",
		"molestie", "mollis", "montes", "morbi", "mus", "nam", "nascetur",
		"natoque", "nec", "neque", "netus", "nibh", "nisi", "nisl", "non",
		"nostra", "nulla", "nullam", "nunc", "odio", "orci", "ornare",
		"parturient", "pellentesque", "penatibus", "per", "pharetra",
		"phasellus", "placerat", "platea", "porta", "porttitor", "posuere",
		"potenti", "praesent", "pretium", "primis", "proin", "pulvinar",
		"purus", "quam", "quis", "quisque", "rhoncus", "ridiculus", "risus",
		"rutrum", "sagittis", "sapien", "scelerisque", "sed", "sem", "semper",
		"senectus", "sociosqu", "sodales", "sollicitudin", "suscipit",
		"suspendisse", "taciti", "tellus", "tempor", "tempus", "tincidunt",
		"torquent", "tortor", "tristique", "turpis", "ullamcorper", "ultrices",
		"ultricies", "urna", "ut", "varius", "vehicula", "vel", "velit",
		"venenatis", "vestibulum", "vitae", "vivamus", "viverra", "volutpat",
		"vulputate"}
	b := make([]string, length)
	for i := range b {
		b[i] = wordset[rand.Intn(len(wordset))]
	}
	return strings.Join(b, " ")
}

func generateRandomInt(min, max int, isNegative bool) int {
	randomInt := rand.Intn(max-min) + min
	if isNegative {
		return -randomInt
	}
	return randomInt
}

func generateRandomFloat(min, max float64, isNegative bool) float64 {
	randomFloat := min + rand.Float64()*(max-min)
	if isNegative {
		return -randomFloat
	}
	return randomFloat
}

func generateRandomBool() bool {
	return rand.Intn(2) == 1
}

func generateRandomArray(length int, elementTypes []string, explicitElements map[string]any) []any {
	if length <= 0 || len(elementTypes) == 0 || length != len(elementTypes) {
		return []any{}
	}
	randomIndexElements := make(map[int]any)
	for k, v := range explicitElements {
		index, err := strconv.Atoi(strings.Split(k, ":")[0])
		if err == nil && index >= 0 && index < length {
			randomIndexElements[index] = v
		} else if err == nil {
			// assign to a random index
			randIndex := rand.Intn(length)
			randomIndexElements[randIndex] = v
		}
	}
	array := make([]any, length)
	for i := 0; i < length; i++ {
		if explicitElements != nil && explicitElements[strconv.Itoa(i)] != nil {
			array[i] = explicitElements[strconv.Itoa(i)]
		} else if randomIndexElements[i] != nil {
			array[i] = randomIndexElements[i]
		} else {
			array[i] = generateRandomType(elementTypes[i])
		}
	}
	return array
}

func generateRandomMap(length int, keyTypes []string, valueTypes []string, explicitElements map[string]any) map[any]any {
	if length <= 0 || len(keyTypes) == 0 || len(valueTypes) == 0 || length != len(keyTypes) || length != len(valueTypes) || len(keyTypes) != len(valueTypes) {
		return map[any]any{}
	}
	m := make(map[any]any)
	for i := range explicitElements {
		m[i] = explicitElements[i]
	}
	for i := 0; i < length-len(explicitElements); i++ {
		key := generateRandomType(keyTypes[i])
		value := generateRandomType(valueTypes[i])
		m[key] = value
	}
	return m
}

/*

This function generates a random value based on the provided type string. The type string can specify basic types like string, int, float, and bool,
as well as more complex types like arrays and maps with specific lengths and element types.

the supported type strings are:

- "<string>": gives back the same string as is, no parsing is done
- "string|": generates a random string of length 10
- "string|<length>": generates a random string of the specified length
- "string|<string>": gives back the string as is, can be used when parsing is possible but not needed
- "int|": generates a random integer between 0 and 150
- "int|<number>": generates a random integer between 0 and the specified number
- "int|-": generates a random negative integer between 0 and -150
- "int|-<number>": generates a random negative integer between 0 and the specified negative number
- "float|": generates a random float between 0 and 150
- "float|<number>": generates a random float between 0 and the specified number
- "float|-": generates a random negative float between 0 and -150
- "float|-<number>": generates a random negative float between 0 and the specified negative number
- "bool|": generates a random boolean value
- "bool|<boolean>": returns the boolean as is (true or false, case insensitive)
- "{<key>:<valueType>,...}":
	generates a map of given explicit keys and value types.
	- <key> is an explicit string which will be kept as the map key
	- <valueType> is the type of the value to be generated, can be anything in this list
	- there has to be a comma at the end of the map if it contains a single key-value pair
- "[<elementType>,...]":
	generates an array of the given types
	- <elementType> is the type of the element to generate, can be anything in this list
- "array;<length>;<type1>:<type2>:...;ind:<explicit_type>;...":
	generates an array of the specified length with elements of the specified types,
	allowing for explicit element values
- "map;<length>;<key_type1>:<key_type2>:...;<value_type1>:<value_type2>:...;<explicit_key>:<value_type>;...":
	generates a map of the specified length with keys and values of the specified types,
	allowing for explicit key-value pairs
- "array.<length>.<elementType>[<ind1>:<type1>,<ind2>:<type2>...]":
	generates an array of specified length and element types with optional explicit elements at given indices
	- <length> can be a fixed number or a range in the format (min_max)
	- <elementType> can be any valid single or compound primitive type enclosed in parantheses () separated by ':' for compound types
	- explicit elements are specified as comma separated index:type pairs,
		where the index is the position in the array and type is the type of the element
- "map.<length>.<keyType>.<valueType>[<explicit_key1>:<value1>,<explicit_key2>:<value2>...]":
	generates a map of specified length with given key and value types and optional explicit key-value pairs
	- <length> can be a fixed number or a range in the format (min_max)
	- <keyType> and <valueType> can be any valid single or compound primitive type enclosed in parantehese () separated by ':' for compound types
	- explicit key-value pairs are specified as comma separated key:value pairs,
		where key is the explicit key and value is the corresponding value

If the type string does not match any of the supported formats, it is returned as is.

*/
func generateRandomType(typ string) any {
	if strings.HasPrefix(typ, "string|") { // string|length syntax
		return parseTypeString(typ)
	} else if strings.HasPrefix(typ, "array") { // array syntax - could be array; or array.
		return parseTypeArray(typ)
	} else if strings.HasPrefix(typ, "map") { // map syntax - could be map; or map.
		return parseTypeMap(typ)
	} else if strings.HasPrefix(typ, "int|-") { // int|-number syntax
		return parseTypeInt(typ, true)
	} else if strings.HasPrefix(typ, "int|") { // int|number syntax
		return parseTypeInt(typ, false)
	} else if strings.HasPrefix(typ, "float|-") { // float|-number syntax
		return parseTypeFloat(typ, true)
	} else if strings.HasPrefix(typ, "float|") { // float|number syntax
		return parseTypeFloat(typ, false)
	} else if strings.HasPrefix(typ, "bool|") { // bool|boolean syntax
		return parseTypeBool(typ)
	}

	return typ
}

/*
This function checks if the provided type string is a valid string type.
*/
func isValidTypeString(typ string) bool {
	return strings.HasPrefix(typ, "string|")
}

/*
This function parses an explicit type string and returns the corresponding value.
It panics if the type string is not valid (does not start with 'string|').
*/
func parseTypeString(typ string) any {
	if isValidTypeString(typ) {
		lengthPart := typ[7:]
		if lengthPart != "" {
			if v, err := strconv.Atoi(lengthPart); err == nil {
				return generateRandomString(v)
			}
		}
		return generateRandomString(10)
	}
	panic("[String parser] Invalid type string: " + typ)
}

/*
This function checks if the provided type string is a valid boolean type.
*/
func isValidTypeBool(typ string) bool {
	return strings.HasPrefix(typ, "bool|")
}

/*
This function parses an explicit type string and returns the corresponding boolean value.
It panics if the type string is not valid (does not start with 'bool|').
*/
func parseTypeBool(typ string) bool {
	if isValidTypeBool(typ) {
		booleanPart := strings.ToLower(typ[5:])
		if booleanPart == "true" {
			return true
		} else if booleanPart == "false" {
			return false
		}
		return generateRandomBool()
	} else {
		panic("[Bool Parser] Invalid type string: " + typ)
	}
}

/*
This function checks if the provided type string is a valid integer type.
*/
func isValidTypeInt(typ string) bool {
	return strings.HasPrefix(typ, "int|")
}

/*
This function parses an explicit type string and returns the corresponding integer value.
It panics if the type string is not valid (does not start with 'int|' or 'int|-').
*/
func parseTypeInt(typ string, isNegative bool) int {
	if isValidTypeInt(typ) {
		numberPart := typ[4:]
		if isNegative {
			numberPart = typ[5:]
		}
		if numberPart != "" {
			if v, err := strconv.Atoi(numberPart); err == nil {
				return generateRandomInt(0, v, isNegative)
			}
		}
		return generateRandomInt(0, 150, isNegative)
	}
	panic("[Int Parser] Invalid type string: " + typ)
}

/*
This function checks if the provided type string is a valid float type.
*/
func isValidTypeFloat(typ string) bool {
	return strings.HasPrefix(typ, "float|")
}

/*
This function parses an explicit type string and returns the corresponding float value.
It panics if the type string is not valid (does not start with 'float|' or 'float|-').
*/
func parseTypeFloat(typ string, isNegative bool) float64 {
	if isValidTypeFloat(typ) {
		numberPart := typ[6:]
		if isNegative {
			numberPart = typ[7:]
		}
		if numberPart != "" {
			if v, err := strconv.ParseFloat(numberPart, 64); err == nil {
				return generateRandomFloat(0, v, isNegative)
			}
		}
		return generateRandomFloat(0, 150, isNegative)
	}
	panic("[Float Parser] Invalid type string: " + typ)
}

/*
This function checks if the provided type string is a valid array type.
return values:
	-1 means invalid
	1 means array; syntax
	2 means array. syntax
*/
func isValidTypeArray(typ string) int {
	if strings.HasPrefix(typ, "array;") {
		return 1
	} else if strings.HasPrefix(typ, "array.") {
		return 2
	}
	return -1
}

/*
This function parses an explicit type string and returns the corresponding array or map value.
It returns typ as is if the type string is not a valid array type.
It panics if the type string is invalid.
*/
func parseTypeArray(typ string) any {
	if isValidTypeArray(typ) == -1 {
		panic("[Array Parser] Invalid type string: " + typ)
	} else if isValidTypeArray(typ) == 1 { // array; syntax
		parts := strings.Split(typ, ";")
		if len(parts) >= 3 {
			lengthPart := parts[1]
			explicitElements := parts[3:]
			// parse the explicit elements
			explicitMap := make(map[string]any)
			for _, elem := range explicitElements {
				elemParts := strings.SplitN(elem, ":", 2)
				if len(elemParts) >= 2 {
					explicitMap[elemParts[0]] = generateRandomType(elemParts[1])
				}
			}
			if lengthPart != "" {
				if v, err := strconv.Atoi(lengthPart); err == nil {
					return generateRandomArray(v, strings.Split(parts[2], ":"), explicitMap)
				}
			}
			return generateRandomArray(5, strings.Split(parts[2], ":"), explicitMap)
		}
		return generateRandomArray(5, []string{"string|2"}, nil)
	} else if isValidTypeArray(typ) == 2 { // array. syntax
		parts := strings.Split(typ, ".")
		if len(parts) == 3 {
			lengthPart := parts[1]
			typeAndElements := strings.SplitN(parts[3], "[", 2)
			elementTypes := strings.Split(strings.Trim(typeAndElements[0], "()"), ":") // compound types enclosed in ()
			explicitElements := []string{}
			if len(typeAndElements) == 2 {
				explicitPart := strings.TrimSuffix(typeAndElements[1], "]")
				explicitElements = strings.Split(explicitPart, ",")
			}
			explicitMap := parseKeyValPairs(explicitElements)
			if lengthPart != "" {
				if v, err := strconv.Atoi(lengthPart); err == nil {
					return generateRandomArray(v, elementTypes, explicitMap)
				}
			}
			return generateRandomArray(5, elementTypes, explicitMap)
		}
		return generateRandomArray(5, []string{"string|2"}, nil)
	}
	return typ
}

/*
This function checks if the provided type string is a valid map type.
return values:
	-1 means invalid
	1 means map; syntax
	2 means map. syntax
*/
func isValidTypeMap(typ string) int {
	if strings.HasPrefix(typ, "map;") {
		return 1
	} else if strings.HasPrefix(typ, "map.") {
		return 2
	}
	return -1
}

/*
This function parses an explicit type string and returns the corresponding map value.
It returns typ as is if the type string is not a valid map type.
It panics if the type string is invalid.
*/
func parseTypeMap(typ string) any {
	if isValidTypeMap(typ) == -1 {
		panic("[Map Parser] Invalid type string: " + typ)
	} else if isValidTypeMap(typ) == 1 { // map; syntax
		parts := strings.Split(typ, ";")
		if len(parts) >= 4 {
			lengthPart := parts[1]
			keyTypes := strings.Split(parts[2], ":")
			valueTypes := strings.Split(parts[3], ":")
			explicitElements := parts[4:]
			explicitMap := parseKeyValPairs(explicitElements)
			if lengthPart != "" {
				if v, err := strconv.Atoi(lengthPart); err == nil {
					return generateRandomMap(v, keyTypes, valueTypes, explicitMap)
				}
			}
			return generateRandomMap(5, keyTypes, valueTypes, explicitMap)
		}
		return generateRandomMap(5, []string{"string|2"}, []string{"string|2"}, nil)
	} else if isValidTypeMap(typ) == 2 { // map. syntax
		parts := strings.Split(typ, ".")
		if len(parts) == 4 {
			lengthPart := parts[1]
			keyTypes := strings.Split(strings.Trim(parts[2], "()"), ":") // compound types enclosed in ()
			typeAndElements := strings.SplitN(parts[3], "[", 2)
			valueTypes := strings.Split(strings.Trim(typeAndElements[0], "()"), ":")
			explicitElements := []string{}
			if len(typeAndElements) == 2 {
				explicitPart := strings.TrimSuffix(typeAndElements[1], "]")
				explicitElements = strings.Split(explicitPart, ",")
			}
			explicitMap := parseKeyValPairs(explicitElements)
			if lengthPart != "" {
				if v, err := strconv.Atoi(lengthPart); err == nil {
					return generateRandomMap(v, keyTypes, valueTypes, explicitMap)
				}
			}
			return generateRandomMap(5, keyTypes, valueTypes, explicitMap)
		}
		return generateRandomMap(5, []string{"string|2"}, []string{"string|2"}, nil)
	}
	return typ
}

/*
This function parses a list of key-value pair strings and returns a map.
Each string in the list should be in the format "key:value".
*/
func parseKeyValPairs(pairs []string) map[string]any {
	result := make(map[string]any)
	for _, pair := range pairs {
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) >= 2 {
			result[parts[0]] = parts[1]
		}
	}
	return result
}

/*
example inputs:
{
	'ok': 'some',
	'another': ['some 1', 'some[]'],
	'another2': {
		'key1': '{}',
		'key2': 'ok\'s',
	}
}

What will this format:
- extra whitespaces, tab and newline characters between any combination of '{', '\'' or '"', '[', ']', '}', ':', ','

result of the given example after testing:
{'ok':'some','another':['some 1','some[]'],'another2':{'key1':'{}','key2':'ok\\'s',}}
*/
func formatMapString(mapString string) string {
	var formattedString strings.Builder
	var inQuote rune = 0
	escape := false
	for _, ele := range mapString {
		if ele == '\'' || ele == '"' {
			if inQuote == 0 {
				inQuote = ele
			} else if !escape {
				inQuote = 0
			}
		} else if ele == '\\' {
			println("writing: " + string(ele) + ", in condition 1, " + formattedString.String())
			escape = true
			formattedString.WriteRune(ele)
			continue
		}
		if (ele != ' ' && ele != '\n' && ele != '\t' && ele != '\r') || inQuote != 0 {
			println("writing: " + string(ele) + ", in condition 2, " + formattedString.String())
			formattedString.WriteRune(ele)
		}
		escape = false
	}
	println("final string: " + formattedString.String())
	return formattedString.String()
}
