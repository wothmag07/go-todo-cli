package tasks

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

type Task struct {
	ID         int
	Name       string
	IsComplete string
	CreatedAt  time.Time
}

var Tasks []Task
var nextID int

const (
	tasksFile        = "tasks.csv"
	StatusIncomplete = "Incomplete"
	StatusCompleted  = "Completed"
)

func getTasksFilePath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// Check if we are in the cmd/todo-list directory
	if filepath.Base(wd) == "todo-list" {
		return filepath.Join(filepath.Dir(wd), "..", tasksFile), nil
	}
	return tasksFile, nil
}

func LoadTasks() error {
	path, err := getTasksFilePath()
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(tasksFile)
			if err != nil {
				return err
			}
			writer := csv.NewWriter(file)
			defer writer.Flush()
			writer.Write([]string{"ID", "Task", "IsComplete", "CreatedAt"})
			return nil
		}
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(records) > 1 {
		Tasks = []Task{}
		for i, record := range records {
			if i == 0 { // Skip header
				continue
			}
			id, _ := strconv.Atoi(record[0])
			createdAt, _ := time.Parse(time.RFC3339, record[3])
			Tasks = append(Tasks, Task{
				ID:         id,
				Name:       record[1],
				IsComplete: record[2],
				CreatedAt:  createdAt,
			})
			if id >= nextID {
				nextID = id + 1
			}
		}
	} else {
		nextID = 1
	}

	return nil
}

func SaveTasks() error {
	path, err := getTasksFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Task", "IsComplete", "CreatedAt"})
	for _, task := range Tasks {
		writer.Write([]string{
			strconv.Itoa(task.ID),
			task.Name,
			task.IsComplete,
			task.CreatedAt.Format(time.RFC3339),
		})
	}
	return nil
}

func AddTask(name string) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("task name cannot be empty")
	}

	task := Task{
		ID:         nextID,
		Name:       strings.TrimSpace(name),
		IsComplete: StatusIncomplete,
		CreatedAt:  time.Now(),
	}
	Tasks = append(Tasks, task)
	nextID++
	fmt.Printf("Added task: \"%s\"\n", name)
	return nil
}

func ListTasks() {
	if len(Tasks) == 0 {
		fmt.Println("No tasks to show.")
		return
	}
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', 0)
	fmt.Fprintln(w, "ID\tTask\tIsComplete\tCreatedAt")
	fmt.Fprintln(w, "--\t----\t------\t---------")
	for _, task := range Tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Name, task.IsComplete, task.CreatedAt.Format(time.RFC822))
	}
	w.Flush()
}

func DeleteTask(id int) error {
	for i, task := range Tasks {
		if task.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			fmt.Printf("Deleted task with ID: %d\n", id)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func CompleteTask(id int) error {
	for i, task := range Tasks {
		if task.ID == id {
			Tasks[i].IsComplete = StatusCompleted
			fmt.Printf("Completed task with ID: %d\n", id)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}
func UpdateTask(id int, newName string) error {
	if strings.TrimSpace(newName) == "" {
		return fmt.Errorf("task name cannot be empty")
	}

	for i, task := range Tasks {
		if task.ID == id {
			Tasks[i].Name = strings.TrimSpace(newName)
			fmt.Printf("Updated task with ID: %d\n", id)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}
