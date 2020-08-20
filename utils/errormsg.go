package utils
const (
	SUCCESS = 200
	ERROR = 500
	REPEAT_NAME = 501
	USER_NOT_EXIST = 502
	REPEAT_CATE = 503
	POST_NOT_FOUND=504
	TOKEN_ERROR=505
	PASSWORD_ERROR=506
)
var ErrorType = map[int]string{
	SUCCESS: "SUCCESS",
	ERROR: "FAIL",
	REPEAT_NAME:"用户名重复",
	USER_NOT_EXIST: "用户不存在",
	REPEAT_CATE:"文章分类重复",
	POST_NOT_FOUND:"文章不存在",
	TOKEN_ERROR:"token错误",
	PASSWORD_ERROR:"密码错误",
}

func ErrorMsg(errorCode int)string  {
	return ErrorType[errorCode]
}