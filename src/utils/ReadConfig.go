package utils

import (
	"io/ioutil"
	"encoding/json"
)

/**
	读取数据库的配置文件
 */
func ReadBaseConfig(path string) (*Config, error) {
	res, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config

	err = json.Unmarshal(res, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

type Config struct {
	Url0 string
	Url1 string
	Url3 string
}
