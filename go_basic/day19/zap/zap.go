package main

import (
	"net/http"

	"github.com/natefinch/lumberjack"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	Init()
	//在函数即将执行结束时，把日志写到磁盘
	defer logger.Sync()
	for i := 0; i < 10000; i++ {
		logger.Info("test log to ...")
	}
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
	//SugaredLogger和Logger
	_ = logger.Sugar()
	simpleHttpGet("http://www.baidu.com")
}

//func Init() (err error) {
//	logger, err = zap.NewProduction()
//	zap.NewDevelopment()
//	if err != nil {
//		return err
//	}
//	return
//}
//定制化logger
func Init() {
	//Encoder 编码器（如何写入日志）
	//将jsonEncoder更改为普通的Encoder，为此我们将newJSONEncoder改为newConsoleEncoder
	//encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	encoder := getEncoder()
	//指定日志写到哪，使用zapcore.AddSync()函数并且将打开的文件句柄传进去
	writerSyncer := getLogWriter()
	//logLevel 哪种级别的日志将被写入
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	//定制logger
	logger = zap.New(core, zap.AddCaller())

}

//func getLogWriter() zapcore.WriteSyncer {
//	file, _ := os.OpenFile("./text.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0744)
//	return zapcore.AddSync(file)
//}
func getEncoder() zapcore.Encoder {
	//日志编码配置
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encodeConfig)

}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("error fetching url...", zap.String("url", url), zap.Error(err))
	} else {
		logger.Info("Success ...", zap.String("url", url), zap.Error(err))
		resp.Body.Close()
	}
}
func getLogWriter() zapcore.WriteSyncer {
	lumberjackLog := &lumberjack.Logger{
		Filename:   "./text.log",
		MaxSize:    1,
		MaxAge:     30,
		MaxBackups: 5,
		Compress:   false,
	}
	return zapcore.AddSync(lumberjackLog)
}
