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
	// http://localhost:8081/
	e.Logger.Fatal(e.Start(":8081"))

}

/*
```js
GET /symbol_info?group=<group_name>
GET /symbols?symbol=<symbol>
GET /search?query=<query>&type=<type>&exchange=<exchange>&limit=<limit>
GET /history?symbol=<ticker_name>&from=<unix_timestamp>&to=<unix_timestamp>&resolution=<resolution>
GET /marks?symbol=<ticker_name>&from=<unix_timestamp>&to=<unix_timestamp>&resolution=<resolution>
GET /timescale_marks?symbol=<ticker_name>&from=<unix_timestamp>&to=<unix_timestamp>&resolution=<resolution>
GET /quotes?symbols=<ticker_name_1>,<ticker_name_2>,...,<ticker_name_n>
```
*/
