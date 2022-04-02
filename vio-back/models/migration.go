package models

import "time"

type Migration struct {
	ID   int64     `gorm:"primaryKey;autoIncrement;column:mi_id"`
	Name string    `gorm:"column:mi_name"`
	Date time.Time `gorm:"column:mi_execution_date"`
}

func (Migration) TableName() string {
	return "mi_migrations"
}
