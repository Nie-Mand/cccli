package committype_model

import (
	"fmt"

	"github.com/Nie-Mand/cccli/internal/core/domain"
	"github.com/Nie-Mand/cccli/internal/utils"
)

func (m CommitTypeModel) View() string {
	s := "\nPress q to quit.\n"
	s += fmt.Sprintf("Select the type of change that you're commiting ~ %s\n\n", m.query)

	filtered := make([]domain.CommitType, 0, len(m.choices))

	if m.query != "" {
		for _, choice := range m.choices {
			if m.query != "" && choice.Matches(m.query) {
				filtered = append(filtered, choice)
			}
		}
	} else {
		filtered = m.choices
	}

	windowStart, windowEnd := utils.GetWindow(len(filtered), m.cursor, m.pageSize)
	for i, choice := range filtered[windowStart:windowEnd] {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		// Render the row
		s += fmt.Sprintf("%s %s: %s\n\t%s\n", cursor, choice.Id, choice.Name, choice.Description)
	}

	// Send the UI for rendering
	return s
}
