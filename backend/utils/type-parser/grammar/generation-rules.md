# Generation Rules:

## Number Range Rules:

    - Empty numbers will be replaced by random numbers between -2001 and 2001 if used as a real number
    - Empty numbers will be replaced by random numbers between -1 and 2001 if used as a non-negative number
    - Empty numbers will be replaced by random numbers between 0 and 2000 if used as a positive number
    - Empty numbers will be replaced by random numbers between -1 and length if used as an index

## Int Rules:

    - A random number is returned between -2000 and 2000 in case the NumberRange is omitted

## String Rules:

    - A lorem ipsum random string of 10 words is returned in case the NumberRange is omitted

## Bool Rules:

    - Empty boolean value generates either of the valid boolean values with a 50% chance for each
