package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"time"
)

// Config holds the database configuration settings.
type Config struct {
	Host            string
	Port            string
	User            string
	Password        string
	Database        string
	SSLMode         string
	Driver          string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime time.Duration
}

type Conn struct {
	*sql.DB
}

// NewConn initializes a new database connection with the provided configuration.
func NewConn(c Config) (*Conn, error) {
	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode)

	db, err := sql.Open(c.Driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if c.MaxOpenConn > 0 {
		db.SetMaxOpenConns(c.MaxOpenConn)
	}
	if c.MaxIdleConn > 0 {
		db.SetMaxIdleConns(c.MaxIdleConn)
	}
	if c.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(c.ConnMaxLifetime)
	}

	log.Printf("Database connection established with host: %s, database: %s", c.Host, c.Database)
	return &Conn{db}, nil
}

func (c *Conn) HealthCheck() error {
	return c.Ping()
}

func (c *Conn) Close() error {
	log.Println("Closing database connection")
	return c.DB.Close()
}

var instance *Conn
var once sync.Once

func GetInstance(c Config) (*Conn, error) {
	var err error
	once.Do(func() {
		instance, err = NewConn(c)
	})
	if instance == nil {
		return nil, fmt.Errorf("failed to initialize database instance: %w", err)
	}
	return instance, nil
}

func (c *Config) Validate() error {
	if c.Host == "" || c.Port == "" || c.User == "" || c.Password == "" || c.Database == "" || c.SSLMode == "" || c.Driver == "" {
		return fmt.Errorf("missing required database configuration")
	}
	return nil
}
