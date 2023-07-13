package service

type IEmail interface {
	ParseFiles() (code string, body string, err error)
}

var (
	localIEmail IEmail
)

func Email() IEmail {
	if localIEmail == nil {
		panic("implement not found for interface IEmail, forgot register?")
	}
	return localIEmail
}
func RegisterEmail(i IEmail) {
	localIEmail = i
}
