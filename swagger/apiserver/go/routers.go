/*
 * Chain Query
 *
 * The LBRY blockchain is read into SQL where important structured information can be extracted through the Chain Query API.
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"net/http"
	"strings"

	. "github.com/lbryio/chainquery/apiactions"
	. "github.com/lbryio/lbry.go/api"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc Handler
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/",
		IndexAction,
	},

	Route{
		"AutoUpdate",
		strings.ToUpper("Get"),
		"/api/autoupdate",
		AutoUpdateAction,
	},

	Route{
		"SQLQuery",
		strings.ToUpper("Get"),
		"/api/sql",
		SQLQueryAction,
	},

	Route{
		"AddressSummary",
		strings.ToUpper("Get"),
		"/api/addresssummary",
		AddressSummaryAction,
	},

	Route{
		"ChainQueryStatus",
		strings.ToUpper("Get"),
		"/api/status",
		ChainQueryStatusAction,
	},
}
