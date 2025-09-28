# Task Tracker 
A lightweight command-line tool to manage your tasks.  
Add, update, delete, and track progress — all without leaving the terminal.

Data is saved automatically in `task.json`.

---

## Installation

```bash
git clone https://github.com/ludikfaineant/task-tracker-cli.git
cd task-tracker-cli &&  go build -o task-cli ./cmd/main.go
```

## Usage

### Adding a new task
```bash
./task-cli add "Buy groceries"
```
### Updating and deleting tasks
```bash
./task-cli update 1 "Buy groceries and cook dinner"
./task-cli delete 1
```
### Marking a task as in progress or done
```bash
./task-cli mark-in-progress 1
./task-cli mark-done 1
```
### Listing all tasks
```bash
./task-cli list
```
### Listing tasks by status
```bash
./task-cli list done
./task-cli list todo
./task-cli list in-progress
```

## Testing
```bash
go test -v ./...
```


