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
	client := NewDefault()

	got, err := client.Merchant.Me(context.Background())

	require.Nil(t, err, fmt.Sprintf("unexpected error %s", err))

	want := new(paypack.Merchant)

	raw, _ := ioutil.ReadFile("testdata/merchant.json.golden")
	_ = json.Unmarshal(raw, want)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected Results")
		t.Log(diff)
	}

}
