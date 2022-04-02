package ocr

import (
	"gorm.io/gorm"
)

type Repository interface {
	// Create(newLink *models.Links) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db}
}

func (r repositoryImpl) Create() error {
	// if err := r.db.Create(&newLink).Error; err != nil {
	// 	return err
	// }

	return nil
}
