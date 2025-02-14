# OneBrick API for Golang

A Go implementation of the OneBrick API, providing various payment integration features.

## Features

### Payment Methods

#### Virtual Account
- ✅ Closed VA
- ✅ Closed VA Status
- ✅ Open VA
- ✅ Open VA Status

#### QRIS
- ❌ QRIS Creation
- ❌ QRIS Status

#### Payment Link
- ❌ Payment Link Creation
- ❌ Payment Link Status
- ❌ Payment Link Cancel

#### E-Wallet
- ✅ E-Wallet Integration
- ❌ E-Wallet Status

#### Bank Transfer
- ❌ BCA Unique Code
- ❌ BCA Unique Code Status

### Disbursement
- ❌ Payment Out
- ❌ BI-Fast Payment Out

## Installation

```bash
go get github.com/adexaja/go-onebrick
```

## Usage

```go
package main

import (
    	onebrickapi "github.com/adexaja/go-onebrick/coreapi"
      "github.com/adexaja/go-onebrick"
)

func main() {
    client := new(onebrickapi.Client)
    client.New(os.Getenv("ONEBRICK_CLIENT_ID"), os.Getenv("ONEBRICK_CLIENT_SECRET"), onebrick.Sandbox)

    // Example: Create a Closed VA
  	token, err := client.GetPublicAccessToken()
  	if err != nil {
  		panic(err)
  	}
  
  	paymentRequest := onebrickapi.CloseVARequest{
  		PublicAccessToken: token.Data.AccessToken,
  		PaymentMethodType: "virtual_bank_account",
  		Amount:            15000,
  		Description:       "Test Initial VA Close",
  		ExpiredAt:         60,
  		BankShortCode:     "MANDIRI",
  		ReferenceId:       "brick12368",
  		DisplayName:       "BRICK",
  	}
  	response, err := client.CreateClosedVA(paymentRequest)
  	if err != nil {
  		panic(err)
  	}
  
  	// handle response 

}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)

## Support

For support, please contact me@rezki.dev
