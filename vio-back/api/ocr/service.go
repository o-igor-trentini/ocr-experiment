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
	"vio-back/services"
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

func (s serviceImpl) ConvertBase64ToStruct(b64 []models.ImgB64) (models.Vio, vioerror.ResponseError) {
	toStruct := models.Vio{}

	for i, value := range b64 {
		text, err := s.getTextInImage(value.B64)
		if err != nil {
			fmt.Printf("erro ao converter as imagens; erro: %s", err)
			return toStruct, vioerror.NewError(services.ErrConvertBase64ToStruct, services.ErrConvertBase64ToStructMsg, nil)
		}

		switch i {
		case 0:
			s.clearAndConvertFirstText(text, &toStruct)

		case 1:
			s.clearAndConvertSecondText(text, &toStruct)

		default:
			fmt.Print("mais de duas imagens foram recebidas")
			return toStruct, nil
		}
	}

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
	// removendo quebras de linha e possível interpretação errada de "S"
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "$", "S", -1)

	// removendo texto desnecessário
	text = text[strings.Index(text, string(appconst.Nome)):]

	textLen := len(text)

	helpers.ClearTerminal()
	fmt.Println(text)

	// pegando índices dos campos na string
	iNome := strings.Index(text, string(appconst.Nome))
	iDoc := strings.Index(text, string(appconst.DocIdentidadeOrgEmissorUf))
	iCpf := strings.Index(text, string(appconst.Cpf))
	iNascimento := strings.Index(text, string(appconst.DataDeNascimento))
	iPai := strings.Index(text, string(appconst.FiliacaoPai))
	iMae := strings.Index(text, string(appconst.FiliacaoMae))
	iPermissao := strings.Index(text, string(appconst.Permissao))
	iAcc := strings.Index(text, string(appconst.Acc))

	fmt.Println(text[iNome+len(appconst.Nome) : iDoc])
	toStruct.Nome = text[iNome+len(appconst.Nome) : iDoc]

	fmt.Println(text[iDoc+len("Doc. Identidade/Org. Emissor/UF") : iCpf])
	toStruct.DocIdentidadeOrgEmissorUf = text[iDoc+len("Doc. Identidade/Org. Emissor/UF") : iCpf]

	fmt.Println(text[iCpf+len(appconst.Cpf) : iNascimento])
	toStruct.Cpf = text[iCpf+len(appconst.Cpf) : iNascimento]

	fmt.Println(text[iNascimento+len("Data de Nascimento") : iPai-len("filiacao ")])
	toStruct.DataDeNascimento = text[iNascimento+len("Data de Nascimento") : iPai-len("Filiacao ")]

	fmt.Println(text[iPai+len(appconst.FiliacaoPai) : iMae-len("Filiacao ")])
	toStruct.FiliacaoPai = text[iPai+len(appconst.FiliacaoPai) : iMae-len("Filiacao ")]

	fmt.Println(text[iMae+len(appconst.FiliacaoMae) : iPermissao])
	toStruct.FiliacaoMae = text[iMae+len(appconst.FiliacaoMae) : iPermissao]

	fmt.Println(text[iPermissao+len(appconst.Permissao) : iAcc])
	toStruct.Permissao = text[iPermissao+len(appconst.Permissao) : iAcc]

	fmt.Println(text[iAcc+len(appconst.Acc) : textLen-len("eee")])
	toStruct.Acc = text[iAcc+len(appconst.Acc) : textLen-len("eee")]
}

func (s serviceImpl) clearAndConvertSecondText(text string, toStruct *models.Vio) {
	// removendo quebras de linha e possível interpretação errada de "S"
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "$", "S", -1)

	// removendo texto desnecessário
	text = text[strings.Index(text, string(appconst.CatHabilitacao)):]

	textLen := len(text)

	// helpers.ClearTerminal()
	// fmt.Println(text)

	// pegando índices dos campos na string
	iCatHab := strings.Index(text, string(appconst.CatHabilitacao))
	iNumRegistro := strings.Index(text, string(appconst.NumRegistro))
	iValidade := strings.Index(text, string(appconst.Validade))
	iPrimHab := strings.Index(text, string(appconst.PrimHabilitacao))
	iObs := strings.Index(text, string(appconst.Observacoes))
	iLocal := strings.Index(text, string(appconst.Local))
	iUf := strings.Index(text, string(appconst.Uf))
	iEmissao := strings.Index(text, string(appconst.DataEmissao))
	iNumValidacaoCnh := strings.Index(text, string(appconst.NumValidacaoCnh))
	iNumFormRenach := strings.Index(text, string(appconst.NumFormRenach))

	fmt.Println(text[iCatHab+len("Cat. Hab.") : iNumRegistro])
	toStruct.CatHabilitacao = text[iCatHab+len("Cat. Hab.") : iNumRegistro]

	fmt.Println(text[iNumRegistro+len("N° Registro") : iValidade])
	toStruct.NumRegistro = text[iNumRegistro+len("N° Registro") : iValidade]

	fmt.Println(text[iValidade+len(appconst.Validade) : iPrimHab-len("cao")])
	toStruct.Validade = text[iValidade+len(appconst.Validade) : iPrimHab-len("cao")]

	fmt.Println(text[iPrimHab+len("1a Habilita") : iObs])
	toStruct.PrimHabilitacao = text[iPrimHab+len("1a Habilita") : iObs]

	fmt.Println(text[iObs+len("Observacoess") : iLocal])
	toStruct.Observacoes = text[iObs+len("Observacoess") : iLocal]

	fmt.Println(text[iLocal+len(appconst.Local) : iUf])
	toStruct.Local = text[iLocal+len(appconst.Local) : iUf]

	fmt.Println(text[iUf+len(appconst.Uf) : iEmissao-len("Data de ")])
	toStruct.Uf = text[iUf+len(appconst.Uf) : iEmissao-len("Data de ")]

	fmt.Println(text[iEmissao+len(appconst.DataEmissao) : iNumValidacaoCnh-len("Numero validacao  ")])
	toStruct.DataEmissao = text[iEmissao+len(appconst.DataEmissao) : iNumValidacaoCnh-len("Numero validacao  ")]

	fmt.Println(text[iNumValidacaoCnh+len(appconst.NumValidacaoCnh) : iNumFormRenach-len("Numero Formulario ")])
	toStruct.NumValidacaoCnh = text[iNumValidacaoCnh+len(appconst.NumValidacaoCnh) : iNumFormRenach-len("Numero Formulario ")]

	fmt.Println(text[iNumFormRenach+len(appconst.NumFormRenach) : textLen-len("eee")])
	toStruct.NumFormRenach = text[iNumFormRenach+len(appconst.NumFormRenach) : textLen-len("eee")]
}
