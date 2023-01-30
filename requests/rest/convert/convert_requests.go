package convert

type (
	GetCurrencyPair struct {
		FromCcy string `json:"fromCcy"`
		ToCcy   string `json:"toCcy"`
	}
	PostEstimateQuote struct {
		BaseCcy  string `json:"baseCcy"`
		QuoteCcy string `json:"quoteCcy"`
		Side     string `json:"side"`
		RfqSz    string `json:"rfqSz"`
		RfqSzCcy string `json:"rfqSzCcy"`
		ClQReqId string `json:"clQReqId"`
		Tag      string `json:"tag"`
	}
	PostTrade struct {
		QuoteId  string `json:"quoteId"`
		BaseCcy  string `json:"baseCcy"`
		QuoteCcy string `json:"quoteCcy"`
		Side     string `json:"side"`
		Sz       string `json:"sz"`
		SzCcy    string `json:"szCcy"`
		ClTReqId string `json:"clTReqId"`
		Tag      string `json:"tag"`
	}
)
