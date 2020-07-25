package main

import (
	"net/http"
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
	e.GET("/time", api.GetTime)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	// http://localhost:8081/
	e.Logger.Fatal(e.Start(":8081"))

}

/*
```js
GET /config
GET /symbol_info?group=<group_name>
GET /symbols?symbol=<symbol>
GET /search?query=<query>&type=<type>&exchange=<exchange>&limit=<limit>
GET /history?symbol=<ticker_name>&from=<unix_timestamp>&to=<unix_timestamp>&resolution=<resolution>
GET /marks?symbol=<ticker_name>&from=<unix_timestamp>&to=<unix_timestamp>&resolution=<resolution>
GET /timescale_marks?symbol=<ticker_name>&from=<unix_timestamp>&to=<unix_timestamp>&resolution=<resolution>
GET /time
GET /quotes?symbols=<ticker_name_1>,<ticker_name_2>,...,<ticker_name_n>
```
*/
