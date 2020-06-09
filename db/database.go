package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"rating-service/leonidas"
	"time"
)

var conn *sql.DB

func InitDatabase () error {
	var err error
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "ankur", "password", "mysql", "3306", "test")
	conn, err = sql.Open("mysql", url)
	if err != nil {
		leonidas.Logging(leonidas.ERROR,nil,"Unable to established connection to DB "+ err.Error())
		return err
	}
	if err = conn.Ping(); err!=nil {
		leonidas.Logging(leonidas.ERROR,nil,"Unable to ping connection to DB "+ err.Error() )
		return err
	}
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5*time.Minute)
	leonidas.Logging(leonidas.INFO,nil,"Connection successful with DB !!")
	return nil
}

func GetDBSource() (*sql.DB, error) {
	if conn != nil {
		return conn, nil
	}
	return nil, errors.New("no connection available ")
}

func InsertData(db *sql.DB) error {
	stmtIns, err := db.Prepare("INSERT INTO user_movies VALUES(? ,?, ?)")
	if err != nil {
		panic(err)
	}

	// Close the statement when we leave main() / the program terminates
	defer stmtIns.Close()

	// our id field auto increments so we don't need to pass actual value for it.
	result, err := stmtIns.Exec("Pompi", "Dragon King", 2)
	if err != nil {
		log.Println(err)
		return err
	}
	id, err := result.LastInsertId()
	log.Println(id)
	log.Println(err)
	return nil
}

func GetDataFromDB(db *sql.DB, query string) (*sql.Rows, error) {
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error while fetching data from DB ", err.Error())
		return nil, err
	}
	//defer rows.Close()
	//var ratings []model.Ratings
	//var name string
	//var rating int
	//
	//for rows.Next() {
	//	err = rows.Scan( &name,&rating)
	//	r := model.Ratings{name,rating}
	//	ratings = append(ratings,r)
	//}
	//data, _ := json.Marshal(&ratings)
	//fmt.Println(string(data))
	return rows, nil
}
