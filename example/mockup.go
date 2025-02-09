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

func PaymentLinkParams(token string) coreapi.PaymentLinkRequest {
	return coreapi.PaymentLinkRequest{
		PublicAccessToken:  token,
		Files:              nil,
		ReferenceId:        "brick12345",
		Amount:             "15000",
		Description:        "Test Initial VA Close",
		EndUserPhoneNumber: "08123456789",
		EndUserEmail:       "brick@brick.com",
		EndUserAddress:     "Jl. Raya Postar No. 123, Jakarta",
		PIN:                "123456",
		RedirectURL:        "https://www.brick.com",
	}
}

func EwalletParams(token string) coreapi.EWalletRequest {
	return coreapi.EWalletRequest{
		PublicAccessToken: token,
		Amount:            15000,
		ReferenceId:       "brick12345",
		EwalletCode:       "123456",
		ExpiryTime:        time.Now().Add(time.Hour * 24 * 7),
		ReturnUrl:         "https://www.brick.com",
	}
}

func BCAUniqueCodeParams(token string) coreapi.BCAUniqueCodeRequest {
	return coreapi.BCAUniqueCodeRequest{
		PublicAccessToken: token,
		ReceiveAmount:     15000,
	}
}
