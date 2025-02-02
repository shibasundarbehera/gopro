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
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		currentDir = "."
	}

	// Try to find the goAPIs directory by looking for the data folder
	dataPath := filepath.Join(currentDir, "data", "Users.json")
	if _, err := os.Stat(dataPath); err != nil {
		// If not found, try going up one level
		parentDir := filepath.Dir(currentDir)
		dataPath = filepath.Join(parentDir, "data", "Users.json")
		if _, err := os.Stat(dataPath); err != nil {
			// If still not found, try the original goAPIs directory
			dataPath = "/Users/ssbehera/Documents/Misc-RnD/2025/gopro/goAPIs/data/Users.json"
		}
	}

	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Host: getEnv("SERVER_HOST", "localhost"),
		},
		Database: DatabaseConfig{
			FilePath: getEnv("DATA_FILE_PATH", dataPath),
		},
	}
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
