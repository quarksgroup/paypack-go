package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/quarksgroup/paypack-go/paypack"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

// TestMe is a test function for the api package
func TestMe(t *testing.T) {
	defer gock.Off()

	gock.New("https://payments.paypack.rw/api").
		Get("/merchants/me").
		Reply(200).
		Type("application/json").
		File("testdata/merchant.json")
	cli := NewDefault()

	got, err := cli.Profile(context.Background())

	require.Nil(t, err, fmt.Sprintf("unexpected error %s", err))

	want := new(paypack.Merchant)

	raw, _ := ioutil.ReadFile("testdata/merchant.json.golden")

	err = json.Unmarshal(raw, want)
	if err != nil {
		t.Log(err)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected Results")
		t.Log(diff)
	}

}

// TestChekout is a test function for testing checkout find endpoint
func TestChekout(t *testing.T) {
	defer gock.Off()

	gock.New("https://payments.paypack.rw/api").
		Get("/checkouts/find/").
		Reply(200).
		Type("application/json").
		File("testdata/checkout.json")
	cli := NewDefault()

	got, err := cli.Checkout(context.Background(), "xxxxxx", "xxxxxx")

	require.Nil(t, err, fmt.Sprintf("unexpected error %s", err))

	want := new(paypack.Checkout)

	raw, _ := ioutil.ReadFile("testdata/checkout.json.golden")

	err = json.Unmarshal(raw, want)
	if err != nil {
		t.Log(err)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected Results")
		t.Log(diff)
	}

}
