package config

import "io/ioutil"

func Load(path string) (*Configure, error) {
	_, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

type Configure struct {
	DBConfig DBConfig `yaml:"db"`
}

// DBConfig InfluxDB相关配置, 见配置文件conf/server.yaml
type DBConfig struct {
	Org    string `yaml:"org"`
	Bucket string `yaml:"bucket"`
	Url    string `yaml:"url"`
	Token  string `yaml:"token"`
}
