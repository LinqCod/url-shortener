package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
)

func initRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/new", CreateShortUrl).Methods("POST")
	r.HandleFunc("/api/v1/:url", RedirectOriginalUrl).Methods("GET")
	r.HandleFunc("/api/v1/urls", GetAllUrls).Methods("GET")

	return r
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	req := r.Body

	var originalUrl OriginalUrl
	if err := json.NewDecoder(req).Decode(&originalUrl); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := len(urls) + rand.Int()
	shortUrl := encode(id)
	response := Response{
		OriginalUrl:  originalUrl.url,
		ShortenedUrl: shortUrl,
	}
	urls = append(urls, response)

	w.Write([]byte(shortUrl))
}

func RedirectOriginalUrl(w http.ResponseWriter, r *http.Request) {
	shortenedUrl := mux.Vars(r)["url"]

	if originalUrl, ok := urls[shortenedUrl]; !ok {
		http.Error(w, "There is no originalUrl for this url", http.StatusBadRequest)
	} else {
		http.Redirect(w, r, originalUrl, http.StatusPermanentRedirect)
	}
}

func GetAllUrls(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(urls); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
