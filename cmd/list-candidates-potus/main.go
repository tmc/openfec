package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/tmc/openfec"
)

var (
	verbose = flag.Bool("v", false, "verbose output")
	year    = flag.Int("year", 2016, "Election cycle to list candidates from")
	format  = flag.String("f", "{{.CandidateID}} - {{.Name}} {{.Party}}", "Formatting string")
	party   = flag.String("party", "", "Political party (default: all)")
)

func main() {
	apiKey := os.Getenv("DATA_GOV_API_KEY")
	if apiKey == "" {
		apiKey = "DEMO_KEY"
	}
	flag.Parse()
	client, err := openfec.NewClient(apiKey)
	if err != nil {
		log.Fatalln(client)
	}
	if *verbose {
		client.TraceOn(log.New(os.Stderr, "openfec: ", log.LstdFlags))
	}
	query := &openfec.CandidateQuery{
		Sort:            "name",
		Office:          []openfec.Office{openfec.President},
		CandidateStatus: []openfec.CandidateStatus{openfec.PresentCandidate},
		Cycle:           []int{*year},
		Party:           *party,
	}
	candidates, err := client.GetCandidates(query)
	if err != nil {
		if err == openfec.ErrUnauthorized {
			fmt.Println("Authorization failure. Check the value of the 'DATA_GOV_API_KEY' environment variable.")
			os.Exit(1)
		} else {
			log.Fatalln(err)
		}
	}
	tmpl, err := template.New("format string").Parse(*format)
	if err != nil {
		log.Fatalln(err)
	}
	for candidates.Next() {
		value := candidates.Value()
		if *format == "json" {
			buf, _ := json.Marshal(value)
			fmt.Println(string(buf))
		} else {
			tmpl.Execute(os.Stdout, value)
			fmt.Println()
		}
	}
	if candidates.Err() != nil {
		fmt.Println("Issue iterating candidates:", candidates.Err())
	}
}
