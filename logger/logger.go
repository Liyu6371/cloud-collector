package logger

import (
	"cloud-collector/config"
	"cloud-collector/define"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

var (
	defaultLog      *logrus.Logger
	defaultLogLevel = logrus.WarnLevel

	DefaultLogMaxSize    = 10   // 每个日志文件最大10MB
	DefaultLogMaxBackups = 3    // 保留最近的3个日志文件
	DefaultLogMaxAge     = 7    // 保留最近7天的日志
	DefaultLogCompress   = true // 是否压缩旧日志

	LogLevelMap = map[string]logrus.Level{
		"ERROR": logrus.ErrorLevel,
		"WARN":  logrus.WarnLevel,
		"INFO":  logrus.InfoLevel,
		"DEBUG": logrus.DebugLevel,
	}

	LogPathMap = map[string]string{
		"darwin":  "/tmp/" + define.NameSpace + ".log",
		"linux":   "/var/log/gse/" + define.NameSpace + ".log",
		"windows": "C:\\gse\\logs\\" + define.NameSpace + ".log",
	}

	defaultLogPath = "/var/log/gse/" + define.NameSpace + ".log"
)

func init() {
	// 根据系统来判断存储的日志文件的路径
	system := strings.ToLower(runtime.GOOS)
	if v, ok := LogPathMap[system]; ok {
		defaultLogPath = v
	}
	defaultLog = logrus.New()
	defaultLog.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	defaultLog.SetLevel(defaultLogLevel)
	defaultLog.SetOutput(&lumberjack.Logger{
		Filename:   defaultLogPath,
		MaxSize:    DefaultLogMaxSize,
		MaxBackups: DefaultLogMaxBackups,
		MaxAge:     DefaultLogMaxAge,
		Compress:   DefaultLogCompress,
	})
}

func SetLogByConfig(c *config.Logger) {
	if c == nil {
		return
	}
	if level, ok := LogLevelMap[c.Level]; ok {
		defaultLog.SetLevel(level)
	}
	if c.Path != "" {
		defaultLog.SetOutput(
			&lumberjack.Logger{
				Filename:   c.Path,
				MaxSize:    DefaultLogMaxSize,
				MaxBackups: DefaultLogMaxBackups,
				MaxAge:     DefaultLogMaxAge,
				Compress:   DefaultLogCompress,
			})
	}
}

// Debug 相关
func Debug(args ...interface{}) {
	defaultLog.Debug(args...)
}

func Debugln(args ...interface{}) {
	defaultLog.Debugln(args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLog.Debugf(format, args...)
}

// Info 相关
func Info(args ...interface{}) {
	defaultLog.Info(args...)
}

func Infoln(args ...interface{}) {
	defaultLog.Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	defaultLog.Infof(format, args...)
}

// Warn 相关
func Warn(args ...interface{}) {
	defaultLog.Warn(args...)
}

func Warnln(args ...interface{}) {
	defaultLog.Warnln(args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLog.Warnf(format, args...)
}

// Error 相关
func Error(args ...interface{}) {
	defaultLog.Error(args...)
}

func Errorln(args ...interface{}) {
	defaultLog.Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLog.Errorf(format, args...)
}
