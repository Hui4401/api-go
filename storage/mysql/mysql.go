package mysql

import (
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "api-go/util/logs"
)

var client *gorm.DB

func InitMySQL(url string) {
    db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
    if err != nil {
        logs.PanicKvs("connect to mysql error", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        logs.PanicKvs("get client error", err)
    }
    // 连接池中空闲连接的最大数量
    sqlDB.SetMaxIdleConns(300)
    // 打开数据库连接的最大数量
    sqlDB.SetMaxOpenConns(500)
    // 连接可复用的最大时间
    sqlDB.SetConnMaxLifetime(time.Second * 30)

    client = db
}

func GetClient() *gorm.DB {
    return client
}
