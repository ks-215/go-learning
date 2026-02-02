package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var (
	tasks  []Task
	nextID int = 1
)

func getTitle() string {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nタイトルを入力してください：\n")
		sc.Scan()
		title := strings.TrimSpace(sc.Text())
		if title != "" {
			return title
		}
		fmt.Print("タイトルが空です。タイトルを入力してください。\n")
	}
}

func addTask(title string) {
	task := Task{ID: nextID, Title: title, Completed: false}
	tasks = append(tasks, task)
	nextID++
}

func getTask(id int) Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return tasks[i]
		}
	}
	return Task{}
}

func getAllTasks() []Task {
	if len(tasks) == 0 {
		return []Task{}
	}
	return tasks
}

func testTitles() []string {
	var testTitle []string
	for i := range 10 {
		title := fmt.Sprintf("test title %d", i)
		testTitle = append(testTitle, title)
	}
	return testTitle
}

func showMenu() {
	fmt.Println("\n========TODO========")
	fmt.Println("1.タスクの追加")
	fmt.Println("2.タスク一覧")
	fmt.Println("3.タスクの取得")
	fmt.Println("4. 終了")
}

func showTasks() {
	tasks := getAllTasks()
	fmt.Println("####タスク一覧####")

	if len(tasks) == 0 {
		fmt.Println("タスクが存在していません")
		return
	}

	for _, task := range tasks {
		check := " "
		if task.Completed {
			check = "X"
		}
		fmt.Printf("[%d] [%s] %s\n", task.ID, check, task.Title)
	}
}

func showTask() {
	for {
		sc := bufio.NewScanner(os.Stdin)
		fmt.Println("\nタスクIDを入力してください")
		sc.Scan()

		taskIds := strings.TrimSpace(sc.Text())
		taskIdi, err := strconv.Atoi(taskIds)

		if err != nil {
			fmt.Println("数字を入力してください")
			continue
		}

		if taskIdi < 1 || taskIdi >= nextID {
			fmt.Println("タスクIDが存在しません。")
			return
		} else {
			task := getTask(taskIdi)
			check := " "
			if task.Completed {
				check = "X"
			}
			fmt.Printf("[%d] [%s] %s\n", task.ID, check, task.Title)
			return
		}
	}
}

func selectMenu() int {
	for {
		sc := bufio.NewScanner(os.Stdin)
		fmt.Println("\n選択：")
		sc.Scan()

		choiceS := strings.TrimSpace(sc.Text())
		choiceI, err := strconv.Atoi(choiceS)

		if err != nil {
			fmt.Println("数字を入力してください")
			continue
		}

		switch choiceI {
		case 1, 2, 3, 4:
			return choiceI
		default:
			fmt.Println("１～４の数字をを入力してください。")
		}
	}
}

func runTodo() {
	for {
		showMenu()
		choice := selectMenu()

		switch choice {
		case 1:
			title := getTitle()
			addTask(title)
		case 2:
			showTasks()
		case 3:
			showTask()
		case 4:
			fmt.Println("\n【TODOを終了します】")
			return
		}
	}
}

// func testRun() {
// 	titles := testTitles()

// 	for _, title := range titles {
// 		addTask(title)
// 	}

// 	oneTask := getTask(3)
// 	fmt.Printf("ID: %d, Title: %s, Completed: %v\n\n", oneTask.ID, oneTask.Title, oneTask.Completed)

// 	allTasks := getAllTasks()
// 	for _, task := range allTasks {
// 		fmt.Println(task)
// 	}
// }

func main() {
	runTodo()
}
