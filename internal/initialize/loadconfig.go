package initialize

import (
	"fmt"

	"github.com/TS0906/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config") //path config
	viper.SetConfigName("local")    //ten file
	viper.SetConfigType("yaml")     //loai file

	//read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config: %w", err))
	}
	// read server config
	fmt.Println("server port: ", viper.GetInt("server.port"))
	fmt.Println("server host: ", viper.GetString("security.jwt.key"))

	//config struct
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode config, %v", err)
	}

}
