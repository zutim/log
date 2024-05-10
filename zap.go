package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerOptions struct {
	Path       string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type Option func(options *LoggerOptions)

// InitLogger 初始化日志记录器
func initLogger(opts ...Option) *zap.Logger {

	defaultOptions := LoggerOptions{Path: "log/run.log", MaxSize: 100, MaxBackups: 10, MaxAge: 30, Compress: false}

	for _, setter := range opts {
		setter(&defaultOptions)
	}

	// 配置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.DebugLevel)

	// 配置日志输出位置和格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 创建日志文件的配置
	logger := &lumberjack.Logger{
		Filename:   defaultOptions.Path,
		MaxSize:    defaultOptions.MaxSize,    // 单位：MB，达到该大小时触发分割
		MaxBackups: defaultOptions.MaxBackups, // 最多保留的旧日志文件数量
		MaxAge:     defaultOptions.MaxAge,     // 单位：天，旧日志文件保留的时间
		Compress:   defaultOptions.Compress,   // 分割后的日志文件是否压缩
	}

	// 配置日志核心
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(logger),
		atomicLevel,
	)
	defer logger.Close()

	// 创建 Logger 对象
	loggerObj := zap.New(core, zap.AddCaller())

	return loggerObj
}

type Writer struct {
	Log *zap.SugaredLogger
}

func (w Writer) Printf(format string, args ...interface{}) {
	w.Log.Infof(format, args...)
}
