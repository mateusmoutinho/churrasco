package standard

import (
	argvdeps "github.com/mateusmoutinho/churrasco/adapters/libs/argvdeps"
	database "github.com/mateusmoutinho/churrasco/adapters/libs/database"
	embeddeps "github.com/mateusmoutinho/churrasco/adapters/libs/embeddeps"
	envdeps "github.com/mateusmoutinho/churrasco/adapters/libs/envdeps"
	hashdeps "github.com/mateusmoutinho/churrasco/adapters/libs/hashdeps"
	jwtdeps "github.com/mateusmoutinho/churrasco/adapters/libs/jwtdeps"
	passworddeps "github.com/mateusmoutinho/churrasco/adapters/libs/passworddeps"
	randdeps "github.com/mateusmoutinho/churrasco/adapters/libs/randdeps"
	ratelimitdeps "github.com/mateusmoutinho/churrasco/adapters/libs/ratelimitdeps"
	reflectdeps "github.com/mateusmoutinho/churrasco/adapters/libs/reflectdeps"
	serializables "github.com/mateusmoutinho/churrasco/adapters/libs/serializables"
	serverdeps "github.com/mateusmoutinho/churrasco/adapters/libs/serverdeps"
	signaldeps "github.com/mateusmoutinho/churrasco/adapters/libs/signaldeps"
	sortdeps "github.com/mateusmoutinho/churrasco/adapters/libs/sortdeps"
	std "github.com/mateusmoutinho/churrasco/adapters/libs/std"
	stringsdeps "github.com/mateusmoutinho/churrasco/adapters/libs/stringsdeps"
	timedeps "github.com/mateusmoutinho/churrasco/adapters/libs/timedeps"
	deps "github.com/mateusmoutinho/churrasco/sandbox/deps"
)

func New() deps.Deps {
	deps := deps.Deps{}
	argvdeps.Bind(&deps)
	database.Bind(&deps)
	embeddeps.Bind(&deps)
	envdeps.Bind(&deps)
	hashdeps.Bind(&deps)
	jwtdeps.Bind(&deps)
	passworddeps.Bind(&deps)
	randdeps.Bind(&deps)
	ratelimitdeps.Bind(&deps)
	reflectdeps.Bind(&deps)
	serializables.Bind(&deps)
	serverdeps.Bind(&deps)
	signaldeps.Bind(&deps)
	sortdeps.Bind(&deps)
	std.Bind(&deps)
	stringsdeps.Bind(&deps)
	timedeps.Bind(&deps)
	return deps
}
