package ocr

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
	"vio-back/helpers"
	"vio-back/services/vioerror"

	"github.com/otiai10/gosseract"
)

type Service interface {
	ConvertBase64ToStruct(b64 string) vioerror.ResponseError
}

type serviceImpl struct {
	ocrRepository Repository
}

func NewService(ocrRepository Repository) Service {
	return &serviceImpl{ocrRepository}
}

func (s serviceImpl) ConvertBase64ToStruct(b64 string) vioerror.ResponseError {
	// if err := s.ocrRepository.Create(newLink); err != nil {
	// 	return encerror.NewError(services.ErrGenerateLink, services.ErrGenerateLinkMsg, nil)
	// }

	ocr := gosseract.NewClient()
	defer ocr.Close()

	img, errConvert := convertBase64ToImage(b64)
	if errConvert != nil {
		fmt.Printf("erro: %s", errConvert)
	}

	imgBytes, errImgBytes := getImageBytes(img)
	if errImgBytes != nil {
		fmt.Print(errImgBytes)
	}

	ocr.SetImageFromBytes(imgBytes.Bytes())

	text, errToText := ocr.Text()
	if errToText != nil {
		fmt.Print(errToText)
	}

	helpers.ClearTerminal()
	fmt.Print(text)

	return nil
}

func convertBase64ToImage(b64 string) (*image.Image, error) {
	coI := strings.Index(b64, ",")
	rawImage := b64[coI+1:]

	unbased, err := base64.StdEncoding.DecodeString(rawImage)
	if err != nil {
		return nil, fmt.Errorf("erro: %s", err)
	}

	res := bytes.NewReader(unbased)
	jpgImg, errConvert := jpeg.Decode(res)
	if errConvert != nil {
		return nil, fmt.Errorf("erro: %s", errConvert)
	}

	return &jpgImg, nil
}

func getImageBytes(jpg *image.Image) (*bytes.Buffer, error) {
	if jpg == nil {
		imgfile, err := os.Open("/home/igor/projetos/pessoal/vio/vio-back/images/frase.jpg")
		if err != nil {
			return &bytes.Buffer{}, fmt.Errorf("erro: %s", err)
		}
		defer imgfile.Close()

		img, errDecode := jpeg.Decode(imgfile)
		if errDecode != nil {
			return &bytes.Buffer{}, fmt.Errorf("erro: %s", errDecode)
		}

		jpg = &img
	}

	buf := bytes.NewBuffer(nil)

	if err := png.Encode(buf, *jpg); err != nil {
		return &bytes.Buffer{}, fmt.Errorf("erro: %s", err)
	}

	return buf, nil
}
