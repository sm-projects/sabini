//***************************************************************************
// Copyright 2017 smtechnocrat. All rights reserved.
//  
// Package lexer implements a scanner for the Sabini source code
// It takes a string as source which can then be tokenized
// through repeated calls to the NextToken method.
//***************************************************************************
package lexer

import (
	"fmt"
	"github.com/smazumder05/sabini/token"
)

type Lexer struct {
	input        string
	position     int  //points to the current char in input string
	readPosition int  //current read position after current char
	ch           byte //current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = '0'
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LBRACE, l.ch)
	case ')':
		tok = newToken(token.RBRACE, l.ch)
	case '{':
		tok = newToken(token.LPAREN, l.ch)
	case '}':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '0':
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		tok = newToken(token.ILLEGAL, l.ch)
		fmt.Printf("My Current token: %s", string(l.ch))
	}
	l.readChar()
	return tok
}

func newToken(ttype token.TokenType, ch byte) token.Token {
	return token.Token{Type: ttype, Literal: string(ch)}
}
