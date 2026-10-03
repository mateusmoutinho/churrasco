package deps

import (
	argvdeps "github.com/mateusmoutinho/churrasco/sandbox/deps/argvdeps"
	database "github.com/mateusmoutinho/churrasco/sandbox/deps/database"
	embeddeps "github.com/mateusmoutinho/churrasco/sandbox/deps/embeddeps"
	envdeps "github.com/mateusmoutinho/churrasco/sandbox/deps/envdeps"
	hashdeps "github.com/mateusmoutinho/churrasco/sandbox/deps/hashdeps"
	jwtdeps "github.com/mateusmoutinho/churrasco/sandbox/deps/jwtdeps"
	passworddeps "github.com/mateusmoutinho/churrasco/sandbox/deps/passworddeps"
	randdeps "github.com/mateusmoutinho/churrasco/sandbox/deps/randdeps"
	ratelimitdeps "github.com/mateusmoutinho/churrasco/sandbox/deps/ratelimitdeps"
	reflectdeps "github.com/mateusmoutinho/churrasco/sandbox/deps/reflectdeps"
	serializables "github.com/mateusmoutinho/churrasco/sandbox/deps/serializables"
	serverdeps "github.com/mateusmoutinho/churrasco/sandbox/deps/serverdeps"
	signaldeps "github.com/mateusmoutinho/churrasco/sandbox/deps/signaldeps"
	sortdeps "github.com/mateusmoutinho/churrasco/sandbox/deps/sortdeps"
	std "github.com/mateusmoutinho/churrasco/sandbox/deps/std"
	stringsdeps "github.com/mateusmoutinho/churrasco/sandbox/deps/stringsdeps"
	timedeps "github.com/mateusmoutinho/churrasco/sandbox/deps/timedeps"
)

// Deps is every capability the sandbox needs from the outside world, one field
// per sub-contract directory of sandbox/deps/. An adapter fills the fields; the
// sandbox only calls them, which is what keeps it free of OS packages.
type Deps struct {
	Argvdeps      argvdeps.Sandbox
	Database      database.Sandbox
	Embeddeps     embeddeps.Sandbox
	Envdeps       envdeps.Sandbox
	Hashdeps      hashdeps.Sandbox
	Jwtdeps       jwtdeps.Sandbox
	Passworddeps  passworddeps.Sandbox
	Randdeps      randdeps.Sandbox
	Ratelimitdeps ratelimitdeps.Sandbox
	Reflectdeps   reflectdeps.Sandbox
	Serializables serializables.Sandbox
	Serverdeps    serverdeps.Sandbox
	Signaldeps    signaldeps.Sandbox
	Sortdeps      sortdeps.Sandbox
	Std           std.Sandbox
	Stringsdeps   stringsdeps.Sandbox
	Timedeps      timedeps.Sandbox
}
