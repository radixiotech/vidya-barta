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
	Secret string
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

func NewVBConfig() *VBApiConfig {
	return &VBApiConfig{
		DB: DB{
			MaxIdleConns: GetEnvInt("DB_MAX_IDLE_CONN", 5),
			MaxOpenConns: GetEnvInt("DB_MAX_OPEN_CONN", 20),
			User:         GetEnvString("DB_USER", "postgres"),
			Host:         GetEnvString("DB_HOST", "localhost"),
			Password:     GetEnvString("DB_PASSWORD", "password"),
			Name:         GetEnvString("DB_NAME", "vidyaBartaDB"),
			DisableTLS:   GetEnvString("DB_TLS", "disable") == "disable",
		},
		Web: Web{
			APIHost:         GetEnvString("WEB_API_HOST", "localhost:3000"),
			ReadTimeout:     time.Duration(GetEnvInt("WEB_READ_TIMEOUT", int(time.Second*5))),
			IdleTimeout:     time.Duration(GetEnvInt("WEB_IDLE_TIMEOUT", int(time.Second*120))),
			WriteTimeout:    time.Duration(GetEnvInt("WEB_WRITE_TIMEOUT", int(time.Second*10))),
			ShutdownTimeout: time.Duration(GetEnvInt("WEB_SHUTDOWN_TIMEOUT", int(time.Second*20))),
		},
		Auth: Auth{
			Issuer: GetEnvString("AUTH_ISSUER", "vidya-barta-backend"),
			Secret: GetEnvString(
				"AUTH_SECRET", "92c3ba3f929dc49a3468c0ff6b7340997c04d522af3a5216f43d71a3b5c97788c64484",
			),
		},
	}
}
