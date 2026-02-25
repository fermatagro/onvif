package networking

import (
	"bytes"
	"context"
	"net/http"

	"github.com/juju/errors"
)

// SendSoap send soap message
// Deprecated: use SendSoapContext instead
func SendSoap(httpClient *http.Client, endpoint, message string) (*http.Response, error) {
	return SendSoapContext(context.Background(), httpClient, endpoint, message)
}

func SendSoapContext(ctx context.Context, httpClient *http.Client, endpoint, message string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBufferString(message))
	if err != nil {
		return nil, errors.Annotate(err, "NewRequestWithContext")
	}
	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")

	resp, err := httpClient.Do(req)
	if err != nil {
		return resp, errors.Annotate(err, "Do")
	}

	return resp, nil
}
