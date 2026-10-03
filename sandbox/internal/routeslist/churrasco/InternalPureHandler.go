package churrasco

import (
	"github.com/mateusmoutinho/churrasco/sandbox/api"
	"github.com/mateusmoutinho/churrasco/sandbox/deps/serverdeps"
	estatisticasdb "github.com/mateusmoutinho/churrasco/sandbox/internal/databases/estatisticas"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/routeprops"
)

// InternalPureHandler answers GET /churrasco. Every value the route
// declares is already on entries, read off the request by the generic
// RequestHandler, and the response already carries the route's response-type.
//
// Answering — setting a status, or writing a byte, which sends a 200 — is what
// ends the chain. A handler that does neither has declined, and the next route
// matching this request runs — which is how a route becomes a middleware.
// What a middleware in front set on props — the request's
// routeprops.RouteProps, whose parts sandbox/internal/routeprops/ declares —
// is there to read. Refuse a request
// by returning routeio.Fail; nil means "done" or "not mine".
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	carne := entries.Adultos*400 + entries.Criancas*200 // gramas
	cervejas := entries.Adultos * 6                     // latas
	json := sandbox.Deps.Std.Sprintf(`{"carne_g": %d, "cervejas": %d}`, carne, cervejas)

	// Registra o cálculo para as estatísticas do backoffice. Uma falha aqui
	// não muda a resposta: o cálculo é devolvido igual e o erro vai para o log.
	_, err := estatisticasdb.New(sandbox).AddCalculo(estatisticasdb.CalculoNew{
		Adultos:  int64(entries.Adultos),
		Criancas: int64(entries.Criancas),
		Carneg:   int64(carne),
		Cervejas: int64(cervejas),
		Criadoem: sandbox.Deps.Std.Now() / 1_000_000_000,
	})
	if err != nil {
		sandbox.Deps.Std.Error("churrasco: não foi possível registrar o cálculo: %s\n", err.Error())
	}

	response.Write([]byte(json))
	return nil
}
