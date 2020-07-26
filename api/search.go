package api

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"tradingview_udf_binance_go/crawler"
	"tradingview_udf_binance_go/model"

	"github.com/labstack/echo"
)

// func Filter(arr []model.SymbolInfo, f func(model.SymbolInfo) bool) []model.SymbolInfo {
// 	result := make([]model.SymbolInfo, 0)
// 	for _, v := range arr {
// 		if f(v) {
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }

// Filter Collection의 filter 역할
func Filter(arr *[]model.SymbolInfo, f func(model.SymbolInfo) bool) {
	result := make([]model.SymbolInfo, 0)
	for _, v := range *arr {
		if f(v) {
			result = append(result, v)
		}
	}
	*arr = result
}

// Search 트뷰 좌측 심볼 검색 api 지원
func Search(c echo.Context) error {
	query := c.QueryParam("query")
	ptype := c.QueryParam("type")
	exchange := c.QueryParam("exchange")
	limit := c.QueryParam("limit")

	symbols := crawler.GetExchangeInfo().Symbols
	if query != "" {
		Filter(&symbols, func(v model.SymbolInfo) bool {
			return strings.Contains(v.Symbol, query)
		})
	}
	if ptype != "" {
		// Filter(&symbols, func(v model.SymbolInfo) bool {
		// 	return v.type ==ptype
		// })
	}
	if exchange != "" {
		// Filter(&symbols, func(v model.SymbolInfo) bool {
		// 	return strings.Contains(v.Symbol, query)
		// })
	}
	if limit != "" {
		lm, err := strconv.Atoi(limit)
		if err != nil {
			log.Println("limit : " + limit)
			return c.JSON(http.StatusConflict, model.UDFError{S: "error", Errmsg: "limit parse error"})
		}
		return c.JSON(http.StatusOK, symbols[:lm])
	}

	return c.JSON(http.StatusOK, symbols)
	// result := make([]model.SymbolInfo)
	// for _, symbol := range symbols {
	// 	if strings.Contains(symbol.Symbol, query) {

	// 	}
	// 	result = append(result, symbol)
	// }
}
