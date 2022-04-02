package vioerror

import "vio-back/services"

type ResponseError interface {
	GetCode() services.StatusCode
	GetMsg() services.StatusMsg
	GetCause() error
}

type vioErrorImpl struct {
	StatusCode services.StatusCode `json:"status_code"`
	StatusMsg  services.StatusMsg  `json:"status_msg"`
	CauseErr   error               `json:"cause_err"`
}

func NewError(statusCode services.StatusCode, statusMsg services.StatusMsg, err error) ResponseError {
	return &vioErrorImpl{StatusCode: statusCode, StatusMsg: statusMsg, CauseErr: err}
}

func (p vioErrorImpl) GetCode() services.StatusCode {
	return p.StatusCode
}
func (p vioErrorImpl) GetMsg() services.StatusMsg {
	return p.StatusMsg
}
func (p vioErrorImpl) GetCause() error {
	return p.CauseErr
}
