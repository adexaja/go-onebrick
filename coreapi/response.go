package coreapi

import (
	"time"
)

type StatusEnum string

var (
	StatusPending   StatusEnum = "Pending"
	StatusPaid      StatusEnum = "Paid"
	StatusCompleted StatusEnum = "Completed"
	StatusExpired   StatusEnum = "Expired"
)

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

type OpenVAResponseData struct {
	ID            string `json:"id"`
	BankShortCode string `json:"bankShortCode"`
	ReferenceId   string `json:"referenceId"`
	DisplayName   string `json:"displayName"`
	AccountNo     string `json:"accountNo"`
}

type OpenVAStatusResponseData struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	BankShortCode string `json:"bankShortCode"`
	ReferenceId   string `json:"referenceId"`
	DisplayName   string `json:"displayName"`
	AccountNo     string `json:"accountNo"`
}

type CloseVAResponseData struct {
	ID            string  `json:"id"`
	Type          string  `json:"type"`
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
	CreatedAt     int     `json:"createdAt"`
	ExpiredAt     int     `json:"expiredAt"`
	Status        string  `json:"status"`
	BankShortCode string  `json:"bankShortCode"`
	ReferenceId   string  `json:"referenceId"`
	DisplayName   string  `json:"displayName"`
	AccountNo     string  `json:"accountNo"`
}
type CloseVAStatusResponseData struct {
	ID            string     `json:"id"`
	Type          string     `json:"type"`
	Amount        float64    `json:"amount"`
	Description   string     `json:"description"`
	CreatedAt     int        `json:"createdAt"`
	ExpiredAt     int        `json:"expiredAt"`
	BankShortCode string     `json:"bankShortCode"`
	ReferenceId   string     `json:"referenceId"`
	DisplayName   string     `json:"displayName"`
	AccountNo     string     `json:"accountNo"`
	Status        StatusEnum `json:"status"`
}

type (
	DynamicQRISResponseData struct{}
	StaticQRISResponseData  struct{}
	PaymentLinkResponseData struct{}
)
