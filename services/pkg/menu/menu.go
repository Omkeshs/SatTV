package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"

	handler "std/omkesh/Practice/SatTV/pkg/handlers"

	"github.com/fatih/color"
)

//Menu **
func Menu(handler *handler.HandlerStruct) {
user:
	color.Magenta("\nPlease Select Your Account:")
	users, _ := handler.ListUsers()
	color.Magenta("\n\nEnter Your ID : _")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	id, _ := strconv.Atoi(string(char))
	if id > 0 {
		id -= 1
	} else {
		fmt.Println("\nUser ID Cannot be zero")
		goto user
	}
	user := handler.SelectUser(users[id].ID)
	fmt.Println("\nWelcome to D2H:\n", "\n1.View current balance in the account", "\n2.Recharge Account", "\n3.View available packs, channels and services", "\n4.Subscribe to base packs", "\n5.Add channels to an existing subscription", "\n6.Subscribe to special services", "\n7.View current subscription details", "\n8.Update email and phone number for notifications", "\n0.Exit")
	fmt.Printf("\nEnter the option:")
	reader = bufio.NewReader(os.Stdin)
	char, _, err = reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	menu, _ := strconv.Atoi(string(char))
	fmt.Printf("You choose option number: %d\n\n", menu)

	switch menu {
	case 1:
		//View current balance in the account
		handler.ViewBalance(user)
		MenuPrompt(handler)
	case 2:
		//Recharge Account
		handler.Recharge(user)
		MenuPrompt(handler)
	case 3:
		//View available packs
		handler.ViewPacks()
		MenuPrompt(handler)
	case 4:
		//Subscribe to base packs
		handler.SubscribeBasePack(*user)
		MenuPrompt(handler)
	case 5:
		//Add channels
		handler.AddChannel(*user)
		MenuPrompt(handler)
	case 6:
		//Subscribe to special services
		handler.SubscribeSpecialService(*user)
		MenuPrompt(handler)
	case 7:
		//View current subscription details
		handler.ViewSubscription(*user)
		MenuPrompt(handler)
	case 8:
		//Update email and phone number
		handler.UpdateInfo(*user)
		MenuPrompt(handler)
	case 0:
		fmt.Println("Bye!")
		os.Exit(0)
	}

}

//MenuPrompt **
func MenuPrompt(handler *handler.HandlerStruct) {
menu:
	color.Set(color.FgRed)
	fmt.Printf("\nDo Your Want to exit(Y|y to exit or N|n to continue) : ")
	color.Unset()
	// color.Red("\nDo Your Want to exit(Y|y to exit or N|n to continue):")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	switch unicode.ToUpper(char) {
	case 'N':
		Menu(handler)
	case 'Y':
		color.Green("Bye!")
		os.Exit(0)
	default:
		fmt.Println("Enter A Valid Input!")
		goto menu
	}
}
