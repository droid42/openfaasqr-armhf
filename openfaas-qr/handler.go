package function

import (
	"net/http"

	"github.com/openfaas-incubator/go-function-sdk"
	"github.com/skip2/go-qrcode"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	var png []byte
	png, err = qrcode.Encode(string(req.Body), qrcode.Medium, 256)
	if err != nil {
		return handler.Response{
			Body:       []byte(err.Error()),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	var header = http.Header{}
	header.Add("Content-Type", "image/png")

	return handler.Response{
		Header:     header,
		Body:       png,
		StatusCode: http.StatusOK,
	}, err

	//message := fmt.Sprintf("Hello world, input was: %s", string(req.Body))
	//return handler.Response{
	//	Body:       []byte(message),
	//	StatusCode: http.StatusOK,
	//}, err
}
