package test

import (
	"github.com/naoki-maeda/zaif-go/private"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPositions(t *testing.T) {
	expected := `{
    "success": 1,
    	"return": {
        	"182": {
            	"group_id": 1,
            	"currency_pair": "btc_jpy",
            	"action": "bid",
            	"leverage": 2.5,
            	"price": 110005,
            	"limit": 130000,
            	"stop": 90000,
            	"amount": 0.03,
            	"fee_spent": 0,
            	"timestamp": "1402018713",
            	"term_end": "1404610713",
            	"timestamp_closed": "1402019000",
            	"deposit": 35.76 ,
            	"deposit_jpy": 35.76,
            	"refunded": 35.76 ,
            	"refunded_jpy": 35.76,
            	"swap": 0
        	}
    	}
	}`
	request := NewRequest("get_positions", expected)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", request.ContentType)
		w.Header().Set("Key", request.Key)
		w.Header().Set("Sign", request.Sign)
		if r.Method != "POST" {
			t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
		}
		io.WriteString(w, request.ExpectedBody)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := private.NewApiClient(apiKey, apiSecret, server.URL)

	res, err := client.GetPositions(private.GetPositionsParams{Type: "futures", GroupID: 1, CurrencyPair: "btc_jpy"})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "btc_jpy", res["182"].CurrencyPair)
	assert.Equal(t, 35.76, res["182"].RefundedJPY)
}

func TestPositionHistory(t *testing.T) {
	expected := `{
    "success": 1,
    	"return": {
        	"182": {
            	"group_id": 1,
            	"currency_pair": "btc_jpy",
            	"action": "bid",
            	"amount": 0.0001,
            	"price": 499000,
            	"timestamp": "1504251232",
            	"your_action": "bid",
            	"bid_leverage_id": 182
        	},
        	"183": {
            	"group_id": 1,
            	"currency_pair": "btc_jpy",
            	"action": "ask",
            	"amount": 0.0001,
            	"price": 450000,
            	"timestamp": "1504251267",
            	"your_action": "ask",
            	"ask_leverage_id": 182
        	}
		}
	}`
	request := NewRequest("position_history", expected)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", request.ContentType)
		w.Header().Set("Key", request.Key)
		w.Header().Set("Sign", request.Sign)
		if r.Method != "POST" {
			t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
		}
		io.WriteString(w, request.ExpectedBody)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := private.NewApiClient(apiKey, apiSecret, server.URL)

	res, err := client.PositionHistory(private.PositionHistoryParams{Type: "margin"})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "bid", res["182"].Action)
	assert.Equal(t, 182, res["183"].AskLeverageId)
}

func TestActivePositions(t *testing.T) {
	expected := `{
    "success": 1,
    	"return": {
        	"184": {
            	"group_id": "1",
            	"currency_pair": "btc_jpy",
            	"action": "ask",
            	"amount": 0.0001,
            	"price": 450000,
            	"timestamp": "1402021125",
            	"term_end": "1404613125",
            	"leverage": 1,
            	"fee_spent": 0.0015,
            	"price_avg": 450000,
            	"amount_done": 0.0001,
            	"deposit_jpy": 48.72
        	}
    	}
	}`
	request := NewRequest("active_positions", expected)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", request.ContentType)
		w.Header().Set("Key", request.Key)
		w.Header().Set("Sign", request.Sign)
		if r.Method != "POST" {
			t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
		}
		io.WriteString(w, request.ExpectedBody)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := private.NewApiClient(apiKey, apiSecret, server.URL)

	res, err := client.ActivePositions(private.ActivePositionsParams{Type: "futures", GroupID: 1, CurrencyPair: "btc_jpy"})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "btc_jpy", res["184"].CurrencyPair)
	assert.Equal(t, float64(450000), res["184"].PriceAvg)
}

func TestCreatePosition(t *testing.T) {
	expected := `{
    "success": 1,
    	"return": {
        	"leverage_id": 22258,
        	"timestamp": "1504253833",
        	"term_end": "1506845833",
        	"price_avg": 118000,
        	"amount_done": 0.0001,
        	"deposit_jpy": 11.92,
        	"funds": {
            	"jpy": 325,
            	"btc": 1.392,
            	"mona": 2600
        	}
    	}
	}`
	request := NewRequest("create_position", expected)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", request.ContentType)
		w.Header().Set("Key", request.Key)
		w.Header().Set("Sign", request.Sign)
		if r.Method != "POST" {
			t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
		}
		io.WriteString(w, request.ExpectedBody)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := private.NewApiClient(apiKey, apiSecret, server.URL)

	res, err := client.CreatePosition(private.CreatePositionParams{Type: "futures", GroupID: 1, CurrencyPair: "btc_jpy"})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 22258, res.LeverageId)
	assert.Equal(t, 0.0001, res.AmountDone)
}

func TestChangePosition(t *testing.T) {
	expected := `{
    "success": 1,
    	"return": {
        	"leverage_id": 22258,
        	"price_avg": 118000,
        	"amount_done": 0.0001
    	}
	}`
	request := NewRequest("create_position", expected)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", request.ContentType)
		w.Header().Set("Key", request.Key)
		w.Header().Set("Sign", request.Sign)
		if r.Method != "POST" {
			t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
		}
		io.WriteString(w, request.ExpectedBody)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := private.NewApiClient(apiKey, apiSecret, server.URL)

	res, err := client.ChangePosition(private.ChangePositionParams{Type: "futures", GroupID: 1, LeverageId: 22258, Price: 119000})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 22258, res.LeverageId)
	assert.Equal(t, 0.0001, res.AmountDone)
}

func TestChancelPosition(t *testing.T) {
	expected := `{
    "success": 1,
    	"return": {
        	"leverage_id": 2072,
        	"refunded_jpy": 645.96,
        	"funds": {
            	"btc": 0.496,
            	"jpy": 1564.96,
            	"xem": 0.0,
            	"mona": 10.0
        	},
        	"fee_spent": 0.0,
        	"timestamp_closed": "1508384951",
        	"swap": 0.0
    	}
	}`
	request := NewRequest("cancel_position", expected)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", request.ContentType)
		w.Header().Set("Key", request.Key)
		w.Header().Set("Sign", request.Sign)
		if r.Method != "POST" {
			t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
		}
		io.WriteString(w, request.ExpectedBody)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := private.NewApiClient(apiKey, apiSecret, server.URL)

	res, err := client.CancelPosition(private.CancelPositionParams{Type: "futures", GroupID: 1, LeverageId: 2072})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2072, res.LeverageId)
	assert.Equal(t, 0.0, res.Swap)
}
