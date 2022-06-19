package logs

import (
    "log"

    "go.uber.org/zap"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap/zapcore"
    "os"
)

var Logger *zap.SugaredLogger

func Init() {
    var zapConfig zap.Config

    if gin.Mode() == gin.DebugMode || gin.Mode() == gin.TestMode {
        zapConfig = zap.NewDevelopmentConfig()
    } else if gin.Mode() == gin.ReleaseMode {
        zapConfig = zap.NewProductionConfig()
    } else {
        log.Panicf("get gin mode error, mode: %s", gin.Mode())
    }

    zapConfig.Encoding = "console"
    zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
    zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")

    // 同时输出到文件
    logPath := os.Getenv("LOG_PATH")
    if logPath != "" {
        if _, err := os.Stat(logPath); os.IsNotExist(err) {
            if err = os.MkdirAll(logPath, 0666); err != nil {
                log.Panicf("creat log path error: %s", err.Error())
            }
        }
        zapConfig.OutputPaths = append(zapConfig.OutputPaths, logPath+"/info.log")
        zapConfig.ErrorOutputPaths = append(zapConfig.ErrorOutputPaths, logPath+"/error.log")
    }

    _logger, err := zapConfig.Build(
        zap.AddCallerSkip(1),
        zap.AddStacktrace(zap.PanicLevel),
    )
    if err != nil {
        log.Panicf("init zap Logger error: %s", err.Error())
    }
    Logger = _logger.Sugar()
}

func DebugKvs(kvs ...interface{}) {
    Logger.Debugw("", kvs...)
}

func InfoKvs(kvs ...interface{}) {
    Logger.Infow("", kvs...)
}

func WarnKvs(kvs ...interface{}) {
    Logger.Warnw("", kvs...)
}

func ErrorKvs(kvs ...interface{}) {
    Logger.Errorw("", kvs...)
}

func PanicKvs(kvs ...interface{}) {
    Logger.Panicw("", kvs...)
}
