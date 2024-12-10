package env

import (
	"os"
	"testing"
)

const (
	testInvalidNodeEnv = "invalid"
)

func TestValidationNodeEnv(t *testing.T) {
	t.Run("NODE_ENV=production", func(t *testing.T) {
		ok := validateNodeEnv(productionNodeEnv)
		if !ok {
			t.Errorf("expected %v, got %v", true, ok)
		}
	})

	t.Run("NODE_ENV=development", func(t *testing.T) {
		ok := validateNodeEnv(developmentNodeEnv)
		if !ok {
			t.Errorf("expected %v, got %v", true, ok)
		}
	})

	t.Run("NODE_ENV=local", func(t *testing.T) {
		ok := validateNodeEnv(localNodeEnv)
		if !ok {
			t.Errorf("expected %v, got %v", true, ok)
		}
	})

	t.Run("NODE_ENV=invalid", func(t *testing.T) {
		ok := validateNodeEnv(testInvalidNodeEnv)
		if ok {
			t.Errorf("expected %v, got %v", false, ok)
		}
	})
}

func TestGettingNodeEnv(t *testing.T) {
	t.Run("NODE_ENV=production", func(t *testing.T) {
		_ = os.Setenv(nodeEnvKey, productionNodeEnv)

		nodeEnv := getNodeEnv()

		if nodeEnv != productionNodeEnv {
			t.Errorf("expected %v, got %v", productionNodeEnv, nodeEnv)
		}
	})

	t.Run("NODE_ENV=development", func(t *testing.T) {
		_ = os.Setenv(nodeEnvKey, developmentNodeEnv)

		nodeEnv := getNodeEnv()

		if nodeEnv != developmentNodeEnv {
			t.Errorf("expected %v, got %v", developmentNodeEnv, nodeEnv)
		}
	})

	t.Run("NODE_ENV=local", func(t *testing.T) {
		_ = os.Setenv(nodeEnvKey, localNodeEnv)

		nodeEnv := getNodeEnv()

		if nodeEnv != localNodeEnv {
			t.Errorf("expected %v, got %v", localNodeEnv, nodeEnv)
		}
	})

	t.Run("NODE_ENV=invalid", func(t *testing.T) {
		_ = os.Setenv(nodeEnvKey, testInvalidNodeEnv)

		nodeEnv := getNodeEnv()

		if nodeEnv != productionNodeEnv {
			t.Errorf("expected %v, got %v", productionNodeEnv, nodeEnv)
		}
	})
}

func TestLoad(t *testing.T) {
	t.Cleanup(resetEnv)

	t.Run("Valid env file", func(t *testing.T) {
		_ = os.Setenv(nodeEnvKey, localNodeEnv)
		err := Load()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("Missing env file", func(t *testing.T) {
		_ = os.Setenv(nodeEnvKey, "nonexistent")
		err := Load()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func resetEnv() {
	_ = os.Unsetenv(nodeEnvKey)
}
