package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"rating-service/config"
	"rating-service/consul"
	"rating-service/controller"
	"rating-service/db"
	"rating-service/leonidas"
)
// Start is to start the rest api
func Start() error{
	var err error
	if err = consul.GetSystemProperties(); err!=nil {
		return err
	}
	if err = consul.GetServiceProperties(); err!=nil {
		return err
	}
	if err = db.InitDatabase(); err !=nil {
		return err
	}
	if err :=consul.RegisterServiceWithConsul(); err !=nil {
		return err
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", controller.Health).Methods("GET")
	router.HandleFunc("/rating/{userId}", controller.GetRatingDataByUser).Methods("GET")
	leonidas.Logging(leonidas.INFO,nil, fmt.Sprintf("Starting rating service on port[%s]",
		os.Getenv(config.HttpPort)))
	if err := http.ListenAndServe(os.Getenv(config.HttpServer)+":"+os.Getenv(config.HttpPort), router); err != nil {
		leonidas.Logging("ERROR",nil,"Error while starting server: "+ err.Error())
		return err
	}
	return nil
}
