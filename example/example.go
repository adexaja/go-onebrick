package example

import (
	"fmt"

	"github.com/adexaja/go-onebrick"
	"github.com/adexaja/go-onebrick/coreapi"
)

var (
	c            coreapi.Client
	ClientID     string = "client_id"
	ClientSecret string = "client_secret"
	Token        string = "token"
)

func init() {
	c.New(ClientID, ClientSecret, onebrick.Sandbox)
	initToken()
}

func initToken() {
	resp, err := c.GetPublicAccessToken()
	if err != nil {
		fmt.Println(err)
	}

	Token = resp.Data.AccessToken
}

func CreateClosedVA() {
	params := ClosedVaParams(Token)
	resp, err := c.CreateClosedVA(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func CreateOpenVA() {
	params := OpenVaParams(Token)
	resp, err := c.CreateOpenVA(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func CreateQRIS() {
	params := QRISParams(Token)
	resp, err := c.CreateQRIS(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
