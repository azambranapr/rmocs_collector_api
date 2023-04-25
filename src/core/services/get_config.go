package services

import "github.com/spf13/viper"

type DbConfig struct {
	Password    string       `json:"password"`
	User        string       `json:"user"`
	Database    string       `json:"database"`
	Collections []Collection `json:"collections"`
}

type Collection struct {
	Name string `json:"name"`
}

type ServerConfig struct {
	Drive   string `json:"drive"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type Config struct {
	Db     DbConfig     `json:"db"`
	Server ServerConfig `json:"server"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	vp = viper.New()
	var config Config

	vp.SetConfigName("config_mongo")
	vp.SetConfigType("json")
	vp.AddConfigPath("./config")
	err := vp.ReadInConfig()

	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
