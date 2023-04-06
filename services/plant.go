package services

import (
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/repository"
	"github.com/bivek/fmt_backend/utils"

	"gorm.io/gorm"
)

type PlantService struct {
	repository repository.PlantRepository
}

func NewPlantService(repository repository.PlantRepository) PlantService {
	return PlantService{
		repository: repository,
	}
}

func (c PlantService) WithTrx(trxHandle *gorm.DB) PlantService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

// createplant -> call to create the plant
func (c PlantService) CreatePlant(plant models.Plant) error {
	err := c.repository.CreatePlant(plant)
	return err
}

//GetAllPlant -> call to get all the user

func (c PlantService) GetAllPlant(pagination utils.Pagination) ([]models.Plant, int64, error) {
	return c.repository.GetAllPlants(pagination)
}

//Get plant by ID

func (c PlantService) GetPlantByID(plantID string) (models.Plant, error) {
	return c.repository.GetPlantById(plantID)
}

//update plantdata

func (c PlantService) UpdatePlant(Plant models.Plant) error {
	err := c.repository.UpdatePlant(Plant)
	return err
}
