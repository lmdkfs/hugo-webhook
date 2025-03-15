package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v69/github"
	"github.com/lmdkfs/hugo-webhook/internal/service"
	"github.com/lmdkfs/hugo-webhook/pkg/helper/resp"
)

func NewHugoWebhookHandler(handler *Handler, hugoWebhook service.HugoWebHookService) *HugoWebhookHandler {
	return &HugoWebhookHandler{
		Handler:            handler,
		hugoWebhookService: hugoWebhook,
	}
}

type HugoWebhookHandler struct {
	*Handler
	hugoWebhookService service.HugoWebHookService
}

func (h *HugoWebhookHandler) UpdateWebSite(c *gin.Context) {
	fmt.Printf("request: %+v\n", c.Request)
	payload, err := github.ValidatePayload(c.Request, []byte("PkheGsB3fOypw6kw"))
	if err != nil {
		resp.HandleError(c, http.StatusBadRequest, 1, err.Error(), nil)
	}
	event, err := github.ParseWebHook(github.WebHookType(c.Request), payload)
	if err != nil {
		resp.HandleError(c, http.StatusBadRequest, 1, err.Error(), nil)
	}
	switch event := event.(type) {
	case *github.PushEvent:
		fmt.Printf("PushEvent : %v\n", event)
		err := h.hugoWebhookService.UpdateWebSite()
		if err != nil {
			resp.HandleError(c, http.StatusInternalServerError, 1, err.Error(), nil)
		}
		resp.HandleSuccess(c, event)
	case *github.PullRequestEvent:
		fmt.Printf("PullRequestEvent: %v\n", event)
		resp.HandleSuccess(c, event)
	}

}
