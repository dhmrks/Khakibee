package config

import (
	"log"

	"github.com/Tkanos/gonfig"
)

// Config :configuration data type
type Config struct {
	Conn        string `json:"conn"`
	Port        int    `json:"port"`
	Alloworigin string `json:"alloworigin"`
	Subpath     string `json:"subpath"`
	Websocket   string `json:"websocket"`
	Email       struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
	} `json:"email"`
}

//Conf : application config stored as global variable
var Conf *Config

func init() {
	Conf = readConfig()
}

func readConfig() *Config {

	conf := Config{}
	err := gonfig.GetConf("config.json", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Configuration loaded")

	return &conf
}
