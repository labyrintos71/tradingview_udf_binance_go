package api

import (
	"log"
	"net/http"
	"strconv"
	"tradingview_udf_binance_go/crawler"
	"tradingview_udf_binance_go/model"

	"github.com/labstack/echo"
)

var ResolrutionMap = map[string]string{
	"1":   "1m",
	"3":   "3m",
	"5":   "5m",
	"15":  "15m",
	"30":  "30m",
	"60":  "1h",
	"120": "2h",
	"240": "4h",
	"360": "6h",
	"480": "8h",
	"720": "12h",
	"D":   "1d",
	"1D":  "1d",
	"3D":  "3d",
	"W":   "1w",
	"1W":  "1w",
	"M":   "1M",
	"1M":  "1M",
}

// History klines를 이용한 심볼 조회
func History(c echo.Context) error {
	symbol := c.QueryParam("symbol")
	resolution := ResolrutionMap[c.QueryParam("resolution")]
	from, err := strconv.ParseInt(c.QueryParam("from"), 0, 64)
	if err != nil {
		log.Println("limit : " + c.QueryParam("from"))
		return c.JSON(http.StatusConflict, model.UDFError{S: "error", Errmsg: "from parse error"})
	}
	to, err := strconv.ParseInt(c.QueryParam("to"), 0, 64)
	if err != nil {
		log.Println("limit : " + c.QueryParam("to"))
		return c.JSON(http.StatusConflict, model.UDFError{S: "error", Errmsg: "to parse error"})
	}

	klines := crawler.GetKlines(symbol, resolution, from*1000, to*1000)

	result := make([]model.Bar, len(*klines))
	for i, data := range *klines {
		result[i] = model.Bar{
			S: "ok",
			T: data.OpenTime / 1000,
			C: data.Close,
			O: data.Open,
			H: data.High,
			L: data.Low,
			V: data.Volume,
		}
	}

	return c.JSON(http.StatusOK, result)

	//symbol=ETHBTC&resolution=1D&from=1594992018&to=1595856018

}
