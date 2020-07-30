package crawler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

func GetKlines(symbol string, interval string, startTime int64, endTime int64) *[]model.Kline {
	body := requestAPI(klinesURL + "?symbol=" + symbol + "&interval=" + interval + "&startTime=" + strconv.FormatInt(startTime, 10) + "&endTime=" + strconv.FormatInt(endTime, 10))
	//log.Println(klinesURL + "?symbol=" + symbol + "&interval=" + interval + "&startTime=" + strconv.FormatInt(startTime, 10) + "&endTime=" + strconv.FormatInt(endTime, 10))
	var arr [][]interface{}
	json.Unmarshal(body, &arr)

	exchangeInfo := make([]model.Kline, len(arr))
	for i, data := range arr {
		exchangeInfo[i] = model.Kline{
			OpenTime:         int64(data[0].(float64) / 1000),
			Open:             data[1].(float64),
			High:             data[2].(float64),
			Low:              data[3].(float64),
			Close:            data[4].(float64),
			Volume:           data[5].(float64),
			CloseTime:        int64(data[6].(float64) / 1000),
			QuoteVolume:      data[7].(float64),
			NumberOfTrades:   data[8].(int64),
			TakerBaseVolume:  data[9].(float64),
			TakerQuoteVolume: data[10].(float64),
			Ignore:           data[11].(float64),
		}
	}

	return &exchangeInfo
}

func requestAPI(endpoint string) []byte {
	resp, _ := http.Get(BaseURL + endpoint)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body
}
