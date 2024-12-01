package main

import (
	"demo/password/account"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("__Менеджер паролей__")
	vault := account.NewVault()
Menu:
	for {
		variant := getMenu()
		fmt.Scanln(&variant)
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант: ")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.Vault) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccountByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
		return
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promptData("Введите URL для поиска и удаления")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
		return
	}
	color.Red("Не найдено")
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
