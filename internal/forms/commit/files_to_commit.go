package commit

import (
	"github.com/charmbracelet/huh"
)

func (f *CommitForm) makeFilesToCommitInput() *huh.MultiSelect[string] {
	list := f.ChangedFilesGateway.List()

	options := make([]huh.Option[string], 0, len(list)+1)

	options = append(options, huh.NewOption[string]("All files", "."))

	for _, file := range list {
		options = append(options, huh.NewOption[string](
			file,
			file,
		))
	}

	return huh.NewMultiSelect[string]().
		Title("Check the files you want to commit").
		Options(options...).
		Value(&f.Commit.FilesChanged)
}
