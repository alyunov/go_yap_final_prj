package api

import (
	"encoding/json"
	"net/http"

	"github.com/alyunov/go_yap_final_prj/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := db.DeleteTask(id)
	if err != nil {
		http.Error(w, `{"error":"задача с таким id не найдена"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{}); err != nil {
		http.Error(w, `{"error":"ошибка кодирования JSON"}`, http.StatusInternalServerError)
		return
	}
}
