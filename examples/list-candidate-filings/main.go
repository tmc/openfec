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
	verbose   = flag.Bool("v", false, "verbose output")
	candidate = flag.String("candidate", "", "Candidate ID to fetch filings for")
	format    = flag.String("f", "{{.}}", "Formatting string")
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
	filings, err := client.GetCandidateFilings(*candidate)
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
	for filings.Next() {
		value := filings.Value()
		if *format == "json" {
			buf, _ := json.Marshal(value)
			fmt.Println(string(buf))
		} else {
			tmpl.Execute(os.Stdout, value)
			fmt.Println()
		}
	}
	if filings.Err() != nil {
		fmt.Println("Issue iterating filings:", filings.Err())
	}
}
