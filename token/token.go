package token

// A TokenType is a string
type TokenType string

// A token is defined by its type
// and literal value. For example type
// can be INTEGER then literal could be "3"
type Token struct {
    Type TokenType
    Literal string
}

// All possible token types
const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    IDENT = "IDENT"
    INT = "INT"

    ASSIGN = "="
    PLUS = "+"

    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = "("
    LBRACE = "{"
    RBRACE = "}"

    FUNCTION = "FUNCTION"
    LET = "LET"
)
