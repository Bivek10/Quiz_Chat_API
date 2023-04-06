package services

import (
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/repository"
	"gorm.io/gorm"
)

// ClientService -> struct
type ClientService struct {
	repository repository.ClientRepository
}

// NewClientService -> creates a new ClientService
func NewClientService(repository repository.ClientRepository) ClientService {
	return ClientService{
		repository: repository,
	}
}

// WithTrx -> enables repository with transaction
func (c ClientService) WithTrx(trxHandle *gorm.DB) ClientService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

// CreateUser -> call to create the User
func (c ClientService) CreateClient(client models.Clients) error {
	err := c.repository.CreateClient(client)
	return err
}

func (c ClientService) LoginClient(Email string) (models.Clients, error) {
	clients, err := c.repository.LoginClient(Email)
	return clients, err
}
