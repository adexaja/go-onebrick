package coreapi

import "github.com/adexaja/go-onebrick"

// Client : CoreAPI Client struct
type Client struct {
	ClientID     string
	ClientSecret string
	Env          onebrick.EnvironmentType
	HttpClient   onebrick.HttpClient
	Options      *onebrick.ConfigOptions
}

// New : this function will always be called when the CoreApi is initiated
func (c *Client) New(clientID string, clientSecret string, env onebrick.EnvironmentType) {
	c.Env = env
	c.ClientID = clientID
	c.ClientSecret = clientSecret
	c.Options = &onebrick.ConfigOptions{}
	c.HttpClient = onebrick.GetHttpClient(env)
}

// getDefaultClient : internal function to get default Client
func getDefaultClient() *Client {
	return &Client{
		ClientID:     onebrick.ClientID,
		ClientSecret: onebrick.ClientSecret,
		Env:          onebrick.Environment,
		HttpClient:   onebrick.GetHttpClient(onebrick.Environment),
	}
}
