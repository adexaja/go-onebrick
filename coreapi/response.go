package coreapi

import (
	"encoding/json"
	"fmt"
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
	Status   int             `json:"status"`
	Error    json.RawMessage `json:"error"`
	MetaData *MetaData       `json:"metaData"`
	Data     *T              `json:"data"`
}

// GetError returns either string error or ResponseError
func (r *Response[T]) GetError() (*ResponseError, string, error) {
	if r.Error == nil {
		return nil, "", nil
	}

	// Try to unmarshal as string first
	var strError string
	if err := json.Unmarshal(r.Error, &strError); err == nil {
		return nil, strError, nil
	}

	// If not string, try to unmarshal as ResponseError
	var structError ResponseError
	if err := json.Unmarshal(r.Error, &structError); err != nil {
		return nil, "", fmt.Errorf("failed to parse error: %w", err)
	}

	return &structError, "", nil
}

type TokenResponseData struct {
	Message     string        `json:"message"`
	AccessToken string        `json:"accessToken"`
	IssuedAt    BrickDateTime `json:"issuedAt"`
	ExpiresAt   BrickDateTime `json:"expiresAt"`
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
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Amount        float64       `json:"amount"`
	Description   string        `json:"description"`
	CreatedAt     BrickDateTime `json:"createdAt"`
	ExpiredAt     BrickDateTime `json:"expiredAt"`
	Status        string        `json:"status"`
	BankShortCode string        `json:"bankShortCode"`
	ReferenceId   string        `json:"referenceId"`
	DisplayName   string        `json:"displayName"`
	AccountNo     string        `json:"accountNo"`
}
type CloseVAStatusResponseData struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Amount        float64       `json:"amount"`
	Description   string        `json:"description"`
	CreatedAt     BrickDateTime `json:"createdAt"`
	ExpiredAt     BrickDateTime `json:"expiredAt"`
	BankShortCode string        `json:"bankShortCode"`
	ReferenceId   string        `json:"referenceId"`
	DisplayName   string        `json:"displayName"`
	AccountNo     string        `json:"accountNo"`
	Status        StatusEnum    `json:"status"`
}
type DynamicQRISResponseData struct {
	ID          string        `json:"id"`
	ReferenceId string        `json:"referenceId"`
	Amount      float64       `json:"amount"`
	CreatedAt   BrickDateTime `json:"createdAt"`
	ExpiredAt   BrickDateTime `json:"expiredAt"`
	QRData      string        `json:"qrData"`
}
type DynamicQRISStatusResponseData struct {
	ID          string        `json:"id"`
	ReferenceId string        `json:"referenceId"`
	Amount      float64       `json:"amount"`
	Status      string        `json:"status"`
	QRData      string        `json:"qrData"`
	PaidAt      BrickDateTime `json:"paidAt"`
}
type DynamicQRISCancelResponseData struct {
	Message     string        `json:"message"`
	ID          string        `json:"id"`
	ReferenceId string        `json:"referenceId"`
	CancelledAt BrickDateTime `json:"cancelledAt"`
}

type PaymentLinkResponseData struct {
	PaymentLinkPath string        `json:"paymentLinkPath"`
	ExpiredAt       BrickDateTime `json:"expiredAt"`
	Amount          float64       `json:"amount"`
	ReferenceId     string        `json:"referenceId"`
	Status          string        `json:"status"`
}
type BcaUniqueCodeResponseData struct {
	Message       string  `json:"message"`
	TransactionId string  `json:"transactionId"`
	BankShortCode string  `json:"bankShortCode"`
	AccountNo     string  `json:"accountNo"`
	DisplayName   string  `json:"displayName"`
	Amount        float64 `json:"amount"`
	UniqueCode    float64 `json:"uniqueCode"`
	AmountUnique  float64 `json:"amountUnique"`
}
type EwalletResponseData struct {
	Message     string    `json:"message"`
	ID          string    `json:"id"`
	ReferenceId string    `json:"referenceId"`
	Amount      Amount    `json:"amount"`
	Status      string    `json:"status"`
	EwalletCode string    `json:"ewalletCode"`
	RedirectURL string    `json:"redirectUrl"`
	ExpiredAt   time.Time `json:"expiredAt"`
	CreatedAt   time.Time `json:"createdAt"`
}
