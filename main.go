package main

import (
	"log"
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"time"
	"os/exec"
)

type Task struct {
	Period int32
	Command string
	Output string
}


type TaskArray struct {
	Tasks []Task
}

func main() {
	file, err := os.Open("cronjob.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	var taskArray TaskArray
	//var tasks []Task
	err_json := json.Unmarshal(b, &taskArray)
	if err_json != nil {
		fmt.Println("error:", err_json)
	}
	for {
		for _, element := range taskArray.Tasks {
			if (needToRunNow(element.Period)) {
				fmt.Println(time.Now())
				fmt.Println(element.Command)
				go runCommand(element.Command, element.Output)
			}
		}
			time.Sleep(time.Millisecond*1000)
	}
}

func runCommand(command string, output string) {
	runcmd(command+" >> "+output, true)
}

func runcmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			fmt.Println(err)
		}
		return out
	}
	out, err := exec.Command(cmd).Output()
	if err != nil {
		fmt.Println(err)
	}
	return out
}

func needToRunNow(period int32) bool {
	current:=int32(time.Now().Unix())
	if (current % period == 0) {
		return true
	} else {
		return false
	}

}
