package Route

import (
	"net/http"

	"github.com/FACELESS-GOD/URLShortnerService.git/Package/Controller"
	"github.com/FACELESS-GOD/URLShortnerService.git/Package/Middleware"
	"github.com/gorilla/mux"
)

func CustomRouter(Router *mux.Router) {

	if Router == nil {
		panic("Router is null")
	}

	Router.Use(Middleware.MiddlewareHandler)

	Router.HandleFunc("/", Controller.AddURL).Methods(http.MethodGet)

}
