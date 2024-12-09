package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	nodeEnvKey     = "NODE_ENV"
	defaultNodeEnv = "production"
)

func Load() {
	nodeEnv := getNodeEnv()
	envFileName := getEnvFileName(nodeEnv)

	if err := godotenv.Load(envFileName); err != nil {
		log.Printf("Environment file '%s' not found, continuing with default environment variables.", envFileName)
	}
}

func Get(key string) string {
	return os.Getenv(key)
}

func Set(key, value string) error {
	err := os.Setenv(key, value)
	if err != nil {
		return fmt.Errorf("error setting environment variable %s: %w", key, err)
	}
	return nil
}

func getNodeEnv() string {
	nodeEnv := os.Getenv(nodeEnvKey)

	if nodeEnv == "" {
		setDefaultNodeEnv()
		nodeEnv = defaultNodeEnv
	}

	return nodeEnv
}

func setDefaultNodeEnv() {
	if err := os.Setenv(nodeEnvKey, defaultNodeEnv); err != nil {
		log.Fatalf("Error setting NODE_ENV to %s: %v", defaultNodeEnv, err)
	}

	log.Printf("%s set to %s successfully.", nodeEnvKey, defaultNodeEnv)
}

func getEnvFileName(nodeEnv string) string {
	return fmt.Sprintf(".%s.env", nodeEnv)
}
