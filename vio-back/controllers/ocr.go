package controllers

import (
	"fmt"
	"vio-back/api/ocr"

	"github.com/gin-gonic/gin"
)

type OCRController struct {
	ocrService ocr.Service
}

func Links(ocrService ocr.Service) OCRController {
	return OCRController{ocrService}
}

func (co OCRController) ToText(c *gin.Context) {
	type imgB64 struct {
		B64 string `json:"img_base64"`
	}

	var b64 []imgB64
	if err := c.ShouldBindJSON(&b64); err != nil {
		fmt.Printf("erro: %s", err)
		return
	}

	if err := co.ocrService.ConvertBase64ToStruct(b64[0].B64); err != nil {
		fmt.Print("deu ruim")
		return
	}

	c.JSON(200, "deu boa")
}
