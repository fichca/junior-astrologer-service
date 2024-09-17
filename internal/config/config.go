package config

import (
	"log"
	"os"
)

type Config struct {
	App *App
	DB  *DB
}

type App struct {
	HTTP  *HTTP
	Admin *Admin
}

type Admin struct {
	Password string
}

type DB struct {
	Postgre *Postgre
}

type Postgre struct {
	URL      string
	Username string
	Password string
}

type HTTP struct {
	Port string
}

func (c *Config) Parse() {
	c.App = &App{
		HTTP: parseAppHttpEnv(),
	}
	c.DB = &DB{
		Postgre: parsePostgreEnv(),
	}
}

func parsePostgreEnv() *Postgre {
	cfg := Postgre{}
	cfg.URL = os.Getenv("POSTGRE_URL")
	if cfg.URL == "" {
		log.Fatal("missing POSTGRE_URL")
	}

	cfg.Username = os.Getenv("POSTGRE_USERNAME")
	if cfg.Username == "" {
		log.Fatal("missing POSTGRE_USERNAME")
	}

	cfg.Password = os.Getenv("POSTGRE_PASSWORD")
	if cfg.Password == "" {
		log.Fatal("missing POSTGRE_PASSWORD")
	}

	return &cfg
}

func parseAppHttpEnv() *HTTP {
	cfg := HTTP{}
	cfg.Port = os.Getenv("HTTP_PORT")
	if cfg.Port == "" {
		log.Printf("missing HTTP_PORT, using default 8080")
		cfg.Port = "8080"
	}

	return &cfg
}
