package repository

import (
	"gorm.io/gorm"

	"github.com/bivek/fmt_backend/constants"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/utils"
)

type FriendRequestRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewFriendRequestRepository(db infrastructure.Database, logger infrastructure.Logger) FriendRequestRepository {
	return FriendRequestRepository{
		db:     db,
		logger: logger,
	}
}

func (fr FriendRequestRepository) WithTrx(trxHandle *gorm.DB) FriendRequestRepository {
	if trxHandle == nil {
		fr.logger.Zap.Error("Transaction Database not found")
		return fr
	}
	fr.db.DB = trxHandle
	return fr
}

// send request
func (fr FriendRequestRepository) SendRequest(friendRequest models.FriendRequest) error {
	return fr.db.DB.Create(&friendRequest).Error
}

// acceptrequest
func (fr FriendRequestRepository) AcceptRequest(friendRequest models.FriendRequest) error {
	return fr.db.DB.Model(&models.FriendRequest{}).
		Where("sender = ?", friendRequest.Sender).
		Updates(map[string]interface{}{
			"status": friendRequest.Status,
		}).Error
}

//get accepted friends list

func (fr FriendRequestRepository) GetAcceptedFriend(pagination utils.Pagination, clientID int) ([]models.Clients, int64, error) {
	var friendlist []models.Clients

	var count int64

	//queryBuilder := e.db.DB.Model(&models.EventParticipation{}).
	// Joins("join events on events.id = event_participations.event_id").
	// Joins("join residents on residents.id = event_participations.resident_id")

	queryBuilder := fr.db.DB.Model(&models.Clients{}).Joins("join friendrequest on friendrequest.sender = ? ", clientID).Where("status = ?", constants.Accepted)

	queryBuilder = queryBuilder.Offset(pagination.Offset).Order("created_at desc")

	if !pagination.All {
		queryBuilder = queryBuilder.Limit(pagination.PageSize)
	}

	err := queryBuilder.Find(&friendlist).Error

	return friendlist, count, err

}

// get pending friendlist
func (fr FriendRequestRepository) GetPendingFriend(pagination utils.Pagination, clientID int) ([]models.ClientResponse, int64, error) {
	var friendlist []models.ClientResponse

	var count int64

	queryBuilder := fr.db.DB.Model(&models.Clients{}).
		Joins("left join friendrequest on clients.id = friendrequest.receiver").
		Where("friendrequest.status = ? ", constants.Pending).
		Where("friendrequest.sender= ?", clientID)

	queryBuilder = queryBuilder.Offset(pagination.Offset).Order("created_at desc")

	if !pagination.All {
		queryBuilder = queryBuilder.Limit(pagination.PageSize)
	}

	err := queryBuilder.Find(&friendlist).Error

	return friendlist, count, err
	// 	select * from clients left join friendrequest on
	// clients.id=friendrequest.receiver
	// where friendrequest.sender =8 and friendrequest.status="pending";

}

// cancle request
func (fr FriendRequestRepository) CancleRequest(clientID int) error {

	err := fr.db.DB.Where("receiver = ? ", clientID).Delete(&models.FriendRequest{}).Error

	return err

}

//get All Un-Friend list

func (fr FriendRequestRepository) GetUnFriend(pagination utils.Pagination, clientID int) ([]models.ClientResponse, int64, error) {
	var unfriendlist []models.ClientResponse

	var count int64

	queryBuilder := fr.db.DB.
		Table("clients").
		Select("*").
		Joins("LEFT JOIN friendrequest ON clients.id = friendrequest.receiver").
		Where("clients.id NOT IN (?) AND clients.id NOT IN (?)",
			fr.db.DB.Table("friendrequest").Select("receiver").Where("sender = ?", clientID),
			fr.db.DB.Table("friendrequest").Select("sender").Where("receiver = ?", clientID),
		)

	queryBuilder = queryBuilder.Offset(pagination.Offset).Order("clients.created_at desc ")

	if !pagination.All {
		queryBuilder = queryBuilder.Limit(pagination.PageSize)
	}

	err := queryBuilder.Find(&unfriendlist).Error

	return unfriendlist, count, err

	// select * from clients left join friendrequest on
	// clients.id = friendrequest.receiver
	// where clients.id not in
	// (select receiver from friendrequest where sender=8) AND
	// clients.id not in
	// (select sender from friendrequest where receiver=8);

}
