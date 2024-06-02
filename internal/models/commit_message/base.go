package commitmessage_model

import "fmt"

type CommitMessageModel struct {
	message string
}

func InitCommitMessageModel() CommitMessageModel {
	return CommitMessageModel{
		message: "",
	}
}

func (m CommitMessageModel) Value() string {
	return m.message
}

func (m CommitMessageModel) Summary() string {
	s := fmt.Sprintf("Write a short, imperative tense description of the change\n~ %s\n", m.message)
	return s
}
