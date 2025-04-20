package main

import (
	"net/http"

	Config "github.com/FACELESS-GOD/URLShortnerService.git/Helper/MetaData"
	"github.com/FACELESS-GOD/URLShortnerService.git/Package/Route"
	"github.com/FACELESS-GOD/URLShortnerService.git/Package/Utility"
	"github.com/gorilla/mux"
)

func main() {
	Config.InitiateConfig()
	Utility.InitiateDatabase()
	Utility.InitaiteRedisInstance()
	Utility.InitiateLogs()

	MuxRouter := mux.NewRouter()
	Route.CustomRouter(MuxRouter)

	err := http.ListenAndServe(Config.MainURL, MuxRouter)

	if err != nil {
		panic(err)
	}
}