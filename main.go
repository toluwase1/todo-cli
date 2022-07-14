package main

import (
	"fmt"
)

var todolist []string

func main() {
	welcomePage()

	for {
		var input int
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Your input is not valid, please try again")
			fmt.Println(err)
			continue
		}
		fmt.Println(input)
		switch input {
		case 1:
			title, description := getUserInputToCreateTodo()
			createTodo(title, description)
			fmt.Println(todolist)
			continue
		case 2:
			fmt.Println("block 2")
			id := getUserInputToDeleteTodo()
			deleteTodoById(id)
		case 3:
			fmt.Println("block 3")
			id := getUserInputToDeleteTodo()
			updateTodo(id)
		case 4:
			fmt.Println("block 4")
			ListAllTodo()
		default:
			fmt.Println("enter a correct input value")
		}

	}

}
func welcomePage() {
	fmt.Println("Welcome to our TodoList Application")
	fmt.Println("Please enter 1 to add to a new todo")
	fmt.Println("Please enter 2 to remove from the list")
	fmt.Println("Please enter 3 to update the list")
	fmt.Println("Please enter 4 to view all todos")
}

func getUserInputToCreateTodo() (string, string) {
	var title string
	var description string
	fmt.Println("Enter the title of your todo: ")
	fmt.Scan(&title)

	fmt.Println("Enter the description of your todo: ")
	fmt.Scan(&description)

	return title, description
}

func createTodo(title string, description string) []string {
	todo := title + " " + description
	todolist = append(todolist, todo)
	return todolist
}

func deleteTodoById(value int) {
	for i := 0; i < len(todolist); i++ {
		if i == value {
			//Deletes the item with such index
			todolist = append(todolist[:i], todolist[i+1:]...)
			//left := todolist[:value+1]
			//right := todolist[value:]
			//todolist = append(todolist, left...)
			//todolist = append(todolist, right...)
		}
	}
}

func getUserInputToDeleteTodo() int {
	var id int
	fmt.Println("Enter the id of any todo: ")
	fmt.Scan(&id)
	return id
}

func updateTodo(id int) {
	for i, _ := range todolist {
		if todolist[i] == string(id) {
			todolist = append(todolist[:i], todolist[i+1:]...)
		}
	}
}

func ListAllTodo() {
	for i := 0; i < len(todolist); i++ {
		numberedList := fmt.Sprintf("%d: %s \n", i+1, todolist[i])
		fmt.Println(numberedList)
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
