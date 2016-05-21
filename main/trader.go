package main

import (
	"fmt"
	"github.com/mkokho/trader"
	"github.com/gpmgo/gopm/modules/log"
)

func main() {
    client := trader.Client{
	    ApiKey: "0410abc34660b2158e7bcc90dd1a04fbb64484d3",
        Venue: "GWHKEX",
	    Stock: "EFAY",
    }

	var qr trader.QuoteResponse
	if err := client.Quote(&qr); err != nil {
		log.Error(err.Error())
	}

	fmt.Printf("Respose: %+v", qr)
}