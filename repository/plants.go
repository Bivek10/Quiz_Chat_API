package repository

import (
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/utils"

	"gorm.io/gorm"
)

type PlantRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewPlantRepository(db infrastructure.Database, logger infrastructure.Logger) PlantRepository {
	return PlantRepository{

		db:     db,
		logger: logger,
	}
}

//withTrx enables repository with transation

func (p PlantRepository) WithTrx(trxHandle *gorm.DB) PlantRepository {
	if trxHandle == nil {
		p.logger.Zap.Error("Transtion Database not found in gin context")
		return p
	}
	p.db.DB = trxHandle
	return p
}

// save ->plant
func (p PlantRepository) CreatePlant(Plant models.Plant) error {
	return p.db.DB.Create(&Plant).Error
}

//GetAll plants

func (c PlantRepository) GetAllPlants(pagination utils.Pagination) ([]models.Plant, int64, error) {
	var plants []models.Plant
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("create_at desc")
	queryBuilder = queryBuilder.Model(&models.Plant{})

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`plants`.`name` LIKE ?", searchQuery))
	}
	err := queryBuilder.Find(&plants).Offset(-1).Limit(-1).Count(&totalRows).Error
	return plants, totalRows, err
}

//Getplant by ID

func (c PlantRepository) GetPlantById(plantID string) (models.Plant, error) {
	plant := models.Plant{}
	queryBuilder := c.db.DB
	queryBuilder = queryBuilder.Model(&models.Plant{})
	queryBuilder.Where(&models.Plant{PlantId: plantID})
	err := queryBuilder.Find(&plant).Error
	return plant, err
}

// updateplant data by ID
func (c PlantRepository) UpdatePlant(Plant models.Plant) error {
	queryBuilder := c.db.DB.Updates(&Plant).Error
	return queryBuilder
}
