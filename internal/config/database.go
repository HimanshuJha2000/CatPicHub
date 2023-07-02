package config

import (
	"fmt"
	"os"
)

// Database : struct to hold live Database config
type Database struct {
	Dialect  string `toml:"dialect"`
	Protocol string `toml:"protocol"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Name     string `toml:"name"`
	Username string `env:"username"`
	Password string `env:"password"`
}

// PostgresqlConnectionDSNFormat : DNS for connecting postgresql
const PostgresqlConnectionDSNFormat = "%s://%s:%s@%s:%d/%s?sslmode=disable"

// URL : gives formatted postgresql url.
func (c Database) URL() string {

	c.Username = os.Getenv("DATABASE_USERNAME")
	c.Password = os.Getenv("DATABASE_PASSWORD")

	return fmt.Sprintf(
		PostgresqlConnectionDSNFormat,
		"postgresql",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Name)
}
