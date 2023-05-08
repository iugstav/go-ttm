package main

type token int

const (
	INT token = iota

	ADD
	SUB
	MUL
	DIV

	ILLEGAL
)

type Token struct {
	_type token
	Value string
}
