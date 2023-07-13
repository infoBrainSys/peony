package consts

import "errors"

// 注册/登录
const (
	RegisterFailed   = "注册失败"
	RegisterSuccess  = "注册成功"
	LoginFailed      = "登录失败，账号或密码错误"
	LoginSuccess     = "登录成功"
	UserNotExist     = "用户不存在"
	PasswordNotMatch = "两次输入的密码不一致"
	VerifyCodeWrong  = "验证码错误"
)

// JWT
var (
	TokenIsNotValid = errors.New("token 无效")
	TokenIsExpired  = errors.New("token 已过期")
)

// 响应码
const (
	Success = iota
	Failed
)

// 帖子状态
const (
	PublishSuccess = "发布成功"
	PublishFailed  = "发布失败"
)

// 邮件
const (
	SendEmailSuccess = "发送邮件成功"
	SendEmailFailed  = "发送邮件失败"
	EmailIsNotValid  = "邮箱不合法"
)
