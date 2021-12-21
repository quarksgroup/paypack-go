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

var token = &paypack.Token{
	Access:  "YYUWYUSJJKSisudiosJ",
	Refresh: "XXXXXXiuisduidsX",
	Expires: 163973646,
}

// TestLogin tests the login functionality service
func TestLogin(t *testing.T) {
	defer gock.Off()

	gock.New("https://payments.paypack.rw/api").
		Post("/auth/agents/authorize").
		Reply(200).
		Type("application/json").
		File("testdata/token.json")
	client := NewDefault()

	got, err := client.Auth.Login(context.Background(), "client_id", "client_secret")

	require.Nil(t, err, fmt.Sprintf("unexpected error %s", err))

	want := new(paypack.Token)
	raw, _ := ioutil.ReadFile("testdata/token.json.golden")
	_ = json.Unmarshal(raw, want)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected Results")
		t.Log(diff)
	}

}

// TestRefresh tests for the refresh token functionality service
func TestRefersh(t *testing.T) {
	defer gock.Off()

	gock.New("https://payments.paypack.rw/api").
		Get(fmt.Sprintf("/auth/refresh/%s", token.Refresh)).
		Reply(200).
		Type("application/json").
		File("testdata/token.json")
	client := NewDefault()

	got, err := client.Auth.Refresh(context.Background(), token)

	require.Nil(t, err, fmt.Sprintf("unexpected error %s", err))

	want := new(paypack.Token)
	raw, _ := ioutil.ReadFile("testdata/token.json.golden")
	_ = json.Unmarshal(raw, want)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected Results")
		t.Log(diff)
	}
}
