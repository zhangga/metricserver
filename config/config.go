package config

import (
	"github.com/zhangga/logs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Load(path string) (*Configure, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		logs.Errorf("Load server config file err: %v", err)
		return nil, err
	}
	config := &Configure{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		logs.Errorf("Unmarshal server config file err: %v", err)
		return nil, err
	}
	return config, nil
}

type Configure struct {
	ServerConfig ServerConfig `yaml:"Server"`
	DBConfig     DBConfig     `yaml:"DB"`
}

// ServerConfig 服务相关配置
type ServerConfig struct {
	Address string `yaml:"Address"` // eg: ":port", "ip:port"
}

// DBConfig InfluxDB相关配置, 见配置文件conf/server.yaml
type DBConfig struct {
	Org    string `yaml:"Org"`
	Bucket string `yaml:"Bucket"`
	Url    string `yaml:"Url"`
	Token  string `yaml:"Token"`
}
