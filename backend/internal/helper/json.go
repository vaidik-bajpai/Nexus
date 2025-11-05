package helper

import (
	"encoding/json"
	"net/http"

	"github.com/vaidik-bajpai/Nexus/backend/internal/types"
)

func ReadJSON(r *http.Request, payload any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&payload); err != nil {
		// remember to implement proper error handling
		return err
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, statusCode int64, response *types.Response) {
	w.WriteHeader(int(statusCode))
	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
