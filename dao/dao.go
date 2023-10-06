package dao

import (
	"uyouii.cool/shortlink/dao/db_base"
	"uyouii.cool/shortlink/dao/mysql_db"
)

func GetMysqlShortLinkDb(config *mysql_db.MysqlDbConfig) db_base.ShortLinkDbInterface {
	return mysql_db.GetNewShortLinkDb(config)
}
