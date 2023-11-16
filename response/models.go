package response

type RespDataPayload interface{}
type AckBody struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data RespDataPayload `json:"data,omitempty"`
}
