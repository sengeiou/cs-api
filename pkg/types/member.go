package types

type MemberType int8

const (
	MemberTypeNormal MemberType = iota + 1
	MemberTypeGuest
)
