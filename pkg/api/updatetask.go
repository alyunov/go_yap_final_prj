package api

import (
	"encoding/json"
	"net/http"

	"github.com/alyunov/go_yap_final_prj/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t db.Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, `{"error":"Ошибка десериализации JSON"}`, http.StatusBadRequest)
		return
	}
	err = db.UpdateTask(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{}); err != nil {
		http.Error(w, `{"error":"Ошибка кодирования JSON"}`, http.StatusInternalServerError)
		return
	}
}
