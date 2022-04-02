package routes

import (
	"vio-back/api/ocr"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB, r *gin.Engine) {
	ocrRepository := ocr.NewRepository(db)
	ocrService := ocr.NewService(ocrRepository)

	OCR(r, ocrService)
}
