package logging

import (
	"sync"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func InitLogger() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"} // or specify a log file path
	l, _ := cfg.Build()
	logger = l
}

func GetLogger() *zap.Logger {
	once.Do(func() {
		InitLogger() // Initialize the logger only once
	})
	return logger
}
