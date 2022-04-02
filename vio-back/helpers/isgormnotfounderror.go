package helpers

import (
	"errors"

	"gorm.io/gorm"
)

func IsGormRecordNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
