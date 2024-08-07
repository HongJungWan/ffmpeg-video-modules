package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func loadConfig() bool {
	_, err := os.Stat(file)
	if err != nil {
		return false
	}

	viper.SetConfigFile(file)
	viper.SetConfigType("toml")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(conf.DBHost)
	err = viper.GetViper().Unmarshal(&conf)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func parseConfig() bool {
	flag.StringVar(&file, "c", "config.toml", "config file")
	flag.Parse()
	if !loadConfig() {
		return false
	}

	return true
}
