package main

import (
	"fmt"
	"os"

	"github.com/Nie-Mand/cccli/internal/models"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(models.Init())
	p.SetWindowTitle("Conventional Commits CLI")
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
