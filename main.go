package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
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
		log.Fatal("Index out of range")
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

func (tl *taskList) remove(index int) {
	if index < 0 && index >= len(tl.Tasks) {
		log.Fatalf("Index out of range")
	}
	tl.Tasks = append(tl.Tasks[:index], tl.Tasks[index+1:]...)
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
			log.Fatal("Not enough arguments")
		}
		time, err := strconv.Atoi(args[3])
		if err != nil {
			log.Fatalf("Failed to convert time-string to integer: %s", err)
		}
		todoList.addTask(task{Title: args[1], Description: args[2], IsDone: false, TimeInMinute: time})
		err = todoList.saveToFile("data.json")
		if err != nil {
			log.Fatalf("Failed to save in file with error: %s", err)
		}
	case "list":
		if len(args) != 1 {
			log.Fatal("False arguments")
		}
		todoList.listTasks()
	case "done":
		if len(args) != 2 {
			log.Fatal("False arguments")
		}
		index, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("Failed to convert index-string to integer: %s", err)
		}
		err = todoList.markDone(index - 1)
		if err != nil {
			log.Fatal(err)
		}
		err = todoList.saveToFile("data.json")
		if err != nil {
			log.Fatal(err)
		}
	case "remove":
		index, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}
		todoList.remove(index - 1)
		err = todoList.saveToFile("data.json")
		if err != nil {
			log.Fatal(err)
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
}
