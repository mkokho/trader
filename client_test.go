package trader_test

import (
	"testing"
	"github.com/mkokho/trader"
)

var client = trader.Client{
	ApiKey: "0410abc34660b2158e7bcc90dd1a04fbb64484d3",
	Venue: "TESTEX",
	Stock: "FOOBAR",
}

func TestClient_Quote(t *testing.T) {
	qr, err := client.Quote()
	if err != nil {
		t.Errorf("Quote method returned unexpected error %v", err)
	}

	if (!qr.Ok) {
		t.Errorf("Quote response: ok is false, expected true", err)
	}
}

func TestClient_PostOrder(t *testing.T) {
	order := trader.Order{
		Account: "EXB123456",
		Direction: "buy",
		OrderType: "limit",
		Qty: 10,
		Price: 20,
	}

	resp, err := client.PostOrder(&order)
	if err != nil {
		t.Errorf("PostOrder method returned unexpected error %v", err)
	}

	if (!resp.Ok) {
		t.Errorf("PostOrder response: ok is false, expected true", err)
	}
}