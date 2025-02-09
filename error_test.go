package onebrick

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorStruct(t *testing.T) {
	err := &Error{
		Message:        "Error Test Message",
		StatusCode:     200,
		RawError:       errors.New("TEST FROM GO ERROR"),
		RawApiResponse: nil,
	}
	var midError *Error
	assert.Error(t, err)
	assert.True(t, true, errors.Is(err, err))
	assert.True(t, true, errors.As(err, &midError))
	assert.Equal(t, `Error Test Message`, err.GetMessage())
	assert.Equal(t, 200, err.GetStatusCode())
	assert.Equal(t, "Error Test Message: TEST FROM GO ERROR", err.Error())
}

func TestErrorResponse(t *testing.T) {
	c := GetHttpClient(Environment)
	jsonReq, _ := json.Marshal("{\"amount\": 15000, \"referenceId\": \"brick12345\", \"description\": \"Test Initial VA Close\", \"bankShortCode\": \"MANDIRI\", \"displayName\": \"BRICK\"}")

	err := c.Call(http.MethodPost, "https://sandbox.onebrick.io/v2/payments/gs/va/close", "", nil, bytes.NewBuffer(jsonReq), nil)

	var midError *Error
	assert.True(t, true, errors.Is(err, err))
	assert.True(t, true, errors.As(err, &midError))
	assert.Error(t, err)
	assert.Equal(t, 401, err.StatusCode)
	assert.Equal(t, "sandbox.onebrick.io", err.RawApiResponse.Request.Host)
}
