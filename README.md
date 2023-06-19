# log
一个基于zap的多模块日志

example:


```package main

import (
	"go.uber.org/zap"
	logutil "github.com/zutim/log"
)

func main() {

	NewApp()

	App.Logs("").Info("你好a")
	App.Logs("order").Error("cao gege")
	App.Logs("user").Debug("I was a wang lei")
}

type Apps struct {
	Log map[string]*zap.SugaredLogger
}

func (a *Apps) Logs(key string) *zap.SugaredLogger {
	if _, ok := a.Log[key]; ok {
		return a.Log[key]
	}
	return a.Log["log"]
}

var App *Apps

func NewApp() {
	App = &Apps{
		Log: map[string]*zap.SugaredLogger{
			"log":   NewDefaultLogger(),
			"order": NewOrderLogger(),
			"user":  NewUserLogger(),
		},
	}
}

func NewDefaultLogger() *zap.SugaredLogger {
	logger := logutil.InitLogger(func(options *logutil.LoggerOptions) {
		options.Path = "log/common.log"
	})
	defer logger.Sync()
	return logger.Sugar()
}

func NewOrderLogger() *zap.SugaredLogger {
	logger := logutil.InitLogger(func(options *logutil.LoggerOptions) {
		options.Path = "log/order.log"
	})
	defer logger.Sync()
	return logger.Sugar()
}

func NewUserLogger() *zap.SugaredLogger {
	logger := logutil.InitLogger(func(options *logutil.LoggerOptions) {
		options.Path = "log/user.log"
	})
	defer logger.Sync()
	return logger.Sugar()
}
```

