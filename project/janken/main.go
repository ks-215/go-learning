package main

import (
	"fmt"
	"math/rand"
)

func getUserHand() string {
	var userhand string

	for {
		fmt.Print("\nグー, チョキ, パー を入力してください:\n")
		fmt.Scan(&userhand)

		switch userhand {
		case "グー", "チョキ", "パー":
			return userhand
		default:
			fmt.Print("不正な入力です。もう一度入力してください。\n\n")
		}
	}
}

func getComputerHand() string {
	hands := []string{"グー", "チョキ", "パー"}
	hand := hands[rand.Intn(3)]
	return hand
}

func judgeWin(userHand, computerHand string) string {
	if userHand == computerHand {
		return "Draw"
	}

	if (userHand == "グー" && computerHand == "チョキ") ||
		(userHand == "チョキ" && computerHand == "パー") ||
		(userHand == "パー" && computerHand == "グー") {
		return "Win"
	}

	return "Loss"
}

type Record struct {
	Win  int
	Draw int
	Loss int
}

func playGame() Record {
	var record Record

	for {
		userHand := getUserHand()
		computerHand := getComputerHand()

		fmt.Print("========================")
		fmt.Printf("\nあなた: %s\n", userHand)
		fmt.Printf("あいて: %s\n", computerHand)

		judge := judgeWin(userHand, computerHand)

		switch judge {
		case "Win":
			fmt.Print("\nあなたの勝ちです！\n")
			record.Win++
		case "Draw":
			fmt.Print("\n引き分けです！\n")
			record.Draw++
		case "Loss":
			fmt.Print("\nあなたの負けです！\n")
			record.Loss++
		}
		fmt.Print("========================\n")

		var rematch string
		for {
			fmt.Print("\n再戦しますか？\nyes or no で入力してください：\n")
			fmt.Scan(&rematch)

			if rematch == "yes" || rematch == "no" {
				break
			}
		}

		if rematch == "no" {
			return record
		}
	}
}

func main() {
	result := playGame()
	fmt.Print("\n################\n")
	fmt.Print("＜対戦成績＞\n")
	fmt.Printf("勝ち：%d\n", result.Win)
	fmt.Printf("分け：%d\n", result.Draw)
	fmt.Printf("負け：%d\n", result.Loss)
	fmt.Print("################\n")
}
