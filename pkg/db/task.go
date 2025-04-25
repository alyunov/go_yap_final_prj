package db

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alyunov/go_yap_final_prj/pkg/model"
)

const DateFormat = "20060102"

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}
type Response struct {
	ID    string `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

func AddTask(task *Task) (int64, error) {
	var id int64
	// определите запрос
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`
	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}

func Tasks(limit int) ([]*Task, error) {
	var tasks []*Task
	rows, err := db.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT ?", limit)
	if err != nil {
		return []*Task{}, fmt.Errorf(`{"error":"ошибка запроса"}`)
	}
	defer rows.Close()
	for rows.Next() {
		t := new(Task)
		err := rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err = rows.Err(); err != nil {
			return []*Task{}, fmt.Errorf(`{"error":"ошибка распознавания данных"}`)
		}
		tasks = append(tasks, t)
	}
	if len(tasks) == 0 {
		tasks = []*Task{}
	}

	return tasks, nil
}

func GetTask(id string) (*Task, error) {
	var t Task
	if id == "" {
		return nil, fmt.Errorf(`{"error":"не указан id"}`)
	}
	row := db.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?", id)
	err := row.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
	if err != nil {
		return nil, fmt.Errorf(`{"error":"Задача не найдена"}`)
	}
	return &t, nil
}

func UpdateTask(task *Task) error {
	if task.ID == "" {
		return fmt.Errorf(`{"error":"не указан id"}`)
	}

	if task.Title == "" {
		return fmt.Errorf(`{"error":"не указан заголовок задачи"}`)
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

	if model.afterNow(now, t) {
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

	query := `UPDATE scheduler SET date=?, title=?, comment=?, repeat=? WHERE id=?`
	res, err := db.Exec(query, t.Date, t.Title, t.Comment, t.Repeat, t.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}
