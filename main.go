package main

import (
	"github.com/dytlan/moonlay-todo-list/server"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error load config : " + err.Error())
	}
	if err := server.Run(); err != nil {
		log.Fatal("Running Failed : " + err.Error())
	}
}
