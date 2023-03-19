package services

import (
	"gorm.io/gorm"

	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/repository"
)

type FriendRequestService struct {
	repository repository.FriendRequestRepository
}

func NewFriendRequestService(repo repository.FriendRequestRepository) FriendRequestService {
	return FriendRequestService{
		repository: repo,
	}
}

//WithTrx -> enable repository with transaction

func (f FriendRequestService) WithTrx(trxHandle *gorm.DB) FriendRequestService {
	f.repository = f.repository.WithTrx(trxHandle)
	return f
}

func (f FriendRequestService) SendRequest(friendRequest models.FriendRequest) error {
	err := f.repository.SendRequest(friendRequest)
	return err
}

func (f FriendRequestService) GetAcceptedFriend(pagination Utils.Pagination, clientID) 
