package initialize

import (
	"awesomeProject/global"
	"fmt"
	viper2 "github.com/spf13/viper"
)

func loadConfig() {
	viper := viper2.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("config loi %w \n", err))

	}
	fmt.Println("server port:: ", viper.GetInt("server.port"))

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println("loi config %n", err)
	}

}
