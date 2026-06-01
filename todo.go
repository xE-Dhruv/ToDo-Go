package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Todo struct
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

// storage file
const fileName = "todos.json"

// load todos from file
func loadTodos() []Todo {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return []Todo{}
	}
	var todos []Todo
	json.Unmarshal(file, &todos)
	return todos
}

// save todos to file
func saveTodos(todos []Todo) {
	data, _ := json.MarshalIndent(todos, "", "  ")
	os.WriteFile(fileName, data, 0644)
}

// add a new todo
func addTodo(todos []Todo, task string) []Todo {
	id := 1
	if len(todos) > 0 {
		id = todos[len(todos)-1].ID + 1
	}
	todos = append(todos, Todo{ID: id, Task: task, Done: false})
	fmt.Printf("✅ Added: \"%s\"\n", task)
	return todos
}

// list all todos
func listTodos(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("📭 No todos yet! Add one with: add <task>")
		return
	}
	fmt.Println("\n📋 Your Todos:")
	fmt.Println(strings.Repeat("-", 40))
	for _, t := range todos {
		status := "[ ]"
		if t.Done {
			status = "[✓]"
		}
		fmt.Printf("%s %d. %s\n", status, t.ID, t.Task)
	}
	fmt.Println(strings.Repeat("-", 40))
}

// mark todo as done
func completeTodo(todos []Todo, id int) []Todo {
	for i, t := range todos {
		if t.ID == id {
			todos[i].Done = true
			fmt.Printf("🎉 Completed: \"%s\"\n", t.Task)
			return todos
		}
	}
	fmt.Printf("❌ Todo with ID %d not found\n", id)
	return todos
}

// delete a todo
func deleteTodo(todos []Todo, id int) []Todo {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Printf("🗑️  Deleted: \"%s\"\n", t.Task)
			return todos
		}
	}
	fmt.Printf("❌ Todo with ID %d not found\n", id)
	return todos
}

// print help
func printHelp() {
	fmt.Println("\n📌 Commands:")
	fmt.Println("  add <task>       → Add a new todo")
	fmt.Println("  list             → Show all todos")
	fmt.Println("  done <id>        → Mark todo as complete")
	fmt.Println("  delete <id>      → Delete a todo")
	fmt.Println("  help             → Show commands")
	fmt.Println("  exit             → Quit the app")
	fmt.Println()
}

func main() {
	todos := loadTodos()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("================================")
	fmt.Println("   📝 Go CLI Todo App")
	fmt.Println("================================")
	printHelp()

	for {
		fmt.Print(">> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		// split command and arguments
		parts := strings.SplitN(input, " ", 2)
		command := strings.ToLower(parts[0])

		switch command {
		case "add":
			if len(parts) < 2 || parts[1] == "" {
				fmt.Println("❌ Usage: add <task>")
			} else {
				todos = addTodo(todos, parts[1])
				saveTodos(todos)
			}

		case "list":
			listTodos(todos)

		case "done":
			if len(parts) < 2 {
				fmt.Println("❌ Usage: done <id>")
			} else {
				id, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("❌ ID must be a number")
				} else {
					todos = completeTodo(todos, id)
					saveTodos(todos)
				}
			}

		case "delete":
			if len(parts) < 2 {
				fmt.Println("❌ Usage: delete <id>")
			} else {
				id, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("❌ ID must be a number")
				} else {
					todos = deleteTodo(todos, id)
					saveTodos(todos)
				}
			}

		case "help":
			printHelp()

		case "exit", "quit":
			fmt.Println("👋 Bye!")
			os.Exit(0)

		default:
			fmt.Printf("❌ Unknown command: %s (type 'help' for commands)\n", command)
		}
	}
}