openfec
=======

Documentation: http://godoc.org/github.com/tmc/openfec

Status: some endpoints missing, pull requests welcome!

License: ISC

Example: 

list-candidates-potus
---------------------

```sh
⚛ ~$ go get github.com/tmc/openfec/examples/...
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
CHAFEE, LINCOLN DAVENPORT MR. DEM
CLINTON, HILLARY RODHAM DEM
KELSO, LLOYD THOMAS DEM
O'MALLEY, MARTIN JOSEPH DEM
SANDERS, BERNARD DEM
WELLS, ROBERT CARR JR DEM
WILLIAMS, ELAINE WHIGHAM DEM
WILSON, WILLIE DEM
WINSLOW, BRAD MR. DEM
```

