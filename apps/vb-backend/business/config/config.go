package config

import (
	"time"
)

type Web struct {
	APIHost         string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

type Auth struct {
	Issuer string
}

type DB struct {
	MaxIdleConns int
	MaxOpenConns int
	DisableTLS   bool
	Name         string
	User         string
	Host         string
	Password     string
}

type VBApiConfig struct {
	DB   DB
	Web  Web
	Auth Auth
}

func New() *VBApiConfig {
	return &VBApiConfig{}
}
