openfec
=======

[![GoDoc](https://godoc.org/github.com/tmc/openfec?status.svg)](http://godoc.org/github.com/tmc/openfec)


This package provides programmatic access to the OpenFEC api.

Included under cmd are example command line utilities showing consumption of this api.

Documentation: http://godoc.org/github.com/tmc/openfec

Status: some endpoints missing, pull requests welcome!

License: ISC

Examples: 

list-candidates-potus
---------------------

```sh
⚛ ~$ go get github.com/tmc/openfec/cmd/...
⚛ ~$ 
⚛ ~$ list-candidates-potus -h
Usage of list-candidates-potus:
  -f string
	  Formatting string (default "{{.Name}} {{.Party}}")
  -party string
	  Political party (default: all)
  -v	verbose output
  -year int
	  Election cycle to list candidates from (default 2016)

⚛ ~$ 
⚛ ~$ export DATA_GOV_API_KEY=(YOUR KEY HERE)
⚛ ~$ list-candidates-potus -party DEM
P80000268 - CARTER, WILLIE FELIX DEM
P60008075 - CHAFEE, LINCOLN DAVENPORT MR. DEM
P00003392 - CLINTON, HILLARY RODHAM DEM
P60007267 - KELSO, LLOYD THOMAS DEM
P60007671 - O'MALLEY, MARTIN JOSEPH DEM
P60007168 - SANDERS, BERNARD DEM
P20004065 - WELLS, ROBERT CARR JR DEM
P60005204 - WILLIAMS, ELAINE WHIGHAM DEM
P60007515 - WILSON, WILLIE DEM
P60007754 - WINSLOW, BRAD MR. DEM
```

list-candidate-filings
---------------------

```sh
⚛ ~$ go get github.com/tmc/openfec/cmd/...
⚛ ~$ 
⚛ ~$ list-candidate-filings -h
$ list-candidate-filings -h
Usage of list-candidate-filings:
  -candidate string
      Candidate ID to fetch filings for
  -f string
      Formatting string (default "{{.}}")
  -v	verbose output
⚛ ~$ 
⚛ ~$ export DATA_GOV_API_KEY=(YOUR KEY HERE)
⚛ ~$ list-candidate-filings -f json -candidate P60007168 |jq .
{
  "amendment_indicator": "N",
  "beginning_image_number": "15031422533",
  "candidate_id": "P60007168",
  "candidate_name": "SANDERS, BERNARD",
  "committee_id": "P60007168",
  "cycle": 2016,
  "document_description": "Statement of candidacy 2015",
  "election_year": 2016,
  "ending_image_number": "15031422535",
  "file_number": -9144826,
  "form_type": "F2",
  "pages": 3,
  "pdf_url": "http://docquery.fec.gov/pdf/533/15031422533/15031422533.pdf",
  "receipt_date": "2015-04-30",
  "report_year": 2015,
  "sub_id": 1050120150017429800
}
```

