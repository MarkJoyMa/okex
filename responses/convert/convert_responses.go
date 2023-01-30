package convert

import (
	models "github.com/amir-the-h/okex/models/convert"
	"github.com/amir-the-h/okex/responses"
)

type (
	GetCurrencies struct {
		responses.Basic
		Currencies []*models.Currency `json:"data"`
	}
	GetCurrencyPair struct {
		responses.Basic
		Pairs []*models.CurrencyPair `json:"data"`
	}
	PostEstimateQuote struct {
		responses.Basic
		EstimateQuotes []*models.EstimateQuote `json:"data"`
	}
	PostTrade struct {
		responses.Basic
		Trades []*models.Trade `json:"data"`
	}
)
