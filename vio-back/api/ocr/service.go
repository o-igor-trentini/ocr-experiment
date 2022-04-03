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
	"vio-back/appconst"
	"vio-back/helpers"
	"vio-back/models"
	"vio-back/services/vioerror"

	"github.com/otiai10/gosseract"
)

type Service interface {
	ConvertBase64ToStruct(b64 []models.ImgB64) (models.Vio, vioerror.ResponseError)
}

type serviceImpl struct {
	ocrRepository Repository
}

func NewService(ocrRepository Repository) Service {
	return &serviceImpl{ocrRepository}
}

// if err := s.ocrRepository.Create(newLink); err != nil {
// 	return encerror.NewError(services.ErrGenerateLink, services.ErrGenerateLinkMsg, nil)
// }
func (s serviceImpl) ConvertBase64ToStruct(b64 []models.ImgB64) (models.Vio, vioerror.ResponseError) {
	toStruct := models.Vio{}

	var texts []string
	for i, value := range b64 {
		text, err := s.getTextInImage(value.B64)
		if err != nil {
			fmt.Printf("erro ao converter as imagens; erro: %s", err)
			return toStruct, nil
		}

		switch i {
		case 0:
			s.clearAndConvertFirstText(text, &toStruct)

		case 1:

		default:
			fmt.Print("mais de duas imagens foram recebidas")
			return toStruct, nil
		}

		texts = append(texts, text)
	}

	helpers.ClearTerminal()
	fmt.Print(texts)

	// quando pegar o base64, separar o que tem na string antes do ACC, lógica inversa no segundo base64
	// text = strings.Replace(text, "\n", "", -1)
	// text = text[strings.Index(text, "Nome"):]

	// helpers.ClearTerminal()
	// fmt.Print(text)

	// iName := strings.Index(text, "Nome")
	// iDoc := strings.Index(text, "Doc")

	// data := models.Vio{}

	// helpers.ClearTerminal()
	// fmt.Printf("nome: %s", text[iName+4:iDoc])

	// data.Nome = text[iName+4 : iDoc]

	return toStruct, nil
}

func (s serviceImpl) convertBase64ToImage(b64 string) (*image.Image, error) {
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

func (s serviceImpl) getImageBytes(jpg *image.Image) (*bytes.Buffer, error) {
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

func (s serviceImpl) getTextInImage(b64 string) (string, error) {
	ocr := gosseract.NewClient()
	defer ocr.Close()

	img, errConvert := s.convertBase64ToImage(b64)
	if errConvert != nil {
		fmt.Printf("erro: %s", errConvert)
	}

	imgBytes, errImgBytes := s.getImageBytes(img)
	if errImgBytes != nil {
		fmt.Print(errImgBytes)
	}

	ocr.SetImageFromBytes(imgBytes.Bytes())

	text, errToText := ocr.Text()
	if errToText != nil {
		fmt.Print(errToText)
	}

	return text, nil
}

func (s serviceImpl) clearAndConvertFirstText(text string, toStruct *models.Vio) {
	// removendo quebras de linha
	text = strings.Replace(text, "\n", "", -1)

	// removendo texto desnecessário
	text = text[strings.Index(text, string(appconst.Nome)):]

	helpers.ClearTerminal()
	fmt.Println(text)

	// pegando índices dos campos na string
	iNome := strings.Index(text, string(appconst.Nome))
	iDoc := strings.Index(text, string(appconst.DocIdentidadeOrgEmissorUf))
	// iCpf

	fmt.Printf("nome: %s", text[iNome+len(appconst.Nome):iDoc])
	toStruct.Nome = text[iNome+len(appconst.Nome) : iDoc]
}
