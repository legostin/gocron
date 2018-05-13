package timechecker

import (
	"time"
	"gocron/types"
	"strings"
	"strconv"
	"fmt"
)
// Функция для проверки необходимости запуска команды в данный момент времени
func NeedToRunNow(element types.Task) bool {
	period         :=element.Period
	current        :=int32(time.Now().Unix())
	current_date   := time.Now().Format("2006-01-02")
	current_time   := time.Now().Format("15:04:05")
	current_weekday:= time.Now().Format("Mon")
	need:=true
	// Если задан период, то проверяем подходит ли текущее время для выполнения задачи
	if (period>0) {
		if (!checkPeriod(period,current)) {
			need=false
		}
	}

	// Проверяем, не входит ли текущее время в период "сна" задания
	if (need) {
		for _, stime := range element.SleepTime {
			if (checkInSleepTime(stime, current_time)) {
				need = false
			}
		}
	}

	// Проверяем, не входит ли текущая дата в период "сна" задания
	if (need) {
		for _, sday := range element.SleepDays {
			if (checkInSleepingDays(sday, current_weekday)) {
				need = false
			}
		}
	}

	// Если задано, что задача должна выполняться в определенное время, то проверяем это условие
	if (need && len(element.Time)>0) {
		temp_need:=false
		for _, stime := range element.Time {
			if (checkInTime(stime, current_time)) {
				temp_need=true
			}
		}
		if (!temp_need) {
			need=false
		}
	}

	// Если задано, что задача должна выполняться в определенный день, то проверяем это условие
	if (need  && len(element.DateTime)>0) {
		temp_need_datetime:=false
		for _, sdatetime := range element.DateTime {
			if (checkInTime(sdatetime, current_date+" "+current_time)) {
				temp_need_datetime=true
			}
		}
		if (!temp_need_datetime) {
			need=false
		}
	}

	return need

}

func checkPeriod(period int32, current int32) bool {
	if (current % period == 0) {
		return true
	} else {
		return false
	}
}

func checkInSleepTime(sleeprange string, current_time string) bool{
	sleeprange_start:=toInteger(strings.Split(sleeprange, "-")[0])
	sleeprange_end  :=toInteger(strings.Split(sleeprange, "-")[1])
	current_seconds :=toInteger(current_time)
	if (current_seconds>=sleeprange_start && current_seconds<=sleeprange_end) {
		return true
	} else {
		return false
	}
}

func checkInSleepingDays(day string, current_day string) bool {
	if (day==current_day) {
		return true
	} else {
		return false
	}
}
func checkInTime(time string, current_time string) bool {

	if (time==current_time) {
		fmt.Println(time,current_time)
		return true
	} else {
		return false
	}
}

func toInteger(time string) int{
	hours,_:=strconv.Atoi(strings.Split(time,":")[0])
	minutes,_:= strconv.Atoi(strings.Split(time,":")[1])
	seconds,_:=strconv.Atoi(strings.Split(time,":")[2])
	return hours*3600+minutes*60+seconds
}