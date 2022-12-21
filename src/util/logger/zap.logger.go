package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Logger /** Logger Struct **/
type Logger struct {
	Zap   *zap.Logger
	Sugar *zap.SugaredLogger
}

// InitLogger /** Logger Initialisation **/
func InitLogger(prod bool, logName string) *zap.Logger {
	var config zapcore.EncoderConfig
	if prod {
		config = zap.NewProductionEncoderConfig()
	} else {
		config = zap.NewDevelopmentEncoderConfig()
	}
	config.EncodeTime = zapcore.ISO8601TimeEncoder //ISO8601 time.
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFile, _ := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //CHMOD style permissions.
	writer := zapcore.AddSync(logFile)
	var defaultLogLevel zapcore.Level // This is our logging threshold system. Error and above for Prod, Debug and above for Dev.
	if prod {
		defaultLogLevel = zapcore.ErrorLevel
	} else {
		defaultLogLevel = zapcore.DebugLevel
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger
}

// NewLogger /** Returns our logger struct. We use this so we can expand the logger in future. **/
func NewLogger(prod bool, logName string) Logger {
	zap := InitLogger(prod, logName)
	sugar := zap.Sugar()
	return Logger{Zap: zap, Sugar: sugar}
}
