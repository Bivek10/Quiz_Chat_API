package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/utils"
)

// UserRepository -> database structure
type UserRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewUserRepository -> creates a new User repository
func NewUserRepository(db infrastructure.Database, logger infrastructure.Logger) UserRepository {
	return UserRepository{
		db:     db,
		logger: logger,
	}
}

// WithTrx enables repository with transaction
func (c UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}

// Save -> User
func (c UserRepository) Create(User models.User) error {

	return c.db.DB.Create(&User).Error
}

// GetAllUser -> Get All users
func (c UserRepository) GetAllUsers(pagination utils.Pagination) ([]models.User, int64, error) {
	var users []models.User
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.User{})

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`users`.`name` LIKE ?", searchQuery))
	}
	err := queryBuilder.
		Find(&users).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
	return users, totalRows, err
}

func (c UserRepository) LoginUser(email string, password string) (models.User, error) {
	users := models.User{}

	queryBuilder := c.db.DB
	queryBuilder = queryBuilder.Model(&models.User{})
	queryBuilder.Where(&models.User{Email: email})
	err := queryBuilder.Find(&users).Error
	fmt.Printf("hased password %v", users.Password)
	fmt.Printf("unhased password %v", password)
	isTrue := utils.DecryptPassword([]byte(users.Password), []byte(password))
	if isTrue {
		fmt.Println("password matched")

		//fmt.Println(err)

	}
	return users, err

}
