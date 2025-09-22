# Go To-Do CLI 📝

A simple, command-line to-do list application written in Go. This tool allows you to manage your tasks directly from the terminal.



---

## Features

* **Add** a new task to your list.
* **List** all current tasks.
* **Mark** a task as complete.
* **Remove** a specific task by its number.
* **Clear** the entire to-do list.

---

## Installation

You need to have **Go** (version 1.18 or higher) installed on your system.

1.  **Clone the repository** (replace with your actual repository URL):
    ```sh
    git clone [https://github.com/your-username/todo-cli.git](https://github.com/your-username/todo-cli.git)
    cd todo-cli
    ```

2.  **Build the application:**
    ```sh
    go build .
    ```

3.  **Move the executable** to a directory in your system's PATH to make it globally accessible:
    ```sh
    sudo mv todo /usr/local/bin/
    ```

---

## Usage

Here are the basic commands for using the tool.

### Add a Task
```sh
todo add "Buy milk and bread"
```

### List all tasks
```sh
todo list
```
### Output
1. Buy milk and bread
2. Finish the README file   (✓)

### Complete a task

Mark a task as complete using its number.

```sh
todo complete 1
```

### Remove a task

Remove a task from the list using its number.

```sh
todo remove 2
```
### Clear All Tasks
This will remove all tasks from your list.

```sh
todo clear
```

### Get Help
Shows a list of all available commands.

```sh
todo
```