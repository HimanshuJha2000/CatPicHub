package config

import (
	"fmt"
	"os"
)

// Database : struct to hold live / test Database config
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
	// charset=utf8: uses utf8 character set data format
	// parseTime=true: changes the output type of DATE and DATETIME values to time.Time instead of []byte / strings
	// loc=Local: Sets the location for time.Time values (when using parseTime=true). "Local" sets the system's location
	return fmt.Sprintf(
		PostgresqlConnectionDSNFormat,
		"postgresql",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Name)
}
