package sandbox

import (
	api "github.com/mateusmoutinho/churrasco/sandbox/api"
	cli "github.com/mateusmoutinho/churrasco/sandbox/constructors/cli"
	config "github.com/mateusmoutinho/churrasco/sandbox/constructors/config"
	server "github.com/mateusmoutinho/churrasco/sandbox/constructors/server"
	deps "github.com/mateusmoutinho/churrasco/sandbox/deps"
)

// New builds the whole library: one call per package under
// sandbox/constructors/, each filling the field of the Sandbox it owns. The
// list is the directories themselves, so a constructor written by hand is
// called exactly like a generated one — this file is rendered around what is
// there, never the other way round.
func New(deps *deps.Deps) *api.Sandbox {
	self := api.Sandbox{Deps: deps}

	cli.Constructor(&self)
	config.Constructor(&self)
	server.Constructor(&self)

	return &self
}
