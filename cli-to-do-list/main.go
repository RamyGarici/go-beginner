package main
import ("time"
"fmt"
"os"
"bufio")

type Task struct{
    ID int `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	
}

func createTask(taskName string)error{
	newTask := Task{
		ID: len(taskList)+1,
		Name: taskName,
		Done : false,
		CreatedAt: time.Now(),

	}
	taskList = append(taskList, newTask)
	fmt.Println("Task successfully added")
    return nil


}


var taskList = make([]Task,0)


func main(){
	scanner:= bufio.NewScanner(os.Stdin)
for{
	fmt.Println("Welcome to your ToDo list")
	fmt.Println("Enter 1 to create a task")
	fmt.Println("Enter 2 to list tasks")
	fmt.Println("Enter 3 to delete a task")
	fmt.Println("Enter 4 to mark a finished task")
	fmt.Println("Enter 5 to stop the program")
	if !scanner.Scan(){
		continue
	}
	i := scanner.Text()
    switch i {
	    case "1":
			fmt.Println("Enter your task name:")
			var taskName string
			scan := scanner.Scan()
			if scan == false{
				fmt.Printf("%v",scanner.Err())
				continue
			}
			taskName = scanner.Text()
			createTask((taskName))
			
		case "2":
			fmt.Println("Task list")
		case "3":
			fmt.Println("Which task do you want to delete?")
		case "4":
			fmt.Println("Which task do you want to mark as finished?")
		case "5":
			fmt.Println("Thank you!")
			os.Exit(0)
		default:
			fmt.Println("Wrong option")

	}

}

}