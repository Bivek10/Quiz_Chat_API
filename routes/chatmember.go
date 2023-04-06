package routes

import (
	"github.com/bivek/fmt_backend/controllers"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/middlewares"
)

// ChatMemberRoutes -> struct
type ChatMemberRoutes struct {
	logger               infrastructure.Logger
	router               infrastructure.Router
	chatMemberController controllers.ChatMemberController
	middleware           middlewares.FirebaseAuthMiddleware
}

// NewChatMemberRoutes -> creates new chatMember controller
func NewChatMemberRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	chatMemberController controllers.ChatMemberController,
	middleware middlewares.FirebaseAuthMiddleware,
) ChatMemberRoutes {
	return ChatMemberRoutes{
		router:               router,
		logger:               logger,
		chatMemberController: chatMemberController,
		middleware:           middleware,
	}
}

// Setup chatMember routes
func (c ChatMemberRoutes) Setup() {
	c.logger.Zap.Info(" Setting up chatMember routes")
	chatMember := c.router.Gin.Group("/chatmember")
	{
		chatMember.POST("", c.chatMemberController.CreateChatMember)
		chatMember.GET("", c.chatMemberController.GetAllChatMember)
		chatMember.GET("/:id", c.chatMemberController.GetOneChatMember)
		chatMember.PUT("/:id", c.chatMemberController.UpdateOneChatMember)
		chatMember.DELETE("/:id", c.chatMemberController.DeleteOneChatMember)
	}
}
