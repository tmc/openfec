package openfec_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/tmc/openfec"
)

func TestGetCandidates(t *testing.T) {
	apiKey := os.Getenv("DATA_GOV_API_KEY")
	if apiKey == "" {
		t.Skip("missing DATA_GOV_API_KEY")
	}
	client, err := openfec.NewClient(apiKey)
	if err != nil {
		t.Fatal(err)
	}
	//client.TraceOn(log.New(os.Stderr, "openfec", log.LstdFlags))
	candidates, pagination, err := client.GetCandidates()
	//spew.Dump(candidates, pagination, err)
	_, _ = candidates, pagination
	if err != nil {
		t.Fatal(err)
	}
}

func ExampleNewClient() {
	apiKey := os.Getenv("DATA_GOV_API_KEY")
	client, err := openfec.NewClient(apiKey)
	if err != nil {
		log.Fatalln(client)
	}
	candidates, pagination, err := client.GetCandidates()
	fmt.Printf("err:%v\ntype1:%T\ntype2:%T\n", err, candidates, pagination)
	// output:
	// err:<nil>
	// type1:[]*openfec.Candidate
	// type2:*openfec.Pagination
}

func ExampleClient_GetCandidates() {
	apiKey := os.Getenv("DATA_GOV_API_KEY")
	client, err := openfec.NewClient(apiKey)
	if err != nil {
		log.Fatalln(client)
	}
	candidates, pagination, err := client.GetCandidates()
	fmt.Printf("err:%v\ntype1:%T\ntype2:%T\n", err, candidates, pagination)
	// output:
	// err:<nil>
	// type1:[]*openfec.Candidate
	// type2:*openfec.Pagination
}
