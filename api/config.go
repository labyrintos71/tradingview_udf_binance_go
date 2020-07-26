package api

import (
	"net/http"
	"tradingview_udf_binance_go/model"

	"github.com/labstack/echo"
)

// GetConfig 시간을가져오는 메소드
func GetConfig(c echo.Context) error {
	config := &model.Config{
		SupportsSearch:         true,
		SupportsGroupRequest:   false,
		SupportsMarks:          false,
		SupportsTimescaleMarks: false,
		SupportsTime:           true,
		SupportedResolutions:   []string{"1", "3", "5", "15", "30", "60", "120", "240", "360", "480", "720", "1D", "3D", "1W", "1M"},
		Exchanges:              []model.Info{{Value: "Binance", Name: "Binance Exchange", Desc: "바이낸스"}},
		SymbolsTypes:           []model.Info{{Value: "BTC", Name: "비트코인"}, {Value: "usdt", Name: "테더코인"}},
	}

	return c.JSON(http.StatusOK, config)
}
