package types

type MemberType int8

const (
	MemberTypeNormal MemberType = iota + 1
	MemberTypeGuest
)

type MemberStatus int8

const (
	MemberStatusOnline MemberStatus = iota + 1
	MemberStatusOffline
)
