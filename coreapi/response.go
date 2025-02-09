package coreapi

import "time"

type MetaData struct {
	Source string `json:"source"`
	Entity string `json:"entity"`
}
type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Action  string `json:"action"`
	Reason  string `json:"reason"`
}

type Response[T any] struct {
	Status   int            `json:"status"`
	Error    *ResponseError `json:"error"`
	MetaData *MetaData      `json:"metaData"`
	Data     T              `json:"data"`
}

type TokenResponseData struct {
	Message     string    `json:"message"`
	AccessToken string    `json:"accessToken"`
	IssuedAt    time.Time `json:"issuedAt"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

type (
	OpenVAResponseData  struct{}
	CloseVAResponseData struct{}
)

type (
	DynamicQRISResponseData struct{}
	StaticQRISResponseData  struct{}
	PaymentLinkResponseData struct{}
)
