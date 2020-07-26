package crawler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"tradingview_udf_binance_go/model"
)

const (
	// BaseURL 바이낸스 API 주소
	BaseURL         = "https://api.binance.com/api/v3"
	exchangeInfoURL = "/exchangeInfo"
	klinesURL       = "/klines"
)

func GetExchangeInfo() *model.ExchangeInfo {
	body := requestAPI(exchangeInfoURL)
	exchangeInfo := new(model.ExchangeInfo)
	if err := json.Unmarshal(body, exchangeInfo); err != nil {
		log.Println(err.Error())
	}
	return exchangeInfo
}
func requestAPI(endpoint string) []byte {
	resp, _ := http.Get(BaseURL + endpoint)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}
