package routes

import (
	"github.com/bivek/fmt_backend/controllers"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/middlewares"
)

// ChatMessageRoutes -> struct
type ChatMessageRoutes struct {
	logger                infrastructure.Logger
	router                infrastructure.Router
	chatMessageController controllers.ChatMessageController
	middleware            middlewares.FirebaseAuthMiddleware
}

// NewChatMessageRoutes -> creates new chatMessage controller
func NewChatMessageRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	chatMessageController controllers.ChatMessageController,
	middleware middlewares.FirebaseAuthMiddleware,
) ChatMessageRoutes {
	return ChatMessageRoutes{
		router:                router,
		logger:                logger,
		chatMessageController: chatMessageController,
		middleware:            middleware,
	}
}

// Setup chatMessage routes
func (c ChatMessageRoutes) Setup() {
	c.logger.Zap.Info(" Setting up chatMessage routes")
	chatMessage := c.router.Gin.Group("/chatmessage")
	{
		chatMessage.POST("", c.chatMessageController.CreateChatMessage)
		chatMessage.GET(":roomid", c.chatMessageController.GetAllChatMessage)

		chatMessage.DELETE(":id", c.chatMessageController.DeleteOneChatMessage)
	}
}
