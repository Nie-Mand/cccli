package commitmessage_model

import (
	"fmt"
)

func (m CommitMessageModel) View() string {
	s := fmt.Sprintf("Write a short, imperative tense description of the change\n~ %s\n\n", m.message)
	return s
}
