package main

import (
	"golangqatestdesu/config"
	"golangqatestdesu/internal/app"
	"golangqatestdesu/pkg/logger"
)

func main() {
	Log := logger.NewLogger()
	Cfg := config.NewCondig()
	if Cfg == nil {
		Log.Info("bad config")
		return
	}
	App := app.NewApp(*Cfg, *Log)
	App.Run()
}
