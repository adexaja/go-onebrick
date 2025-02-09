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

func initToken() {
	resp, err := c.GetPublicAccessToken()
	if err != nil {
		fmt.Println(err)
	}

	Token = resp.Data.AccessToken
}

func CreateClosedVA() {
	params := example.ClosedVaParams(Token)
	resp, err := c.CreateClosedVA(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func CreateOpenVA() {
	params := example.OpenVaParams(Token)
	resp, err := c.CreateOpenVA(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func CreateQRIS() {
	params := example.QRISParams(Token)
	resp, err := c.CreateQRIS(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func CreatePaymentLink() {
	params := example.PaymentLinkParams(Token)
	resp, err := c.CreatePaymentLink(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func CreateEwallet() {
	params := example.EwalletParams(Token)
	resp, err := c.CreateEwallet(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func CreateBCAUniqueCode() {
	params := example.BCAUniqueCodeParams(Token)
	resp, err := c.CreateBCAUniqueCode(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func main() {
	c.New(ClientID, ClientSecret, onebrick.Sandbox)
	initToken()

	CreateClosedVA()
	CreateOpenVA()
	CreateQRIS()
	CreatePaymentLink()
	CreateEwallet()
	CreateBCAUniqueCode()
}
