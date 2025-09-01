package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

var db *gorm.DB
var once sync.Once

// initMysql 初始化 MySQL 数据库连接
func initMysql() error {
	dsn := "root:xianxiao@tcp(192.168.5.100:3306)/ad_check?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "xiaojing:asdf*UH3oj@tcp(mysqla002e904e7a3.rds.ivolces.com)/ad_check?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 设置为 true 表示使用单数形式的表名
		},
	})
	if err != nil {
		return err
	}

	// 连接池配置
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetConnMaxLifetime(time.Second * 10)
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(50)

	return nil
}

// GetDb 获取数据库连接
func GetDb() *gorm.DB {
	once.Do(func() {
		_ = initMysql()
	})

	return db
}
