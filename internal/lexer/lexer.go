package lexer

import (
	"bytes"
	"fmt"
	"io"
	"unicode"

	"github.com/0xErwin1/gosp/internal/token"
)

type Lexer struct {
	Tokens     []*token.Token
	source     *bytes.Buffer
	sourceRune []rune
	reader     io.Reader
	current    int
	line       int
	start      int
	errorList  []error
}

func NewLexer(reader io.Reader, source *bytes.Buffer) *Lexer {
	return &Lexer{
		Tokens:     []*token.Token{},
		source:     source,
		sourceRune: bytes.Runes(source.Bytes()),
		reader:     reader,
		current:    0,
		line:       1,
		start:      0,
		errorList:  []error{},
	}
}

func (lexer *Lexer) Lex() ([]*token.Token, []error) {
	for !lexer.isAtEnd() {
		err := lexer.scanToken()

		if err != nil {
			lexer.errorList = append(lexer.errorList, err)
		}
	}
	lexer.addToken(token.EOF, "")

	return lexer.Tokens, lexer.errorList
}

func (lexer *Lexer) scanToken() error {
	c, err := lexer.next()

	if err != nil {
		return err
	}

	switch {
	case c == '\n':
		lexer.line++
	case lexer.isWhitespace(c):
		break
	case c == '(':
		lexer.addToken(token.LEFT_PAREN, string(c))
	case c == ')':
		lexer.addToken(token.RIGHT_PAREN, string(c))
	case c == '!':
		if lexer.match('=') {
			lexer.addToken(token.NOT_EQUAL, string(c))
		} else {
			lexer.addToken(token.BANG, string(c))
		}
	case c == '+':
		lexer.addToken(token.PLUS, string(c))
	case c == '-':
		lexer.addToken(token.MINUS, string(c))
	case c == '*':
		lexer.addToken(token.STAR, string(c))
	case c == '/':
		if lexer.match('=') {
			lexer.addToken(token.DIFF, string(c))
		} else {
			lexer.addToken(token.SLASH, string(c))
		}
	case c == '=':
		lexer.addToken(token.EQUAL, string(c))
	case unicode.IsLetter(c) || c == '_' || c == ':':
		return lexer.identifier()
	case c == ':':
		lexer.addToken(token.COLON, string(c))
	case c == '.':
		lexer.addToken(token.DOT, string(c))
	case c == ';':
		lexer.addToken(token.SEMICOLON, string(c))
	case c == ',':
		if lexer.match('@') {
			lexer.addToken(token.COMMA_AT, string(c))
		} else {
			lexer.addToken(token.COMMA, string(c))
		}
  case c == '\'':
    lexer.addToken(token.QUOTE, string(c))
	case c == '`':
		lexer.addToken(token.BACKTRICK, string(c))
	case c == '#':
		if lexer.match('\'') {
			lexer.addToken(token.SAHRP_QUOTE, string(c))
		} else if lexer.match('f') || lexer.match('t') {
			lexer.addToken(token.BOOLEAN, string(c))
		} else {
			return fmt.Errorf("unexpected character: %c", c)
		}
	case c == ';':
		lexer.comment()
	case c == '"':
		return lexer.string()
	case unicode.IsDigit(c):
		return lexer.number()
	default:
		return fmt.Errorf("unexpected character: %c", c)
	}

	return nil
}

func (lexer *Lexer) comment() {
	for lexer.peek() != '\n' && !lexer.isAtEnd() {
		_, err := lexer.next()

		if err != nil {
			return
		}
	}
}

func (lexer *Lexer) string() error {
	lexer.start = lexer.current - 1
	var value []rune
	isEscaped := false

	for !lexer.isAtEnd() {
		c := lexer.peek()

		if isEscaped {
			switch c {
			case 'n':
				value = append(value, '\n')
			case 'r':
				value = append(value, '\r')
			case 't':
				value = append(value, '\t')
			case '"':
				value = append(value, '"')
			case '\\':
				value = append(value, '\\')
			default:
				return fmt.Errorf("invalid escape sequence: \\%c", c)
			}
			isEscaped = false
		} else {
			if c == '\\' {
				isEscaped = true
			} else if c == '"' {
				break
			} else {
				value = append(value, c)
			}
		}

		_, err := lexer.next()

		if err != nil {
			return err
		}
	}

	if lexer.isAtEnd() {
		return fmt.Errorf("unterminated string")
	}

	_, err := lexer.next()

	if err != nil {
		return err
	}

	lexer.addToken(token.STRING, string(value))

	return nil
}

func (lexer *Lexer) identifier() error {
	lexer.start = lexer.current - 1

	for lexer.isAlphanumeric(lexer.peek()) {
		_, err := lexer.next()

		if err != nil {
			return err
		}
	}

	text := string(lexer.sourceRune[lexer.start:lexer.current])
	if _type, ok := token.Keywords[text]; ok {
		lexer.addToken(_type, text)
	} else {
		lexer.addToken(token.IDENTIFIER, text)
	}

	return nil
}

func (lexer *Lexer) number() error {
	lexer.start = lexer.current - 1

	for unicode.IsDigit(lexer.peek()) {
		_, err := lexer.next()

		if err != nil {
			return err
		}
	}

	text := string(lexer.sourceRune[lexer.start:lexer.current])
	lexer.addToken(token.NUMBER, text)
	return nil
}

func (lexer *Lexer) match(expected rune) bool {
	if lexer.isAtEnd() {
		return false
	}

	if lexer.peek() != expected {
		return false
	}

	_, err := lexer.next()

	return err == nil
}

func (lexer *Lexer) next() (rune, error) {
	if lexer.isAtEnd() {
		return rune(0), io.EOF
	}

	c := lexer.sourceRune[lexer.current]
	lexer.current++
	return c, nil
}

func (lexer *Lexer) peek() rune {
	if lexer.isAtEnd() {
		return rune(0)
	}

	return lexer.sourceRune[lexer.current]
}

func (lexer *Lexer) isAtEnd() bool {
	return lexer.current >= len(lexer.sourceRune)
}

func (lexer *Lexer) isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func (lexer *Lexer) isWhitespace(c rune) bool {
	return unicode.IsSpace(c)
}

func (lexer *Lexer) isDigit(c rune) bool {
	return unicode.IsDigit(c)
}

func (lexer *Lexer) addToken(tokenType token.TokenType, literal string) {
	lexer.Tokens = append(lexer.Tokens,
		token.NewToken(tokenType, literal, lexer.line, lexer.current-lexer.start, lexer.start, lexer.current, "repl"),
	)
}
