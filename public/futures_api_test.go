package public

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGroups(t *testing.T) {
	groupId := "1"
	response := &Response{
		path: "/groups" + "/" + groupId,
		body: `[{
					"id": 1,
					"currency_pair": "btc_jpy",
					"start_timestamp": 1490972400,
					"end_timestamp": 4102412399,
					"use_swap": true
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

	res, err := client.Groups(groupId)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, res[0].UseSwap)
}

func TestFuturesLastPrice(t *testing.T) {
	groupId := "1"
	currencyPair := "btc_jpy"
	response := &Response{
		path: "/last_price" + "/" + groupId + "/" + currencyPair,
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

	res, err := client.FuturesLastPrice(groupId, currencyPair)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 787070.0, res.LastPrice)
}

func TestFuturesTicker(t *testing.T) {
	groupId := "1"
	currencyPair := "btc_jpy"
	response := &Response{
		path: "/ticker" + "/" + groupId + "/" + currencyPair,
		body: `{"last": 784000.0, 
				"high": 791195.0, 
				"low": 773110.0, 
				"vwap": 784234.6928, 
				"volume": 11625.9544, 
				"bid": 784000.0, 
				"ask": 784400.0, 
				"swap_rate_bid": 0.08051154, 
				"swap_rate_ask": -0.08051154
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

	res, err := client.FuturesTicker(groupId, currencyPair)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 784234.6928, res.Vwap)
	assert.Equal(t, -0.08051154, res.SwapRateAsk)
}

func TestFuturesTrades(t *testing.T) {
	groupId := "1"
	currencyPair := "btc_jpy"
	response := &Response{
		path: "/trades" + "/" + groupId + "/" + currencyPair,
		body: `[{
        			"date": 1491756592,
        			"price": 135340.0,
        			"amount": 0.02,
        			"tid": 102659,
        			"currency_pair": "btc_jpy",
        			"trade_type": "ask"
    			},
    			{
        			"date": 1491756591,
        			"price": 135345.0,
					"amount": 0.01,
        			"tid": 102658,
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

	res, err := client.FuturesTrades(groupId, currencyPair)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "ask", res[0].TradeType)
	assert.Equal(t, 135345.0, res[1].Price)
}

func TestFuturesDepth(t *testing.T) {
	groupId := "1"
	currencyPair := "btc_jpy"
	response := &Response{
		path: "/depth" + "/" + groupId + "/" + currencyPair,
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

	res, err := client.FuturesDepth(groupId, currencyPair)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 134875.0, res.Asks[0][0])
	assert.Equal(t, 134865.0, res.Bids[1][0])
}

func TestSwapHistory(t *testing.T) {
	groupId := "1"
	currencyPair := "btc_jpy"
	response := &Response{
		path: "/swap_history" + "/" + groupId + "/" + currencyPair,
		body: `[{
        			"timestamp": 1504008000,
        			"swap_rate_bid": 0.1,
        			"swap_rate_ask": -0.1
    			},
    			{
        			"timestamp": 1504000800,
					"swap_rate_bid": 0.375,
        			"swap_rate_ask": -0.375
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

	res, err := client.SwapHistory(groupId, currencyPair)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, -0.1, res[0].SwapRateAsk)
	assert.Equal(t, 1504000800, res[1].Timestamp)
}
