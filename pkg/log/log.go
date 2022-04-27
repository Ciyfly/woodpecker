/*
 * @Date: 2022-04-13 16:45:32
 * @LastEditors: recar
 * @LastEditTime: 2022-04-14 18:48:44
 */

package log

import (
	"fmt"
	"os"
	"woodpecker/pkg/conf"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zlog *zap.Logger
var alarm *zap.Logger

type Message struct {
	Type     int
	Msg      string
	Request  string
	Response string
}

type ConnMessage struct {
	SrcIp    string
	DstIp    string
	Service  string
	Protocol string
}

func Init() {
	// log log
	var coreArrLog []zapcore.Core

	//获取编码器
	encoderConfigLog := zap.NewProductionEncoderConfig()                                 //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfigLog.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000") //指定时间格式
	encoderConfigLog.EncodeLevel = zapcore.CapitalColorLevelEncoder                      //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	encoderLog := zapcore.NewConsoleEncoder(encoderConfigLog)

	lowPriorityLog := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev > zap.DebugLevel

	})

	//info文件writeSyncer
	infoFileWriteSyncerLog := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/log.log",                        //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    conf.GlobalConfig.LogConfig.MaxSize,    //文件大小限制,单位MB
		MaxBackups: conf.GlobalConfig.LogConfig.MaxBackups, //最大保留日志文件数量
		MaxAge:     conf.GlobalConfig.LogConfig.MaxAge,     //日志文件保留天数
		Compress:   conf.GlobalConfig.LogConfig.Compress,   //是否压缩处理
	})
	infoFileCoreLog := zapcore.NewCore(encoderLog, zapcore.NewMultiWriteSyncer(infoFileWriteSyncerLog, zapcore.AddSync(os.Stdout)), lowPriorityLog)
	coreArrLog = append(coreArrLog, infoFileCoreLog)
	zlog = zap.New(zapcore.NewTee(coreArrLog...))
	defer zlog.Sync()

}

func Info(info string) {
	zlog.Info(info)
}

func Infof(format string, args ...interface{}) {
	zlog.Info(fmt.Sprintf(format, args...))
}

func Debug(debug string) {
	zlog.Debug(debug)
}

func Debugf(format string, args ...interface{}) {
	zlog.Debug(fmt.Sprintf(format, args...))
}

func Error(error string) {
	zlog.Error(error)
}

func Errorf(format string, args ...interface{}) {
	zlog.Error(fmt.Sprintf(format, args...))
}
