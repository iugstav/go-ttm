package main

type token int

const (
	INT token = iota

	ADD
	SUB
	MUL
	DIV
)

type Token struct {
	_type token
	Value string
}
