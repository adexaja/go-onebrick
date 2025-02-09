package coreapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adexaja/go-onebrick"
)

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
