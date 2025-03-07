package config

import (
	"time"
)

type Config struct {
	SMTPHost       string
	SMTPPort       int
	SMTPUser       string
	SMTPPassword   string
	CodeExpiration time.Duration
}

func GetConfig() *Config {
	return &Config{
		SMTPHost:       "smtp.gmail.com",
		SMTPPort:       587,
		SMTPUser:       "sengsmtp@gmail.com",
		SMTPPassword:   "hqdi nila ljbd swlz",
		CodeExpiration: time.Minute * 1,
	}
}
