package commit

import (
	"github.com/Nie-Mand/cccli/internal/core/domain"
	"github.com/charmbracelet/huh"
)

func (f *CommitForm) makeCommitEmojiInput() *huh.Select[domain.CommitEmoji] {
	list := f.CommitEmojiGateway.List()

	options := make([]huh.Option[domain.CommitEmoji], 0, len(list))

	for _, emoji := range list {
		options = append(options, huh.NewOption[domain.CommitEmoji](
			emoji.Emoji+" "+emoji.Description,
			emoji,
		))
	}

	return huh.NewSelect[domain.CommitEmoji]().
		Title("Choose an gitmoji.").
		Options(options...).
		Value(&f.Commit.CommitEmoji)
}
