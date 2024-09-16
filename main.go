package main

import (
	"fmt"
	"strconv"
)

var todoItemList []string

func addNewTodo() {
	fmt.Println("adding a new todo")
	todo()

}

func deleteElement(deleteTodo []string, index int) []string {
	return append(deleteTodo[:index], deleteTodo[index+1:]...)
}
func deleteTodo(todoList []string) {

	var choice string

	for i := 0; i < len(todoList); i++ {

		fmt.Println(i, todoList[i])
	}

	fmt.Println("which element is done :")
	fmt.Scan(&choice)
	index, _ := strconv.Atoi(choice)
	if index > len(todoList) {
		fmt.Println("index is not in the list")
	}
	todoItemList = deleteElement(todoList, index)
	fmt.Println(todoItemList)
}

func todo() {
	var todoTask string
	var choice string

	var addTodo, deleteTodoItem = 1, 2
	for {
		fmt.Print("Insert Task ")
		fmt.Scan(&todoTask)
		todoItemList = append(todoItemList, todoTask)
		fmt.Println(todoTask)
		fmt.Println(todoItemList)
		fmt.Println("Choice 1 to add 2 for delete ")
		fmt.Scan(&choice)
		compare, _ := strconv.Atoi(choice)
		if addTodo == compare {
			addNewTodo()
		} else if deleteTodoItem == 2 {
			deleteTodo(todoItemList)
		} else {
			break
		}
	}
}

func main() {
	fmt.Println("ToDo")

	todo()
}
