package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"database"`
}

func main() {
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
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode config, %v", err)
	}

	fmt.Println("Config Port: ", config.Server.Port)

	for _, db := range config.Database {
		fmt.Printf("database User: %s, Password: %s, Host: %s\n", db.User, db.Password, db.Host)
	}
}
