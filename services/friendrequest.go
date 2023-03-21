package services

import (
	"gorm.io/gorm"

	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/repository"
	"github.com/bivek/fmt_backend/utils"
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

func (f FriendRequestService) AcceptRequest(friendRequest models.FriendRequest) error {
	err := f.repository.AcceptRequest(friendRequest)
	return err
}

func (f FriendRequestService) GetAcceptedFriend(pagination utils.Pagination, clientID int) ([]models.Clients, int64, error) {
	friendlist, count, err := f.repository.GetAcceptedFriend(pagination, clientID)
	return friendlist, count, err
}

func (f FriendRequestService) GetPendingFriend(pagination utils.Pagination, clientID int) ([]models.Clients, int64, error) {
	friendlist, count, err := f.repository.GetPendingFriend(pagination, clientID)
	return friendlist, count, err
}

func (f FriendRequestService) CancleRequest(clientID int) error {
	err := f.repository.CancleRequest(clientID)
	return err
}

func (f FriendRequestService) GetUnFriend(pagination utils.Pagination, clientID int) ([]models.Clients, int64, error) {
	unfriendlist, count, err := f.repository.GetUnFriend(pagination, clientID)
	return unfriendlist, count, err
}
