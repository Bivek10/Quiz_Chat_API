package routes

import (
	"github.com/bivek/fmt_backend/controllers"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/middlewares"
)

type ConversationRoutes struct {
	logger                 infrastructure.Logger
	router                 infrastructure.Router
	conversationController controllers.ConversationController
	middlewares            middlewares.FirebaseAuthMiddleware
	trxMiddleware          middlewares.DBTransactionMiddleware
}

//setup quiz routes

func (i ConversationRoutes) Setup() {
	i.logger.Zap.Info("setting up Conversation routes")
	quizs := i.router.Gin.Group("/conversation")
	{

		quizs.GET("/cancle/:id", i.conversationController.CancleMessage)
		quizs.GET("/getmessage:senderid:receiverid", i.conversationController.GetOldMessage)

		quizs.POST("", i.trxMiddleware.DBTransactionHandle(), i.conversationController.SaveMessage)
	}
}

//NewQuizRoutes -> creates new quiz

func NewConversationRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	ctrl controllers.ConversationController,
	middlewares middlewares.FirebaseAuthMiddleware,
	trxMiddleware middlewares.DBTransactionMiddleware,
) ConversationRoutes {
	return ConversationRoutes{
		logger:                 logger,
		router:                 router,
		conversationController: ctrl,
		middlewares:            middlewares,
		trxMiddleware:          trxMiddleware,
	}
}
