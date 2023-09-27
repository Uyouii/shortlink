package mysql_db

import (
	"context"
	"fmt"
	"time"

	"uyouii.cool/shortlink/common"
	"uyouii.cool/shortlink/dao/db_base"
	"xorm.io/xorm"
)

type MysqlDbConfig struct {
	User         string
	Password     string
	Addr         string
	Port         int64
	DatabaseName string
}

type ShortLinkDao struct {
	Config MysqlDbConfig
	Engine *xorm.Engine
}

func GetNewShortLinkDao(config *MysqlDbConfig) db_base.ShortLinkDbInterface {
	shorLinkDao := &ShortLinkDao{}
	err := shorLinkDao.Init(config)
	if err != nil {
		panic(err)
	}
	return shorLinkDao
}

func (d *ShortLinkDao) Init(config *MysqlDbConfig) error {
	infof, errorf := common.GetLogFuns(context.Background())

	if config.Addr == "" && config.Port == 0 {
		config.Addr = "127.0.0.1"
		config.Port = 3306
	}

	infof("init mysql with config: %+v", config)

	datasource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
		d.Config.User, d.Config.Password, d.Config.Addr, d.Config.Port, d.Config.DatabaseName)
	engine, err := xorm.NewEngine("mysql", datasource)
	if err != nil {
		errorf("init mysql failed, err: %v", err)
		return err
	}
	d.Config = *config
	d.Engine = engine

	infof("init mysql db engine success")

	return nil
}

func (d *ShortLinkDao) GenShortLink(rawLink string) (*db_base.ShortLink, error) {
	return nil, nil
}

func (d *ShortLinkDao) GenShortLinkExpire(rawLink string, expireAt time.Time) (*db_base.ShortLink, error) {
	return nil, nil
}

func (d *ShortLinkDao) GetByShortLink(shortLink string) (*db_base.ShortLink, error) {
	return nil, nil
}

func (d *ShortLinkDao) GetByRawLink(rawLink string) (*db_base.ShortLink, error) {
	return nil, nil
}

func (d *ShortLinkDao) SetShortLinkExpire(shorLink string, expireAt time.Time) error {
	return nil
}
