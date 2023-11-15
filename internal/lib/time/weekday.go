package time

import (
	"fmt"
	"time"
)

var daysOfWeek = map[string]time.Weekday{}

func init() {
	for d := time.Sunday; d <= time.Saturday; d++ {
		name := d.String()
		daysOfWeek[name] = d
		daysOfWeek[name[:3]] = d
	}
}

func ParseWeekday(v string) (time.Weekday, error) {
	if d, ok := daysOfWeek[v]; ok {
		return d, nil
	}

	return time.Sunday, fmt.Errorf("invalid weekday '%s'", v)
}

func ToRus(weekday time.Weekday) string {
	var res string

	switch weekday {
	case time.Sunday:
		res = "Воскресенье"
	case time.Monday:
		res = "Понедельник"
	case time.Tuesday:
		res = "Вторник"
	case time.Wednesday:
		res = "Среда"
	case time.Thursday:
		res = "Четверг"
	case time.Friday:
		res = "Пятница"
	case time.Saturday:
		res = "Суббота"
	}

	return res
}

func NextWeekday(start time.Time, targetWeekday time.Weekday) time.Time {
	daysToAdd := (int(targetWeekday) - int(start.Weekday()) + 7) % 7
	nextWeekday := start.Add(time.Duration(daysToAdd) * 24 * time.Hour)
	return nextWeekday
}
