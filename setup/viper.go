package setup

import (
	"fmt"

	"github.com/spf13/viper"
	"server/global"
)

func InitializeViper(path string)  *viper.Viper{
	vip := viper.New()
	vip.SetConfigName("config")
	vip.SetConfigType("yaml")
	vip.AddConfigPath(path)

	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := vip.Unmarshal(&global.GL_CONFIG); err != nil {
		panic(err)
	}

	fmt.Println("viper initialized")

	return vip
}