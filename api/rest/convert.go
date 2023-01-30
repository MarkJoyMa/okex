package rest

import (
	"encoding/json"
	"net/http"

	"github.com/amir-the-h/okex"
	requests "github.com/amir-the-h/okex/requests/rest/convert"
	responses "github.com/amir-the-h/okex/responses/convert"
)

type Convert struct {
	client *ClientRest
}

// NewConvert returns a pointer to a fresh Convert
func NewConvert(c *ClientRest) *Convert {
	return &Convert{c}
}

func (c *Convert) GetCurrencies() (response responses.GetCurrencies, err error) {
	p := "/api/v5/asset/convert/currencies"

	res, err := c.client.Do(http.MethodGet, p, true)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

func (c *Convert) GetCurrencyPair(req requests.GetCurrencyPair) (response responses.GetCurrencyPair, err error) {
	p := "/api/v5/asset/convert/currency-pair"
	m := okex.S2M(req)

	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

func (c *Convert) PostEstimateQuote(req requests.PostEstimateQuote) (response responses.PostEstimateQuote, err error) {
	p := "/api/v5/asset/convert/estimate-quote"
	m := okex.S2M(req)

	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

func (c *Convert) PostTrade(req requests.PostTrade) (response responses.PostTrade, err error) {
	p := "/api/v5/asset/convert/trade"
	m := okex.S2M(req)

	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}
