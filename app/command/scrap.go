package command

import (
	"log"
	"strconv"
	"strings"

	"github.com/freekup/product-scrapper/app"
	"github.com/freekup/product-scrapper/app/tools/dwritter"
	"github.com/freekup/product-scrapper/app/tools/scrapper"
	"github.com/freekup/product-scrapper/app/usecase"
	"github.com/spf13/cobra"
)

var defaultSaveTo = "db,csv"
var defaultMaxData = 100
var scrapCmd = &cobra.Command{
	Use:     "scrap",
	Short:   "Scrapping tokopedia's product",
	Long:    `This command used to scrap products data from Tokopedia platform. `,
	Example: "scrap -s=db,csv",
	Run: func(cmd *cobra.Command, args []string) {
		saveTo := cmd.Flag("saveto").Value.String()
		maxData, err := strconv.Atoi(cmd.Flag("max").Value.String())
		if err != nil {
			log.Fatal(err)
		}

		scrappingUsecase := createScrappingUsecase(maxData, saveTo)
		scrappingUsecase.ProcessScrap()
	},
}

func init() {
	scrapCmd.Flags().StringP("saveto", "s", defaultSaveTo, "Use to set location to save")
	scrapCmd.Flags().IntP("max", "m", defaultMaxData, "Use to set location to save")

	rootCmd.AddCommand(scrapCmd)
}

func createScrappingUsecase(maxData int, saveTo string) *usecase.ScrappingUsecase {
	scrapperData := scrapper.CreateScrapper(maxData)

	dataWritter := dwritter.NewDataWritter()
	dataWritter.RegisterWritter("db", dwritter.NewDBWritter(app.App.Repository))
	dataWritter.RegisterWritter("csv", dwritter.NewCSVWritter())

	return usecase.NewScrappingUsecase(strings.Split(saveTo, ","), scrapperData, dataWritter)
}
