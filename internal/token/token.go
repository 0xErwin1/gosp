package token

import "fmt"

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	WHITESPACE
	COMMENT // # comment

	//
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	LEFT_BRACKET
	RIGHT_BRACKET
	SEMICOLON
	COLON
	DOT
	COMMA
	COMMA_AT
	BANG
	BACKTRICK
	SAHRP_QUOTE
	QUOTE

	// Operators
	PLUS
	MINUS
	STAR
	SLASH
	EQUAL

	EQUAL_EQUAL
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL
	NOT_EQUAL
	DIFF
	AND
	OR

	// Literals
	NUMBER  // 123
	STRING  // "hello"
	BOOLEAN // #t | #f

	// Keywords
	DEFUN
	DEFVAR
	IF
	THEN
	ELSE
	LAMBDA
	LET
	DO
	NIL

	INT_TYPE
	FLOAT_TYPE
	CHAR_TYPE
	STRING_TYPE
	BOOLEAN_TYPE
	NIL_TYPE
	LIST_TYPE
	FUNCTION_TYPE

	IDENTIFIER // hello
)

var TokenTypeNames = map[TokenType]string{
	BOOLEAN:       "BOOLEAN",
	BOOLEAN_TYPE:  "BOOLEAN_TYPE",
	CHAR_TYPE:     "CHAR_TYPE",
	COMMA:         "COMMA",
	COLON:         "COLON",
	COMMENT:       "COMMENT",
	DEFUN:         "DEFUN",
	DEFVAR:        "DEFVAR",
	SLASH:         "SLASH",
	DO:            "DO",
	DOT:           "DOT",
	ELSE:          "ELSE",
	EOF:           "EOF",
	EQUAL:         "EQUAL",
	FLOAT_TYPE:    "FLOAT_TYPE",
	FUNCTION_TYPE: "FUNCTION_TYPE",
	IDENTIFIER:    "IDENTIFIER",
	IF:            "IF",
	ILLEGAL:       "ILLEGAL",
	INT_TYPE:      "INT_TYPE",
	LAMBDA:        "LAMBDA",
	LEFT_BRACE:    "LEFT_BRACE",
	LEFT_BRACKET:  "LEFT_BRACKET",
	LEFT_PAREN:    "LEFT_PAREN",
	LET:           "LET",
	LIST_TYPE:     "LIST_TYPE",
	MINUS:         "MINUS",
	STAR:          "STAR",
	NIL:           "NIL",
	NIL_TYPE:      "NIL_TYPE",
	NUMBER:        "NUMBER",
	PLUS:          "PLUS",
	RIGHT_BRACE:   "RIGHT_BRACE",
	RIGHT_BRACKET: "RIGHT_BRACKET",
	RIGHT_PAREN:   "RIGHT_PAREN",
	SEMICOLON:     "SEMICOLON",
	STRING:        "STRING",
	STRING_TYPE:   "STRING_TYPE",
	THEN:          "THEN",
	WHITESPACE:    "WHITESPACE",
	BANG:          "BANG",
	BACKTRICK:     "BACKTRICK",
	SAHRP_QUOTE:   "SAHRP_QUOTE",
	QUOTE:         "QUOTE",
	COMMA_AT:      "COMMA_AT",
}

var Keywords = map[string]TokenType{
	"defun":  DEFUN,
	"defvar": DEFVAR,
	"if":     IF,
	"then":   THEN,
	"else":   ELSE,
	"lambda": LAMBDA,
	"let":    LET,
	"do":     DO,
	"nil":    NIL,
	"int":    INT_TYPE,
	"float":  FLOAT_TYPE,
	"char":   CHAR_TYPE,
	"string": STRING_TYPE,
	"bool":   BOOLEAN_TYPE,
}

func (t TokenType) String() string {
	return TokenTypeNames[t]
}

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
	Start   int
	End     int
	File    string
}

func NewToken(tokenType TokenType, literal string, line int, column int, start int, end int, file string) *Token {
	return &Token{
		Type:    tokenType,
		Literal: literal,
		Line:    line,
		Column:  column,
		Start:   start,
		End:     end,
		File:    file,
	}
}

func (t Token) String() string {
	return fmt.Sprintf(
		"Token: %s, Literal: %s, Line: %d, Column: %d, Start: %d, End: %d, File: %s",
		TokenTypeNames[t.Type],
		t.Literal,
		t.Line,
		t.Column,
		t.Start,
		t.End,
		t.File,
	)
}
