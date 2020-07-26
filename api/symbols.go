package api

import (
	"log"
	"net/http"
	"strconv"
	"tradingview_udf_binance_go/crawler"
	"tradingview_udf_binance_go/model"

	"github.com/labstack/echo"
)

func findFilter(filters *[]model.SymbolFilter) string {
	for _, f := range *filters {
		if f.FilterType == "PRICE_FILTER" {
			return f.TickSize
		}
	}
	return ""
}

// Symbols binance의 심볼을 조회해서 가져다줌
func Symbols(c echo.Context) error {
	symbol := c.QueryParam("symbol")
	if symbol == "" {
		return c.JSON(http.StatusBadRequest, model.UDFError{S: "error", Errmsg: "symbol" + model.ParameterNotFound})
	}
	exchangeInfo := crawler.GetExchangeInfo()
	for _, symbolInfo := range exchangeInfo.Symbols {
		if symbolInfo.Symbol == symbol {
			ticksize := findFilter(&symbolInfo.Filters)

			pscale, err := strconv.ParseFloat(ticksize, 64)
			if err != nil {
				log.Println("Symbols : " + symbolInfo.Symbol)
				log.Println("TickSize : " + ticksize)
				return c.JSON(http.StatusConflict, model.UDFError{S: "error", Errmsg: "ticksize calc ERROR"})
			}

			a := &model.Symbol{
				Symbol:               symbolInfo.Symbol,
				Ticker:               symbolInfo.Symbol,
				Name:                 symbolInfo.BaseAsset + " Coin",
				FullName:             "[" + symbolInfo.QuoteAsset + "]" + symbolInfo.BaseAsset + " Coin",
				Description:          symbolInfo.BaseAsset + " / " + symbolInfo.QuoteAsset,
				Exchange:             "BINANCE",
				ListedExchange:       "BINANCE",
				Type:                 "crypto",
				CurrencyCode:         symbolInfo.QuoteAsset,
				Session:              "24x7",
				Timezone:             "UTC",
				Minmovent:            1,
				Minmov:               1,
				Minmovement2:         0,
				Minmov2:              0,
				Pricescale:           int64(1 / pscale),
				SupportedResolutions: []string{"1", "3", "5", "15", "30", "60", "120", "240", "360", "480", "720", "1D", "3D", "1W", "1M"},
				HasIntraday:          true,
				HasDaily:             true,
				HasWeeklyAndMonthly:  true,
				DataStatus:           "streaming",
			}
			return c.JSON(http.StatusOK, a)
		}

	}
	return c.JSON(http.StatusOK, model.UDFError{S: "error", Errmsg: model.SymbolNotFound})
}
