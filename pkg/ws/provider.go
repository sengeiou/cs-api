package ws

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewNotifier,
		NewStaffDispatcher,
		NewClientManager,
		NewHandler,
		NewRedisWorker,
	),
	fx.Invoke(
		InitClientManager,
		InitWebSocket,
		InitRedisSubscriber,
	),
)

func InitWebSocket(engine *gin.Engine, h *Handler) {
	engine.Any("/ws/chat", h.ChatHandler)
	//engine.Any("/test/ws/chat", h.TestChatHandler)
}
