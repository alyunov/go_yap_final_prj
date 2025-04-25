package api

import (
	"encoding/json"
	"net/http"

	"github.com/alyunov/go_yap_final_prj/pkg/db"
)

func taskDoneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Метод не поддерживается"}`, http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	err := db.TaskDone(id)
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
