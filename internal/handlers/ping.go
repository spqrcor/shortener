package handlers

import (
	"net/http"
	"shortener/internal/config"
	"shortener/internal/db"
)

func PingHandler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		conf := config.NewConfig()
		if req.Method != http.MethodGet || conf.DatabaseDSN == "" {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err := db.Connect(conf.DatabaseDSN)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusOK)
	}
}
