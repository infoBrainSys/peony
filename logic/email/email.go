package email

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"math/rand"
	"peony/service"
	"peony/utils"
	"strings"
)

type sEmail struct {
	Code []string
}

func init() {
	service.RegisterEmail(New())
}

func New() service.IEmail {
	return &sEmail{}
}

func (e *sEmail) generateCode() *sEmail {
	// 大小写+数字验证码："ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	codeStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 大写+数字组合
	code := make([]string, 4)
	for i := 0; i < 4; i++ {
		index := rand.Intn(len(codeStr))
		code[i] = string(codeStr[index])
	}
	e.Code = code
	return e
}

func (e *sEmail) ParseFiles() (code string, body string, err error) {
	tpl, err := template.ParseFiles("utils/email/mail.html")
	if err != nil {
		fmt.Println(err)
	}
	var buf bytes.Buffer

	err = tpl.Execute(&buf, gin.H{
		"name": utils.V.GetString("app.name"),
		"code": e.generateCode().Code,
	})
	if err != nil {
		return "", "", err
	}
	c := strings.Join(e.Code, "")
	return c, buf.String(), nil
}
