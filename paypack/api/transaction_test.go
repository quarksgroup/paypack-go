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

//TestCashin - Test Cashin transaction with Paypack API
func TestCashin(t *testing.T) {
	defer gock.Off()

	gock.New("https://payments.paypack.rw/api").
		Post("/transactions/cashin").
		Reply(200).
		Type("application/json").
		File("testdata/cashin.json")
	client := NewDefault()

	in := &paypack.TransactionRequest{
		Amount: 100,
		Number: "07898989898",
	}

	got, err := client.Cashin(context.Background(), in)

	require.Nil(t, err, fmt.Sprintf("unexpected error %s", err))

	want := new(paypack.TransactionResponse)

	raw, _ := ioutil.ReadFile("testdata/cashin.json.golden")

	err = json.Unmarshal(raw, want)

	if err != nil {
		t.Log(err)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected Results")
		t.Log(diff)
	}
}

//TestCashout - Test Cashout transaction with Paypack API
func TestCashout(t *testing.T) {
	defer gock.Off()

	gock.New("https://payments.paypack.rw/api").
		Post("/transactions/cashout").
		Reply(200).
		Type("application/json").
		File("testdata/cashout.json")
	client := NewDefault()

	in := &paypack.TransactionRequest{
		Amount: 100,
		Number: "07898989898",
	}

	got, err := client.Cashout(context.Background(), in)

	require.Nil(t, err, fmt.Sprintf("unexpected error %s", err))
	want := new(paypack.TransactionResponse)
	raw, _ := ioutil.ReadFile("testdata/cashout.json.golden")

	err = json.Unmarshal(raw, want)

	if err != nil {
		t.Log(err)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected Results")
		t.Log(diff)
	}
}

// TestFind - Test Find transaction with Paypack API using transaction reference
func TestFind(t *testing.T) {
	defer gock.Off()
	gock.New("https://payments.paypack.rw/api").
		Get("transactions/find/").
		Reply(200).
		Type("application/json").
		File("testdata/find.json")
	client := NewDefault()

	got, err := client.FindTx(context.Background(), "xxxxx")

	require.Nil(t, err, fmt.Sprintf("unexpected error %s", err))

	want := new(paypack.Transaction)

	raw, _ := ioutil.ReadFile("testdata/find.json.golden")

	err = json.Unmarshal(raw, want)

	if err != nil {
		t.Log(err)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected Results")
		t.Log(diff)
	}
}
