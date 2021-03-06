//***************************************************************************
// Copyright 2017 smtechnocrat. All rights reserved.
//  
// Package token contains all the Identifers supported by the language.
//
//***************************************************************************
package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers and literals
	IDENT = "IDENT"
	INT   = "IDENT"

	//Operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimitters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "{"
	RPAREN = "}"
	LBRACE = "("
	RBRACE = ")"

	//keywords
	FUNCTION = "func"
	LET      = "let"
)
