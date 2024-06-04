package commit

import (
	"errors"

	"github.com/Nie-Mand/cccli/internal/core"
	"github.com/Nie-Mand/cccli/internal/core/domain"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

type CommitForm struct {
	Confirm bool

	domain.Commit

	Form *huh.Form

	core.CommitTypeGateway
	core.CommitEmojiGateway
	core.ChangedFilesGateway
}

func NewCommitForm(opts ...CommitFormOption) (*CommitForm, error) {
	f := &CommitForm{
		Confirm: true,
		Commit: domain.Commit{
			CommitType: domain.CommitType{
				Id: "",
			},
			CommitMessage: "",
		},
	}

	for _, opt := range opts {
		opt(f)
	}

	if f.CommitTypeGateway == nil {
		return nil, errors.New("CommitTypeGateway is required")
	} else if f.CommitEmojiGateway == nil {
		return nil, errors.New("CommitEmojiGateway is required")
	} else if f.ChangedFilesGateway == nil {
		return nil, errors.New("ChangedFilesGateway is required")
	}

	f.Form = f.makeForm()

	return f, nil
}

func (f *CommitForm) Run() error {
	if err := f.Form.Run(); err != nil {
		return err
	}

	if err := f.makeConfirmCommitInput().Run(); err != nil {
		return err
	}

	if !f.Confirm {
		return errors.New("Commit aborted")
	} else {

		log.Info("Commit confirmed")
	}
	return nil
}

func (f *CommitForm) makeForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			f.makeFilesToCommitInput(),
		),
		huh.NewGroup(
			f.makeCommitTypeInput(),
		),
		huh.NewGroup(
			f.makeCommitMessageInput(),
		),
		huh.NewGroup(
			f.makeCommitEmojiInput(),
		),
	)
}

type CommitFormOption func(*CommitForm)

func WithCommitTypeGateway(g core.CommitTypeGateway) func(*CommitForm) {
	return func(f *CommitForm) {
		f.CommitTypeGateway = g
	}
}

func WithCommitEmojiGateway(g core.CommitEmojiGateway) func(*CommitForm) {
	return func(f *CommitForm) {
		f.CommitEmojiGateway = g
	}
}

func WithChangedFilesGateway(g core.ChangedFilesGateway) func(*CommitForm) {
	return func(f *CommitForm) {
		f.ChangedFilesGateway = g
	}
}
