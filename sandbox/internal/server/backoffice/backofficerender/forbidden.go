package backofficerender

import (
	"github.com/mateusmoutinho/churrasco/sandbox/api"
	"github.com/mateusmoutinho/churrasco/sandbox/deps/serverdeps"
	"github.com/mateusmoutinho/churrasco/sandbox/internal/databases/backofficedb"
)

// ForbiddenPage is what backoffice/forbidden.html is rendered with.
type ForbiddenPage struct {
	Viewer Viewer
}

// Forbidden answers, under a 403, the page telling user that what they asked
// for needs the root role.
func Forbidden(sandbox *api.Sandbox, response *serverdeps.Response, user *backofficedb.BackofficeuserItem) error {
	return Html(sandbox, response, api.StatusForbidden, "backoffice/forbidden.html", ForbiddenPage{Viewer: viewerOf(sandbox, user)})
}
