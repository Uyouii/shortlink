package shortlink

import (
	"context"
	"time"

	"github.com/uyouii/shortlink/dao"
	"github.com/uyouii/shortlink/dao/db_base"
	"github.com/uyouii/shortlink/dao/mysql_db"
)

// TODO: add redis cache and local cache
type ShortLinkGenerator struct {
	db db_base.ShortLinkDbInterface
}

func GetShortLinkGenerator() *ShortLinkGenerator {
	return &ShortLinkGenerator{}
}

func (g *ShortLinkGenerator) WithMysql(config *mysql_db.MysqlDbConfig) *ShortLinkGenerator {
	g.db = dao.GetMysqlShortLinkDb(config)
	return g
}

func (g *ShortLinkGenerator) GenShortLink(ctx context.Context, rawLink string) (*db_base.ShortLink, error) {
	return g.db.GenShortLink(ctx, rawLink)
}

func (g *ShortLinkGenerator) GenShortLinkWithExpire(ctx context.Context, rawLink string, expireAt time.Time) (*db_base.ShortLink, error) {
	return g.db.GenShortLinkWithExpire(ctx, rawLink, expireAt)
}

func (g *ShortLinkGenerator) GetByShortLinkPath(ctx context.Context, shortLinkPath string) (*db_base.ShortLink, error) {
	return g.db.GetByShortLinkPath(ctx, shortLinkPath)
}

func (g *ShortLinkGenerator) GetByRawLink(ctx context.Context, rawLink string) (*db_base.ShortLink, error) {
	return g.db.GetByRawLink(ctx, rawLink)
}

func (g *ShortLinkGenerator) SetShortLinkExpire(ctx context.Context, shorLinkPath string, expireAt time.Time) error {
	return g.db.SetShortLinkExpire(ctx, shorLinkPath, expireAt)
}

func (g *ShortLinkGenerator) DeleteShortLink(ctx context.Context, shortLinkPath string) error {
	return g.db.DeleteShortLink(ctx, shortLinkPath)
}
