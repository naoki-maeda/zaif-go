package test

import (
	"github.com/naoki-maeda/zaif-go"
	"github.com/naoki-maeda/zaif-go/private"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	apiKey    = "API_KEY"
	apiSecret = "API_SECRET"
)

type RequestItem struct {
	Key          string
	Sign         string
	ContentType  string
	ExpectedBody string
}

func NewRequest(method string, body string) *RequestItem {
	return &RequestItem{
		Key:          apiKey,
		Sign:         zaif.Sign(method, apiSecret),
		ContentType:  "application/x-www-form-urlencoded",
		ExpectedBody: body,
	}
}

func (i *RequestItem) CreateHandler(t *testing.T, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", i.ContentType)
	w.Header().Set("Key", i.Key)
	w.Header().Set("Sign", i.Sign)
	if r.Method != "POST" {
		t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
	}
	io.WriteString(w, i.ExpectedBody)
}

func TestGetInfo(t *testing.T) {
	expected := `{
	"success":1,
	"return":{
		"funds":{
			"jpy":15320,
			"btc":1.389,
			"xem":100.2,
			"mona":2600,
			"pepecash":0.1
		},
		"deposit":{
			"jpy":20440,
			"btc":1.479,
			"xem":100.2,
			"mona":3200,
			"pepecash":0.1
		},
		"rights":{
			"info":1,
			"trade":1,
			"withdraw":0,
			"personal_info":0,
			"id_info":0
		},
		"trade_count":18,
		"open_orders":3,
		"server_time":1401950833
		}
	}`
	request := NewRequest("get_info", expected)
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

	res, err := client.GetInfo()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1.389, res.Funds.BTC)
	assert.Equal(t, 18, res.TradeCount)
}
func TestGetInfo2(t *testing.T) {
	expected := `{
	"success":1,
	"return":{
		"funds":{
			"jpy":15320,
			"btc":1.389,
			"xem":100.2,
			"mona":2600,
			"pepecash":0.1
		},
		"deposit":{
			"jpy":20440,
			"btc":1.479,
			"xem":100.2,
			"mona":3200,
			"pepecash":0.1
		},
		"rights":{
			"info":1,
			"trade":1,
			"withdraw":0,
			"personal_info":0
		},
		"open_orders":3,
		"server_time":1401950833
		}
	}`
	request := NewRequest("get_info2", expected)
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
	res, err := client.GetInfo2()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 100.2, res.Funds.XEM)
	assert.Equal(t, 3, res.OpenOrders)
}

func TestGetPersonalInfo(t *testing.T) {
	expected := `{
    "success": 1,
    	"return": {
        	"ranking_nickname": "サトシ・ナカモト",
        	"icon_path": "https://abs.twimg.com/sticky/default_profile_images/default_profile_0_normal.png"
    	}
	}`
	request := NewRequest("get_personal_info", expected)
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

	res, err := client.GetPersonalInfo()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "サトシ・ナカモト", res.RankingNickname)
}

func TestGetIdInfo(t *testing.T) {
	expected := `{
    "success": 1,
    	"return": {
        	"id": 100,
        	"email": "satoshinakamoto@gmail.com",
			"name": "哲史",
			"kana": "サトシ",
			"certified": true
    	}
	}`
	request := NewRequest("get_id_info", expected)
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

	res, err := client.GetIdInfo()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "哲史", res.Name)
	assert.Equal(t, true, res.Certified)
}

func TestTradeHistory(t *testing.T) {
	expected := `{
    "success": 1,
    	"return": {
        	"182": {
            	"currency_pair": "btc_jpy",
            	"action": "bid",
            	"amount": 0.03,
            	"price": 56000,
            	"fee": 0,
            	"your_action": "ask",
            	"bonus": 1.6,
            	"timestamp": 1402018713,
            	"comment" : "demo"
        	}
    	}
	}`
	request := NewRequest("trade_history", expected)
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

	res, err := client.TradeHistory(private.TradeHistoryParams{CurrencyPair: "btc_jpy"})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "btc_jpy", res["182"].CurrencyPair)
	assert.Equal(t, 0.03, res["182"].Amount)
}

