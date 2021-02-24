package plog

import (
	"github.com/gogf/gf/os/gfile"
	"github.com/lestrrat-go/file-rotatelogs"
	"go-dev-web/pkg/timeutils"
	"go-dev-web/pkg/types"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path"
	"time"
)

var (
	logger      *zap.Logger
	loggerSugar *zap.SugaredLogger
)

func Init(filepath string, envType string) (err error) {
	var (
		core        zapcore.Core
		logHookFile io.Writer
	)

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(timeutils.FormatZap(t))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	minLevelFn := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		if envType == types.EnvTypeProduction {
			return lvl >= zapcore.InfoLevel
		} else {
			return lvl >= zapcore.DebugLevel
		}
	})

	if envType == types.EnvTypeLocal {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), minLevelFn),
		)
	} else {
		if !gfile.Exists(filepath) {
			if err = gfile.Mkdir(filepath); err != nil {
				return
			}
		}
		if logHookFile, err = getWriter(filepath, "log"); err != nil {
			return
		}
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), minLevelFn),
			zapcore.NewCore(encoder, zapcore.AddSync(logHookFile), minLevelFn),
		)
	}

	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	loggerSugar = logger.Sugar()
	return
}

func getWriter(filepath, filename string) (hook io.Writer, err error) {
	hook, err = rotatelogs.New(
		path.Join(filepath, filename+".%Y%m%d"),
		rotatelogs.WithMaxAge(time.Hour*24*100),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	return
}

func Logger() *zap.Logger {
	return logger
}

func LoggerSugar() *zap.SugaredLogger {
	return loggerSugar
}
