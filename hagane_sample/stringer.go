package main

//go:generate stringer -type=MyStatus
type MyStatus int

const (
	A MyStatus = iota
	B
	C
)
