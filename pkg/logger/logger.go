package logger

import (
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/imdario/mergo"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var zlog *zap.Logger

type LogCfg struct {
	Filename string
	Level    int
	MaxDays  int
	LogDir   string
	Maxsize  int64
}

func DefaultLogCfg() *LogCfg {
	return &LogCfg{
		Filename: "go-ops.log",
		LogDir:   "logs",
		Maxsize:  1024,
		MaxDays:  30,
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func init() {

	encoder := getEncoder()
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	zlog = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func InitLog(reqCfg *LogCfg) {

	cfg := DefaultLogCfg()
	if reqCfg != nil {
		err := mergo.Merge(cfg, reqCfg, mergo.WithOverride)
		if err != nil {
			panic(err)
		}
	}

	baseLogPath := path.Join(cfg.LogDir, cfg.Filename)
	writter, err := rotatelogs.New(
		baseLogPath+".%Y-%m-%d",
		rotatelogs.WithLinkName(baseLogPath),
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(cfg.MaxDays)),
		rotatelogs.WithRotationSize(1024*1024*cfg.Maxsize),
	)
	if err != nil {
		panic(err)
	}

	encoder := getEncoder()
	core := zapcore.NewCore(encoder, zapcore.AddSync(writter), zapcore.DebugLevel)
	zlog = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

}

func Sync() {
	zlog.Sync()
}

func Debug(args ...interface{}) {
	zlog.Sugar().Debug(args...)
}

func Info(args ...interface{}) {
	zlog.Sugar().Debug(args...)
}

func Warn(args ...interface{}) {
	zlog.Sugar().Warn(args...)
}

func Error(args ...interface{}) {
	zlog.Sugar().Error(args...)
}

func Debugf(f string, args ...interface{}) {
	zlog.Sugar().Debugf(f, args...)
}

func Infof(f string, args ...interface{}) {
	zlog.Sugar().Infof(f, args...)
}

func Warnf(f string, args ...interface{}) {
	zlog.Sugar().Warnf(f, args...)
}

func Errorf(f string, args ...interface{}) {
	zlog.Sugar().Errorf(f, args...)
}
