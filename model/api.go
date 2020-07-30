package model

//Info 기본적인 정보를 갖는 구조체
type Info struct {
	Value string `json:"value"`
	Name  string `json:"name"`
	Desc  string `json:"desc ,omitempty"`
}

//Config jsapi config data
type Config struct {
	Exchanges              []Info   `json:"exchanges"`
	SymbolsTypes           []Info   `json:"symbols_types"`
	SupportedResolutions   []string `json:"supported_resolutions"`
	SupportsSearch         bool     `json:"supports_search"`
	SupportsGroupRequest   bool     `json:"supports_group_request"`
	SupportsMarks          bool     `json:"supports_marks"`
	SupportsTimescaleMarks bool     `json:"supports_timescale_marks"`
	SupportsTime           bool     `json:"supports_time"`
}

//Symbol symbols 출력 데이터
type Symbol struct {
	Symbol               string   `json:"symbol"`
	Ticker               string   `json:"ticker"`
	Name                 string   `json:"name"`
	FullName             string   `json:"full_name"`
	Description          string   `json:"description"`
	Exchange             string   `json:"exchange"`
	ListedExchange       string   `json:"listed_exchange"`
	Type                 string   `json:"type"`
	CurrencyCode         string   `json:"currency_code"`
	Session              string   `json:"session"`
	Timezone             string   `json:"timezone"`
	Minmovent            int64    `json:"minmovement"`
	Minmov               int64    `json:"minmov"`
	Minmovement2         int64    `json:"minmovement2"`
	Minmov2              int64    `json:"minmov2"`
	Pricescale           int64    `json:"pricescale"`
	SupportedResolutions []string `json:"supported_resolutions"`
	HasIntraday          bool     `json:"has_intraday"`
	HasDaily             bool     `json:"has_daily"`
	HasWeeklyAndMonthly  bool     `json:"has_weekly_and_monthly"`
	DataStatus           string   `json:"data_status"`
}

//Symbol historys 출력 데이터
type Bar struct {
	S      string  `json:"s"`
	Errmsg string  `json:"errnsg ,omitempty"`
	T      int64   `json:"t"`
	C      float64 `json:"c"`
	O      float64 `json:"o"`
	H      float64 `json:"h"`
	L      float64 `json:"l"`
	V      float64 `json:"v"`
}
