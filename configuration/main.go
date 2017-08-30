package configuration

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/pawski/go-xchange/logger"
	"sync"
)

type config struct {
	InfluxDbHost    string `yaml:"influx_host"`
	WalutomatUrl   string `yaml:"walutomat_url"`
	InfluxDbDatabase string `yaml:"influx_database"`
}

var cfg config
var once sync.Once

func Get() (config)  {
	once.Do(func() {
		cfg = loadConfiguration()
	})

	return cfg
}

func loadConfiguration() (config)  {
	source, err := ioutil.ReadFile("config.yml")

	if err != nil {
		logger.Get().Fatalf("error: %v", err)
	}

	var config config

	err = yaml.Unmarshal([]byte(source), &config)

	if err != nil {
		panic(err)
	}

	return config
}