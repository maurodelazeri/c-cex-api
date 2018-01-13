package ccex

import (
	"encoding/json"
	"fmt"
)

type Ticks []*Tick

type Tick struct {
	MarketName     string  `json:"MarketName"`
	High           float64 `json:"High"`
	Low            float64 `json:"Low"`
	Volume         float64 `json:"Volume"`
	Last           float64 `json:"Last"`
	BaseVolume     float64 `json:"BaseVolume"`
	TimeStamp      int64   `json:"TimeStamp,string"`
	Bid            float64 `json:"Bid"`
	Ask            float64 `json:"Ask"`
	OpenBuyOrders  float64 `json:"OpenBuyOrders"`
	OpenSellOrders float64 `json:"OpenSellOrders"`
	PrevDay        float64 `json:"PrevDay"`
	Created        int64   `json:"Created,string"`
}

type jsonResponse struct {
	Result json.RawMessage `json:"result"`
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

	resp, err := client.do("api_pub.html?a=getmarketsummaries", nil)
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
