package main

import (
	"encoding/json"
	"fmt"
	"gocron/timechecker"
	"gocron/types"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {

	for {
		file, err := os.Open("cronjob.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		b, err := ioutil.ReadAll(file)
		var taskArray types.TaskArray
		//var tasks []Task
		err_json := json.Unmarshal(b, &taskArray)
		if err_json != nil {
			fmt.Println("error:", err_json)
		}
		for _, element := range taskArray.Tasks {
			if timechecker.NeedToRunNow(element) {
				fmt.Println(time.Now())
				go runCommand(element.Command, element.Output)
			}
		}
		time.Sleep(time.Millisecond * 1000)
	}
}

func runCommand(command string, output string) {
	cmnd := command + " >> " + output
	gocronLog("START", cmnd)
	runcmd(command+" >> "+output, true)
}

func runcmd(cmd string, shell bool) []byte {
	start_time := int(time.Now().Unix())
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			gocronLog("ERROR", cmd+"("+err.Error()+")")
		}
		finish_time := int(time.Now().Unix())
		t := strconv.Itoa(finish_time - start_time)
		gocronLog("FINISH in "+t+"s", cmd)
		return out
	}
	out, err := exec.Command(cmd).Output()
	if err != nil {
		gocronLog("ERROR", cmd+"("+err.Error()+")")

	}
	finish_time := int(time.Now().Unix())
	t := strconv.Itoa(finish_time - start_time)
	gocronLog("FINISH in "+t, cmd)
	return out
}

func gocronLog(message_type string, message string) {
	format_message := "[" + time.Now().String() + "]" + " " + message_type + ": " + message + "\n"
	fmt.Println(format_message)
	f, _ := os.OpenFile("./logs/gocron.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	f.WriteString(format_message)
	f.Close()
	if message_type == "ERROR" {
		f, _ := os.OpenFile("./logs/error.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		f.WriteString(format_message)
		f.Close()
	}
}
