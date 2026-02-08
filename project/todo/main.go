package main

import (
	"bufio"
	"encoding/json"
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
	menus      = []string{
		"\n========TODO========",
		"1.タスクの追加",
		"2.タスク一覧",
		"3.タスクの取得",
		"4.タスク完了",
		"5.タスク削除",
		"6. 終了",
	}
)

func inputTitle() string {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nタイトルを入力してください：")
		sc.Scan()
		title := strings.TrimSpace(sc.Text())
		if title != "" {
			return title
		}
		fmt.Println("タイトルが空です。タイトルを入力してください。")
	}
}

func inputID() int {
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
			fmt.Println("タスクIDが存在しません")
			return 0
		} else {
			return taskIdi
		}
	}
}

func readJson() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("JSONの読み込みに失敗: ", err)
		return
	}

	if len(data) == 0 {
		fmt.Println("data.jsonが空です。新規データで開始します。")
		return
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("JSONのパースに失敗: ", err)
		return
	}

	for _, task := range tasks {
		if task.ID >= nextID {
			nextID = task.ID + 1
		}
	}
}

func outputJson() {
	data, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		fmt.Println("JSONの作成に失敗しました: ", err)
		return
	}

	err = os.WriteFile("data.json", data, 0644)
	if err != nil {
		fmt.Println("JSONファイルの保存に失敗しました: ", err)
	}

	fmt.Println("データを保存しました")
}

func addTask(title string) bool {
	task := Task{ID: nextID, Title: title, Completed: false}
	tasks = append(tasks, task)
	nextID++
	return true
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

func showMenu() {
	for i := range menus {
		fmt.Println(menus[i])
	}
}

func showTasks() {
	tasks := getAllTasks()
	fmt.Println("\n####タスク一覧####")

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

func showTask(id int) {
	task := getTask(id)
	check := " "
	if task.Completed {
		check = "X"
	}
	fmt.Printf("[%d] [%s] %s\n", task.ID, check, task.Title)
}

func deleteTask(id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("タスクを削除しました")
			return true
		}
	}
	fmt.Println("タスクの削除に失敗しました")
	return false
}

func completedTask(id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			if tasks[i].Completed {
				fmt.Println("タスクは既に完了しています")
				return true
			}
			tasks[i].Completed = true
			fmt.Println("タスクを完了しました")
			return true
		}
	}
	fmt.Println("タスクを完了できませんでした")
	return false
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

		if choiceI >= 1 && choiceI <= len(menus) {
			return choiceI
		}
		fmt.Printf("1～%d の数字をを入力してください。\n", len(menus))
	}
}

func runTodo() {
	if _, err := os.Stat("data.json"); err == nil {
		readJson()
		fmt.Println("データを読み込みました")
	}

	for {
		showMenu()
		choice := selectMenu()

		switch choice {
		case 1:
			title := inputTitle()
			addTask(title)
		case 2:
			showTasks()
		case 3:
			id := inputID()
			showTask(id)
		case 4:
			id := inputID()
			completedTask(id)
		case 5:
			id := inputID()
			deleteTask(id)
		case 6:
			fmt.Println("\n【TODOを終了します】")
			outputJson()
			return
		}
	}
}

func main() {
	runTodo()
}
