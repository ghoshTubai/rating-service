package main

import (
	"fmt"
	"os"
	"rating-service/app"
	"rating-service/leonidas"
)

func main() {
	//os.Setenv("SystemProperties","rating-data/system/configs")
	//os.Setenv("ServiceProperties","rating-data/service/configs")
	//os.Setenv(config.ConsulAddr, "127.0.0.1:8500")
	if err := app.Start(); err!=nil {
		leonidas.Logging(leonidas.ERROR,nil, fmt.Sprintf("%s -shutting down rating-service because ")+err.Error())
		os.Exit(1)
	}
}