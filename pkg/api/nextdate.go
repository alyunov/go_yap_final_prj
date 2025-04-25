package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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
	nextDate, err := NextDate(nowTime, date, repeat)
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

func afterNow(date, now time.Time) bool {
	dateD := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	nowD := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return dateD.After(nowD)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", fmt.Errorf("пустой repeat")
	}
	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", fmt.Errorf("неверный формат даты: %v", err)
	}

	repeatSplit := strings.Split(repeat, " ")
	repeatType := repeatSplit[0]
	switch repeatType {
	case "d":
		if len(repeatSplit) < 2 {
			return "", fmt.Errorf("не указано количество дней")
		}
		days, err := strconv.Atoi(repeatSplit[1])
		if err != nil {
			return "", err
		}
		if days > 400 {
			return "", fmt.Errorf("максимально допустимое число дней равно 400")
		}
		for {
			date = date.AddDate(0, 0, days)
			if afterNow(date, now) {
				break
			}
		}

		return date.Format(DateFormat), nil
	case "y":
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format(DateFormat), nil
	default:
		return "", fmt.Errorf("неправильный формат правила")
	}

}
