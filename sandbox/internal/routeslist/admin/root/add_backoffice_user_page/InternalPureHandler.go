package add_backoffice_user_page

import (
	"github.com/mateusmoutinho/churrasco/sandbox/api"
	"github.com/mateusmoutinho/churrasco/sandbox/deps/serverdeps"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/generated/routeio"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/routeprops"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/server/backoffice/backofficeauth"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/server/backoffice/backofficerender"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/server/backoffice/backofficeusers"
)

// InternalPureHandler answers GET /admin/root/add-backoffice-user with the
// empty form that POST /admin/root/add-backoffice-user reads, the viewer role
// selected.
func InternalPureHandler(sandbox *api.Sandbox, props *routeprops.RouteProps, entries *Entries, response *serverdeps.Response) error {
	if props.User == nil {
		return routeio.Fail(sandbox, api.StatusUnauthorized, "", "no authenticated user")
	}
	fields := backofficeusers.Fields{Role: int64(backofficeauth.RoleViewer)}
	return backofficerender.AddBackofficeUserForm(sandbox, response, api.StatusOk, props.User, fields, "")
}
