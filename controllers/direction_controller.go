package controllers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
	"url_shortener/data"
	"url_shortener/models"

	"github.com/gorilla/mux"
)

// CreateShortURL crea una referencia de una url en la base de datos
func CreateShortURL(w http.ResponseWriter, r *http.Request) {

	dir := models.Direction{}
	err := dir.FromJSON(r.Body)

	dir.ShortURL = randStringRunes(6)
	dir.CreateAt = time.Now()
	dir.UpdateAt = time.Now()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db := data.GetConnection()
	defer db.Close()
	insert, err := db.Query("INSERT INTO urls(url, short_url, create_at, update_at) VALUES(?, ?, ?, ?)", dir.URL, dir.ShortURL, dir.CreateAt, dir.UpdateAt)

	if err != nil {
		log.Println(err.Error())
	}

	res := &responseURL{dir.URL, dir.ShortURL}
	json.NewEncoder(w).Encode(res)

	defer insert.Close()

}

// GetURL redirecciona apartir de una url corta
func GetURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	res := &responseURL{}
	res.ShortURL = vars["url"]
	res.LongURL = findByShortURL(res.ShortURL)

	http.Redirect(w, r, res.LongURL, http.StatusMovedPermanently)
}

func findByShortURL(short string) string {
	db := data.GetConnection()
	defer db.Close()

	var longURL string

	err := db.QueryRow("SELECT url FROM urls WHERE short_url = ?", short).Scan(&longURL)
	if err != nil {
		panic(err.Error())
	}
	return longURL
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type responseURL struct {
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}
