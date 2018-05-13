package types


type Task struct {
	Period int32
	SleepTime []string
	Time []string
	DateTime []string
	SleepDays []string
	Command string
	Output string
}


type TaskArray struct {
	Tasks []Task
}