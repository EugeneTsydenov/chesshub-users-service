package config

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Cleanup(resetEnv)

	t.Run("successfully loads configuration from environment variables", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "testing")

		c, err := Load(context.Background())

		assert.NoError(t, err)

		assert.NotNil(t, c)

		assert.Equal(t, "testing", c.Database.URL)
	})

	t.Run("returns error when required environment variable is missing", func(t *testing.T) {
		c, err := Load(context.Background())

		assert.Error(t, err)
		assert.Nil(t, c)
	})
}

func resetEnv() {
	_ = os.Unsetenv("DATABASE_URL")
}
