package postgres

import (
	"testing"
	"time"

	"github.com/lixiaolongcuit/gokit/pkg/gormx"
	"gorm.io/driver/postgres"
)

/*
CREATE TABLE "public"."user" (
  "id" serial8 NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "username" varchar(255) COLLATE "pg_catalog"."default",
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "create_time" timestamp,
  "update_time" timestamp,
  "delete_time" timestamp,
  CONSTRAINT "user_pkey" PRIMARY KEY ("id")
);
*/

type User struct {
	Id         int64 `gorm:"primaryKey"`
	Name       string
	Username   string
	Password   string
	CreateTime *time.Time
	UpdateTime *time.Time
	DeleteTime *time.Time
}

func (User) TableName() string {
	return "user"
}

func TestConnect(t *testing.T) {
	cfg := &Config{
		Postgres: postgres.Config{
			DSN: "host=127.0.0.1 port=5432 user=postgres password=123456 dbname=test sslmode=disable TimeZone=Asia/Shanghai",
		},
		Pool: &gormx.PoolConfig{
			MaxIdleConns:    10,
			MaxOpenConns:    100,
			ConnMaxLifetime: 1,
		},
	}
	db, err := NewPostgresDB(cfg, nil)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	user := &User{
		Name:       "test name",
		Username:   "test",
		Password:   "123456",
		CreateTime: &now,
		UpdateTime: &now,
	}
	if err := db.Create(user).Error; err != nil {
		t.Fatalf("create user error:%v", err)
	}
	var selectUser User
	if err := db.First(&selectUser).Error; err != nil {
		t.Fatalf("select user error:%v", err)
	}
	if user.Id != selectUser.Id {
		t.Fatal("select user id not eq insert user id")
	}
	t.Logf("%+v", selectUser)
}
