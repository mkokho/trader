package trader

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"log"
)

type Client struct {
	ApiKey string
	Venue  string
	Stock  string
}

type Order struct {
	Account   string
	Direction string
	OrderType string
	Qty       int
	Price     int
}

type OrderResponse struct {
	Ok bool
	Qty int
	Fills []OrderFill
	Error string
}

type OrderFill struct {
	Price int
	Qty int
	Ts string
}

type QuoteResponse struct {
	Ok     bool
	Venue  string
	Symbol string
	Error string
}

func (client *Client) Quote() (*QuoteResponse, error) {
	req, err := http.NewRequest("GET", client.url("quote"), nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	if (response != nil) {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var qr QuoteResponse
	if err := json.NewDecoder(response.Body).Decode(&qr); err != nil {
		return nil, err
	}

	log.Printf("Received quote: %+v", qr)
	return &qr, nil
}

func (c *Client) PostOrder(order *Order) (*OrderResponse, error) {
	payload, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", c.url("orders"), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-Starfighter-Authorization", c.ApiKey)

	response, err := http.DefaultClient.Do(request)
	if (response != nil) {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var or OrderResponse
    if err := json.NewDecoder(response.Body).Decode(&or); err != nil {
	    return nil, err
    }

	log.Printf("Received order response: %+v", or)
	return &or, nil
}

func (c *Client) url(end string) string {
	url := fmt.Sprintf("https://api.stockfighter.io/ob/api/venues/%s/stocks/%s/%s", c.Venue, c.Stock, end)
	return url
}