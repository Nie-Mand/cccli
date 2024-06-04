package commit

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func (f *CommitForm) makeConfirmCommitInput() *huh.Confirm {
	return huh.NewConfirm().
		Title("Are you sure you want to commit this change with the following commit message: ").
		Affirmative(fmt.Sprintf("git commit -m \"%s\"", f.Commit.String())).
		Negative("CTR-C").
		Value(&f.Confirm)
}
