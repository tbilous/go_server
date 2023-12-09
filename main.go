package main

import (
	"log"
	"net/http"

	"github.com/tbilous/go_server/router"
	"github.com/tbilous/go_server/services"
	"github.com/tbilous/go_server/utils"
)

func main() {
	log.Println("Hello World")
	var db_conn = utils.GetConnection()
	services.SetDB(db_conn)
	var appRouter = router.CreateRouter()
	log.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}
