package services

type StatusCode uint64

const (
	ErrConvertBase64ToStruct StatusCode = 1
	ErrExceededImgLimit      StatusCode = 2
	ErrConvertBase64ToImage  StatusCode = 3
	ErrConvertImageToBytes   StatusCode = 4
	ErrConvertBytesToText    StatusCode = 5
)
