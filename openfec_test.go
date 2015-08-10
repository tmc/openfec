package openfec_test

import (
	"fmt"
	"log"
	"os"

	"github.com/tmc/openfec"
)

func ExampleNewClient() {
	apiKey := os.Getenv("DATA_GOV_API_KEY")
	client, err := openfec.NewClient(apiKey)
	if err != nil {
		log.Fatalln(client)
	}
	candidates, err := client.GetCandidates(nil)
	fmt.Printf("err:%v\ntype1:%T\n", err, candidates)
	// output:
	// err:<nil>
	// type1:*openfec.CandidateIter
}

func ExampleClient_GetCandidates() {
	apiKey := os.Getenv("DATA_GOV_API_KEY")
	client, err := openfec.NewClient(apiKey)
	if err != nil {
		log.Fatalln(client)
	}
	query := &openfec.CandidateQuery{
		Sort:   "name",
		Office: []openfec.Office{openfec.President},
		Cycle:  []int{2016},
	}
	candidates, err := client.GetCandidates(query)
	if err != nil {
		log.Fatalln(err)
	}
	for candidates.Next() {
		c := candidates.Value()
		_ = c
		// do something with each candidate here
	}
	fmt.Println(err)
	// output:
	// <nil>
}
