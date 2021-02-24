package plog

func Debug(args ...interface{}) {
	loggerSugar.Debug(args...)
}

func Info(args ...interface{}) {
	loggerSugar.Info(args...)
}

func Warn(args ...interface{}) {
	loggerSugar.Warn(args...)
}

func Error(args ...interface{}) {
	loggerSugar.Error(args...)
}

func Fatal(args ...interface{}) {
	loggerSugar.Fatal(args...)
}

func Panic(args ...interface{}) {
	loggerSugar.Panic(args...)
}

func Debugf(template string, args ...interface{}) {
	loggerSugar.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	loggerSugar.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	loggerSugar.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	loggerSugar.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	loggerSugar.Fatalf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	loggerSugar.Panicf(template, args...)
}
