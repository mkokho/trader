package trader

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)

type Order struct {
	orderType string
	qty       int
	direction string
	account   string
}

type QuoteResponse struct {
	Ok bool
	Venue  string
	Symbol string
}

type Client struct {
	ApiKey string
	Venue  string
	Stock  string
}

func (client *Client) Quote(qr *QuoteResponse) error {
	r, err := client.get();
	defer r.Body.Close()

	if err != nil {
		return err
	}

	if err := json.NewDecoder(r.Body).Decode(&qr); err != nil {
		return err
	}

	return nil
}

func (c *Client) get() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.url(), nil)
	req.Header.Add("X-Starfighter-Authorization", c.ApiKey)

	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}

func (c *Client) url() string {
	url := fmt.Sprintf("https://api.stockfighter.io/ob/api/venues/%s/stocks/%s", c.Venue, c.Stock)
	log.Printf("Url: %v", url)
	return url
}
