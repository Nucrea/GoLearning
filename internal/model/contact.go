package model

type ContactType int

const (
    PHONE ContactType = iota
    EMAIL
	LOGIN
)

type Contact struct {
	Type ContactType
	Value string
}