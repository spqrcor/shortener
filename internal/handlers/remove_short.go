package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"shortener/internal/authenticate"
	"shortener/internal/services"
)

// RemoveShortHandler обработчик роута DELETE /api/user/urls
func RemoveShortHandler(b services.BatchRemover) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodDelete {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		var input []string
		var buf bytes.Buffer

		_, err := buf.ReadFrom(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(buf.Bytes(), &input); err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		if len(input) == 0 {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		UserID, ok := req.Context().Value(authenticate.ContextUserID).(uuid.UUID)
		if ok {
			go b.DeleteShortURL(UserID, input)
		}

		res.WriteHeader(http.StatusAccepted)
	}
}
