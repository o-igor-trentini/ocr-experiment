package controllers

import (
	"fmt"
	"vio-back/api/ocr"
	"vio-back/models"

	"github.com/gin-gonic/gin"
)

type OCRController struct {
	ocrService ocr.Service
}

func Links(ocrService ocr.Service) OCRController {
	return OCRController{ocrService}
}

func (co OCRController) ToText(c *gin.Context) {
	var b64 []models.ImgB64
	if err := c.ShouldBindJSON(&b64); err != nil {
		c.JSON(500, "deu ruim")
		fmt.Printf("erro: %s", err)
		return
	}

	toStruct, err := co.ocrService.ConvertBase64ToStruct(b64)
	if err != nil {
		c.JSON(500, "deu ruim")
		return
	}

	fmt.Print(toStruct)

	c.JSON(200, "deu boa")
}
