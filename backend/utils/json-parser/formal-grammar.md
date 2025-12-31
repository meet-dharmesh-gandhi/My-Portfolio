## Grammar in Context Free Language form:

Helper Rules:
U -> U|Φ // Actual Unicode characters
NO -> 0|1|2|3|4|5|6|7|8|9 // Numbers Only
PN -> NOΦ|NO // Positive Number
NN -> -PN // Negative Number
IN -> PN|NN // Integers
DN -> IN.PN // Decimal Number
MN -> // custom number, refer to the english grammar for better details of this number
Q -> '\''
C -> ','
SP -> ' '
CL -> ';'
SMCL -> ':'
LP -> '('
RP -> ')'
LBK -> '['
RBK -> ']'
LBC -> '{'
RBC -> '}'
UD -> '\_'

Actual Rules:
S -> G|S|I|F|B|A|M

G -> U|g U
g -> 'general|'

S -> s|s PN
s -> 'string|'

I -> i|i PN|i LP IN UD IN RP|LP UD IN RP|LP IN UD RP
i -> 'int|'

F -> f|f DN|iLP DN UD DN RP|LP UD DN RP|LP DN UD RP
f -> 'float|'

B -> b|b TR|b FL
TR -> 'true'|'True'|'tRue'|'trUe'|'truE'|'TRue'|'TrUe'|'TruE'|'tRUe'|'tRuE'|'trUE'|'TRUe'|'TRuE'|'TrUE'|'tRUE'|'TRUE'
FL -> 'false'|'falsE'|'falSe'|'falSE'|'faLse'|'faLsE'|'faLSe'|'faLSE'|'fAlse'|'fAlsE'|'fAlSe'|'fAlSE'|'fALse'|'fALsE'|'fALSe'|'fALSE'|'False'|'FalsE'|'FalSe'|'FalSE'|'FaLse'|'FaLsE'|'FaLSe'|'FaLSE'|'FAlse'|'FAlsE'|'FAlSe'|'FAlSE'|'FALse'|'FALsE'|'FALSe'|'FALSE'

A -> LBK (A1)^+ RBK|a CL PN CL S (A2)^\*|a CL PN CL A3 (A2)^\*
A1 -> Q S Q C|Q S Q C SP
a -> 'array'
A2 -> CL MN SMCL S
A3 -> LP S (CT)^+ RP
CT -> CL S

M -> LBC (KV)^+ RBC|m CL PN CL Q S Q CL Q S Q (M1)^\*|m CL PN CL M2 CL Q S Q (M1)^\*|m CL PN CL Q S Q CL M2 (M1)^\*|m CL PN CL M2 CL M2 (M1)^\*
KV -> Q U Q SMCL Q S Q C|Q U Q SMCL Q S Q C SP
m -> 'map'
M1 -> CL Q U Q SMCL Q S Q
M2 -> LP Q S Q (M21)^+ RP
M21 -> CL Q S Q
