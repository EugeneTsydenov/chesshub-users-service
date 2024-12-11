package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Cleanup(resetEnv)

	t.Run("successfully loads configuration from environment variables", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "testing")

		c, err := Load()

		assert.NoError(t, err)

		assert.NotNil(t, c)

		assert.Equal(t, "testing", c.Database.DatabaseUrl)
	})

	t.Run("returns error when required environment variable is missing", func(t *testing.T) {
		c, err := Load()

		assert.Error(t, err)
		assert.Equal(t, "", c.Database.DatabaseUrl)
	})
}

func resetEnv() {
	_ = os.Unsetenv("DATABASE_URL")
}
