package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rating-service/db"
	"rating-service/leonidas"
	"rating-service/model"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	return
}

func GetRatingDataByUser(w http.ResponseWriter, r *http.Request) {

	leonidas.Logging("DEBUG",r.Header.Get("uutid"),"Executing getRatingDataByUser")
	vars := mux.Vars(r)
	id := vars["userId"]
	conn, err := db.GetDBSource()
	if err != nil {
		leonidas.Logging("ERROR",r.Header.Get("uutid"),"Error while getting DB source"+err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	query := fmt.Sprintf("SELECT movie,rating FROM user_movies where user_name = '%s' ", id)
	log.Println(query)
	rows, err := db.GetDataFromDB(conn, query)
	defer rows.Close()
	if err != nil {
		leonidas.Logging("ERROR",r.Header.Get("uutid"),"Unable to get data from database. " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var ratings []model.Rating
	var name string
	var movRating int
	for rows.Next() {
		err = rows.Scan(&name, &movRating)
		r := model.Rating{name, movRating}
		ratings = append(ratings, r)
	}
	resp := model.UserRatingData{id, ratings}
	w.Header().Add("Content-Type", "applicaton/json")
	json.NewEncoder(w).Encode(resp)
	return
}
