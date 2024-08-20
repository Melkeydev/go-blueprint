package flags

import (
	"fmt"
	"slices"
	"strings"
)

type Database string

// These are all the current databases supported. If you want to add one, you
// can simply copy and past a line here. Do not forget to also add it into the
// AllowedDBDrivers slice too!
const (
	MySql    Database = "mysql"
	Postgres Database = "postgres"
	Sqlite   Database = "sqlite"
	Mongo    Database = "mongo"
	Redis    Database = "redis"
	None     Database = "none"
)

var AllowedDBDrivers = []string{string(MySql), string(Postgres), string(Sqlite), string(Mongo), string(Redis), string(None)}

func (f Database) String() string {
	return string(f)
}

func (f *Database) Type() string {
	return "Database"
}

func (f *Database) Set(value string) error {
	if slices.Contains(AllowedDBDrivers, value) {
		*f = Database(value)
		return nil
	}

	return fmt.Errorf("Database to use. Allowed values: %s", strings.Join(AllowedDBDrivers, ", "))
}
