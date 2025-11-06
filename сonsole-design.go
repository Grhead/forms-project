package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func main() {
	printHeader()

	prompt := promptui.Select{
		Label: "Выберите действие",
		Items: []string{"Создать форму", "Добавить вопросы", "Получить ответы", "Выход"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	switch result {
	case "Создать форму":
		action("Создание формы...", color.FgCyan)
	case "Добавить вопросы":
		action("Добавление вопросов...", color.FgYellow)
	case "Получить ответы":
		action("Получение ответов...", color.FgGreen)
	default:
		color.Yellow("Завершение работы.")
		os.Exit(0)
	}
}

func printHeader() {
	header := color.New(color.FgHiCyan, color.Bold)
	header.Println("╔════════════════════════════════════════╗")
	header.Println("║               TUSUR Forms              ║")
	header.Println("╚════════════════════════════════════════╝")
}

func action(msg string, c color.Attribute) {
	clr := color.New(c, color.Bold)
	clr.Printf("→ %s\n", msg)
	//

	fmt.Println()
	color.Green("✔ Готово!\n")
}
