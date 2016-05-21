package main

import (
	"github.com/mkokho/trader"
)

func main() {
	client := trader.Client{
		ApiKey: "0410abc34660b2158e7bcc90dd1a04fbb64484d3",
		Venue: "RLWBEX",
		Stock: "CDO",
	}

	qr, err := client.Quote()
	if err != nil {
		panic(err)
	}

	if qr.Ok {
		_, err := client.PostOrder(&trader.Order{
			Account: "WMG19022764",
			Direction: "buy",
			OrderType: "fill-or-kill",
			Qty: 10,
			Price: 12000,
		})

		if err != nil {
			panic(err)
		}
	}
}