package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 数据库链接单例
var DB *gorm.DB

// 初始化mysql连接
func Database(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("mysql连接异常: " + err.Error())
	}

	// 设置连接池
	sqlDB, _ := db.DB()
	// 连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(300)
	// 打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(500)
	// 连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}

func migration() {
	// 自动迁移模式
	err := DB.AutoMigrate(&User{})
	if err != nil {
		panic("auto migration error: " + err.Error())
	}
}
