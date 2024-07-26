package until

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	APPNAME string `json:"app_name"`
	APPMODE string `json:"app_mode"`
	APPHOST string `json:"app_host"`
	APPPORT string `json:"app_port"`
}

var _cfg *Config = nil

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
