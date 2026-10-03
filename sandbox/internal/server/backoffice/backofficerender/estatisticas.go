package backofficerender

import (
	"github.com/mateusmoutinho/churrasco/sandbox/api"
	"github.com/mateusmoutinho/churrasco/sandbox/deps/serverdeps"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/databases/backofficedb"
)

// EstatisticasPage is what backoffice/estatisticas.html is rendered with.
type EstatisticasPage struct {
	Viewer Viewer
	// Acessos is how many times the site's home page was opened.
	Acessos int
	// Calculos is how many times /churrasco answered a calculation.
	Calculos int
}

// Estatisticas answers the statistics page for user.
func Estatisticas(sandbox *api.Sandbox, response *serverdeps.Response, user *backofficedb.BackofficeuserItem, acessos int, calculos int) error {
	return Html(sandbox, response, api.StatusOk, "backoffice/estatisticas.html", EstatisticasPage{
		Viewer:   viewerOf(sandbox, user),
		Acessos:  acessos,
		Calculos: calculos,
	})
}
