package response

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Response(ctx context.Context, w http.ResponseWriter, code int, msg string, data RespData) {
	var body AckBody
	body.Code = code
	body.Data = data
	body.Msg = msg
	httpx.OkJson(w, body)
}

func Ok(ctx context.Context, w http.ResponseWriter) {
	var body AckBody
	body.Code = 0
	body.Msg = "成功"
	httpx.OkJsonCtx(ctx, w, body)
}

func OkWithMsg(ctx context.Context, w http.ResponseWriter, msg string) {
	var body AckBody
	body.Code = 0
	body.Msg = msg
	httpx.OkJsonCtx(ctx, w, body)
}

func OkWithData(ctx context.Context, w http.ResponseWriter, data RespData) {
	var body AckBody
	body.Code = 0
	body.Msg = "成功"
	body.Data = data
	httpx.OkJsonCtx(ctx, w, body)
}

func OkWithMsgData(ctx context.Context, w http.ResponseWriter, msg string, data RespData) {
	var body AckBody
	body.Code = 0
	body.Msg = msg
	body.Data = data
	httpx.OkJsonCtx(ctx, w, body)
}

func Failed(ctx context.Context, w http.ResponseWriter) {
	var body AckBody
	body.Code = -1
	body.Msg = "失败"
	httpx.OkJsonCtx(ctx, w, body)
}

func FailedWithMsg(ctx context.Context, w http.ResponseWriter, msg string) {
	var body AckBody
	body.Code = -1
	body.Msg = msg
	httpx.OkJsonCtx(ctx, w, body)
}

func FailedWithMsgData(ctx context.Context, w http.ResponseWriter, msg string, data RespData) {
	var body AckBody
	body.Code = 0
	body.Msg = msg
	body.Data = data
	httpx.OkJsonCtx(ctx, w, body)
}

func FailedWithCodeMsgData(ctx context.Context, w http.ResponseWriter, code int, msg string, data RespData) {
	var body AckBody
	body.Code = code
	body.Msg = msg
	body.Data = data
	httpx.OkJsonCtx(ctx, w, body)
}
