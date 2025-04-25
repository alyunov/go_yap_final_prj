package api

import (
	"encoding/json"
	"net/http"

	"github.com/alyunov/go_yap_final_prj/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50) // в параметре максимальное количество записей
	if err != nil {
		http.Error(w, `{"error":"Ошибка получения списка задач"}`, http.StatusInternalServerError)
		return
	}

	response := TasksResp{
		Tasks: tasks,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"error":"Ошибка кодирования JSON"}`, http.StatusInternalServerError)
		return
	}
}
