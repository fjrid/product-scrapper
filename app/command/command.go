package command

import (
	"fmt"
	"os"

	"github.com/freekup/product-scrapper/app/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "brick-scrap",
	Short: "This App used to Scrap Product from Tokopedia",
	Long: `This App use this library to scrap Tokopedia's products:
- Cobra
- Viper
- Chromedp
- Squirrel

This app created for Brick technical test.

To up the migration you can just use 'up' command.
To start doing scrapping you can just use 'scrap' command
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /.env)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	config.InitializeConfig()
}
