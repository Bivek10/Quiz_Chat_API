package controllers

import (
	"net/http"

	"github.com/bivek/fmt_backend/constants"
	"github.com/bivek/fmt_backend/errors"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/responses"
	"github.com/bivek/fmt_backend/services"
	"github.com/bivek/fmt_backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PlantController ->struct
type PlantController struct {
	logger          infrastructure.Logger
	plantService    services.PlantService
	firbaseServices services.FirebaseService
	env             infrastructure.Env
}

func NewPlantController(logger infrastructure.Logger, plantService services.PlantService, firebaseServices services.FirebaseService, env infrastructure.Env) PlantController {
	return PlantController{
		logger:          logger,
		plantService:    plantService,
		firbaseServices: firebaseServices,
		env:             env,
	}
}

// create plant
func (cc PlantController) CreatePlant(c *gin.Context) {
	plant := models.Plant{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&plant); err != nil {
		cc.logger.Zap.Error("Error [create plant] (should Bind Json): ", err)
		responses.HandleError(c, err)
		return
	}

	if err := cc.plantService.WithTrx(trx).CreatePlant(plant); err != nil {
		cc.logger.Zap.Error("Error [create user] [db create user]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create user")

		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Plant data created succesfully")

}

// get allplant  datassssss
func (cc PlantController) GetAllPlant(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	plants, count, err := cc.plantService.GetAllPlant(pagination)
	if err != nil {
		cc.logger.Zap.Error("Error finding plant records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get users data")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, plants, count)
}

// get plant by ID

func (cc PlantController) GetPlantByID(c *gin.Context) {
	plantID := c.Param("id")
	plants, err := cc.plantService.GetPlantByID(plantID)
	if err != nil {
		cc.logger.Zap.Error("Error finding plant records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get plant data by ID")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, plants)
}

func (cc PlantController) UpdatePlant(c *gin.Context) {
	plant := models.Plant{}
	err := cc.plantService.UpdatePlant(plant)
	if err != nil {
		cc.logger.Zap.Error("Error finding plant records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to udpate plant data")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Plant data update successfully")
}
