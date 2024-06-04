package commit

import (
	"errors"
	"fmt"

	"github.com/Nie-Mand/cccli/internal/core"
	"github.com/Nie-Mand/cccli/internal/core/domain"
	"github.com/charmbracelet/huh"
)

type CommitForm struct {
	Confirm bool

	domain.Commit

	Form *huh.Form

	core.CommitTypeGateway
}

func NewCommitForm(opts ...CommitFormOption) (*CommitForm, error) {
	f := &CommitForm{
		Confirm: true,
		Commit:  domain.Commit{},
	}

	for _, opt := range opts {
		opt(f)
	}

	if f.CommitTypeGateway == nil {
		return nil, errors.New("CommitTypeGateway is required")
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
		fmt.Println("Commit confirmed")
	}
	return nil
}

func (f *CommitForm) makeForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			f.makeCommitTypeInput(),
		),
		huh.NewGroup(
			f.makeCommitMessageInput(),
		),
	)
}

type CommitFormOption func(*CommitForm)

func WithCommitTypeGateway(g core.CommitTypeGateway) func(*CommitForm) {
	return func(f *CommitForm) {
		f.CommitTypeGateway = g
	}
}
