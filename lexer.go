package json

import (
	"errors"
	"io"
	"text/scanner"
)

type Token struct {
	token   int
	literal string
}

type Lexer struct {
	scanner.Scanner
	result interface{}
	err    string
}

func NewLexer(reader io.Reader) *Lexer {
	var s scanner.Scanner
	s.Init(reader)
	return &Lexer{s, nil, ""}
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	if token == scanner.Int {
		token = INT
	}
	if token == scanner.String {
		token = STRING
	}
	if token == scanner.Float {
		token = FLOAT
	}
	switch l.TokenText() {
	case "null":
		token = NULL
	case "true":
		token = TRUE
	case "false":
		token = FALSE
	}
	lval.token = Token{token: token, literal: l.TokenText()}
	return token
}

func (l *Lexer) Error(e string) {
	l.err = e
}

func ParseObject(reader io.Reader) (*Object, error) {
	v, err := Parse(reader)
	if err != nil {
		return nil, err
	}
	o, ok := v.(*Object)
	if !ok {
		return nil, errors.New("not an object")
	}
	return o, nil
}

func ParseArray(reader io.Reader) (Array, error) {
	v, err := Parse(reader)
	if err != nil {
		return nil, err
	}
	o, ok := v.(Array)
	if !ok {
		return nil, errors.New("not an array")
	}
	return o, nil
}

func Parse(reader io.Reader) (interface{}, error) {
	l := new(Lexer)
	l.Init(reader)
	yyParse(l)
	if len(l.err) > 0 {
		return nil, errors.New(l.err)
	}
	return l.result, nil
}
