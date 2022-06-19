package model

import (
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "api-go/logs"
)

var DB *gorm.DB

func InitMySQL(url string) {
    db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
    if err != nil {
        logs.PanicKvs("connect to mysql error", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        logs.PanicKvs("get DB error", err)
    }
    // 连接池中空闲连接的最大数量
    sqlDB.SetMaxIdleConns(300)
    // 打开数据库连接的最大数量
    sqlDB.SetMaxOpenConns(500)
    // 连接可复用的最大时间
    sqlDB.SetConnMaxLifetime(time.Second * 30)
    // 自动迁移模式
    if err = db.AutoMigrate(&User{}); err != nil {
        logs.PanicKvs("auto migration error", err)
    }

    DB = db
}
