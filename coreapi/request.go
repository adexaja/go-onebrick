package coreapi

import "time"

type TokenRequest struct {
	ClientID     string `json:"username"`
	ClientSecret string `json:"password"`
}
type OpenVARequest struct {
	PublicAccessToken string `json:"publicAccessToken"`
	BankShortCode     string `json:"bankShortCode"`
	ReferenceId       string `json:"referenceId"`
	DisplayName       string `json:"displayName"`
}
type VAStatusRequest struct {
	PublicAccessToken string `json:"publicAccessToken"`
	VaId              string `json:"vaId"`
}

type CloseVARequest struct {
	PublicAccessToken string    `json:"publicAccessToken"`
	PaymentMethodType string    `json:"paymentMethodType"`
	Amount            float64   `json:"amount"`
	Description       string    `json:"description"`
	ExpiredAt         time.Time `json:"expiredAt"`
	BankShortCode     string    `json:"bankShortCode"`
	ReferenceId       string    `json:"referenceId"`
	DisplayName       string    `json:"displayName"`
}
type VACallbackRequest struct {
	ID          string    `json:"id"`
	ReferenceId string    `json:"referenceId"`
	Amount      float64   `json:"amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	PaymentID   string    `json:"paymentId"`
	AccountNo   string    `json:"accountNo"`
	DisplayName string    `json:"displayName"`
}

type QRISRequest struct {
	PublicAccessToken string    `json:"publicAccessToken"`
	ReferenceId       string    `json:"referenceId"`
	Amount            float64   `json:"amount"`
	ExpiredAt         time.Time `json:"expiredAt"`
}

type QRISCallbackRequest struct {
	ID           string    `json:"id"`
	ReferenceId  string    `json:"referenceId"`
	Amount       float64   `json:"amount"`
	CreatedAt    time.Time `json:"createdAt"`
	PaidAt       time.Time `json:"paidAt"`
	Status       string    `json:"status"`
	QrisType     string    `json:"qrisType"`
	Merchant     string    `json:"merchant"`
	MerchantNmId string    `json:"merchantNmId"`
}
type StatusQRISRequest struct {
	PublicAccessToken string `json:"publicAccessToken"`
	ReferenceId       string `json:"referenceId"`
}
type CancelQRISRequest struct {
	PublicAccessToken string `json:"publicAccessToken"`
	ReferenceId       string `json:"referenceId"`
}

type PaymentLinkRequest struct {
	Files              []string `json:"files"`
	ReferenceId        string   `json:"referenceId"`
	Amount             string   `json:"amount"`
	Description        string   `json:"description"`
	EndUserPhoneNumber string   `json:"endUserPhoneNumber"`
	EndUserEmail       string   `json:"endUserEmail"`
	EndUserAddress     string   `json:"endUserAddress"`
	PIN                string   `json:"pin"`
	RedirectURL        string   `json:"redirectURL"`
}

type EWalletRequest struct {
	Amount      float64   `json:"amount"`
	ReferenceId string    `json:"referenceId"`
	EwalletCode string    `json:"ewalletCode"`
	ExpiryTime  time.Time `json:"expiryTime"`
	ReturnUrl   string    `json:"returnUrl"`
}

type EwalletCallbackRequest struct {
	ID          string    `json:"id"`
	ReferenceId string    `json:"referenceId"`
	Amount      float64   `json:"amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	PaidAt      time.Time `json:"paidAt"`
}

type BCAUniqueCodeRequest struct {
	TransactionId string  `json:"transactionid"`
	BankShortCode string  `json:"bankShortCode"`
	DisplayName   string  `json:"displayName"`
	Amount        float64 `json:"amount"`
	UniqueCode    float64 `json:"uniqueCode"`
	AmountUnique  float64 `json:"amountUnique"`
}
