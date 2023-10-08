package dao

import (
	"github.com/uyouii/shortlink/dao/db_base"
	"github.com/uyouii/shortlink/dao/mysql_db"
)

func GetMysqlShortLinkDb(config *mysql_db.MysqlDbConfig) db_base.ShortLinkDbInterface {
	return mysql_db.GetNewShortLinkDb(config)
}
