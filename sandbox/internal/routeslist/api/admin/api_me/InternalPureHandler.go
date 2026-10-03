package api_me

import (
	"github.com/mateusmoutinho/churrasco/sandbox/api"
	"github.com/mateusmoutinho/churrasco/sandbox/deps/serverdeps"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/generated/routeio"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/routeprops"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/server/backoffice/backofficeapi"
)

// InternalPureHandler answers GET /api/admin/me with the user the
// api-authentication middleware put on props.User.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}
	return routeio.WriteJSON(sandbox, *response, api.StatusOk, backofficeapi.UserDocument(sandbox, *props.User))
}
