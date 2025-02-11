package main

import (
	"fmt"

	"github.com/adexaja/go-onebrick"
	"github.com/adexaja/go-onebrick/coreapi"
	"github.com/adexaja/go-onebrick/example"
)

var (
	c            coreapi.Client
	ClientID     string = "client_id"
	ClientSecret string = "client_secret"
	Token        string = "token"
)

func initToken() string {
	resp, err := c.GetPublicAccessToken()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return resp.Data.AccessToken
}

func CreateClosedVA(Token string) {
	params := example.ClosedVaParams(Token)
	resp, err := c.CreateClosedVA(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Data)
}

func CreateOpenVA(Token string) {
	params := example.OpenVaParams(Token)
	resp, err := c.CreateOpenVA(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Data)
}

func CreateQRIS(Token string) {
	params := example.QRISParams(Token)
	resp, err := c.CreateQRIS(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Data)
}

func CreatePaymentLink(Token string) {
	params := example.PaymentLinkParams(Token)
	resp, err := c.CreatePaymentLink(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Data)
}

func CreateEwallet(Token string) {
	params := example.EwalletParams(Token)
	resp, err := c.CreateEwallet(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Data)
}

func CreateBCAUniqueCode(Token string) {
	params := example.BCAUniqueCodeParams(Token)
	resp, err := c.CreateBCAUniqueCode(params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Data)
}

func main() {
	c.New(ClientID, ClientSecret, onebrick.Sandbox)
	token := initToken()

	// CreateClosedVA(token)
	// CreateOpenVA(token)
	// CreateQRIS(token)
	// CreatePaymentLink(token)
	// CreateEwallet(token)
	CreateBCAUniqueCode(token)
}
