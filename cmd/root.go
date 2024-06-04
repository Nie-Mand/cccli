package cmd

import (
	"os"

	"github.com/charmbracelet/log"

	changedfiles "github.com/Nie-Mand/cccli/pkg/changed_files"
	commitemojis "github.com/Nie-Mand/cccli/pkg/commit_emojis"
	committypes "github.com/Nie-Mand/cccli/pkg/commit_types"

	"github.com/Nie-Mand/cccli/internal/forms/commit"
	"github.com/Nie-Mand/cccli/internal/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cccli",
	Short: "Conventional Commit CLI",
	Long:  `A CLI tool to help you write conventional commits easily.`,
	Run: func(cmd *cobra.Command, args []string) {
		form, err := commit.NewCommitForm(
			commit.WithCommitTypeGateway(committypes.NewCommitTypeRepository()),
			commit.WithCommitEmojiGateway(commitemojis.NewCommitEmojiRepository()),
			commit.WithChangedFilesGateway(changedfiles.NewChangedFilesRepository()),
		)

		if err != nil {
			log.Error(err)
			return
		}

		err = form.Run()
		if err != nil {
			log.Error(err)
			return
		}

		err = utils.Add(form.Commit.FilesChanged...)

		if err != nil {
			log.Error(err)

			return
		}

		err = utils.Commit(form.Commit.String())

		if err != nil {
			log.Error(err)
			return
		}

		log.Info("Commited 🚀")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
