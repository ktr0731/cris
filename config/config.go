package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Meta   *Meta
	Server *Server
	Logger *Logger
}

type Meta struct {
	Version string `default:"v1"`
}

type Server struct {
	Host string `default:""`
	Port string `default:"8080"`
}

type Logger struct {
	Output string `default:"stdout"`
	Prefix string `default:"[cris] "`
}

var config Config

func init() {
	err := envconfig.Process("cris", &config)
	if err != nil {
		panic(err)
	}
}

func Get() *Config {
	return &config
}
