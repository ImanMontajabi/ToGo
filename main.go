package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type task struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	IsDone       bool   `json:"isDone"`
	TimeInMinute int    `json:"timeInMinute"`
}

type taskList struct {
	Tasks []task `json:"tasks"`
}

func (tl *taskList) addTask(t task) {
	tl.Tasks = append(tl.Tasks, t)
}

func (tl *taskList) listTasks() {
	for i, v := range tl.Tasks {
		if v.IsDone {
			fmt.Printf("%d. [X] %s | %s | %d\n", i+1, v.Title, v.Description, v.TimeInMinute)
		} else {
			fmt.Printf("%d. [ ] %s | %s | %d\n", i+1, v.Title, v.Description, v.TimeInMinute)
		}
	}
}

func (tl *taskList) markDone(index int) error {
	if index < 0 || index >= len(tl.Tasks) {
		return fmt.Errorf("Index out of bound!")
	}
	tl.Tasks[index].IsDone = true
	return nil
}

func (tl *taskList) saveToFile(fileName string) error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func (tl *taskList) loadFromFile(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, tl)
}

func main() {

}
