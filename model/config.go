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
