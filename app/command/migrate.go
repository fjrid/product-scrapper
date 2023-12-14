package command

import (
	"errors"
	"log"

	"github.com/freekup/product-scrapper/app"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "This command use to migrate database table",
	Long:  `Use up or down to do migration.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal(errors.New("argument is not valid, please use up or down"))
		}

		command := args[0]

		if command == "up" {
			err := app.App.Repository.CreateTable()
			if err != nil {
				log.Fatal(err)
			}
		} else if command == "down" {
			err := app.App.Repository.Droptable()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(errors.New("invalid argument, please use up or down"))
		}

		log.Println("Success to do migrate")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
