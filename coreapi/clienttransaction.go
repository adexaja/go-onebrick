package coreapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adexaja/go-onebrick"
)

// To get Public Access Token from clientID and ClientSecret
func (c *Client) GetPublicAccessToken() (*Response[TokenResponseData], *onebrick.Error) {
	resp := &Response[TokenResponseData]{}
	tokenRequest := TokenRequest{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
	}
	tokenRequestJSON, _ := json.Marshal(tokenRequest)
	err := c.HttpClient.Call(
		http.MethodGet,
		fmt.Sprintf("%s/payments/auth/token", c.Env.BaseUrl()),
		"",
		c.Options,
		bytes.NewBuffer(tokenRequestJSON),
		resp,
	)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// export GetPublicAccessToken with Default Client
func GetPublicAccessToken() (*Response[TokenResponseData], *onebrick.Error) {
	return getDefaultClient().GetPublicAccessToken()
}

// DO /payments/gs/va/close API
func (c *Client) CreateClosedVA(data CloseVARequest) (*Response[CloseVAResponseData], *onebrick.Error) {
	resp := &Response[CloseVAResponseData]{}
	jsonData, _ := json.Marshal(data)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/payments/gs/va/close", c.Env.BaseUrl()),
		data.PublicAccessToken,
		c.Options,
		bytes.NewBuffer(jsonData),
		resp,
	)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Do /payments/gs/va/close API
func CreateClosedVA(data CloseVARequest) (*Response[CloseVAResponseData], *onebrick.Error) {
	return getDefaultClient().CreateClosedVA(data)
}

// Do /payments/gs/va/status API
func (c *Client) CreateClosedVAStatus(data VAStatusRequest) (*Response[CloseVAStatusResponseData], *onebrick.Error) {
	resp := &Response[CloseVAStatusResponseData]{}
	jsonData, _ := json.Marshal(data)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/payments/gs/va/status", c.Env.BaseUrl()),
		data.PublicAccessToken,
		c.Options,
		bytes.NewBuffer(jsonData),
		resp,
	)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Do /payments/gs/va/status API
func CreateClosedVAStatus(data VAStatusRequest) (*Response[CloseVAStatusResponseData], *onebrick.Error) {
	return getDefaultClient().CreateClosedVAStatus(data)
}

// DO /payments/gs/va/open API
func (c *Client) CreateOpenVA(data OpenVARequest) (*Response[OpenVAResponseData], *onebrick.Error) {
	resp := &Response[OpenVAResponseData]{}
	jsonData, _ := json.Marshal(data)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/payments/gs/va/open", c.Env.BaseUrl()),
		data.PublicAccessToken,
		c.Options,
		bytes.NewBuffer(jsonData),
		resp,
	)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Do /payments/gs/va/open API
func CreateOpenVA(data OpenVARequest) (*Response[OpenVAResponseData], *onebrick.Error) {
	return getDefaultClient().CreateOpenVA(data)
}

// Do /payments/gs/va/status API
func (c *Client) CreateOpenVAStatus(data VAStatusRequest) (*Response[OpenVAStatusResponseData], *onebrick.Error) {
	resp := &Response[OpenVAStatusResponseData]{}
	jsonData, _ := json.Marshal(data)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/payments/gs/va/status", c.Env.BaseUrl()),
		data.PublicAccessToken,
		c.Options,
		bytes.NewBuffer(jsonData),
		resp,
	)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Do /payments/gs/va/status API
func CreateOpenVAStatus(data VAStatusRequest) (*Response[OpenVAStatusResponseData], *onebrick.Error) {
	return getDefaultClient().CreateOpenVAStatus(data)
}

// Do /payments/gs/qris API
func (c *Client) CreateQRIS(data QRISRequest) (*Response[DynamicQRISResponseData], *onebrick.Error) {
	resp := &Response[DynamicQRISResponseData]{}
	jsonData, _ := json.Marshal(data)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/payments/gs/qris", c.Env.BaseUrl()),
		data.PublicAccessToken,
		c.Options,
		bytes.NewBuffer(jsonData),
		resp,
	)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Do /payments/gs/qris API
func CreateQRIS(data QRISRequest) (*Response[DynamicQRISResponseData], *onebrick.Error) {
	return getDefaultClient().CreateQRIS(data)
}

// Do /payments/gs/qris/status API
func (c *Client) CreateQRISStatus(data QRISCallbackRequest, publicAccessToken string) (*Response[DynamicQRISStatusResponseData], *onebrick.Error) {
	resp := &Response[DynamicQRISStatusResponseData]{}
	jsonData, _ := json.Marshal(data)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/payments/gs/qris/status", c.Env.BaseUrl()),
		publicAccessToken,
		c.Options,
		bytes.NewBuffer(jsonData),
		resp,
	)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Do /payments/gs/qris/status API
func CreateQRISStatus(data QRISCallbackRequest, publicAccessToken string) (*Response[DynamicQRISStatusResponseData], *onebrick.Error) {
	return getDefaultClient().CreateQRISStatus(data, publicAccessToken)
}

// Do /payments/gs/qris/cancel API
func (c *Client) CreateCancelQRIS(data CancelQRISRequest) (*Response[DynamicQRISCancelResponseData], *onebrick.Error) {
	resp := &Response[DynamicQRISCancelResponseData]{}
	jsonData, _ := json.Marshal(data)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/payments/gs/qris/cancel", c.Env.BaseUrl()),
		data.PublicAccessToken,
		c.Options,
		bytes.NewBuffer(jsonData),
		resp,
	)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Do /payments/gs/qris/cancel API
func CreateCancelQRIS(data CancelQRISRequest) (*Response[DynamicQRISCancelResponseData], *onebrick.Error) {
	return getDefaultClient().CreateCancelQRIS(data)
}

// Do /payments/gs/link API
func (c *Client) CreatePaymentLink(data PaymentLinkRequest, publicAccessToken string) (*Response[PaymentLinkResponseData], *onebrick.Error) {
	resp := &Response[PaymentLinkResponseData]{}
	jsonData, _ := json.Marshal(data)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/payments/gs/link", c.Env.BaseUrl()),
		publicAccessToken,
		c.Options,
		bytes.NewBuffer(jsonData),
		resp,
	)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Do /payments/gs/link API
func CreatePaymentLink(data PaymentLinkRequest, publicAccessToken string) (*Response[PaymentLinkResponseData], *onebrick.Error) {
	return getDefaultClient().CreatePaymentLink(data, publicAccessToken)
}

// Do /payments/gs/ewallet API
func (c *Client) CreateEwallet(data EWalletRequest, publicAccessToken string) (*Response[EwalletResponseData], *onebrick.Error) {
	resp := &Response[EwalletResponseData]{}
	jsonData, _ := json.Marshal(data)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/payments/gs/ewallet", c.Env.BaseUrl()),
		publicAccessToken,
		c.Options,
		bytes.NewBuffer(jsonData),
		resp,
	)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Do /payments/gs/ewallet API
func CreateEwallet(data EWalletRequest, publicAccessToken string) (*Response[EwalletResponseData], *onebrick.Error) {
	return getDefaultClient().CreateEwallet(data, publicAccessToken)
}

// Do /payments/gs/bca/unique-code API
func (c *Client) CreateBCAUniqueCode(data BCAUniqueCodeRequest, publicAccessToken string) (*Response[BcaUqniqueCodeResponseData], *onebrick.Error) {
	resp := &Response[BcaUqniqueCodeResponseData]{}
	jsonData, _ := json.Marshal(data)
	err := c.HttpClient.Call(
		http.MethodPost,
		fmt.Sprintf("%s/payments/gs/bca/unique-code", c.Env.BaseUrl()),
		publicAccessToken,
		c.Options,
		bytes.NewBuffer(jsonData),
		resp,
	)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Do /payments/gs/bca/unique-code API
func CreateBCAUniqueCode(data BCAUniqueCodeRequest, publicAccessToken string) (*Response[BcaUqniqueCodeResponseData], *onebrick.Error) {
	return getDefaultClient().CreateBCAUniqueCode(data, publicAccessToken)
}
