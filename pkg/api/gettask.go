package api

import (
	"encoding/json"
	"net/http"

	"github.com/alyunov/go_yap_final_prj/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	task, err := db.GetTask(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error":"задача с таким id не найдена"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, `{"error":"Ошибка кодирования JSON"}`, http.StatusInternalServerError)
		return
	}
}
