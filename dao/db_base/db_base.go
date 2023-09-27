package db_base

import "time"

type ShortLinkType int

const (
	Temporary ShortLinkType = iota + 1
	Permanent
)

type ShortLink struct {
	Type        ShortLinkType
	ShortLink   string
	RawLinkData string
	ExpireTime  time.Time
	CreateTime  time.Time
	UpdateTime  time.Time
}

type ShortLinkDbInterface interface {
	GenShortLink(rawLink string) (*ShortLink, error)
	GenShortLinkExpire(rawLink string, expireAt time.Time) (*ShortLink, error)
	GetByShortLink(shortLink string) (*ShortLink, error)
	GetByRawLink(rawLink string) (*ShortLink, error)
	SetShortLinkExpire(shorLink string, expireAt time.Time) error
}
