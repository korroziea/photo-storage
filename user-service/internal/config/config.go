package config

type Config struct {
	HTTPPort string `env:"HTTP_PORT, default=3000"`
	
	DB Postgres
}

type Postgres struct {
	
}
