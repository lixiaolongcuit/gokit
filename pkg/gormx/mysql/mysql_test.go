package mysql

import (
	"testing"
	"time"

	"github.com/lixiaolongcuit/gokit/pkg/gormx"
	"gorm.io/driver/mysql"
)

type User struct {
	Id         int64 `gorm:"primaryKey"`
	Name       string
	Username   string
	Password   string
	CreateTime *time.Time
	UpdateTime *time.Time
	DeleteTime *time.Time
}

func TestConnect(t *testing.T) {
	cfg := &Config{
		Mysql: mysql.Config{
			DSN: "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
		},
		Pool: &gormx.PoolConfig{
			MaxIdleConns:    10,
			MaxOpenConns:    100,
			ConnMaxLifetime: 1,
		},
	}
	db, err := NewMysqlDB(cfg, nil)
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
