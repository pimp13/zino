package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special tokens
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	NEWLINE = "NEWLINE"

	// Identifiers + literals
	IDENT    = "IDENT"
	NUMBER   = "NUMBER"
	STRING   = "STRING"
	BOOL     = "BOOL"
	NULL     = "NULL"
	ARRAY    = "ARRAY"
	OBJECT   = "OBJECT"
	FUNCTION = "FUNCTION"
	CLASS    = "CLASS"
	INSTANCE = "INSTANCE"
	METHOD   = "METHOD"
	PROPERTY = "PROPERTY"
	INDEX    = "INDEX"

	// Operators
	ASSIGN                = "="
	PLUS                  = "+"
	MINUS                 = "-"
	MULTIPLY              = "*"
	DIVIDE                = "/"
	MODULO                = "%"
	EQUAL                 = "=="
	NOT_EQUAL             = "!="
	GREATER_THAN          = ">"
	LESS_THAN             = "<"
	GREATER_THAN_OR_EQUAL = ">="
	LESS_THAN_OR_EQUAL    = "<="
	AND                   = "&&"
	OR                    = "||"
	NOT                   = "!"
	IF                    = "IF"
	ELSE                  = "ELSE"
	WHILE                 = "WHILE"
	FOR                   = "FOR"
	RETURN                = "RETURN"

	// Delimiters
	COMMA         = ","
	SEMICOLON     = ";"
	COLON         = ":"
	LEFT_PAREN    = "("
	RIGHT_PAREN   = ")"
	LEFT_BRACE    = "{"
	RIGHT_BRACE   = "}"
	LEFT_BRACKET  = "["
	RIGHT_BRACKET = "]"
)
