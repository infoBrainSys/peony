package service

type (
	IJWT interface {
		IssueToken(email string) (tokenStr string, err error)
		AuthJwtToken(tokenStr string) error
		GetEmail(tokenStr string) (email string, err error)
	}
)

var (
	localIJWT IJWT
)

func JWT() IJWT {
	if localIJWT == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localIJWT
}

func RegisterJWT(i IJWT) {
	localIJWT = i
}
