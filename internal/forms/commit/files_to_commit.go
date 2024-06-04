package commit

import "github.com/charmbracelet/huh"

func (f *CommitForm) makeFilesToCommitInput() *huh.Input {
	return huh.NewInput().
		Title("Write a short, imperative tense description of the change").
		CharLimit(45).
		Prompt("? ").
		Value(&f.Commit.CommitMessage)
}
