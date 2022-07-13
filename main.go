package main

var todolist []string

func main() {

}

func createTodo(todolist []string) []string {
	for i := 0; i < len(todolist); i++ {
		todolist = append(todolist, todolist[i])
	}
	return todolist
}

func deleteTodo(todolist []string, value int) {
	for i := 0; i < len(todolist); i++ {
		if i == value {
			left := todolist[:value+1]
			right := todolist[value:]
			todolist = append(todolist, left...)
			todolist = append(todolist, right...)
		}
	}
}

func updateTodo() {

}

func ListAllTodo() {

}