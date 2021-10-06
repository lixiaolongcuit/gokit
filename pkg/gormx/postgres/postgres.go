package postgres

import (
	"time"

	"github.com/lixiaolongcuit/gokit/pkg/gormx"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Postgres postgres.Config
	Pool     *gormx.PoolConfig
}

func NewPostgresDB(c *Config, log *logrus.Entry) (*gorm.DB, error) {
	//TODO 使用logrus打印日志
	db, err := gorm.Open(postgres.New(c.Postgres), &gorm.Config{})
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
