package openfec_test

import (
	"log"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/tmc/openfec"
)

func TestGetCandidates(t *testing.T) {
	apiKey := os.Getenv("DATA_GOV_API_KEY")
	if apiKey == "" {
		t.Skip("missing DATA_GOV_API_KEY")
	}
	client, err := openfec.NewClient(apiKey)
	client.TraceOn(log.New(os.Stderr, "openfec", log.LstdFlags))
	if err != nil {
		t.Fatal(err)
	}
	candidates, err := client.GetCandidates()
	spew.Dump(candidates, err)
}
