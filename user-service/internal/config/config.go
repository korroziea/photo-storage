package config

import "fmt"

type Config struct {
	HTTPPort int `env:"HTTP_PORT, default=3000"`

	Hashing Hashing

	DB Postgres
}

type Hashing struct {
	Memory      uint32 `env:"HASHING_MEMORY, required"`
	Iterations  uint32 `env:"HASHING_ITERATIONS, required"`
	Parallelism uint8  `env:"HASHING_PARALLELISM, required"`
	SaltLength  uint32 `env:"HASHING_SALT_LEN, required"`
	KeyLength   uint32 `env:"HASHING_KEY_LEN, required"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     int    `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DATABASE"`
	SSLMode  string `env:"SSLMODE, default=disable"`
}

func (p *Postgres) PostgresURL() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.Host,
		p.Port,
		p.User,
		p.Password,
		p.Database,
		p.SSLMode,
	)
}
