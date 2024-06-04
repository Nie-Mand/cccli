package cmd

import (
	"fmt"
	"log"
	"os"

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
		)

		if err != nil {
			log.Fatal(err)
		}

		err = form.Run()
		if err != nil {
			log.Fatal(err)
		}

		err = utils.Commit(form.Commit.String())

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Commit successful")

		os.Exit(0)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
