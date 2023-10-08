package db_base

import (
	"context"
	"time"

	"github.com/uyouii/shortlink/common"
)

type ShortLink struct {
	Type          common.ShortLinkType
	ShortLinkPath string
	RawLink       string
	ExpireTime    time.Time
	CreateTime    time.Time
	UpdateTime    time.Time
}

type ShortLinkDbInterface interface {
	GenShortLink(ctx context.Context, rawLink string) (*ShortLink, error)
	GenShortLinkWithExpire(ctx context.Context, rawLink string, expireAt time.Time) (*ShortLink, error)
	GetByShortLinkPath(ctx context.Context, shortLinkPath string) (*ShortLink, error)
	GetByRawLink(ctx context.Context, rawLink string) (*ShortLink, error)
	SetShortLinkExpire(ctx context.Context, shorLinkPath string, expireAt time.Time) error
	DeleteShortLink(ctx context.Context, shortLinkPath string) error
}
