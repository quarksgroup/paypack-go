// Package paypack implements the payment.Client for the (https://payments.paypack.rw/paypack)
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/rehttp"
	"github.com/quarksgroup/paypack-go/paypack"
	"github.com/quarksgroup/paypack-go/paypack/transport/oauth"
)

const (
	baseURL = "https://payments.paypack.rw/api"
	retries = 3
)

type Client struct {
	inner *paypack.Client
}

// New creates a new payment.Client instance backed by the paypack.DriverPaypack
func New(uri string, tr http.RoundTripper) (*Client, error) {
	base, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	if !strings.HasSuffix(base.Path, "/") {
		base.Path = base.Path + "/"
	}

	if tr == nil {
		tr = &oauth.Transport{
			Scheme: oauth.SchemeBearer,
			Source: oauth.ContextTokenSource(),
			Base:   http.DefaultTransport,
		}
	}

	retryTransport := rehttp.NewTransport(
		tr,
		rehttp.RetryAll(
			rehttp.RetryMaxRetries(retries),
			rehttp.RetryAny(
				rehttp.RetryTemporaryErr(),
				rehttp.RetryStatuses(502, 503),
			),
		),
		rehttp.ExpJitterDelay(100*time.Millisecond, 1*time.Second),
	)

	httpClient := &http.Client{
		Transport: retryTransport,
	}

	inner := &paypack.Client{
		Http:    httpClient,
		BaseURL: base,
		Driver:  paypack.DriverPaypack,
	}

	client := new(Client)

	client.inner = inner

	return client, nil
}

// NewDefault returns a new paypack-payments connection for client using the`
// default "https://payments.paypack.rw/api" address.
func NewDefault() *Client {
	client, _ := New(baseURL, nil)
	return client
}

// do wraps the Client.Do function by creating the Request and
// unmarshalling the response.
func (c *Client) do(ctx context.Context, method, path string, in, out interface{}, headers http.Header) (*paypack.Response, error) {
	req := &paypack.Request{
		Method: method,
		Path:   path,
	}

	//Make all http request to be json content-type
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}
	// if we are posting or putting data, we need to
	// write it to the body of the request.
	if in != nil {
		buf := new(bytes.Buffer)
		_ = json.NewEncoder(buf).Encode(in)
		req.Body = buf
	}

	for k, v := range headers {
		req.Header[k] = v
	}

	// execute the http request
	res, err := c.inner.Do(ctx, req)
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
