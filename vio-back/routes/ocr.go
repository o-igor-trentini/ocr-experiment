package routes

import (
	"vio-back/api/ocr"
	"vio-back/controllers"

	"github.com/gin-gonic/gin"
)

func OCR(r *gin.Engine, ocrService ocr.Service) {
	controller := controllers.Links(ocrService)

	route := r.Group("/ocr")
	{
		route.POST("/", controller.ToText)
	}
}
