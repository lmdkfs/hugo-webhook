package main

import (
	"fmt"

	"github.com/lmdkfs/hugo-webhook/cmd/server/wire"
	"github.com/lmdkfs/hugo-webhook/pkg/config"
	"github.com/lmdkfs/hugo-webhook/pkg/http"
	"github.com/lmdkfs/hugo-webhook/pkg/log"
	"go.uber.org/zap"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	logger.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))

	app, cleanup, err := wire.NewWire(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	http.Run(app, fmt.Sprintf("0.0.0.0:%d", conf.GetInt("http.port")))
}
