package violog

import (
	"log"

	"gorm.io/gorm"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func InitLogger(db *gorm.DB) {
	out := NewLogDatabase(db)

	InfoLogger = log.New(out, "[VIO-INFO] ", log.Ldate|log.Ltime|log.Llongfile)
	WarningLogger = log.New(out, "[VIO-WARNING] ", log.Ldate|log.Ltime|log.Llongfile)
	ErrorLogger = log.New(out, "[VIO-ERROR] ", log.Ldate|log.Ltime|log.Llongfile)
}
