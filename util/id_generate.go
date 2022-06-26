package util

import "github.com/sony/sonyflake"

var sf *sonyflake.Sonyflake = nil

// GenerateID 使用 sonyflake 算法生成全局唯一ID
func GenerateID() (uint64, error) {
    if sf == nil {
        // sonyflake设置
        st := sonyflake.Settings{
            // 每台机器指定唯一的机器号
            MachineID: func() (uint16, error) {
                return 0, nil
            },
        }
        sf = sonyflake.NewSonyflake(st)
    }
    return sf.NextID()
}
