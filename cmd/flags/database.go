package flags

import (
	"fmt"
	"strings"
)

type Database string

// These are all the current databases supported. If you want to add one, you
// can simply copy and paste a line here. Do not forget to also add it into the
// AllowedDBDrivers slice too!
const (
	MySql    Database = "mysql"
	Postgres Database = "postgres"
	Sqlite   Database = "sqlite"
	Mongo    Database = "mongo"
	Redis    Database = "redis"
	Scylla   Database = "scylla"
	None     Database = "none"
)

var AllowedDBDrivers = []string{string(MySql), string(Postgres), string(Sqlite), string(Mongo), string(Redis), string(Scylla), string(None)}

func (f Database) String() string {
	return string(f)
}

func (f *Database) Type() string {
	return "Database"
}

func (f *Database) Set(value string) error {
	// Contains isn't available in 1.20 yet
	// if AllowedDBDrivers.Contains(value) {
	for _, database := range AllowedDBDrivers {
		if database == value {
			*f = Database(value)
			return nil
		}
	}

	return fmt.Errorf("Database to use. Allowed values: %s", strings.Join(AllowedDBDrivers, ", "))
}
