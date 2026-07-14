# Todo CLI

A lightweight command-line task manager written in Go. The application provides a simple interface for managing tasks from the terminal, with support for persistent storage using JSON.

## Features

* Add new tasks
* List all tasks
* Edit existing tasks
* Delete tasks
* Toggle task completion status
* Persistent storage using a local JSON file
* Built entirely with Go's standard library

## Project Structure

```text
.
├── main.go          # Application entry point
├── command.go       # Command-line flag parsing and command execution
├── todo.go          # Todo data model and operations
├── storage.go       # JSON file read/write operations
├── todos.json       # Persistent task storage
├── go.mod
└── README.md
```

## Installation

Clone the repository:

```bash
git clone https://github.com/HardikPrakhar/Learning-Go-.git
cd Learning-Go-
```

Run the application:

```bash
go run .
```

Or build an executable:

```bash
go build -o todo
```

## Usage

### Add a task

```bash
todo -add "Complete Go project"
```

### List all tasks

```bash
todo -list
```

### Edit a task

The edit command expects the format:

```text
index:new title
```

Example:

```bash
todo -edit "2:Practice Go concurrency"
```

### Toggle task completion

```bash
todo -toggle 2
```

### Delete a task

```bash
todo -del 3
```

## Storage

Tasks are stored in `todos.json`. Any changes made through the CLI are automatically saved, ensuring that tasks persist across application runs.

## Technologies

* Go
* `flag`
* `encoding/json`
* `os`
* `fmt`
* `strconv`
* `strings`

## Future Enhancements

* Due dates
* Task priorities
* Search and filtering
* Categories and tags
* Colored terminal output
* Interactive terminal interface (TUI)

## License

This project is licensed under the MIT License.
