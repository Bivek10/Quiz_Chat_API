package routes

import (
	"github.com/bivek/fmt_backend/controllers"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/middlewares"
)

type FriendRequestRoutes struct {
	logger                  infrastructure.Logger
	router                  infrastructure.Router
	friendRequestController controllers.FriendRequestController
	middlewares             middlewares.FirebaseAuthMiddleware
	trxMiddleware           middlewares.DBTransactionMiddleware
}

//setup quiz routes

func (i FriendRequestRoutes) Setup() {
	i.logger.Zap.Info("setting up FriendRequest routes")
	quizs := i.router.Gin.Group("/request")
	{
		quizs.GET("/accepted/:id", i.friendRequestController.GetAcceptedFriend)
		quizs.GET("/pending/:id", i.friendRequestController.GetPendingFriend)
		quizs.GET("/unfriend/:id", i.friendRequestController.GetUnFriend)
		quizs.GET("/cancel:id", i.friendRequestController.CancleRequest)

		quizs.POST("", i.trxMiddleware.DBTransactionHandle(), i.friendRequestController.SendRequest)
	}
}

//NewQuizRoutes -> creates new quiz

func NewFriendRequestRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	friendRequest controllers.FriendRequestController,
	middlewares middlewares.FirebaseAuthMiddleware,
	trxMiddleware middlewares.DBTransactionMiddleware,
) FriendRequestRoutes {
	return FriendRequestRoutes{
		logger:                  logger,
		router:                  router,
		friendRequestController: friendRequest,
		middlewares:             middlewares,
		trxMiddleware:           trxMiddleware,
	}
}
