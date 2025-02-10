package coreapi

type TokenRequest struct {
	ClientID     string `json:"username"`
	ClientSecret string `json:"password"`
}
type OpenVARequest struct {
	PublicAccessToken string `json:"-"`
	BankShortCode     string `json:"bankShortCode"`
	ReferenceId       string `json:"referenceId"`
	DisplayName       string `json:"displayName"`
}
type VAStatusRequest struct {
	PublicAccessToken string `json:"-"`
	VaId              string `json:"vaId"`
}

type CloseVARequest struct {
	PublicAccessToken string  `json:"-"`
	PaymentMethodType string  `json:"paymentMethodType"`
	Amount            float64 `json:"amount"`
	Description       string  `json:"description"`
	ExpiredAt         int     `json:"expiredAt"`
	BankShortCode     string  `json:"bankShortCode"`
	ReferenceId       string  `json:"referenceId"`
	DisplayName       string  `json:"displayName"`
}

type QRISRequest struct {
	PublicAccessToken string        `json:"-"`
	ReferenceId       string        `json:"referenceId"`
	Amount            float64       `json:"amount"`
	ExpiredAt         BrickDateTime `json:"expiredAt"`
}

type StatusQRISRequest struct {
	PublicAccessToken string `json:"-"`
	ReferenceId       string `json:"referenceId"`
}
type CancelQRISRequest struct {
	PublicAccessToken string `json:"-"`
	ReferenceId       string `json:"referenceId"`
}

type PaymentLinkRequest struct {
	PublicAccessToken  string   `json:"-"`
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
	PublicAccessToken string        `json:"-"`
	Amount            float64       `json:"amount"`
	ReferenceId       string        `json:"referenceId"`
	EwalletCode       string        `json:"ewalletCode"`
	ExpiryTime        BrickDateTime `json:"expiryTime"`
	ReturnUrl         string        `json:"returnUrl"`
}

type BCAUniqueCodeRequest struct {
	PublicAccessToken string  `json:"-"`
	ReceiveAmount     float64 `json:"receiveAmount"`
}
