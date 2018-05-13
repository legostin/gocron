package timechecker

import (
	"fmt"
	"gocron/types"
	"strconv"
	"strings"
	"time"
)

//NeedToRunNow Функция для проверки необходимости запуска команды в данный момент времени
func NeedToRunNow(element types.Task) bool {
	period := element.Period
	current := int32(time.Now().Unix())
	currentDate := time.Now().Format("2006-01-02")
	currentTime := time.Now().Format("15:04:05")
	currentWeekday := time.Now().Format("Mon")
	need := true
	// Если задан период, то проверяем подходит ли текущее время для выполнения задачи
	if period > 0 {
		if !checkPeriod(period, current) {
			need = false
		}
	}
	// Проверяем, не входит ли текущее время в период "сна" задания
	if need {
		for _, stime := range element.SleepTime {
			if checkInSleepTime(stime, currentTime) {
				need = false
			}
		}
	}

	// Проверяем, не входит ли текущая дата в период "сна" задания
	if need {
		for _, sday := range element.SleepDays {
			if checkInSleepingDays(sday, currentWeekday) {
				need = false
			}
		}
	}

	// Если задано, что задача должна выполняться в определенное время, то проверяем это условие
	if need && len(element.Time) > 0 {
		tempNeed := false
		for _, stime := range element.Time {
			if checkInTime(stime, currentTime) {
				tempNeed = true
			}
		}
		if !tempNeed {
			need = false
		}
	}

	// Если задано, что задача должна выполняться в определенный день, то проверяем это условие
	if need && len(element.DateTime) > 0 {
		tempNeedDatetime := false
		for _, sdatetime := range element.DateTime {
			if checkInTime(sdatetime, currentDate+" "+currentTime) {
				tempNeedDatetime = true
			}
		}
		if !tempNeedDatetime {
			need = false
		}
	}

	return need

}

//checkPeriod функция для проверки принадлежности к периоду в секундах
func checkPeriod(period int32, current int32) bool {
	if current%period == 0 {
		return true
	} else {
		return false
	}
}

//checkInSleepTime функция, которая определяет не входит ли текущее время в промежуток "сна" для задачи
func checkInSleepTime(sleeprange string, currentTime string) bool {
	sleeprangeStart := toInteger(strings.Split(sleeprange, "-")[0])
	sleeprangeEnd := toInteger(strings.Split(sleeprange, "-")[1])
	currentSeconds := toInteger(currentTime)
	if currentSeconds >= sleeprangeStart && currentSeconds <= sleeprangeEnd {
		return true
	} else {
		return false
	}
}

//checkInSleepingDays Проверка, не входил ли текущее время в дни "сна" задачи
func checkInSleepingDays(day string, currentDay string) bool {
	if day == currentDay {
		return true
	} else {
		return false
	}
}

//checkInTime Функция проверяет, не является ли текущее время моментом запуска задачи
func checkInTime(time string, currentTime string) bool {

	if time == currentTime {
		fmt.Println(time, currentTime)
		return true
	} else {
		return false
	}
}

//toInteger функция для преобразования времени в секунды с полуночи
func toInteger(time string) int {
	hours, _ := strconv.Atoi(strings.Split(time, ":")[0])
	minutes, _ := strconv.Atoi(strings.Split(time, ":")[1])
	seconds, _ := strconv.Atoi(strings.Split(time, ":")[2])
	return hours*3600 + minutes*60 + seconds
}
