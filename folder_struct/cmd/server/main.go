package main

import (
	"krriya/internal/adapter/persistence/db"
	"krriya/internal/config"
	"krriya/internal/interfaces/input/api"
	"krriya/pkg/logging"
	"log"
	"net/http"
)

func main() {

	//Load Config
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Connection Failed", err)
	}
	//Create DB Conn
	dbConn := db.NewMongoConnection(conf.DB)
	client, ctx, cnlCtx, err := dbConn.Connect()
	if err != nil {
		log.Fatal("something wrong with database", err)
	}

	_ = logging.NewService(".logs/logger.log")

	_ = client.Database(conf.DB.DBDatabase)
	defer dbConn.Close(client, ctx, cnlCtx)

	//Create repo instance

	//Create use case instance

	r := api.SetUpRoutes(conf.Server)

	//listen to server
	http.ListenAndServe(":"+conf.Server.Port, r)

}
