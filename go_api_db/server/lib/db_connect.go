package lib

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	User    string `json: "user"`
	Passwd  string `json: "passwd"`
	Host    string `json: "host"`
	Db      string `json: "db"`
	Port    string `json: "port"`
	Charset string `json: "charset"`
}

type pkgData struct {
	Name      string `json:"name"`
	MovieId   string `json:"movieId"`
	Rating    string `json:"rating"`
	TimeStamp string `json:"timestamp"`
}

func setMysqlURL() (string, error) {
	var config Config
	path, _ := os.Getwd()
	configPath := path + "/conf/conf.json"

	file, err := os.Open(configPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&config)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		config.User,
		config.Passwd,
		config.Host,
		config.Port,
		config.Db,
		config.Charset), nil
}

func MysqlConnection() ([]pkgData, error) {

	url, err := setMysqlURL()
	if err != nil {
		return []pkgData{}, err
	}
	db, err := sql.Open("mysql", url)
	if err != nil {
		return []pkgData{}, err
	}
	fmt.Printf("DB 연동: %+v\n", db.Stats())

	rows, err := db.Query("SELECT name, movie_id, rating, timestamp FROM movies")
	if err != nil {
		return []pkgData{}, err
	}
	defer rows.Close()

	var packages []pkgData

	for rows.Next() {
		var pkg pkgData
		if err := rows.Scan(&pkg.Name, &pkg.MovieId, &pkg.Rating, &pkg.TimeStamp); err != nil {
			return []pkgData{}, err
		}
		packages = append(packages, pkg)
	}
	if err := rows.Err(); err != nil {
		return []pkgData{}, err
	}

	db.Close()
	fmt.Printf("DB 연동 종료: %+v\n", db.Stats())
	return packages, nil
}
