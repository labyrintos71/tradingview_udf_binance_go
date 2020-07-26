package model

const (
	SymbolNotFound = "심볼을 찾을 수 없습니다."
)

type UDFError struct {
	S      string `json:"s"`
	Errmsg string `json:"errmsg"`
}
