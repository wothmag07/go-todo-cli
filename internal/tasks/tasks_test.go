package tasks

import (
	"os"
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	// Reset state
	Tasks = []Task{}
	nextID = 1

	// Test adding a valid task
	err := AddTask("Test task")
	if err != nil {
		t.Errorf("AddTask failed: %v", err)
	}

	if len(Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(Tasks))
	}

	if Tasks[0].Name != "Test task" {
		t.Errorf("Expected task name 'Test task', got '%s'", Tasks[0].Name)
	}

	if Tasks[0].IsComplete != StatusIncomplete {
		t.Errorf("Expected status '%s', got '%s'", StatusIncomplete, Tasks[0].IsComplete)
	}

	// Test adding empty task
	err = AddTask("")
	if err == nil {
		t.Error("Expected error for empty task name")
	}

	// Test adding whitespace-only task
	err = AddTask("   ")
	if err == nil {
		t.Error("Expected error for whitespace-only task name")
	}
}

func TestCompleteTask(t *testing.T) {
	// Reset state
	Tasks = []Task{}
	nextID = 1

	// Add a task
	AddTask("Test task")

	// Test completing existing task
	err := CompleteTask(1)
	if err != nil {
		t.Errorf("CompleteTask failed: %v", err)
	}

	if Tasks[0].IsComplete != StatusCompleted {
		t.Errorf("Expected status '%s', got '%s'", StatusCompleted, Tasks[0].IsComplete)
	}

	// Test completing non-existent task
	err = CompleteTask(999)
	if err == nil {
		t.Error("Expected error for non-existent task")
	}
}

func TestDeleteTask(t *testing.T) {
	// Reset state
	Tasks = []Task{}
	nextID = 1

	// Add tasks
	AddTask("Task 1")
	AddTask("Task 2")

	if len(Tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(Tasks))
	}

	// Test deleting existing task
	err := DeleteTask(1)
	if err != nil {
		t.Errorf("DeleteTask failed: %v", err)
	}

	if len(Tasks) != 1 {
		t.Errorf("Expected 1 task after deletion, got %d", len(Tasks))
	}

	// Test deleting non-existent task
	err = DeleteTask(999)
	if err == nil {
		t.Error("Expected error for non-existent task")
	}
}

func TestUpdateTask(t *testing.T) {
	// Reset state
	Tasks = []Task{}
	nextID = 1

	// Add a task
	AddTask("Original task")

	// Test updating existing task
	err := UpdateTask(1, "Updated task")
	if err != nil {
		t.Errorf("UpdateTask failed: %v", err)
	}

	if Tasks[0].Name != "Updated task" {
		t.Errorf("Expected task name 'Updated task', got '%s'", Tasks[0].Name)
	}

	// Test updating non-existent task
	err = UpdateTask(999, "New task")
	if err == nil {
		t.Error("Expected error for non-existent task")
	}

	// Test updating with empty name
	err = UpdateTask(1, "")
	if err == nil {
		t.Error("Expected error for empty task name")
	}
}

func TestTaskStruct(t *testing.T) {
	now := time.Now()
	task := Task{
		ID:        1,
		Name:      "Test task",
		IsComplete: StatusIncomplete,
		CreatedAt: now,
	}

	if task.ID != 1 {
		t.Errorf("Expected ID 1, got %d", task.ID)
	}

	if task.Name != "Test task" {
		t.Errorf("Expected name 'Test task', got '%s'", task.Name)
	}

	if task.IsComplete != StatusIncomplete {
		t.Errorf("Expected status '%s', got '%s'", StatusIncomplete, task.IsComplete)
	}

	if !task.CreatedAt.Equal(now) {
		t.Errorf("Expected creation time %v, got %v", now, task.CreatedAt)
	}
}

// Helper function to clean up test files
func cleanupTestFiles() {
	os.Remove("tasks.csv")
}

func TestMain(m *testing.M) {
	// Clean up before and after tests
	cleanupTestFiles()
	
	// Run tests
	code := m.Run()
	
	// Clean up after tests
	cleanupTestFiles()
	
	os.Exit(code)
} 