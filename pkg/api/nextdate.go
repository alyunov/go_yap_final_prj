package api

import (
	"log"
	"net/http"
	"time"

	"github.com/alyunov/go_yap_final_prj/pkg/model"
)

func nextDayHandler(res http.ResponseWriter, req *http.Request) {
	now := req.FormValue("now")
	date := req.FormValue("date")
	repeat := req.FormValue("repeat")

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")

	nowTime, err := time.Parse(DateFormat, now)
	if err != nil {
		http.Error(res, "Некорректный формат даты", http.StatusBadRequest)
		return
	}
	nextDate, err := model.NextDate(nowTime, date, repeat)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = res.Write([]byte(nextDate))
	if err != nil {
		log.Fatal(err)
		return
	}

}
