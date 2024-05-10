package log

import (
	"fmt"
	"go.uber.org/zap"
	"sync"
	"time"
)

type logMap struct {
	sync.RWMutex
	m map[string]*zap.Logger
}

func NewLogMap() *logMap {
	return &logMap{
		m: make(map[string]*zap.Logger),
	}
}

func (l *logMap) addMap(path string, logger *zap.Logger) {
	l.Lock()
	defer l.Unlock()
	if _, ok := l.m[path]; ok {
		l.m[path] = logger
	}
}

func (l *logMap) getMap(path string) *zap.SugaredLogger {
	l.RLock()
	defer l.RUnlock()
	if path == "" {
		path = "log/run.log"
	}
	if _, ok := l.m[path]; ok {
		return l.m[path].Sugar()
	}
	return nil
}

func (l *logMap) deleteMap(path string) {
	l.Lock()
	defer l.Unlock()
	if _, ok := l.m[path]; ok {
		delete(l.m, path)
	}
}

func (l *logMap) WithOptionPath(op LoggerOptions) *zap.SugaredLogger {
	// 加上今天的日期
	path := op.Path + "-" + time.Now().Format("2006-01-02") + ".log"
	fmt.Println(path)
	if suger := l.getMap(path); suger != nil {
		return suger
	}

	// map中删除，不再写入日志
	yesterdayPath := op.Path + "-" + time.Now().AddDate(0, 0, -1).Format("2006-01-02") + ".log"
	l.deleteMap(yesterdayPath)

	op.Path = path
	logger := initLogger(func(options *LoggerOptions) {
		*options = op
	})

	l.addMap(path, logger)

	return logger.Sugar()
}
