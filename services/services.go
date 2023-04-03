package services

import "go.uber.org/fx"

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewFirebaseService),
	fx.Provide(NewStorageBucketService),
	fx.Provide(NewUserService),
	fx.Provide(NewTwilioService),
	fx.Provide(NewGmailService),
	fx.Provide(NewS3BucketService),
	fx.Provide(NewPlantService),
	fx.Provide(NewQuizService),
	fx.Provide(NewQuestionServices),
	fx.Provide(NewChoiceServices),
	fx.Provide(NewQuizHistoryServices),
	fx.Provide(NewClientService),
	fx.Provide(NewJwtService),
	fx.Provide(NewFriendRequestService),
	fx.Provide(NewConversationService),
	fx.Provide(NewChatMemberService),
	fx.Provide(NewChatMessageService),
	fx.Provide(NewChatRoomService),
)
