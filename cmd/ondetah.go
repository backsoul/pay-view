package cmd

import (
	"log"

	"github.com/fatih/color"

	internal "github.com/backsoul/viewer/internal/worker"
	"github.com/spf13/cobra"
)

const (
	// flag names
	TrackingID = "id"
)

var ondetahCmd = &cobra.Command{
	Use:   "ondetah",
	Short: "Backsoul bot command",
	Long:  logo + ``,
	Run: func(cmd *cobra.Command, args []string) {
		TrackingID := cmd.Flag(TrackingID).Value.String()
		RunOndetahView(TrackingID)
	},
}

func init() {
	ondetahCmd.PersistentFlags().StringP(TrackingID, "t", "", "platform for do view, yt,twitch or kick [required]")
	rootCmd.AddCommand(ondetahCmd)
}

func RunOndetahView(trackingID string) ([]string, error) {
	color.Green("getting tracking statuses ondetah: %s", trackingID)
	statuses, err := internal.RunBrowserOndetah("https://ondetah-cliente.dev.uxsolutions.com.br/DF/" + trackingID)
	if err != nil {
		log.Fatalf("Error al obtener status ondetah trackingID: %v, error: %s", trackingID, err)
	}
	return statuses, nil
}
