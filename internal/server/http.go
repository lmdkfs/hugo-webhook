package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lmdkfs/hugo-webhook/internal/handler"
	"github.com/lmdkfs/hugo-webhook/internal/middleware"
	"github.com/lmdkfs/hugo-webhook/pkg/helper/resp"
	"github.com/lmdkfs/hugo-webhook/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	hugoWebhookHandler *handler.HugoWebhookHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)
	r.GET("/", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi Richie!",
		})
	})
	r.POST("/hugo_webhook", hugoWebhookHandler.UpdateWebSite)

	return r
}
