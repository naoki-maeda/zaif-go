package public

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response struct {
	path string
	body string
}

func TestCurrencies(t *testing.T) {
	currency := "btc"
	response := &Response{
		path: "/currencies" + "/" + currency,
		body: `[{"id": 1, "token_id": null, "is_token": false, "name": "btc"}]`,
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		if p, w := r.URL.Path, response.path; p != w {
			t.Errorf("request got path %s, want %s", p, w)
		}
		io.WriteString(w, response.body)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewApiClient(server.URL)

	res, err := client.Currencies(currency)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "btc", res[0].Name)
	assert.Equal(t, false, res[0].IsToken)
}

func TestCurrencyPairs(t *testing.T) {
	currencyPair := "btc_jpy"
	response := &Response{
		path: "/currency_pairs" + "/" + currencyPair,
		body: `[{
					"aux_unit_point": 0,
					"item_japanese": "\u30d3\u30c3\u30c8\u30b3\u30a4\u30f3", 
					"aux_unit_step": 5.0, 
					"description": "\u30d3\u30c3\u30c8\u30b3\u30a4\u30f3\u30fb\u65e5\u672c\u5186\u306e\u53d6\u5f15\u3092\u884c\u3046\u3053\u3068\u304c\u3067\u304d\u307e\u3059", 
					"item_unit_min": 0.001, 
					"event_number": 0, 
					"currency_pair": "btc_jpy", 
					"is_token": false, 
					"aux_unit_min": 5.0, 
					"aux_japanese": "\u65e5\u672c\u5186", 
					"id": 1, 
					"item_unit_step": 0.0001, 
					"name": "BTC/JPY", 	
					"seq": 0, 
					"title": "BTC/JPY"
				}]`,
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		if p, w := r.URL.Path, response.path; p != w {
			t.Errorf("request got path %s, want %s", p, w)
		}
		io.WriteString(w, response.body)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewApiClient(server.URL)

	res, err := client.CurrencyPairs(currencyPair)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "BTC/JPY", res[0].Name)
	assert.Equal(t, 0.001, res[0].ItemUnitMin)
}

func TestLastPrice(t *testing.T) {
	currencyPair := "btc_jpy"
	response := &Response{
		path: "/last_price" + "/" + currencyPair,
		body: `{"last_price": 787070.0}`,
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		if p, w := r.URL.Path, response.path; p != w {
			t.Errorf("request got path %s, want %s", p, w)
		}
		io.WriteString(w, response.body)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewApiClient(server.URL)

	res, err := client.LastPrice(currencyPair)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 787070.0, res.LastPrice)
}

func TestTicker(t *testing.T) {
	currencyPair := "mona_jpy"
	response := &Response{
		path: "/ticker" + "/" + currencyPair,
		body: `{
    			"last": 135875.0,
    			"high": 136000.0,
    			"low": 131570.0,
				"vwap": 133301.7489,
    			"volume": 6889.215,
    			"bid": 135875.0,
    			"ask": 135920.0
				}`,
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		if p, w := r.URL.Path, response.path; p != w {
			t.Errorf("request got path %s, want %s", p, w)
		}
		io.WriteString(w, response.body)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewApiClient(server.URL)

	res, err := client.Ticker(currencyPair)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 133301.7489, res.Vwap)
	assert.Equal(t, 135920.0, res.Ask)
}

func TestTrades(t *testing.T) {
	currencyPair := "btc_jpy"
	response := &Response{
		path: "/trades" + "/" + currencyPair,
		body: `[
		{
        	"date": 1491756592,
        	"price": 135340.0,
        	"amount": 0.02,
        	"tid": 43054307,
        	"currency_pair": "btc_jpy",
        	"trade_type": "ask"
    	},
    	{
        	"date": 1491756591,
        	"price": 135345.0,
        	"amount": 0.01,
        	"tid": 43054306,
        	"currency_pair": "btc_jpy",
        	"trade_type": "bid"
    	}]`,
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		if p, w := r.URL.Path, response.path; p != w {
			t.Errorf("request got path %s, want %s", p, w)
		}
		io.WriteString(w, response.body)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewApiClient(server.URL)

	res, err := client.Trades(currencyPair)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 43054307, res[0].Tid)
	assert.Equal(t, 135345.0, res[1].Price)
}

func TestDepth(t *testing.T) {
	currencyPair := "btc_jpy"
	response := &Response{
		path: "/depth" + "/" + currencyPair,
		body: `{
    		"asks": [[134875.0,0.0063],[134885.0,0.1639]],
    		"bids": [[134870.0,0.01],[134865.0,0.3066]]
		}`,
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		if p, w := r.URL.Path, response.path; p != w {
			t.Errorf("request got path %s, want %s", p, w)
		}
		io.WriteString(w, response.body)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := NewApiClient(server.URL)

	res, err := client.Depth(currencyPair)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 134875.0, res.Asks[0][0])
	assert.Equal(t, 134865.0, res.Bids[1][0])
}
