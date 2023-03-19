package controllers

import "go.uber.org/fx"

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(NewUserController),
	fx.Provide(NewUtilityController),
	fx.Provide(NewPlantController),
	fx.Provide(NewQuizController),
	fx.Provide(NewQuestionController),
	fx.Provide(NewChoiceController),
	fx.Provide(NewHistoryController),
	fx.Provide(NewClientController),
	fx.Provide(NewFriendRequestController),
)
