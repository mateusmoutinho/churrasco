package visita

import (
	"github.com/mateusmoutinho/churrasco/sandbox/api"
	"github.com/mateusmoutinho/churrasco/sandbox/deps/serverdeps"
	estatisticasdb "github.com/mateusmoutinho/churrasco/sandbox/internal/databases/estatisticas"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/routeprops"
)

// InternalPureHandler answers POST /visita, que a página inicial chama ao
// abrir: grava um registro na tabela visita do banco estatisticas, contado
// pela tela de Estatísticas do backoffice.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	_, err := estatisticasdb.New(sandbox).AddVisita(estatisticasdb.VisitaNew{
		Criadoem: sandbox.Deps.Std.Now() / 1_000_000_000,
	})
	if err != nil {
		return err
	}
	response.Write([]byte(`{"ok": true}`))
	return nil
}
