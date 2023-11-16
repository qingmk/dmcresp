package response

type RespDataPayload map[string]interface{}
type AckBody struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data RespDataPayload `json:"data,omitempty"`
}
