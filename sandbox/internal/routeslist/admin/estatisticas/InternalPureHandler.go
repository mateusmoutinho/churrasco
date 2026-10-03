package estatisticas

import (
	"github.com/mateusmoutinho/churrasco/sandbox/api"
	"github.com/mateusmoutinho/churrasco/sandbox/deps/serverdeps"
	estatisticasdb "github.com/mateusmoutinho/churrasco/sandbox/internal/databases/estatisticas"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/generated/routeio"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/routeprops"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/server/backoffice/backofficerender"
)

// InternalPureHandler answers GET /admin/estatisticas, open to every
// backoffice user, with backoffice/estatisticas.html: o total de acessos ao
// site e o total de cálculos de churrasco, lidos do banco estatisticas.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}

	db := estatisticasdb.New(sandbox)
	acessos, err := db.CountVisita()
	if err != nil {
		return err
	}
	calculos, err := db.CountCalculo()
	if err != nil {
		return err
	}
	return backofficerender.Estatisticas(sandbox, response, props.User, acessos, calculos)
}
