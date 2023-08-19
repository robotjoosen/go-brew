package main

import (
	"fmt"
	"github.com/robotjoosen/go-brew/cmd/rj/rj_program"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(rj_program.NewRJProgram())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
