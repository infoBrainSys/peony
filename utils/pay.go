package utils

import (
	"github.com/go-pay/gopay/alipay"
	"os"
)

var AlipayClient *alipay.Client

func InitPay() {
	privateKey, err := os.ReadFile("manifest/cert/privateKey.pem")
	if err != nil {
		panic(err)
	}
	client, err := alipay.NewClient(
		V.GetString("pay.alipay.app_id"),
		string(privateKey),
		false,
	)
	if err != nil {
		return
	}
	AlipayClient = client
}
