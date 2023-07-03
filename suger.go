package log

import (
	"go.uber.org/zap"
	"sync"
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

func (l *logMap) WithOptionPath(path string) *zap.SugaredLogger {

	if suger := l.getMap(path); suger != nil {
		return suger
	}

	logger := initLogger(func(options *LoggerOptions) {
		options.Path = path
	})

	l.addMap(path, logger)

	return logger.Sugar()
}
