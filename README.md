# ⏳ ToGo - Task Manager CLI App in Go

A simple command-line todo manager written in Go, with JSON persistence and a built-in timer with progress bar for each task.

## 🚀 Features

- ✅ Add, list, mark done, and remove tasks
- 🕒 Set estimated time for each task (in minutes)
- ⏳ Timer mode with progress bar that visualizes task duration
- 💾 Data saved in `data.json` file (no DB required)
- 📋 View only pending tasks

---

## 🛠️ Installation

Make sure you have Go installed. Then:

```bash
git clone https://github.com/yourusername/ToGo.git
cd ToGo
go run main.go [command]
```

---

## 🧪 Usage

### ➕ Add a new task

```bash
go run main.go add "Task Title" "Some description" 45
```

### 📋 List all tasks

```bash
go run main.go list
```

### ⏳ Show only pending tasks

```bash
go run main.go pending
```

### ✅ Mark a task as done (by index)

```bash
go run main.go done 2
```

### 🗑️ Remove a task (by index)

```bash
go run main.go remove 3
```

### ▶️ Start task timer (with progress bar)

```bash
go run main.go start 1
```

This simulates a countdown timer for the task with index 1.

---

## 🧩 Example Output

```text
1. [ ] Write report | For school project | 30
2. [X] Buy milk     | From grocery store | 10

Progress: [##########..........] 50% (15/30 min)
```

---

## 📁 Data Storage

All tasks are saved into a JSON file `data.json` in the project root directory.

---

## 📎 License

[MIT License](LICENSE)

---

## ✨ Author

Made with ❤️ by **Iman Montajabi**
