package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/programacao-fortificada/upsec/pkg/log"
	"github.com/spf13/viper"
	"os"
	"time"
)

type DB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Server struct {
	Port         string        `json:"port"`
	IdleTimeout  time.Duration `json:"idle_timeout"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
}

type App struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	LogLevel string `json:"log_level"`
	Location string `json:"location"`
	Env      string `json:"env"`
}

type Config struct {
	App    *App    `json:"app"`
	Server *Server `json:"server"`
	DB     *DB     `json:"db"`
}

func Get() (*Config, error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("json")

	v.AutomaticEnv()

	/*
		BindEnvs
	*/

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func Log(l log.Log) error {
	ll, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		ll = "debug"
	}

	lv := l.Level()
	err := lv.Set(ll)
	if err != nil {
		return err
	}

	return nil
}

func Local() {
	l, ok := os.LookupEnv("LOCATION")
	if !ok {
		l = "America/Sao_Paulo"
	}

	time.Local, _ = time.LoadLocation(l)
}
