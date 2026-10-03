package home

import (
	"github.com/mateusmoutinho/churrasco/sandbox/api"
	"github.com/mateusmoutinho/churrasco/sandbox/deps/serverdeps"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/generated/routeio"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/routeprops"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/server/backoffice/backofficeauth"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/server/backoffice/backofficerender"
)

// InternalPureHandler answers GET /admin/home with backoffice/home.html,
// rendered for the user the authentication middleware put on props.User.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}
	return backofficerender.Home(sandbox, response, props.User, backofficeauth.SessionSeconds/60)
}
