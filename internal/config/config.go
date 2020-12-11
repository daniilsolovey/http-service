package config

import "github.com/kovetskiy/ko"

type Config struct {
	HTTPPort string `toml:"http_port" required:"true"`
}

func Load(path string) (*Config, error) {
	config := &Config{}
	err := ko.Load(path, config, ko.RequireFile(false))
	if err != nil {
		return nil, err
	}

	return config, nil
}
