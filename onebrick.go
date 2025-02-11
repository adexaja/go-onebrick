package onebrick

import (
	"context"
	"net/http"
	"time"
)

// EnvironmentType is global config Environment for OneBrick API
type EnvironmentType int8

const (
	_ EnvironmentType = iota

	// Sandbox : represent sandbox environment
	Sandbox

	// Production : represent production environment
	Production

	// libraryVersion : onebrick go library version
	libraryVersion = "v.0.0.1"
)

// Client ID
var ClientID string

// Client Secret
var ClientSecret string

var (
	// Environment default Environment for Onebrick API
	Environment = Sandbox

	// DefaultHttpTimeout default timeout for go HTTP HttpClient
	DefaultHttpTimeout = 80 * time.Second

	// DefaultGoHttpClient default Go HTTP Client for Onebrick HttpClient API
	DefaultGoHttpClient = &http.Client{Timeout: DefaultHttpTimeout}

	// DefaultLoggerLevel logging level that will be used for config globally by Onebrick logger
	DefaultLoggerLevel = &LoggerImplementation{LogLevel: LogError}

	// defaultHttpClientImplementation
	defaultHttpClientImplementation = &HttpClientImplementation{
		HttpClient: DefaultGoHttpClient,
		Logger:     GetDefaultLogger(Environment),
	}
)

// GetDefaultLogger the default logger that the library will use to log errors, debug, and informational messages.
func GetDefaultLogger(env EnvironmentType) LoggerInterface {
	if env == Sandbox {
		return &LoggerImplementation{LogLevel: LogDebug}
	} else {
		return DefaultLoggerLevel
	}
}

// GetHttpClient : get HttpClient implementation
func GetHttpClient(Env EnvironmentType) *HttpClientImplementation {
	return &HttpClientImplementation{
		HttpClient: DefaultGoHttpClient,
		Logger:     GetDefaultLogger(Env),
	}
}

var typeString = map[EnvironmentType]string{
	Sandbox:    "https://sandbox.onebrick.io/v2",
	Production: "https://api.onebrick.io/v2",
}

var bcaUniqueCodeUrl = map[EnvironmentType]string{
	Sandbox:    "https://stgapi.onebrick.io/stg/v2",
	Production: "https://api.onebrick.io/v2",
}

// BaseUrl To get Onebrick Base URL
func (e EnvironmentType) BaseUrl() string {
	for k, v := range typeString {
		if k == e {
			return v
		}
	}
	return "undefined"
}

func (e EnvironmentType) BcaUniqueCodeUrl() string {
	for k, v := range bcaUniqueCodeUrl {
		if k == e {
			return v
		}
	}
	return "undefined"
}

type ConfigOptions struct {
	Ctx context.Context
}

// SetContext : options to change or add Context for each API request
func (o *ConfigOptions) SetContext(ctx context.Context) {
	o.Ctx = ctx
}
