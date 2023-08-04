package main

import (
	"github.com/programacao-fortificada/upsec/internal/config"
	"github.com/programacao-fortificada/upsec/pkg/log"
	"go.uber.org/zap"
)

func main() {
	l := setupLog()
	defer syncLog(l)

	l.Info("reading configurations")
	cfg, err := setupConfig()
}

func setupLog() log.Log {
	zapLog, _ := zap.NewProduction()

	return log.New(zapLog.Sugar())
}

func syncLog(l log.Log) {
	if err := l.Sync(); err != nil {
		l.Error("failed to sync log", err.Error())
	}
}

func setupConfig(l log.Log) (*config.Config, error) {
	err := config.Log(l)
	if err != nil {
		l.Warn("failed to setup config log", err)
	}

	config.Local()

	cfg, err := config.Get()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
