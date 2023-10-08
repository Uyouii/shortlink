package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/uyouii/shortlink/dao/mysql_db"
)

var testDb = mysql_db.GetNewShortLinkDb(&mysql_db.MysqlDbConfig{
	User:           "root",
	Password:       "asdfgh",
	DatabaseName:   "short_link_db",
	PartitionCount: 16,
})

func TestGetPrometheusShowUrl(t *testing.T) {

	ctx := context.Background()

	fmt.Println(testDb.GenShortLink(ctx, "https://uyouii.cool/"))
}

func TestGetShortLinkByPath(t *testing.T) {

	ctx := context.Background()

	shortLink, _ := testDb.GetByShortLinkPath(ctx, "41")
	fmt.Printf("shotlinkdata: %+v\n", shortLink)
}

func TestGetShortLinkByRawLink(t *testing.T) {

	ctx := context.Background()

	shortLink, err := testDb.GetByRawLink(ctx, "https://uyouii.cool/")
	fmt.Printf("shotlinkdata: %+v\n", shortLink)
	if err != nil {
		fmt.Printf("get by raw link failed, err: %v", err)
		return
	}

	shortLink, _ = testDb.GetByShortLinkPath(ctx, shortLink.ShortLinkPath)
	fmt.Printf("shotlinkdata: %+v\n", shortLink)
}

func TestDeleteByShortLinkPath(t *testing.T) {
	ctx := context.Background()

	err := testDb.DeleteShortLink(ctx, "61")
	fmt.Println(err)
}
