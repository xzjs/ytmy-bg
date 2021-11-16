package lib

import (
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v3"
)

type DB struct {
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	DBName string `yaml:"dbname"`
}

type AES struct {
	Key string `yaml:"key"`
}

type CookieConf struct {
	Name   string `yaml:"name"`
	Domain string `yaml:"domain"`
}

type Conf struct {
	DB     DB         `yaml:"db"`
	AES    AES        `yaml:"aes"`
	Cookie CookieConf `yaml:"cookie"`
}

var conf *Conf
var once sync.Once

func GetConfig() *Conf {
	once.Do(func() {
		f, err := ioutil.ReadFile("conf.yaml")
		if err != nil {
			log.Fatal(err)
		}
		conf = &Conf{}
		if err = yaml.Unmarshal(f, conf); err != nil {
			log.Fatal(err)
		}
	})
	return conf
}
