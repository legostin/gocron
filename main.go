package main

import (
	"encoding/json"
	"fmt"
	"github.com/legostin/gocron/timechecker"
	"github.com/legostin/gocron/types"
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
		b, _ := ioutil.ReadAll(file)
		var taskArray types.TaskArray
		//var tasks []Task
		errJson := json.Unmarshal(b, &taskArray)
		if errJson != nil {
			fmt.Println("error:", errJson)
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
	startTime := int(time.Now().Unix())
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			gocronLog("ERROR", cmd+"("+err.Error()+")")
		}
		finishTime := int(time.Now().Unix())
		t := strconv.Itoa(finishTime - startTime)
		gocronLog("FINISH in "+t+"s", cmd)
		return out
	}
	out, err := exec.Command(cmd).Output()
	if err != nil {
		gocronLog("ERROR", cmd+"("+err.Error()+")")

	}
	finishTime := int(time.Now().Unix())
	t := strconv.Itoa(finishTime - startTime)
	gocronLog("FINISH in "+t, cmd)
	return out
}

func gocronLog(messageType string, message string) {
	formatMessage := "[" + time.Now().String() + "]" + " " + messageType + ": " + message + "\n"
	fmt.Println(formatMessage)
	f, _ := os.OpenFile("./logs/gocron.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	f.WriteString(formatMessage)
	f.Close()
	if messageType == "ERROR" {
		f, _ := os.OpenFile("./logs/error.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		f.WriteString(formatMessage)
		f.Close()
	}
}
