package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type _HttpRspJsonCommonMessage struct {
	Code string
	Msg  string
	Data interface{}
}

type _HttpRspJsonCommon struct {
	Result  string
	Message _HttpRspJsonCommonMessage
}

const (
	HTTP_RSP_JSON_RESULT_SUCCESS = "SUCCESS"
	HTTP_RSP_JSON_RESULT_ERROR = "ERROR"
)

/**
 * 以 JSON 格式返回 http 响应数据
 *
 * @w http.ResponseWriter
 * @httpcode int http响应码
 * @result string 响应结果: "SUCCESS" 或 "ERROR"
 * @code string 操作响应 code,用于区分在 result 为 "ERROR" 时的各种不同情况, result 为 "SUCCESS" 时固定传 "0000"
 * @msg string 操作结果说明
 * @data interface{} 返回的数据
 */
func HttpRspJson(w http.ResponseWriter, httpcode int, result, code, msg string, data interface{}) {
	w.WriteHeader(httpcode)

	if data == nil {
		data = make([]string, 0)
	}
	var resp _HttpRspJsonCommon
	resp.Result = result
	resp.Message = _HttpRspJsonCommonMessage{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	respdata, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintf(w, "{%q:%q, %q:{%q:%q}}", "Result", "Error", "Message", "Msg", "json error "+err.Error())
	} else {
		w.Write(respdata)
	}
}

/**
 * 以 JSON 格式返回 http 响应数据, 生成JSON时不转义 <、>、& 特殊字符
 *
 * @w http.ResponseWriter
 * @httpcode int http响应码
 * @result string 响应结果: "SUCCESS" 或 "ERROR"
 * @code string 操作响应 code,用于区分在 result 为 "ERROR" 时的各种不同情况, result 为 "SUCCESS" 时固定传 "0000"
 * @msg string 操作结果说明
 * @data interface{} 返回的数据
 */
func HttpRspJsonEscapeHTML(w http.ResponseWriter, httpcode int, result, code, msg string, data interface{}) {
	w.WriteHeader(httpcode)

	if data == nil {
		data = make([]string, 0)
	}
	var resp _HttpRspJsonCommon
	resp.Result = result
	resp.Message = _HttpRspJsonCommonMessage{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	respMessage := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(respMessage)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(resp.Message)

	respResult := bytes.NewBuffer([]byte{})
	jsonEncoderResult := json.NewEncoder(respResult)
	jsonEncoderResult.SetEscapeHTML(false)
	err := jsonEncoderResult.Encode(resp)
	if err != nil {
		fmt.Fprintf(w, "{%q:%q, %q:{%q:%q}}", "Result", "Error", "Message", "Msg", "json error "+err.Error())
	} else {
		w.Write(respResult.Bytes())
	}
}
