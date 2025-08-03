package configs

import (
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
	Host string
}

type DatabaseConfig struct {
	FilePath string
}

func LoadConfig() *Config {
	// Find the project root directory by looking for go.mod file
	projectRoot := findProjectRoot()

	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Host: getEnv("SERVER_HOST", "localhost"),
		},
		Database: DatabaseConfig{
			FilePath: getEnv("DATA_FILE_PATH", filepath.Join(projectRoot, "data", "Users.json")),
		},
	}
}

// findProjectRoot finds the directory containing go.mod file
func findProjectRoot() string {
	// Start from current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "."
	}

	// Walk up the directory tree to find go.mod
	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			// Reached root directory
			break
		}
		currentDir = parent
	}

	// Fallback to current directory
	return "."
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
