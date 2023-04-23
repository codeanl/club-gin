package e

const (
	SUCCESS       = 200
	InvalidParams = 201
	ERROR         = 400

	ErrorNotExistUser = 10003
	ErrorNotCompare   = 10004

	ErrorAuthCheckTokenFail    = 30001 //token 错误
	ErrorAuthCheckTokenTimeout = 30002 //token 过期
)
