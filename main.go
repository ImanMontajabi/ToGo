package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
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
		return fmt.Errorf("index out of range")
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

func (tl *taskList) remove(index int) error {
	if index < 0 || index >= len(tl.Tasks) {
		return fmt.Errorf("Index out of range")
	}
	tl.Tasks = append(tl.Tasks[:index], tl.Tasks[index+1:]...)
	return nil
}

func (tl *taskList) listPending() {
	for i, v := range tl.Tasks {
		if !v.IsDone {
			fmt.Printf("%d. [ ] %s | %s | %d\n", i+1, v.Title, v.Description, v.TimeInMinute)
		}
	}
}

func (tl *taskList) runTaskTimer(minutes int) bool {
	total := minutes
	minute := 0
	for i := 0; i <= total*60; i++ {
		percent := float64(i) / float64(total*60)
		progressBar := renderProgressBar(percent, 20)
		if i%60 == 0 && i != 0 {
			minute += 1
		}
		fmt.Printf("\rProgress: %s %d%% (%d/%d min)", progressBar, int(percent*100), minute, total)
		time.Sleep(time.Second)
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		help()
		return
	}

	command := args[0]

	todoList := taskList{}
	_ = todoList.loadFromFile("data.json")

	switch command {
	case "add":
		if len(args) != 4 {
			fmt.Printf("False arguments")
			help()
			os.Exit(1)
		}
		timeInMinute, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Printf("Failed to convert time-string to integer: %s", err)
			os.Exit(1)
		}
		todoList.addTask(task{Title: args[1], Description: args[2], IsDone: false, TimeInMinute: timeInMinute})
		err = todoList.saveToFile("data.json")
		if err != nil {
			fmt.Printf("Failed to save in file with error: %s", err)
			os.Exit(1)
		}
	case "list":
		if len(args) != 1 {
			fmt.Printf("False arguments")
			help()
			os.Exit(1)
		}
		todoList.listTasks()
	case "done":
		if len(args) != 2 {
			fmt.Printf("False arguments")
			help()
			os.Exit(1)
		}
		index, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Failed to convert index-string to integer: %s", err)
			os.Exit(1)
		}
		err = todoList.markDone(index - 1)
		if err != nil {
			fmt.Printf("error in markdone: %s", err)
			os.Exit(1)
		}
		err = todoList.saveToFile("data.json")
		if err != nil {
			fmt.Printf("error in save to file: %s", err)
			os.Exit(1)
		}
	case "remove":
		index, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Failed to convert index-string to integer: %s", err)
			os.Exit(1)
		}
		todoList.remove(index - 1)
		err = todoList.saveToFile("data.json")
		if err != nil {
			fmt.Printf("error in save to file: %s", err)
			os.Exit(1)
		}
	case "pending":
		todoList.listPending()
	case "start":
		if len(args) != 2 {
			fmt.Printf("False arguments")
			help()
			os.Exit(1)
		}
		index, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Failed to convert index-string to integer: %s", err)
			os.Exit(1)
		}
		timeInMinute := todoList.Tasks[index-1].TimeInMinute
		isDone := todoList.runTaskTimer(timeInMinute)
		if isDone {
			todoList.Tasks[index-1].IsDone = true
			err = todoList.saveToFile("data.json")
			if err != nil {
				fmt.Printf("failed to save in file: %s", err)
				os.Exit(1)
			}
		}
	case "help":
		help()
	}
}

func help() {
	fmt.Println("Usage:")
	fmt.Println("  add [title] [description] [time in minute]")
	fmt.Println("  list")
	fmt.Println("  done [index]")
	fmt.Println("  remove [index]")
	fmt.Println("  pending      → list only not-done tasks")
}

func renderProgressBar(percent float64, width int) string {
	done := int(percent * float64(width))
	remaining := width - done
	return fmt.Sprintf("[%s%s]", stringRepeat("#", done), stringRepeat(".", remaining))
}

func stringRepeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
