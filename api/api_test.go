package api

import (
	"context"
	"log"
	"testing"

	"github.com/amir-the-h/okex"
	requests "github.com/amir-the-h/okex/requests/rest/convert"
)

func TestNewClient(t *testing.T) {
	apiKey := "xxx"
	secretKey := "xxx"
	passphrase := "xxxx"
	dest := okex.NormalServer // The main API server
	ctx := context.Background()
	client, err := NewClient(ctx, apiKey, secretKey, passphrase, dest)
	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.Rest.Account.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Account Config %+v", response)

	resp2, err := client.Rest.Convert.GetCurrencies()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Account Config %+v", resp2)

	resp3, err := client.Rest.Convert.GetCurrencyPair(requests.GetCurrencyPair{
		FromCcy: "BTC",
		ToCcy:   "USDT",
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Account Config %+v", resp3)
}
