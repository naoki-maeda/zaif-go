package test

import (
	"github.com/naoki-maeda/zaif-go/public"
	"testing"
)

var futuresApi = public.NewApiClient("https://api.zaif.jp/fapi/1")

func Test_Groups(t *testing.T) {
	_, err := futuresApi.Groups("1")
	if err != nil {
		t.Error(err)
	}
}

func Test_FuturesLastPrice(t *testing.T) {
	_, err := futuresApi.FuturesLastPrice("1", "btc_jpy")
	if err != nil {
		t.Error(err)
	}
}

func Test_FuturesTicker(t *testing.T) {
	_, err := futuresApi.FuturesTicker("1", "btc_jpy")
	if err != nil {
		t.Error(err)
	}
}

func Test_FuturesTrades(t *testing.T) {
	_, err := futuresApi.FuturesTrades("1", "btc_jpy")
	if err != nil {
		t.Error(err)
	}
}

func Test_FuturesDepth(t *testing.T) {
	_, err := futuresApi.FuturesDepth("1", "btc_jpy")
	if err != nil {
		t.Error(err)
	}
}

func Test_SwapHistory(t *testing.T) {
	_, err := futuresApi.SwapHistory("1", "btc_jpy")
	if err != nil {
		t.Error(err)
	}
}
