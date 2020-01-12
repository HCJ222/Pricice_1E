package config

import (
	"fmt"
	"time"
	"go.uber.org/zap"
	"github.com/BurntSushi/toml"
)

const (
)

var (
	ConfigPath string
	Config	configType
)


type configType struct {

	Title	string	`json:"title"`

	APIs struct {
		Dunamu		string	`json:"dunamu"`

		Binance		string	`json:"binance"`
		Bithumb		string	`json:"bithumb"`
		Coinone		string	`json:"coinone"`
		HuobiGlobal	string	`json:"huobiGlobal"`
		Upbit		string	`json:"upbit"`

	}

	Options	struct {
		Interval	time.Duration	`json:"interval"`
		ListenPort	string		`json:"listenPort"`
	}
}


func Init(log *zap.Logger) string  {

	Config = readConfig(log)

	return Config.Options.ListenPort
}

func readConfig(log *zap.Logger) configType {

        var config configType

	// log 
        if _, err := toml.DecodeFile(ConfigPath +"/config.toml", &config); err != nil {
		// handle error
		log.Fatal("Config", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err)))
	} else {
		log.Info("Config", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Read", ConfigPath +"config.toml"))
	}

	return config

}
