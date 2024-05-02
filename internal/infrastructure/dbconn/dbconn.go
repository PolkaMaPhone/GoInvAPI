package dbconn

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"path/filepath"
)

type Config struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
	DbSchema   string
}

func LoadConfigFile() (Config, error) {
	var config Config

	// Get the root directory from the environment variable
	rootDir := os.Getenv("PROJECT_ROOT")
	if rootDir == "" {
		return config, fmt.Errorf("PROJECT_ROOT environment variable is not set")
	}

	// Check if we are in a test environment
	testEnv := os.Getenv("TEST_ENV")

	// Construct the path to the config.json file
	var configPath string
	if testEnv == "true" {
		configPath = filepath.Join(rootDir, "config.json.sample")
	} else {
		configPath = filepath.Join(rootDir, "config.json")
	}

	file, err := os.Open(configPath)
	if err != nil {
		return config, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			middleware.WarningLogger.Printf("Failed to close the file: %v", err)
		}
	}(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

type DB interface {
	Connect(connectionString string) error
}

type PgxDB struct {
	Pool *pgxpool.Pool
}

func (db *PgxDB) Connect(connectionString string) error {
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return fmt.Errorf("unable to create connection pool: %v", err)
	}

	db.Pool = pool
	return nil
}

func New(config Config, db DB) (*DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName, config.DbSchema)

	err := db.Connect(connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &db, nil
}
