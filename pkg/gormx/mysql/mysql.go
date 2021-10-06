package mysql

import (
	"time"

	"github.com/lixiaolongcuit/gokit/pkg/gormx"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Mysql mysql.Config
	Pool  *gormx.PoolConfig
}

func NewMysqlDB(c *Config, log *logrus.Entry) (*gorm.DB, error) {
	//TODO 使用logrus打印日志
	db, err := gorm.Open(mysql.New(c.Mysql), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if c.Pool != nil {
		sqlDB.SetMaxIdleConns(c.Pool.MaxIdleConns)
		sqlDB.SetMaxOpenConns(c.Pool.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(c.Pool.ConnMaxLifetime * time.Hour)
	}

	return db, nil
}
