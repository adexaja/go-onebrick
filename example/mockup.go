package example

import (
	"time"

	"github.com/adexaja/go-onebrick/coreapi"
)

func ClosedVaParams(token string) coreapi.CloseVARequest {
	return coreapi.CloseVARequest{
		PublicAccessToken: token,
		PaymentMethodType: "virtual_bank_account",
		Amount:            15000,
		Description:       "Test Initial VA Close",
		ExpiredAt:         60,
		BankShortCode:     "MANDIRI",
		ReferenceId:       "brick12368",
		DisplayName:       "BRICK",
	}
}

func OpenVaParams(token string) coreapi.OpenVARequest {
	return coreapi.OpenVARequest{
		PublicAccessToken: token,
		BankShortCode:     "MANDIRI",
		ReferenceId:       "brick123453",
		DisplayName:       "BRICK",
	}
}

func QRISParams(token string) coreapi.QRISRequest {
	return coreapi.QRISRequest{
		PublicAccessToken: token,
		ReferenceId:       "brick12345",
		Amount:            15000,
		ExpiredAt:         coreapi.NewBrickDateTime(time.Now().Add(time.Hour * 24 * 7)),
	}
}

func PaymentLinkParams(token string) coreapi.PaymentLinkRequest {
	return coreapi.PaymentLinkRequest{
		PublicAccessToken:  token,
		Files:              []string{"https://lh7-rt.googleusercontent.com/docsz/AD_4nXfEQoAs79JAwhLjU8XFv6OsxoKhvSnR_FBOx_lwFX85jPJhsIYE9mtuHlYJ69UAogyrl-GtRe03m8i2Ws0b9Y49cIKQ7S4sB6bhJJ3E7gaRz3v8fSC-HB2ipmA6nJyVEg?key=uUDkSXfgqXW1dzwVRqn5LNm4"},
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
		ReferenceId:       "brick12347",
		EwalletCode:       "DANA",
		ExpiryTime:        60,
		ReturnUrl:         "https://www.brick.com",
	}
}

func BCAUniqueCodeParams(token string) coreapi.BCAUniqueCodeRequest {
	return coreapi.BCAUniqueCodeRequest{
		PublicAccessToken: token,
		ReceiveAmount:     15000,
	}
}
