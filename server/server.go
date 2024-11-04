// server.go
// this file will handle server related ops
package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tarsh333/go_db/utils"
)

func StartServer() {
	router := mux.NewRouter()
	dbRouterV1 := router.PathPrefix("/db/v1/").Subrouter()
	dbRouterV1.Use(checkAndAddUserDirectory)
	dbRouterV1.Use(addResponseType)
	var port int = 8000

	fmt.Println("Server started at port ", port)
	// TODO:: Move path and port to constant
	// create a db folder inside which all the data resides
	utils.CreateFolder("db")
	if err := http.ListenAndServe(":8000", router); err != nil {
		panic(err)
	}
}
