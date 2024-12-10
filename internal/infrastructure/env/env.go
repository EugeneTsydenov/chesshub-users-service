package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

const (
	nodeEnvKey         = "NODE_ENV"
	localNodeEnv       = "local"
	developmentNodeEnv = "development"
	productionNodeEnv  = "production"
)

var allowedNodeEnvs = map[string]string{
	localNodeEnv:       "",
	developmentNodeEnv: "",
	productionNodeEnv:  "",
}

func Load() error {
	nodeEnv := getNodeEnv()
	envFileName := getEnvFileName(nodeEnv)

	if err := godotenv.Load(envFileName); err != nil {
		return fmt.Errorf("environment file '%s' not found, continuing with default environment variables: %w", envFileName, err)
	}

	return nil
}

func getNodeEnv() string {
	nodeEnv := os.Getenv(nodeEnvKey)
	if nodeEnv == "" || !validateNodeEnv(nodeEnv) {
		return getDefaultNodeEnv()
	}

	return nodeEnv
}

func getDefaultNodeEnv() string {
	return productionNodeEnv
}

func validateNodeEnv(nodeEnv string) bool {
	_, ok := allowedNodeEnvs[nodeEnv]
	return ok
}

func getEnvFileName(nodeEnv string) string {
	return fmt.Sprintf("C:/Important/For-work/Programming/Backend-development/current-project/chesshub/services/users-service/.env.%s", nodeEnv)
}
