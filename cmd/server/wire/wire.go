//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/lmdkfs/hugo-webhook/internal/handler"
	"github.com/lmdkfs/hugo-webhook/internal/server"
	"github.com/lmdkfs/hugo-webhook/internal/service"
	"github.com/lmdkfs/hugo-webhook/pkg/log"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

// var RepositorySet = wire.NewSet(
// 	repository.NewDb,
// 	repository.NewRepository,
// 	// repository.NewUserRepository,
// )

var ServiceSet = wire.NewSet(
	service.NewService,
	// service.NewUserService,
	service.NewHugoWebHookService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	// handler.NewUserHandler,
	handler.NewHugoWebhookHandler,
)

func NewWire(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		// RepositorySet,
		ServiceSet,
		HandlerSet,
	))
}
