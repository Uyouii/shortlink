package mysql_db

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uyouii/shortlink/common"
	"github.com/uyouii/shortlink/dao/db_base"
	"xorm.io/xorm"
)

type MysqlDbConfig struct {
	User           string
	Password       string
	Addr           string
	Port           int64
	DatabaseName   string
	PartitionCount int
}

type ShortLinkDao struct {
	Config MysqlDbConfig
	Engine *xorm.Engine
}

func GetNewShortLinkDb(config *MysqlDbConfig) db_base.ShortLinkDbInterface {
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
		config.User, config.Password, config.Addr, config.Port, config.DatabaseName)
	engine, err := xorm.NewEngine("mysql", datasource)
	if err != nil {
		errorf("init mysql failed, err: %v", err)
		return err
	}

	d.Config = *config
	d.Engine = engine

	infof("init mysql db engine success, source: %v", datasource)

	return nil
}

func (d *ShortLinkDao) getTableName(talbeNum int) string {
	if d.Config.PartitionCount == 0 {
		return "short_link_tab"
	}
	return fmt.Sprintf("short_link_tab_%08d", talbeNum)
}

func (d *ShortLinkDao) insert(ctx context.Context, shortLink *ShortLinkTab) error {
	infof, errorf := common.GetLogFuns(ctx)

	tableNum := GetTableNumFromRawLinkKey(d.Config.PartitionCount, shortLink.RawLinkKey)
	tableName := d.getTableName(tableNum)
	session := d.Engine.Context(ctx).Table(tableName)

	_, err := session.Insert(shortLink)

	if err != nil {
		errorf("insert failed, err: %v, table: %v, data: %+v", err, tableName, shortLink)
		return err
	}

	infof("insert into %v success, id: %v, raw link key: %v", tableName, shortLink.Id, shortLink.RawLinkKey)

	return nil
}

func (d *ShortLinkDao) update(ctx context.Context, shortLink *ShortLinkTab) error {
	infof, errorf := common.GetLogFuns(ctx)

	tableNum := GetTableNumFromRawLinkKey(d.Config.PartitionCount, shortLink.RawLinkKey)
	tableName := d.getTableName(tableNum)
	session := d.Engine.Context(ctx).Table(tableName)

	_, err := session.Update(shortLink)
	if err != nil {
		errorf("update failed, err: %v, table: %v, data: %+v", err, tableName, shortLink)
		return err
	}

	infof("update %v success, id: %v, raw link key: %v", tableName, shortLink.Id, shortLink.RawLinkKey)

	return nil
}

func (d *ShortLinkDao) get(ctx context.Context, session *xorm.Session) (*ShortLinkTab, error) {
	_, errorf := common.GetLogFuns(ctx)

	res := &ShortLinkTab{}
	exists, err := session.Get(res)
	if err != nil {
		errorf("get by short link path failed, err: %v", err)
		return nil, err
	}

	if !exists {
		errorf("get empty data")
		return nil, common.GetError(common.ERROR_EMPTY)
	}

	return res, nil
}

func (d *ShortLinkDao) delete(ctx context.Context, shortLink *ShortLinkTab) error {
	infof, errorf := common.GetLogFuns(ctx)

	tableNum := GetTableNumFromShortLinkPath(d.Config.PartitionCount, shortLink.ShortLinkPath)
	tableName := d.getTableName(tableNum)
	session := d.Engine.Context(ctx).Table(tableName)

	_, err := session.Delete(shortLink)
	if err != nil {
		errorf("delete failed, err: %v, table: %v, data: %+v", err, tableName, shortLink)
		return err
	}

	infof("delete %v success, shortlinkpath: %v", tableName, shortLink.ShortLinkPath)

	return nil
}

func (d *ShortLinkDao) convert(dbData *ShortLinkTab) *db_base.ShortLink {
	res := &db_base.ShortLink{
		Type:          common.ShortLinkType(dbData.ShortLinkType),
		ShortLinkPath: dbData.ShortLinkPath,
		RawLink:       dbData.RawLink,
		ExpireTime:    MsTimeStampToTime(dbData.ExpireTimestamp),
		CreateTime:    MsTimeStampToTime(dbData.CreateTimestamp),
		UpdateTime:    MsTimeStampToTime(dbData.UpdateTimestamp),
	}
	return res
}

func (d *ShortLinkDao) GenShortLink(ctx context.Context, rawLink string) (*db_base.ShortLink, error) {
	return d.GenShortLinkWithExpire(ctx, rawLink, time.Time{})
}

