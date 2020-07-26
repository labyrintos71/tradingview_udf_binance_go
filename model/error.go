package model

const (
	SymbolNotFound    = "심볼을 찾을 수 없습니다."
	ParameterNotFound = "파라미터는 필수 파라미터 입니다."
)

type UDFError struct {
	S      string `json:"s"`
	Errmsg string `json:"errmsg"`
}
