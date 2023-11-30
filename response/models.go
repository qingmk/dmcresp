package response

type RespData interface{}
type RespDataJson map[string]interface{}
type AckBody struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data RespData `json:"data,omitempty"`
}
