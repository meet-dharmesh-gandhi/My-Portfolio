## Grammar for the dummy input string:

1. There can be only one **type** in one input string
2. A **type** can be one of the following:
    1. general
    2. string
    3. integer
    4. float
    5. boolean
    6. array
    7. map
3. A **general** is written as:
    - optionally the word 'general|'
4. A **string** is written as:
    - the word 'string'
    - followed by '|'
    - optionally followed by a number
        - the digits of the number should be less than 8
        - the number should be greater than 0
5. A **integer** is written as:
    - the word 'int'
    - followed by '|'
    - exactly one of the two must be present:
        - optionally followed by a number
            - the digits of the number should be less than 8 excluding the negative sign for negative numbers
            - the number can be negative, positive or 0
        - optionally followed by a '('
            - optionally followed by a number
                - the digits of the number should be less than 8 excluding the negative sign for negative numbers
                - the number can be negative, positive or 0
            - followed by a '\_'
            - optionally followed by a number, required if the first number is skipped
                - the digits of the number should be less than 8 excluding the negative sign for negative numbers
                - the number can be negative, positive or 0
            - followed by a ')'
6. A **float** is written as:
    - the word 'float'
    - followed by '|'
    - exactly one of the two must be present:
        - optionally followed by a number
            - the digits of the number's whole part should be less than 8 excluding the negative sign for negative numbers
            - the digits of the number's decimal part should be less than 12 excluding the decimal point
            - the number can be negative, positive or 0
        - optionally followed by a '('
            - optionally followed by a number
                - the digits of the number's whole part should be less than 8 excluding the negative sign for negative numbers
                - the digits of the number's decimal part should be less than 12 excluding the decimal point
                - the number can be negative, positive or 0
            - followed by a '\_'
            - optionally followed by a number, required if the first number is skipped
                - the digits of the number's whole part should be less than 8 excluding the negative sign for negative numbers
                - the digits of the number's decimal part should be less than 12 excluding the decimal point
                - the number can be negative, positive or 0
            - followed by a ')'
7. A **boolean** is written as:
    - the word 'bool'
    - followed by '|'
    - optionally followed by the words 'true' or 'false' in any case
8. An **array** is written as:
    - optionally starts with '['
        - repeat more than 1 number of times:
            - followed by a type string given in this list enclosed in quotes
            - followed by ','
            - optionally followed by a ' '
        - followed by a ']'
    - optionally starts with 'array'
        - followed by ';'
        - followed by a number
            - the digits of the number should be less than 8
            - the number should be a positive integer
        - followed by a ';'
        - exactly one of the two must be present:
            - optionally followed by a type string enclosed in quotes
            - optionally followed by a '('
                - followed by a type string enclosed in quotes
                - repeat more than 1 number of times:
                    - followed by a ':'
                    - followed by a type string enclosed in quotes
                - followed by a ')'
        - repeat any number of times:
            - followed by a ';'
            - followed by a number
                - the number should be between -2 and array length (second argument)
            - followed by a ':'
            - followed by a type string enclosed in quotes
9. A **map** is written as:
    - optionally starts with '{'
        - repeat more than 1 number of times:
            - followed by a string enclosed in quotes
            - followed by a ':'
            - followed by a type string enclosed in quotes
            - followed by a ','
            - optionally followed by a ' '
        - followed by a '}'
    - optionally starts with 'map'
        - followed by ';'
        - followed by a number
            - the digits of the number should be less than 8
            - the number should be a positive integer
        - followed by a ';'
        - exactly one of the two must be present:
            - optionally followed by a type string enclosed in quotes
            - optionally followed by a '('
                - followed by a type string enclosed in quotes
                - repeat more than 1 number of times:
                    - followed by a ':'
                    - followed by a type string enclosed in quotes
                - followed by a ')'
        - followed by a ';'
        - exactly one of the two must be present:
            - optionally followed by a type string enclosed in quotes
            - optionally followed by a '('
                - followed by a type string enclosed in quotes
                - repeat more than 1 number of times:
                    - followed by a ':'
                    - followed by a type string enclosed in quotes
                - followed by a ')'
        - repeat any number of times:
            - followed by a ';'
            - followed by a string enclosed in quotes
            - followed by a ':'
            - followed by a type string enclosed in quotes
