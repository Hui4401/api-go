package model

import (
    "api-go/util/logs"
    "api-go/storage/mysql"
)

// AutoMigrate 注册所有model的自动迁移
func AutoMigrate() {
    client := mysql.GetClient()
    if err := client.AutoMigrate(&User{}); err != nil {
        logs.PanicKvs("auto migration error", err)
    }
}
