package configs

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Postgres struct {
	Host              string `env:"DB_HOST" envDefault:"localhost"`
	Port              string `env:"DB_PORT" envDefault:"5432"`
	User              string `env:"DB_USER" envDefault:"user"`
	Password          string `env:"DB_PASSWORD" envDefault:"password"`
	DBName            string `env:"DB_NAME" envDefault:"speakbuddy"`
	MaxIdleConnection int    `env:"MAX_IDLE_CONNECTION" envDefault:"10"`
	MaxOpenConnection int    `env:"MAX_OPEN_CONNECTION" envDefault:"100"`
}

var PostgresSetting = &Postgres{}

func DatabaseInit() {
	if err := env.Parse(PostgresSetting); err != nil {
		log.Fatalf("[FAIL] Failed to parse PostgresSetting: %v", err)
	}
}
