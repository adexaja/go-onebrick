package example

import (
	"time"

	"github.com/adexaja/go-onebrick/coreapi"
)

func ClosedVaParams(token string) coreapi.CloseVARequest {
	return coreapi.CloseVARequest{
		PublicAccessToken: token,
		PaymentMethodType: "BCA",
		Amount:            15000,
		Description:       "Test Initial VA Close",
		ExpiredAt:         time.Now().Add(time.Hour * 24 * 7),
		BankShortCode:     "BCA",
		ReferenceId:       "brick12345",
		DisplayName:       "BRICK",
	}
}

func OpenVaParams(token string) coreapi.OpenVARequest {
	return coreapi.OpenVARequest{
		PublicAccessToken: token,
		BankShortCode:     "BCA",
		ReferenceId:       "brick12345",
		DisplayName:       "BRICK",
	}
}

func QRISParams(token string) coreapi.QRISRequest {
	return coreapi.QRISRequest{
		PublicAccessToken: token,
		ReferenceId:       "brick12345",
		Amount:            15000,
		ExpiredAt:         time.Now().Add(time.Hour * 24 * 7),
	}
}
