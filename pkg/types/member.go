package types

type MemberType int8

const (
	MemberTypeNormal MemberType = iota + 1
	MemberTypeGuest
)

type MemberOnlineStatus int8

const (
	MemberOnlineStatusOnline MemberOnlineStatus = iota + 1
	MemberOnlineStatusOffline
)
