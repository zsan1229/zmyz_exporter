package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

type NetWorkConfig struct {
	Name  string `mapstructure:"name"`
	Value string `mapstructure:"value"`
}

func ReadNetWorkConfig() []NetWorkConfig {
	viper.SetConfigName("network")
	viper.AddConfigPath("config/")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("read config error: %s \n", err))
	}
	var configs []NetWorkConfig
	err = viper.UnmarshalKey("servers", &configs)
	if err != nil {
		panic(fmt.Errorf("unmarshal config error: %s \n", err))
	}
	// fmt.Printf("%+v\n", configs)
	// fmt.Println(configs[1].Name)
	// fmt.Println(configs[1].Value)
	return configs
}
