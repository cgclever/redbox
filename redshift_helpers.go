package redbox

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Postgres driver
)

// defaultConnectionTimeout is the default timeout, in seconds, for attempting to connect to Redshift
const defaultConnectionTimeout = 300

// Transaction interfaces a SQL transaction
type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Commit() error
	Rollback() error
}

// Redshift represents an interface over the SQL methods
type Redshift interface {
	Close() error
	Begin() (Transaction, error)
}

// RedshiftConfiguration interfaces a configuration for creating a Redshift object
type RedshiftConfiguration interface {
	RedshiftConnection() (Redshift, error)
}

// RedshiftConfig specifies the connection to a Redshift Database
type PostgresRedshiftConfiguration struct {
	Host              string
	Port              string
	User              string
	Password          string
	Database          string
	ConnectionTimeout int
}

// RedshiftConnection returns a direct redshift connection
func (rc *PostgresRedshiftConfiguration) RedshiftConnection() (*sql.DB, error) {
	connectionTimeout := defaultConnectionTimeout
	if rc.ConnectionTimeout > 0 {
		connectionTimeout = rc.ConnectionTimeout
	}

	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s connect_timeout=%d",
		rc.Host, rc.Port, rc.Database, rc.User, rc.Password, connectionTimeout)

	return sql.Open("postgres", connectionString)
}