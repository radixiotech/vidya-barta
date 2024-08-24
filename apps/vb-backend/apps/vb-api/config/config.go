package config

import (
	"time"
)

type web struct {
	APIHost         string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

type auth struct {
	Issuer string
}

type db struct {
	MaxIdleConns int
	MaxOpenConns int
	DisableTLS   bool
	Name         string
	User         string
	Host         string
	Password     string
}

type VBApiConfig struct {
	DB   db
	Web  web
	Auth auth
}
