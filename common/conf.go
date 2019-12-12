package common

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

/*
Conf 配置文件

*/
type Conf struct {
	Host      string `yaml:"host"`
	User      string `yaml:"user"`
	Pwd       string `yaml:"pwd"`
	Dbname    string `yaml:"dbname"`
	JwtSecret string `yaml:"jwtsecret"`
}

/*
GetConf 获取配置文件

*/
func (c *Conf) GetConf() *Conf {
	yamlFile, err := ioutil.ReadFile("conf/conf.prod.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