func TestActiveOrders(t *testing.T) {
	expected := `{
    "success": 1,
    	"return": {
        	"184": {
            	"currency_pair": "btc_jpy",
            	"action": "ask",
            	"amount": 0.03,
            	"price": 56000,
            	"timestamp": 1402021125,
            	"comment" : "demo"
			}
    	}
	}`
	request := NewRequest("active_orders", expected)
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

	res, err := client.ActiveOrders(private.ActiveOrdersParams{CurrencyPair: "btc_jpy"})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "btc_jpy", res["184"].CurrencyPair)
	assert.Equal(t, "ask", res["184"].Action)
}

func TestTrade(t *testing.T) {
	expected := `{
    "success": 1,
    	"return": {
        	"received": 0.1,
        	"remains": 0,
        	"order_id": 0,
        	"funds": {
            	"jpy": 325,
            	"btc": 1.392,
            	"mona": 2600
        	}
    	}
	}`
	request := NewRequest("trade", expected)
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

	res, err := client.Trade(private.TradeParams{CurrencyPair: "mona_jpy", Action: "bid", Price: 500, Amount: 10})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 0.1, res.Received)
	assert.Equal(t, float64(2600), res.Funds.MONA) // floatに変換して無理やり通す
}

func TestCancelOrder(t *testing.T) {
	expected := `{
    "success": 1,
    	"return": {
        	"order_id": 184,
        	"funds": {
            	"jpy": 15320,
            	"btc": 1.392,
            	"mona": 2600,
            	"kaori": 0.1
        	}
    	}
	}`
	request := NewRequest("cancel_order", expected)
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

	res, err := client.CancelOrder(private.CancelOrderParams{OrderID: 184})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 184, res.OrderID)
	assert.Equal(t, 1.392, res.Funds.BTC)
}

func TestWithdraw(t *testing.T) {
	expected := `{
	"success": 1,
		"return": {
      		"id": 23634,
      		"fee": 0.001,
      		"txid":"",
			"funds": {
          		"jpy": 15320,
          		"btc": 1.392,
          		"xem": 100.2,
          		"mona": 2600
      		}
		}
	}`
	request := NewRequest("withdraw", expected)
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

	res, err := client.Withdraw(private.WithdrawParams{Currency: "btc", Address: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", Amount: 0.1})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 0.001, res.Fee)
	assert.Equal(t, 1.392, res.Funds.BTC)
}

func TestDepositHistory(t *testing.T) {
	expected := `{
    "success":1,
    	"return":{
        	"3816":{
				"timestamp":1435745065,
          		"address":"12qwQ3sPJJAosodSUhSpMds4WfUPBeFEM2",
          		"amount":0.001,
          		"txid":"64dcf59523379ba282ae8cd61d2e9382c7849afe3a3802c0abb08a60067a159f"
        	},
        	"3814":{
          		"timestamp":1435548083,
          		"address":"12qwQ3sPJJAosodSUhSpMds4WfUPBeFEM2",
          		"amount":0.001,
          		"txid":"7d012cfff6e67a8938f93215367eef4177604459631ea62c85550980dca71819"
			}
		}
	}`
	request := NewRequest("deposit_history", expected)
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

	res, err := client.DepositHistory(private.DepositHistoryParams{Currency: "btc"})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "12qwQ3sPJJAosodSUhSpMds4WfUPBeFEM2", res["3816"].Address)
	assert.Equal(t, "7d012cfff6e67a8938f93215367eef4177604459631ea62c85550980dca71819", res["3814"].Txid)
}

func TestWithdrawHistory(t *testing.T) {
	expected := `{
    "success":1,
    	"return":{
        	"3816":{
          		"timestamp":1435745065,
          		"address":"12qwQ3sPJJAosodSUhSpMds4WfUPBeFEM2",
          		"amount":0.001,
          		"txid":"64dcf59523379ba282ae8cd61d2e9382c7849afe3a3802c0abb08a60067a159f"
			},
        	"3814":{
          		"timestamp":1435548083,
          		"address":"12qwQ3sPJJAosodSUhSpMds4WfUPBeFEM2",
          		"amount":0.001,
          		"txid":"7d012cfff6e67a8938f93215367eef4177604459631ea62c85550980dca71819"
			}
    	}
	}`
	request := NewRequest("withdraw_history", expected)
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

	res, err := client.WithdrawHistory(private.WithdrawHistoryParams{Currency: "btc"})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 0.001, res["3816"].Amount)
	assert.Equal(t, 1435548083, res["3814"].Timestamp)
}
