package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// giúp đọc qua nhiều db
type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DbName   string `mapstructure:"dbName"`
	} `mapstructure:"database"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/") // path to config file
	viper.SetConfigName("local")      // name file
	viper.SetConfigType("yaml")       // type file

	// read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// read server
	fmt.Println("server.port: ", viper.GetInt("server.port")) // phía file local.yaml có server.port: 8080
	fmt.Println("security.jwt.key: ", viper.GetString("security.jwt.key"))

	// config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("error unmarshal configuration: ", err)
	}
	fmt.Println("Config port: ", config.Server.Port)

	for _, db := range config.Database {
		fmt.Println("db: ", db.User, db.Password, db.Host, db.DbName)
	}
}
