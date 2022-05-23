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
	first := c.Query("first")
	second := c.Query("second")

	fmt.Println(first)
	fmt.Println(second)

	// limpar informações
	// gravar em um arquivo .txt
	// armazenar em uma pasta no S3

	c.JSON(200, "deu boa")
}

// fileExample := []byte(first + second)
// now := time.Now().Format("2006121545")
// fileName := fmt.Sprintf("paramtest_%s.txt", now)

// err := ioutil.WriteFile("./"+fileName, fileExample, 0644)
// if err != nil {
// 	fmt.Print("error")
// }

// var b64 []models.ImgB64
// if err := c.ShouldBindJson(&b64); err != nil {
// 	fmt.Printf("erro: %s", err)
// 	c.JSON(500, "deu ruim")
// 	return
// }

// toStruct, err := co.ocrService.ConvertBase64ToStruct(b64)
// if err != nil {
// 	c.JSON(500, "deu ruim")
// 	return
// }

// fmt.Print(toStruct)
