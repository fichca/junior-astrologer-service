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
	HTTP   *HTTP
	Client *Client
}

type DB struct {
	Postgre *Postgre
}

type Client struct {
	ApiKey      string
	APODBaseURL string
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
		HTTP:   parseAppHttpEnv(),
		Client: parseAppClientEnv(),
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

func parseAppClientEnv() *Client {
	cfg := Client{}
	cfg.ApiKey = os.Getenv("APOD_CLIENT_API_KEY")
	if cfg.ApiKey == "" {
		log.Fatal("missing APOD_CLIENT_API_KEY")
	}
	cfg.APODBaseURL = os.Getenv("APOD_CLIENT_BASE_URL")
	if cfg.APODBaseURL == "" {
		log.Fatal("missing APOD_CLIENT_BASE_URL")
	}

	return &cfg
}
