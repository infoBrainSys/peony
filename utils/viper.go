package utils

import "github.com/spf13/viper"

var V *viper.Viper

func InitViper() {
	V = viper.New()
	V.SetConfigType("yml")
	V.SetConfigName("config")
	V.AddConfigPath("manifest/config")
	err := V.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
