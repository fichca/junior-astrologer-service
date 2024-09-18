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
	Minio   *Minio
}

type Client struct {
	ApiKey      string
	APODBaseURL string
}

type Postgre struct {
	Driver   string
	User     string
	Password string
	Name     string
	SSLMode  string
	Host     string
	Port     string
}
type Minio struct {
	Endpoint  string
	KeyID     string
	SecretKey string
	Bucket    string
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
		Minio:   parseMinioEnv(),
	}
}

func parseMinioEnv() *Minio {
	cfg := Minio{}
	cfg.KeyID = os.Getenv("MINIO_KEYID")
	if cfg.KeyID == "" {
		log.Fatal("missing MINIO_KEYID")
	}

	cfg.Endpoint = os.Getenv("MINIO_ENDPOINT")
	if cfg.Endpoint == "" {
		log.Fatal("missing MINIO_ENDPOINT")
	}

	cfg.SecretKey = os.Getenv("MINIO_SECRET_KEY")
	if cfg.SecretKey == "" {
		log.Fatal("missing MINIO_SECRET_KEY")
	}
	cfg.Bucket = os.Getenv("MINIO_BUCKET")
	if cfg.Bucket == "" {
		log.Fatal("missing MINIO_BUCKET")
	}
	return &cfg
}

func parsePostgreEnv() *Postgre {
	cfg := Postgre{}
	cfg.Driver = os.Getenv("POSTGRE_DRIVER")
	if cfg.Driver == "" {
		log.Fatal("missing POSTGRE_DRIVER")
	}

	cfg.User = os.Getenv("POSTGRE_USER")
	if cfg.User == "" {
		log.Fatal("missing POSTGRE_USER")
	}

	cfg.Password = os.Getenv("POSTGRE_PASSWORD")
	if cfg.Password == "" {
		log.Fatal("missing POSTGRE_PASSWORD")
	}
	cfg.Name = os.Getenv("POSTGRE_NAME")
	if cfg.Name == "" {
		log.Fatal("missing POSTGRE_NAME")
	}
	cfg.SSLMode = os.Getenv("POSTGRE_SSL_MODE")
	if cfg.SSLMode == "" {
		log.Fatal("missing POSTGRE_SSL_MODE")
	}
	cfg.Host = os.Getenv("POSTGRE_HOST")
	if cfg.Host == "" {
		log.Fatal("missing POSTGRE_HOST")
	}
	cfg.Port = os.Getenv("POSTGRE_PORT")
	if cfg.Port == "" {
		log.Fatal("missing POSTGRE_PORT")
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
