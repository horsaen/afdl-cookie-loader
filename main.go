package main

import (
	"flag"
	"fmt"
	"horsaen/afdl-cookie-loader/cookies"
	"horsaen/afdl-cookie-loader/views"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	mode := flag.String("mode", "tui", "Mode")
	username := flag.String("username", "", "Username")
	password := flag.String("password", "", "Password")

	flag.Parse()

	switch *mode {
	case "tui":
		if _, err := tea.NewProgram(views.InitialModel()).Run(); err != nil {
			fmt.Printf("could not start program: %s\n", err)
			os.Exit(1)
		}
	case "afreeca":
		cookies.Afreeca(*username, *password)
	case "flex":
		cookies.Flex(*username, *password)
	case "panda":
		cookies.Panda()
	default:
		fmt.Println("Mode not supported.")
		os.Exit(1)
	}

}
