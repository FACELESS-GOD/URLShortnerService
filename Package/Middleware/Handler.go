package Middleware

import (
	"net/http"
	"time"

	Model "github.com/FACELESS-GOD/URLShortnerService.git/Package/Models"
	"github.com/FACELESS-GOD/URLShortnerService.git/Package/Utility"
	"github.com/gorilla/mux"
)

func MiddlewareHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(writer http.ResponseWriter, Req *http.Request) {
			startTime := time.Now()
			timeLog := "[ INFO ]  Duration:- "
			vars := mux.Vars(Req)

			id := vars["ID"]
			if id != "" {

				old_URL, err := Model.GetURL(id)

				if err != nil {
					writer.WriteHeader(http.StatusInternalServerError)
					return
				}

				timeLog = timeLog + time.Since(startTime).String() + " For " + id
				Utility.LogDuartion(timeLog)
				http.Redirect(writer, Req, old_URL, http.StatusSeeOther)
				return

			} else {
				next.ServeHTTP(writer, Req)
				timeLog = timeLog + time.Since(startTime).String() + " For Addition Operation "
				Utility.LogDuartion(timeLog)
			}

		})

}
