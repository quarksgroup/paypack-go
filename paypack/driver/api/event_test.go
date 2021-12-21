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

// TestEvent tests the event functionality service
func TestEvent(t *testing.T) {
	defer gock.Off()

	gock.New("https://payments.paypack.rw/api").
		Get("/events/transactions").
		Reply(200).
		Type("application/json").
		File("testdata/event.json")
	client := NewDefault()
	got, err := client.Event.List(context.Background(), "limit=10&offset=0&status=failed&kind=cashout")

	require.Nil(t, err, fmt.Sprintf("unexpected error %s", err))

	want := new(paypack.Event)

	raw, _ := ioutil.ReadFile("testdata/event.json.golden")
	_ = json.Unmarshal(raw, want)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected Results")
		t.Log(diff)
	}
}
