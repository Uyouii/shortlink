package common

type DbType int

const (
	MysqlDb DbType = iota + 1
)

type ShortLinkType int

const (
	TemporaryShortLink ShortLinkType = iota + 1
	PermanentShortLink
)
