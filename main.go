package main

import (
	"net/http"
	"strconv"
	"time"
	"tradingview_udf_binance_go/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	//에코 인스턴스 생성
	e := echo.New()
	//미들웨어 선언
	e.Use(middleware.Logger())  //http 요청 기록
	e.Use(middleware.Recover()) //패닉 복구
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/time", func(c echo.Context) error {
		return c.String(http.StatusOK, strconv.FormatInt(time.Now().Unix(), 10))
	})

	e.GET("/config", api.GetConfig)
	e.GET("/symbols", api.Symbols)
	e.GET("/search", api.Search)
	e.GET("/history", api.History)
	// http://localhost:8081/
	e.Logger.Fatal(e.Start(":8081"))

}

/*

GET /config 304 - - 2.889 ms
GET /time 200 10 - 0.352 ms
GET /symbols?symbol=ETHBTC 304 - - 0.483 ms
GET /history?symbol=ETHBTC&resolution=1D&from=1561564601&to=1595779061 200 22649 - 62.088 ms



```js
GET /symbol_info?group=<group_name>

GET /history?symbol=<ticker_name>&from=<unix_timestamp>&to=<unix_timestamp>&resolution=<resolution>
GET /marks?symbol=<ticker_name>&from=<unix_timestamp>&to=<unix_timestamp>&resolution=<resolution>
GET /timescale_marks?symbol=<ticker_name>&from=<unix_timestamp>&to=<unix_timestamp>&resolution=<resolution>
GET /quotes?symbols=<ticker_name_1>,<ticker_name_2>,...,<ticker_name_n>
```
*/
