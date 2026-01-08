# Semantic Rules:

## General Rules:

    - There must be exactly one top level type
    - Any number used should be a real number always unless specified explicitly
    - A real number in this document is integers or floats written out explicitly unless specified otherwise
    - Unless explicitly stated otherwise, Number refers to a real number.
    - When a rule requires an integer, fractional values are invalid.

## Number Range Rules:

    - The left number in the range should be less than or equal to the right number in the range
    - Both the numbers in the range should be real numbers
    - The numbers in the range are optional unless specified otherwise
    - Both the numbers take a default value of 0 if and only if they are optional

## Int Rules:

    - The Number Range Rules apply here along with the below ones

## String Rules:

    - The Number Range Rules apply here along with the below ones
    - The NumberRange should contain only positive integers

## Bool Rules:

    - Either of the valid boolean values will be returned if nothing is found after the UNDERSCORE
    - A valid boolean will be returned as is
    - The boolean values are case insensitive

## Array Rules:

    - The Number Range Rules also apply here along with below ones
    - The NumberRange is the length of the array to be generated hence it has to be a non-negative integer
    - The CompoundType is the element type in the array
    - The second NumberRange is the index or index range of an element whose type will be explicitly defined
    - The second CompoundType is the type of the element/elements at that/those index/indices
    - The second NumberRange cannot have invalid values before and after the UNDERSCORE, unlike what is mentioned in the Number Range Rules
    - Invalid values in the second NumberRange include empty values, negative numbers, floating point numbers, numbers greater than of length of array specified
    - An index appearing multiple times in any range or as a standalone number will be having a merged type of both definitions

## LiteralArray Rules:

    - The InQuoteType will be the type of the generated element at the exact same index as the InQuoteType is

## Map Rules:

    - The Number Range Rules also apply here along with the below ones
    - The NumberRange is the length of the map to be generated hence it has to be a non-negative integer
    - The first CompoundType is the type of the keys to be generated for this map
    - The second CompoundType is the type of the values to be generated for this map
    - The StringLiteral is the explicit key which will be taken as is and should be enclosed in double quotes
    - Duplicate StringLiteral keys will have the type of the last specified (highest index) StringLiteral
    - The order of the generated keys will be the same as the given explicit keys
    - The third CompoundType is the type of the explicit key provided

## LiteralMap Rules:

    - The StringLiteral is the explicit key which will be taken as is and should be enclosed in double quotes
    - The CompoundType is the type of the explicit key provided
    - Duplicate StringLiteral keys will have the type of the last specified (highest index) StringLiteral
    - The order of the generated keys will be the same as the given keys
