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
    input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
    x + y;
}

let result = add(five, ten);
`
    tests:= []token.Token {
        {token.LET, "let"},
        {token.IDENT, "five"},
        {token.ASSIGN, "="},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENT, "ten"},
        {token.ASSIGN, "="},
        {token.INT, "10"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENT, "add"},
        {token.ASSIGN, "="},
        {token.FUNCTION, "fn"},
        {token.LPAREN, "("},
        {token.IDENT, "x"},
        {token.COMMA, ","},
        {token.IDENT, "y"},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.IDENT, "x"},
        {token.PLUS, "+"},
        {token.IDENT, "y"},
        {token.SEMICOLON, ";"},
        {token.RBRACE, "}"},
        {token.LET, "let"},
        {token.IDENT, "result"},
        {token.ASSIGN, "="},
        {token.IDENT, "add"},
        {token.LPAREN, "("},
        {token.IDENT, "five"},
        {token.COMMA, ","},
        {token.IDENT, "ten"},
        {token.RPAREN, ")"},
        {token.SEMICOLON, ";"},
        {token.EOF, ""},
    }

    l := New(input)

    for _, tt := range tests {
        tok := l.NextToken();
        testToken(t, tt, tok)
    }
}
