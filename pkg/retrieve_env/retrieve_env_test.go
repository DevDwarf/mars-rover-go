package retrieve_env_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-go/pkg/retrieve_env"
	"os"
	"testing"
)

var (
	envKey           = "PORT"
	envValue         = "4501"
	envFallbackValue = "4500"
)

func TestRetrieveEnv(t *testing.T) {
	t.Run("succeeds with a valid fallback value", func(t *testing.T) {
		err := os.Unsetenv(envKey)
		assert.NoError(t, err)

		env := retrieve_env.RetrieveEnv(envKey, envFallbackValue)

		assert.Equal(t, envFallbackValue, env)
	})

	t.Run("succeeds with a valid env value", func(t *testing.T) {
		err := os.Setenv(envKey, envValue)
		assert.NoError(t, err)

		env := retrieve_env.RetrieveEnv(envKey, envFallbackValue)

		assert.Equal(t, envValue, env)
	})
}
