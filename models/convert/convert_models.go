package convert

type (
	Currency struct {
		Ccy string `json:"ccy"`
		Min string `json:"min"`
		Max string `json:"max"`
	}
	CurrencyPair struct {
		InstId      string `json:"instId"`
		BaseCcy     string `json:"baseCcy"`
		BaseCcyMax  string `json:"baseCcyMax"`
		BaseCcyMin  string `json:"baseCcyMin"`
		QuoteCcy    string `json:"quoteCcy"`
		QuoteCcyMax string `json:"quoteCcyMax"`
		QuoteCcyMin string `json:"quoteCcyMin"`
	}
	EstimateQuote struct {
		QuoteTime string `json:"quoteTime"`
		TtlMs     string `json:"ttlMs"`
		ClQReqId  string `json:"clQReqId"`
		QuoteId   string `json:"quoteId"`
		BaseCcy   string `json:"baseCcy"`
		QuoteCcy  string `json:"quoteCcy"`
		Side      string `json:"side"`
		OrigRfqSz string `json:"origRfqSz"`
		RfqSz     string `json:"rfqSz"`
		RfqSzCcy  string `json:"rfqSzCcy"`
		CnvtPx    string `json:"cnvtPx"`
		BaseSz    string `json:"baseSz"`
		QuoteSz   string `json:"quoteSz"`
	}
	Trade struct {
		TradeId     string `json:"tradeId"`
		QuoteId     string `json:"quoteId"`
		ClTReqId    string `json:"clTReqId"`
		State       string `json:"state"`
		InstId      string `json:"instId"`
		BaseCcy     string `json:"baseCcy"`
		QuoteCcy    string `json:"quoteCcy"`
		Side        string `json:"side"`
		FillPx      string `json:"fillPx"`
		FillBaseSz  string `json:"fillBaseSz"`
		FillQuoteSz string `json:"fillQuoteSz"`
		Ts          string `json:"ts"`
	}
)
