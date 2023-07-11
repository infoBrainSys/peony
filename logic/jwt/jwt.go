package jwt

import (
	jwtPkg "github.com/golang-jwt/jwt/v5"
	consts "peony/const"
	"peony/service"
	"peony/utils"
	"time"
)

var hs256 = jwtPkg.SigningMethodHS256

type sJWT struct {
	Token *jwtPkg.Token
}

type myClaims struct {
	Email string
	*jwtPkg.RegisteredClaims
}

func init() {
	service.RegisterJWT(New())
}

func New() service.IJWT {
	return &sJWT{}
}

// IssueToken 创建 Token
func (j *sJWT) IssueToken(email string) (tokenStr string, err error) {
	var claims = myClaims{
		Email: email,
		RegisteredClaims: &jwtPkg.RegisteredClaims{
			Issuer:    utils.V.GetString("jwt.iss"),
			ExpiresAt: jwtPkg.NewNumericDate(time.Now().Add(time.Hour * 2400)),
			Subject:   utils.V.GetString("jwt.sub"),
			IssuedAt:  jwtPkg.NewNumericDate(time.Now()),
		},
	}
	token := jwtPkg.NewWithClaims(hs256, claims)
	return token.SignedString([]byte(utils.V.GetString("jwt.secret")))
}

// AuthJwtToken 校验 Token
func (j *sJWT) AuthJwtToken(tokenStr string) error {
	// 裁剪掉 Bearer
	tokenStr = tokenStr[7:]

	token, err := jwtPkg.ParseWithClaims(tokenStr, &myClaims{}, func(token *jwtPkg.Token) (interface{}, error) {
		return []byte(utils.V.GetString("jwt.secret")), nil
	})
	if err != nil {
		return err
	}

	if err != nil {
		return consts.TokenIsNotValid
	}
	if !token.Valid {
		return consts.TokenIsNotValid
	}

	exp, _ := token.Claims.GetExpirationTime()
	// 判断过期时间是否早于当前时间，如果为 true 则 token 已过期
	if ok := exp.After(time.Now()); !ok {
		return consts.TokenIsExpired
	} else {
		return nil
	}
}

func (j *sJWT) GetEmail(tokenStr string) (email string, err error) {
	token, err := jwtPkg.ParseWithClaims(tokenStr, &myClaims{}, func(token *jwtPkg.Token) (interface{}, error) {
		return []byte(utils.V.GetString("jwt.secret")), nil
	})
	if err != nil {
		return "", consts.TokenIsNotValid
	}
	claims := token.Claims.(*myClaims)
	return claims.Email, nil
}
