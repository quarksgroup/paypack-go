// Package paypack implements the payment.Client for the (https://payments.paypack.rw/paypack)
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/quarksgroup/paypack-go/paypack"
)

const (
	// baseURL = "https://payments.paypack.rw/api"
	baseURL = "http://payments.paypack.rw/api"
)

// New creates a new payment.Client instance backed by the paypack.DriverPaypack
func New(uri string) (*paypack.Client, error) {
	base, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(base.Path, "/") {
		base.Path = base.Path + "/"
	}

	// report, err := url.Parse(callback)
	// if err != nil {
	// 	return nil, err
	// }

	client := &wrapper{new(paypack.Client)}

	client.BaseURL = base

	// client.ReportURL = report

	client.Driver = paypack.DriverPaypack

	client.Auth = &authService{client}

	client.Merchant = &merchantService{client}

	client.Event = &eventService{client}

	client.Transaction = &transactionService{client}

	return client.Client, nil
}

type wrapper struct {
	*paypack.Client
}

// NewDefault returns a new paypack-payments connection for client using the`
// default "https://payments.paypack.rw/api" address.
func NewDefault() *paypack.Client {
	client, _ := New(baseURL)
	return client
}

// do wraps the Client.Do function by creating the Request and
// unmarshalling the response.
func (c *wrapper) do(ctx context.Context, method, path string, in, out interface{}) (*paypack.Response, error) {
	req := &paypack.Request{
		Method: method,
		Path:   path,
	}

	// if we are posting or putting data, we need to
	// write it to the body of the request.
	if in != nil {
		buf := new(bytes.Buffer)
		_ = json.NewEncoder(buf).Encode(in)
		req.Header = map[string][]string{
			"Content-Type": {"application/json"},
		}
		req.Body = buf
	}

	// execute the http request
	res, err := c.Client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// if an error is encountered, unmarshal and return the
	// error response.
	if res.Status > 299 && res.Status < 499 {
		err := new(Err)
		_ = json.NewDecoder(res.Body).Decode(err)
		return res, &paypack.Error{Code: res.Status, Message: err.Message}
	}
	if res.Status > 499 {
		return res, &paypack.Error{Code: res.Status, Message: "Something went wrong"}
	}
	if out == nil {
		return res, nil
	}

	// if a json response is expected, parse and return
	// the json response.
	return res, json.NewDecoder(res.Body).Decode(out)
}

// Error represents a PayPack error.
type Err struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
