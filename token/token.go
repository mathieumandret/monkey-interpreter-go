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
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"
    EQ = "=="
    NOT_EQ = "!="

    LT = "<"
    GT = ">"

    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = "("
    LBRACE = "{"
    RBRACE = "}"

    FUNCTION = "FUNCTION"
    LET = "LET"
    TRUE = "TRUE"
    FALSE = "FALSE"
    // TODO: Add ternary operators = let x = cond ? 1 : 0
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"
)

var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}
