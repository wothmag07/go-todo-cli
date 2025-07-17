# Todo List CLI Application

A simple command-line todo list application built in Go using the Cobra CLI framework.

## Features

- ✅ Add new tasks
- 📋 List all tasks with details
- ✅ Mark tasks as complete
- 🗑️ Delete tasks
- ✏️ Update task names
- 💾 Persistent storage using CSV format

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd todo-list
```

2. Build the application:
```bash
go build -o bin/task cmd/todo-list/main.go
```

3. Add the binary to your PATH or run directly:
```bash
./bin/task
```

## Usage

### Add a new task
```bash
./bin/task add "Complete Go project"
./bin/task add "Learn Docker basics"
```

### List all tasks
```bash
./bin/task list
```

### Mark a task as complete
```bash
./bin/task complete 1
```

### Update a task name
```bash
./bin/task update 1 "Complete Go project with tests"
```

### Delete a task
```bash
./bin/task delete 1
```

### Get help
```bash
./bin/task --help
./bin/task add --help
```

## Project Structure

```
todo-list/
├── cmd/                    # CLI commands
│   ├── add.go             # Add task command
│   ├── complete.go        # Complete task command
│   ├── delete.go          # Delete task command
│   ├── list.go            # List tasks command
│   ├── update.go          # Update task command
│   ├── root.go            # Root command configuration
│   └── todo-list/         # Main application
│       └── main.go        # Entry point
├── internal/              # Private application logic
│   └── tasks/             # Task management
│       └── tasks.go       # Core task operations
├── tasks.csv              # Data storage
├── go.mod                 # Go module file
├── go.sum                 # Go module checksums
└── README.md              # This file
```

## Data Storage

Tasks are stored in a CSV file (`tasks.csv`) with the following format:
- ID: Unique identifier
- Task: Task description
- IsComplete: Status (Incomplete/Completed)
- CreatedAt: Creation timestamp

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Viper](https://github.com/spf13/viper) - Configuration management

## Future Improvements

- [ ] Add task priorities
- [ ] Add due dates
- [ ] Add task categories/tags
- [ ] Add task search functionality
- [ ] Add task filtering (by status, date, etc.)
- [ ] Add task export functionality
- [ ] Add configuration file support
- [ ] Add unit tests
- [ ] Add database support (SQLite/PostgreSQL)
- [ ] Add web interface

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE). 