func (d *ShortLinkDao) GenShortLinkWithExpire(ctx context.Context, rawLink string, expireAt time.Time) (*db_base.ShortLink, error) {
	errorf := common.GetErrorf(ctx)

	shortLink := &ShortLinkTab{
		ShortLinkPath:   "",
		ShortLinkType:   int(common.PermanentShortLink),
		RawLinkKey:      GetUrlShortKey(rawLink),
		RawLink:         rawLink,
		CreateTimestamp: time.Now().UnixMilli(),
		UpdateTimestamp: time.Now().UnixMilli(),
	}

	if !expireAt.IsZero() {
		shortLink.ShortLinkType = int(common.PermanentShortLink)
		shortLink.ExpireTimestamp = expireAt.UnixMilli()
	}

	err := d.insert(ctx, shortLink)
	if err != nil {
		errorf("gen shortlink failed, err: %v, rawlink: %v", err, rawLink)
		return nil, err
	}

	tableNum := GetTableNumFromRawLinkKey(d.Config.PartitionCount, shortLink.RawLinkKey)
	linkPathPrefix := GetShortLinkPrefix(d.Config.PartitionCount, tableNum)
	idHexStr := Int64ToHexStr(uint64(shortLink.Id))
	shortLink.ShortLinkPath = linkPathPrefix + idHexStr

	err = d.update(ctx, shortLink)
	if err != nil {
		errorf("gen shortlink failed, err: %v, rawlink: %v", err, rawLink)
		return nil, err
	}

	return d.convert(shortLink), nil
}

func (d *ShortLinkDao) getByShortLinkPath(ctx context.Context, shortLinkPath string) (*ShortLinkTab, error) {
	infof, errorf := common.GetLogFuns(ctx)

	if shortLinkPath == "" {
		errorf("invalid shortlinkpath: %v", shortLinkPath)
		return nil, common.GetErrorWithMsg(common.INVALID_PARAMS, "empty short link path")
	}

	tableNum := GetTableNumFromShortLinkPath(d.Config.PartitionCount, shortLinkPath)
	tableName := d.getTableName(tableNum)

	session := d.Engine.Context(ctx).Table(tableName).Where("short_link_path = ?", shortLinkPath)

	res, err := d.get(ctx, session)
	if err != nil {
		errorf("get by shortlinkpath failed, err: %v, table: %v, shortlinkpath: %v", err, tableName, shortLinkPath)
		return nil, err
	}

	infof("get by shortlink path success, table: %v, shortlinkpath: %v", tableName, shortLinkPath)
	return res, nil
}

func (d *ShortLinkDao) GetByShortLinkPath(ctx context.Context, shortLinkPath string) (*db_base.ShortLink, error) {
	shortLink, err := d.getByShortLinkPath(ctx, shortLinkPath)
	if err != nil {
		return nil, err
	}
	return d.convert(shortLink), nil
}

func (d *ShortLinkDao) getByRawLink(ctx context.Context, rawLink string) (*ShortLinkTab, error) {
	infof, errorf := common.GetLogFuns(ctx)

	if rawLink == "" {
		errorf("invalid rawlink: %v", rawLink)
		return nil, common.GetErrorWithMsg(common.INVALID_PARAMS, "empty raw link")
	}

	rawLinkKey := GetUrlShortKey(rawLink)
	tableNum := GetTableNumFromRawLinkKey(d.Config.PartitionCount, rawLinkKey)

	tableName := d.getTableName(tableNum)
	session := d.Engine.Context(ctx).Table(tableName).Where("raw_link_key = ?", rawLinkKey)

	res, err := d.get(ctx, session)
	if err != nil {
		errorf("get by raw link failed, err: %v, table: %v, raw link: %v", err, tableName, rawLink)
		return nil, err
	}

	infof("get by shortlink path success, table: %v, rawlink: %v", tableName, rawLink)
	return res, nil
}

func (d *ShortLinkDao) GetByRawLink(ctx context.Context, rawLink string) (*db_base.ShortLink, error) {
	_, errorf := common.GetLogFuns(ctx)
	res, err := d.getByRawLink(ctx, rawLink)
	if err != nil {
		errorf("GetByRawLink failed, err: %v", err)
		return nil, err
	}
	return d.convert(res), nil
}

func (d *ShortLinkDao) SetShortLinkExpire(ctx context.Context, shortLinkPath string, expireAt time.Time) error {
	infof, errorf := common.GetLogFuns(ctx)

	shortLink, err := d.getByShortLinkPath(ctx, shortLinkPath)
	if err != nil {
		errorf("getByShortLinkPath failed, err: %v, shortLinkPath: %v", err, shortLinkPath)
		return err
	}

	if expireAt.IsZero() {
		shortLink.ShortLinkType = int(common.PermanentShortLink)
		shortLink.ExpireTimestamp = 0
	} else {
		shortLink.ShortLinkType = int(common.TemporaryShortLink)
		shortLink.ExpireTimestamp = expireAt.UnixMilli()
	}
	shortLink.UpdateTimestamp = time.Now().UnixMilli()

	err = d.update(ctx, shortLink)
	if err != nil {
		errorf("update shortLink failed, err: %v, shortLink: %+v", err, shortLink)
		return err
	}

	infof("SetShortLinkExpire success, path: %v, expireat: %+v", shortLinkPath, expireAt)

	return nil
}

// TODO: add delete table, and reuse the deleted short path
func (d *ShortLinkDao) DeleteShortLink(ctx context.Context, shortLinkPath string) error {
	infof, errorf := common.GetLogFuns(ctx)

	err := d.delete(ctx, &ShortLinkTab{ShortLinkPath: shortLinkPath})
	if err != nil {
		errorf("delete shortlink path failed, err: %v", err)
		return err
	}

	infof("delete by shortlink path success, shortLinkPath: %v", shortLinkPath)

	return nil
}
