package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/robotjoosen/go-brew/cmd/rj/rj_program"
)

func main() {
	p := tea.NewProgram(rj_program.NewRJProgram())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
