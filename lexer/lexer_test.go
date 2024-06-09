package lexer

import (
    "testing"
    "interpreter/token"
)

func testToken(t *testing.T, expected token.Token, actual token.Token) {
    if expected.Type != actual.Type {
        t.Fatalf("token type wrong, expected %q, got %q", expected.Type, actual.Type)
    }

    if expected.Literal != actual.Literal {
        t.Fatalf("literal wrong, expected %q, got %q", expected.Literal, actual.Literal)
    }
}

func TestNextToken(t *testing.T) {
    input := `=+(){},;`

    tests:= []token.Token {
        {token.ASSIGN , "="},
        {token.PLUS, "+"},
        {token.LPAREN, "("},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.RBRACE, "}"},
        {token.COMMA, ","},
        {token.SEMICOLON, ";"},
        {token.EOF, ""},
    }

    l := New(input)

    for _, tt := range tests {
        tok := l.NextToken();
        testToken(t, tok, tt)
    }
}
