package services

type StatusMsg string

const (
	ErrConvertBase64ToStructMsg StatusMsg = "Erro a converter base64 para struct"
	ErrExceededImgLimitMsg      StatusMsg = "Limite de imagens excedido"
	ErrConvertBase64ToImageMsg  StatusMsg = "Erro ao converter base64 para imagem"
	ErrConvertImageToBytesMsg   StatusMsg = "Erro ao converter imagem para bytes"
	ErrConvertBytesToTextMsg    StatusMsg = "Erro ao converter bytes para texto"
)
