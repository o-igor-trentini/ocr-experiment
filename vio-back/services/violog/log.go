package violog

import (
	"time"

	"gorm.io/gorm"
)

type LogError struct {
	ID        int64     `gorm:"column:err_id;primaryKey;autoIncrement"`
	Message   string    `gorm:"column:err_message"`
	CreatedAt time.Time `gorm:"column:err_date"`
}

func (LogError) TableName() string {
	return "log.err_errors"
}

type LogDatabase struct {
	db *gorm.DB
}

func (l LogDatabase) Write(data []byte) (n int, err error) {
	l.db.Create(&LogError{Message: string(data)})

	return 0, nil
}

func NewLogDatabase(db *gorm.DB) LogDatabase {
	log := LogDatabase{db}

	return log

}
