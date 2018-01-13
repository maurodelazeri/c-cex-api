package ccex

import (
	"encoding/json"
	"fmt"
)

type Ticks []*Tick

type Tick struct {
	MarketName     string  `json:"MarketName"`
	High           float64 `json:"AskPrice"`
	Low            float64 `json:"BidPrice"`
	Volume         float64 `json:"Low"`
	Last           float64 `json:"High"`
	BaseVolume     float64 `json:"Volume"`
	TimeStamp      int64   `json:"LastPrice,string"`
	Bid            float64 `json:"BuyVolume"`
	Ask            float64 `json:"SellVolume"`
	OpenBuyOrders  float64 `json:"Change"`
	OpenSellOrders float64 `json:"Open"`
	PrevDay        float64 `json:"Close"`
	Created        int64   `json:"BaseVolume,string"`
}

type jsonResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

// ccex API implementation of Ticker endpoint.
//
// Endpoint:  getmarketsummaries
// Method: GET
//
// Example: https://c-cex.com/t/api_pub.html?a=getmarketsummaries
//
// Sample Response:
//
/*
[
    {
		"MarketName" : "USD-BTC",
		"High" : 0.00007220,
		"Low" : 0.00006897,
		"Volume" : 55886.98455819,
		"Last" : 0.00006994,
		"BaseVolume" : 3.90873570,
		"TimeStamp" : "1515820089",
		"Bid" : 0.00006994,
		"Ask" : 0.00007068,
		"OpenBuyOrders" : 100,
		"OpenSellOrders" : 100,
		"PrevDay" : 0.00007068,
		"Created" : "1515820089",
		"DisplayMarketName" : null
    },
  ]
*/

func (client *Client) GetMarkets() (Ticks, error) {

	resp, err := client.do("getmarketsummaries", nil)
	if err != nil {
		return nil, fmt.Errorf("Client.do: %v", err)
	}

	var response jsonResponse

	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	res := make(Ticks, 0)

	if err := json.Unmarshal(response.Result, &res); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return res, nil
}
