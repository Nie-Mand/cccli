package committype_model

import (
	"fmt"

	"github.com/Nie-Mand/cccli/internal/core/domain"
	"github.com/Nie-Mand/cccli/internal/utils"
)

func (m CommitTypeModel) View() string {
	s := fmt.Sprintf("Select the type of change that you're commiting ~ %s\n\n", m.query)

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

	if len(filtered) == 0 {
		s += "No matches found"
		return s
	}

	windowStart, windowEnd := utils.GetWindow(len(filtered), m.cursor, m.pageSize)
	for idx, choice := range filtered[windowStart:windowEnd] {
		i := idx + windowStart

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %-8s: %s\n", cursor, choice.Id, choice.Description)
	}

	if len(filtered) > m.pageSize {
		s += fmt.Sprintf("\n%d/%d", m.cursor+1, len(filtered))
	}

	return s
}
