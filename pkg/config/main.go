package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var (
	k      = koanf.New(".")
	parser = json.Parser()
)

type Config struct {
	Lemmy
	Server
}

type Server struct {
	BindAddr string `validate:"required,hostname_port"`
}

type Lemmy struct {
	Username string `validate:"required"`
	Url      string `validate:"required,url"`
	Password string `validate:"required"`
}

func New() (*Config, error) {
	validate := validator.New()

	var conf Config
	if err := k.Load(file.Provider("config.json"), parser); err != nil {
		return &conf, err
	}

	err := k.Unmarshal("", &conf)
	if err != nil {
		return &conf, err
	}

	err = validate.Struct(conf)
	return &conf, err
}
