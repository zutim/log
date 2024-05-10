# log
一个基于zap的多模块日志

example:


```package main

import (
	"github.com/zutim/log"
	"go.uber.org/zap"
)

func main() {

	NewApp()
	UserLog().Info("test user")
	OrderLog().Info("test order")

}

type Apps struct {
	Log *zap.SugaredLogger
}

var App *Apps

func NewApp() {
	logsMap := log.NewLogMap().WithOptionPath(log.LoggerOptions{})
	App = &Apps{
		Log: logsMap,
	}
}

func UserLog() *zap.SugaredLogger {
	op := log.LoggerOptions{
		Path: "user",
	}
	return log.NewLogMap().WithOptionPath(op)
}

func OrderLog() *zap.SugaredLogger {
	op := log.LoggerOptions{
		Path: "order",
	}
	return log.NewLogMap().WithOptionPath(op)
}

```

