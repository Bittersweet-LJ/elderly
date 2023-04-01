package controller

import "elderly/models"

// _ResponseSignUpList 注册接口响应数据
type _ResponseSignUpList struct {
	Code    ResCode `json:"code"` // 业务响应状态码
	Message string  `json:"msg"`  // 提示信息
}

// _ResponseLoginList 登录接口响应数据
type _ResponseLoginList struct {
	Code    ResCode     `json:"code"`           // 业务响应状态码
	Message string      `json:"msg"`            // 提示信息
	Data    []*ResLogin `json:"data,omitempty"` // 数据
}

// _ResponseDhtList 温湿度列表接口响应数据
type _ResponseDhtList struct {
	Code    ResCode       `json:"code"` // 业务响应状态码
	Message string        `json:"msg"`  // 提示信息
	Data    []*models.Dht `json:"data"` // 数据
}

// _ResponseDhtList 烟雾灯光列表接口响应数据
type _ResponseLedList struct {
	Code    ResCode       `json:"code"` // 业务响应状态码
	Message string        `json:"msg"`  // 提示信息
	Data    []*models.Led `json:"data"` // 数据
}

type ResLogin struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"user_name"`
	Token    string `json:"token"`
}
