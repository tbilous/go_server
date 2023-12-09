package main

import (
	"go_server/router"
	"go_server/services"
	"go_server/utils"
	"log"
	"net/http"
)

func main() {
	log.Println("Hello World")
	var db_conn = utils.GetConnection()
	services.SetDB(db_conn)
	var appRouter = router.CreateRouter()
	log.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}
