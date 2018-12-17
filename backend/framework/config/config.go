package config

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	TemplateDir string	`json:"templateDir"`
	StaticDir string `json:"staticDir"`
	HttpPort int `json:"httpPort"`
	HttpReadTimeout int `json:"httpReadTimeout"`
	HttpWriteTimeout int `json:"httpWriteTimeout"`
}

var (
	G_Conf *Config
)

// 加载配置
func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf Config
	)

	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}
	G_Conf = &conf
	return
}
