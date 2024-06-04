package commit

import (
	"fmt"

	"github.com/Nie-Mand/cccli/internal/core/domain"
	"github.com/charmbracelet/huh"
)

func (f *CommitForm) makeCommitTypeInput() *huh.Select[domain.CommitType] {
	list := f.CommitTypeGateway.List()

	options := make([]huh.Option[domain.CommitType], 0, len(list))

	for _, commitType := range list {
		options = append(options, huh.NewOption[domain.CommitType](
			fmt.Sprintf("%8s ~ %s", commitType.Id, commitType.Description),
			commitType,
		))
	}

	return huh.NewSelect[domain.CommitType]().
		Title("Select the type of change that you're commiting").
		Options(options...).
		Value(&f.Commit.CommitType)
}
