# ShortLink

a short link go module.

## usage

### init

```go

import (
	"sync"

	"github.com/uyouii/shortlink"
	shortlink_mysql_db "github.com/uyouii/shortlink/dao/mysql_db"
)

var ShortLinkGeneragor *shortlink.ShortLinkGenerator
var once sync.Once

func Init() {
	once.Do(func() {
		ShortLinkGeneragor = shortlink.GetShortLinkGenerator().WithMysql(&shortlink_mysql_db.MysqlDbConfig{
			User:           "root",
			DatabaseName:   "short_link_db",
			Password:       "asdfgh",
			PartitionCount: 16,
		})
	})
}
```

### gen short link
```go
	shortLink, err = global.ShortLinkGeneragor.GenShortLink(ctx, req.Link)
```

### redirect with gin
```go
func RedirectShortLink(c *gin.Context) {
	ctx := context.Background()
	infof, errorf := common.GetLogFuns(ctx)

	shortLinkPath := c.Param("path")
	if shortLinkPath == "" {
		errorf("invalid path")
		SetError(c, common.GetError(common.ERROR_INVALID_PARAMETERS))
		return
	}
	infof("short link path: %v", shortLinkPath)

	shortLinkPath = strings.TrimPrefix(shortLinkPath, "/")

	rawLink, err := getRawLink(ctx, shortLinkPath)
	if err != nil {
		SetError(c, err)
		return
	}

	infof("rediret to %v", rawLink)

	c.Redirect(http.StatusFound, rawLink)
}
```

## Features

- [x] mysql db
    - [x] support table partition
    - [ ] support temperory short link
    - [ ] reuse expired shortlink path
- [ ] cache
    - [ ] redis cache
    - [ ] local cache

