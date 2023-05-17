package vipers

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("./config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("vip init err :", err)
		return
	}
	return
}
