package test

import (
	"testing"
	"github.com/naoki-maeda/zaif-go/public"
)

var api = public.NewApiClient()

func TestPublicData_Currencies(t *testing.T) {
	_, err := api.Currencies("btc")
	if err != nil {
		t.Error(err)
	}
}

func TestPublicData_CurrencyPairs(t *testing.T) {
	_, err := api.CurrencyPairs("btc_jpy")
	if err != nil {
		t.Error(err)
	}
}

func TestPublicData_LastPrice(t *testing.T) {
	_, err := api.LastPrice("mona_jpy")
	if err != nil {
		t.Error(err)
	}
}

func TestPublicData_Ticker(t *testing.T) {
	_, err := api.Ticker("xem_jpy")
	if err != nil {
		t.Error(err)
	}
}

func TestPublicData_Trades(t *testing.T) {
	_, err := api.Trades("eth_jpy")
	if err != nil {
		t.Error(err)
	}
}

func TestPublicData_Depth(t *testing.T) {
	_, err := api.Depth("bch_jpy")
	if err != nil {
		t.Error(err)
	}
}
