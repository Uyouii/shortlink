package mysql_db

type ShortLinkTab struct {
	Id              int64  `xorm:"pk autoincr BIGINT(20)"`
	ShortLinkType   int    `xorm:"INT(11)"`
	ShortLinkPath   string `xorm:"not null VARCHAR(64)"`
	RawLinkKey      string `xorm:"not null unique VARCHAR(64)"` // md5 + fnv128
	RawLink         string `xorm:"not null VARCHAR(2048)"`
	ExpireTimestamp int64  `xorm:"index BIGINT(20)"`
	CreateTimestamp int64  `xorm:"BIGINT(20)"`
	UpdateTimestamp int64  `xorm:"BIGINT(20)"`
}
