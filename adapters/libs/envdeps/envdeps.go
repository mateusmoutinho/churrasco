package envdeps

import (
	"os"

	envdeps "github.com/mateusmoutinho/churrasco/sandbox/deps/envdeps"

	"github.com/mateusmoutinho/churrasco/sandbox/deps"
)

// Bind fills deps.Deps.Envdeps with the standard library's os.
func Bind(deps *deps.Deps) {
	deps.Envdeps = envdeps.Sandbox{
		Getenv: func(key string) string {
			return os.Getenv(key)
		},
	}
}
