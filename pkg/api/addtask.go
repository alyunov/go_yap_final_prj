package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/alyunov/go_yap_final_prj/pkg/db"
	"github.com/alyunov/go_yap_final_prj/pkg/model"
)

func addTaskHandler(res http.ResponseWriter, req *http.Request) {
	var task db.Task

	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		http.Error(res, `{"error":"Ошибка десериализации JSON"}`, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		http.Error(res, `{"error":"Ошибка указания заголовка задачи"}`, http.StatusBadRequest)
		return
	}
	if task.Date == "" {
		task.Date = time.Now().Format(DateFormat)
	}

	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		http.Error(res, `{"error":"Ошибка формата даты"}`, http.StatusBadRequest)
		return
	}
	now := time.Now()

	if model.AfterNow(now, t) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format("20060102")
		} else {
			next, err := model.NextDate(now, task.Date, task.Repeat)
			if err != nil {
				http.Error(res, `{"error":"Ошибка правила повторения"}`, http.StatusBadRequest)
				return
			}
			task.Date = next
		}
	}

	id, err := db.AddTask(&task)
	if err != nil {
		http.Error(res, `{"error":"Ошибка добавления задачи"}`, http.StatusBadRequest)
		return
	}

	//response := db.Response{ID: strconv.FormatInt(id, 10)}
	response := db.Response{ID: strconv.FormatInt(id, 10)}
	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(response); err != nil {
		http.Error(res, `{"error":"Ошибка кодирования JSON"}`, http.StatusInternalServerError)
		return
	}
}
