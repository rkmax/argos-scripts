package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
)

type ExchangeRate struct {
	FromCurrencyCode string  `json:"1. From_Currency Code"`
	FromCurrencyName string  `json:"2. From_Currency Name"`
	ToCurrencyCode   string  `json:"3. To_Currency Code"`
	ToCurrencyName   string  `json:"4. To_Currency Name"`
	ExchangeRate     float32 `json:"5. Exchange Rate,string"`
}

type ExchangeRateResponse struct {
	Rate ExchangeRate `json:"Realtime Currency Exchange Rate"`
}

func currencyExchangeRate(from, to, apiKey string) (*ExchangeRate, error) {
	query := fmt.Sprintf("function=CURRENCY_EXCHANGE_RATE&from_currency=%s&to_currency=%s&apikey=%s", from, to, apiKey)
	requestUrl := getURL(query)

	rs, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	response := &ExchangeRateResponse{}
	err = json.NewDecoder(rs.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	return &response.Rate, nil
}

func getURL(rawQuery string) string {
	theUrl := url.URL{
		Scheme:   "https",
		Host:     "www.alphavantage.co",
		Path:     "/query",
		RawQuery: rawQuery,
	}

	return theUrl.String()
}

func main() {
	flag.Parse()
	from := flag.Arg(0)
	to := flag.Arg(1)
	apiKey := flag.Arg(2)
	rate, err := currencyExchangeRate(from, to, apiKey)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s → %s: %0.0f\n", rate.FromCurrencyCode, rate.ToCurrencyCode, rate.ExchangeRate)
}
