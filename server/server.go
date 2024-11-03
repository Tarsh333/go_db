// server.go
// this file will handle server related ops
package server

import (
	"fmt"
	"net/http"
)

func StartServer() {
	var port int = 8000
	fmt.Println("Server started at port ", port)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
