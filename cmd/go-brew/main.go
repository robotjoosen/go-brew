package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/robotjoosen/go-brew/cmd/go-brew/program"
)

func main() {
	p := tea.NewProgram(program.NewBrewProgram())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
