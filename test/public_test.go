package test

import (
	"github.com/naoki-maeda/zaif-go/public"
	"testing"
)

var publicApi = public.NewApiClient("https://api.zaif.jp/api/1")

func Test_Currencies(t *testing.T) {
	_, err := publicApi.Currencies("btc")
	if err != nil {
		t.Error(err)
	}
}

func Test_CurrencyPairs(t *testing.T) {
	_, err := publicApi.CurrencyPairs("btc_jpy")
	if err != nil {
		t.Error(err)
	}
}

func Test_LastPrice(t *testing.T) {
	_, err := publicApi.LastPrice("mona_jpy")
	if err != nil {
		t.Error(err)
	}
}

func Test_Ticker(t *testing.T) {
	_, err := publicApi.Ticker("xem_jpy")
	if err != nil {
		t.Error(err)
	}
}

func Test_Trades(t *testing.T) {
	_, err := publicApi.Trades("eth_jpy")
	if err != nil {
		t.Error(err)
	}
}

func Test_Depth(t *testing.T) {
	_, err := publicApi.Depth("bch_jpy")
	if err != nil {
		t.Error(err)
	}
}
