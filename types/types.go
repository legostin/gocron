package types
//Task: тип данных для заданий
type Task struct {
	Period    int32
	SleepTime []string
	Time      []string
	DateTime  []string
	SleepDays []string
	Command   string
	Output    string
}
//TaskArray: массив заданий
type TaskArray struct {
	Tasks []Task
}
