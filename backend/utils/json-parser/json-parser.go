package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

/*
Syntax:
- valid types are string, int, float, boolean, array, map
- for string: 'string|<length>'
	- if starts with 'string|' with a number behind it, it indicates length of random string to be generated
	- if no number is present after '|', lorem ipsum style random string of length 10 is generated
	- if no 'string|' prefix is present, treat the whole input as string and return it as is
	- the string should be enclosed in single or double quotes to indicate it is a string
	- but if not enclosed in quotes, it is a fallback for some error in parsing from other types

- for int: 'int|<number>'
	- should start with 'int|', should not contain any characters except 0-9 and '-' (- is kept for negative ints)

- for float: 'float:<number>'
	- should start with 'float|', should not contain any characters except 0-9, '-' and '.' later (- is kept for negative floats, . is kept for decimals)

- for boolean: 'bool:<boolean_value>'
	- should start with 'bool|', can contain 'true' or 'false' in any case
    and may not be present indicating a random value to be filled in while generating output

- for array
	- should start with '[' and end with ']', can contain any type of data with the proper syntax
	- another syntax:
		- should start with 'array;<length>;<type1>:<type2>:...;ind:<explicit_type>;...'
			where length is the length of the array to be generated
			and type is the type of data to be filled in each element of the array
		- length:
			- length can be specified as a range of values if unsure about the exact length
				- the range can be specified with a underscore between two numbers like min_max enclosed in parentheses
				- whenever possible, the mean of them will be taken as the length
		- 'type' values:
			- string, bool, int, float
			- compound types are supported by separating by ':'
			- explicit types can be specified like 'string|4' to generate random strings of length 4
			- explicit element types can be specified instead of leaving things to chance
				- the third part is the compound type specifier
				- there can be fourth and more parts separated by ';' to specify type details
					for each element with the syntax ind:type
			- if the index is out of bounds (negative or greater than length), the element will be kept at some random index
			- if the number of specified elements exceed the array length,
				only the first 'length' elements will be considered
			- the type can be any valid single or compound primitive type
			- complex types like map and array are also supported in this syntax
				- arrays:
					- array;length;array.length.elementType[ind:type,ind:type...]
					- nested arrays are allowed in this syntax
					- this type can also be combined with the other types
					- this array can also specify explicit element types for some or all elements
						in square brackets after the type details section
					- if the ind is out of bounds (negative or greater than length),
						the element will be kept in the array at a random index
					- the elementType can be any valid single or compound primitive type
						- for compound types, the entire type specifier should be enclosed in curly braces {} and separated by ':'
					- length can be specified as a range of values if unsure about the exact length
						- the range can be specified with a underscore between two numbers like min_max enclosed in parentheses
						- whenever possible, the mean of them will be taken as the length
				- maps:
					- array;length;map.length.keyType.valueType[explicit_key:value,explicit_key:value...]
					- keyType and valueType can be any valid single or compound primitive type
					- nested maps are allowed in this syntax
					- this type can also be combined with the other types
					- some or all key values can be specified in square brackets after the type details section
					- if the specified key already exists in the map, it will be made an array of values including the new one
					- if the number of unique specified keys exceed the map length,
						only the first 'length' unique keys will be considered
					- the keyType or valueType can be any valid single or compound primitive type
						- for compound types, the entire type specifier should be enclosed in curly braces {} and separated by ':'
					- length can be specified as a range of values if unsure about the exact length
						- the range can be specified with a underscore between two numbers like min_max enclosed in parentheses
						- whenever possible, the mean of them will be taken as the length
			- example:
				- array;10;string|5:bool|true:int|432:float|-45.34:array.(2_6).{int|:string|}[0:45]:map.(3-8).string|.{int|:bool|:string|:float|-}['fixedKey':int|123,'fixedKey':int|456]

- for map
	- should start with '{' and end with '}', should contain keys in single or double quotes
	- and should have a ':' as the immediate next character followed by value in or not in quotes.
	- another syntax:
		- should start with 'map;<length>;<key_type1>:<key_type2>:...;<value_type1>:<value_type2>:...;<key>:<explicit_value>;...'
			where length is the number of key-value pairs to be generated
			and keyType and valueType are the types of keys and values respectively
		-length:
			- length can be specified as a range of values if unsure about the exact length
				- the range can be specified with a underscore between two numbers like min_max enclosed in parentheses
				- whenever possible, the mean of them will be taken as the length
		- 'keyType' and 'valueType' values:
			- string, bool, int, float
			- compound types are supported by separating by ':'
			- explicit types can be specified like 'string|4' to generate random strings of length 4
			- explicit key-value pairs can be specified instead of leaving things to chance
				- there can be fifth and more parts separated by ';' to specify type details
					for each element with the syntax key:valueType
			- if the specified key already exists in the map, it will be overwritten with the new value
			- if the number of unique specified keys exceed the map length,
				only the first 'length' unique keys will be considered
			- the keyType and valueType can be any valid single or compound primitive type
			- complex types like map and array are also supported in this syntax
				- arrays:
					- array;length;array.length.elementType[ind:type,ind:type...]
					- nested arrays are allowed in this syntax
					- this type can also be combined with the other types
					- this array can also specify explicit element types for some or all elements
						in square brackets after the type details section
					- if the ind is out of bounds (negative or greater than length),
						the element will be kept in the array at a random index
					- the elementType can be any valid single or compound primitive type
						- for compound types, the entire type specifier should be enclosed in curly braces {} and separated by ':'
					- length can be specified as a range of values if unsure about the exact length
						- the range can be specified with a underscore between two numbers like min_max enclosed in parentheses
						- whenever possible, the mean of them will be taken as the length
				- maps:
					- array;length;map.length.keyType.valueType[key:value,ind:type...]
					- keyType and valueType can be any valid single or compound primitive type
					- nested maps are allowed in this syntax
					- this type can also be combined with the other types
					- some or all key values can be specified in square brackets after the type details section
					- if the specified key already exists in the map, it will be made an array of values including the new one
					- if the number of unique specified keys exceed the map length,
						only the first 'length' unique keys will be considered
					- the elementType can be any valid single or compound primitive type
						- for compound types, the entire type specifier should be enclosed in curly braces {} and separated by ':'
					- length can be specified as a range of values if unsure about the exact length
						- the range can be specified with a underscore between two numbers like min_max enclosed in parentheses
						- whenever possible, the mean of them will be taken as the length
			- example:
				- map;10;string|6;int|:string|:float|:bool|:array.(2_5).{int|:string|}:map.(3-7).string|.int|;'key1':int|123;'key2':string|4;'key3':float|-45.67
*/
func parseData(inp string) any {
	if len(inp) == 0 {
		return ""
	}

	fmt.Printf("inp: %s, first char: %c\n", inp, inp[0])

	if inp[0] == '{' && inp[len(inp)-1] == '}' {
		println("map")
		// map
		mapWithoutBraces := inp[1 : len(inp)-1]
		resultMap := make(map[string]any)
		currentKey := ""
		currentValue := ""
		parsingKey := true
		bracketCount := 0
		bracesCount := 0
		for i := 0; i < len(mapWithoutBraces); i++ {
			char := mapWithoutBraces[i]
			if parsingKey {
				if char == ':' {
					parsingKey = false
				} else {
					currentKey += string(char)
				}
			} else {
				if char == ',' && bracketCount == 0 && bracesCount == 0 {
					// end of current key-value pair
					key := strings.TrimSpace(currentKey)
					// strip surrounding quotes from key if present
					if len(key) >= 2 {
						if (key[0] == '\'' && key[len(key)-1] == '\'') || (key[0] == '"' && key[len(key)-1] == '"') {
							key = key[1 : len(key)-1]
						}
					}
					resultMap[key] = parseData(strings.TrimSpace(currentValue))
					currentKey = ""
					currentValue = ""
					parsingKey = true
				} else {
					currentValue += string(char)
				}
				switch char {
				case '[':
					bracketCount++
				case ']':
					bracketCount--
				case '{':
					bracesCount++
				case '}':
					bracesCount--
				}
			}
		}
		// finalize last key-value pair if present
		if strings.TrimSpace(currentKey) != "" {
			key := strings.TrimSpace(currentKey)
			if len(key) >= 2 {
				if (key[0] == '\'' && key[len(key)-1] == '\'') || (key[0] == '"' && key[len(key)-1] == '"') {
					key = key[1 : len(key)-1]
				}
			}
			resultMap[key] = parseData(strings.TrimSpace(currentValue))
		}
		return resultMap
	} else if (inp[:5] == "array;" && strings.Count(inp, ";") >= 2) {
		println("array with array;length;type syntax")
		// array with array;length;type syntax
		parts := strings.Split(inp, ";")
		if len(parts) < 3 {
			return []any{}
		}
		lengthPart := parts[1]
		typePart := parts[2]
		generatedArray := []any{}
		// we take the number of specified elements and subtract those many from the length
		// then we look at the compound types specified
		// then using random numbers we generate the chances of which type comes up from the compound types
		// then we create elements of those types and fill the array skipping the indices which have explicit types specified
		// finally we fill in the explicit types at the specified indices with the actual elements
		// length parsing
		length := 5 // default length
		if strings.HasPrefix(lengthPart, "(") && strings.HasSuffix(lengthPart, ")") {
			rangePart := lengthPart[1 : len(lengthPart)-1]
			rangeBounds := strings.Split(rangePart, "_")
			if len(rangeBounds) == 2 {
				min, err1 := strconv.Atoi(rangeBounds[0])
				max, err2 := strconv.Atoi(rangeBounds[1])
				if err1 == nil && err2 == nil {
					length = (min + max)/2
				} else if err1 == nil && err2 != nil {
					length = min
				} else if err1 != nil && err2 == nil {
					length = max
				}
			} else if len(rangeBounds) == 1 && rangeBounds[0] != "" {
				if v, err := strconv.Atoi(rangeBounds[0]); err == nil {
					length = v
				}
			}
		} else {
			if v, err := strconv.Atoi(lengthPart); err == nil {
				length = v
			}
		}
		// we check the number of specified valid elements and put them in an array
		explicitElements := make(map[int]any)
		givenExplicitElements := parts[3:]
		for _, elemSpec := range givenExplicitElements {
			elemParts := strings.SplitN(elemSpec, ":", 2)
			if len(elemParts) != 2 {
				continue
			}
			indexPart := elemParts[0]
			typeSpecPart := elemParts[1]
			if index, err := strconv.Atoi(indexPart); err == nil {
				if index >= 0 && index < length {
					explicitElements[index] = generateRandomType(typeSpecPart)
				}
			}
		}
	} else if inp[0] == '[' && inp[len(inp)-1] == ']' {
		println("array")
		// array
		arrayWithoutBrackets := inp[1 : len(inp)-1]
		var resultArray []any = []any{}
		currentElement := ""
		bracketCount := 0
		bracesCount := 0
		for i := 0; i < len(arrayWithoutBrackets); i++ {
			char := arrayWithoutBrackets[i]
			if char == ',' && bracketCount == 0 && bracesCount == 0 {
				// end of current element
				resultArray = append(resultArray, parseData(strings.TrimSpace(currentElement)))
				currentElement = ""
			} else {
				currentElement += string(char)
				switch char {
				case '[':
					bracketCount++
				case ']':
					bracketCount--
				case '{':
					bracesCount++
				case '}':
					bracesCount--
				}
			}
		}
		// append last element if any
		if strings.TrimSpace(currentElement) != "" {
			resultArray = append(resultArray, parseData(strings.TrimSpace(currentElement)))
		}
		return resultArray
	} else if strings.HasPrefix(inp, "bool|") {
		println("bool")
		// bool
		var boolPart string = strings.ToLower(inp[5:])
		switch boolPart {
		case "true":
			return true
		case "false":
			return false
		default:
			// generate a random bool
			return rand.Intn(2) == 1
		}
	} else if strings.HasPrefix(inp, "int|") {
		println("int")
		// int
		numPart := inp[4:]
		switch numPart {
		case "":
			return rand.Intn(150)
		case "-":
			// negative random
			return -rand.Intn(150)
		}
		if v, err := strconv.Atoi(numPart); err == nil {
			return v
		}
		// fallback random
		return rand.Intn(150)
	} else if strings.HasPrefix(inp, "float|") {
		println("float")
		// float
		numPart := inp[6:]
		switch numPart {
		case "":
			return rand.Float64() * 150
		case "-":
			// negative random
			return -rand.Float64() * 150
		}
		if v, err := strconv.ParseFloat(numPart, 64); err == nil {
			return v
		}
		return rand.Float64() * 150
	} else if strings.HasPrefix(inp, "string|") {
		println("string with prefix")
		// string
		lengthPart := inp[7:]
		if lengthPart == "" {
			lengthPart = "10"
		}
		if v, err := strconv.Atoi(lengthPart); err == nil {
			return generateRandomString(v)
		}
		return generateRandomString(10)
	}
	println("plain string")
	// plain string
	return strings.Trim(inp, " '\"") // strip surrounding quotes or spaces if any
}
