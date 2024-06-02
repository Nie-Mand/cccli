package commands

import (
	"fmt"

	"github.com/Nie-Mand/cccli/internal/core/domain"
	tea "github.com/charmbracelet/bubbletea"
)

func CommitCommand(commit domain.Commit) func() tea.Msg {
	return func() tea.Msg {
		fmt.Println("| Commit command")
		fmt.Println("Data: ", commit)
		return tea.Quit()
	}
}
