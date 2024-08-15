package cmd

import (
	"log"
	"strconv"
	"sync"

	"github.com/fatih/color"

	internal "github.com/backsoul/viewer/internal/worker"
	pkg "github.com/backsoul/viewer/pkg/proxies"
	"github.com/spf13/cobra"
)

const (
	// flag names
	ID       = "id"
	Platform = "twitch"
	Proxies  = "proxies"
)

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "Backsoul bot command",
	Long:  logo + ``,
	Run: func(cmd *cobra.Command, args []string) {
		platform := cmd.Flag(Platform).Value.String()
		id := cmd.Flag(ID).Value.String()
		numberProxies := cmd.Flag(Proxies).Value.String()
		number, _ := strconv.Atoi(numberProxies)
		RunBotView(id, platform, number)
	},
}

func init() {
	botCmd.PersistentFlags().StringP(ID, "c", "", "id channel [required]")
	botCmd.PersistentFlags().Int(Proxies, 0, "numbers proxies, views intents [required]")
	botCmd.PersistentFlags().StringP(Platform, "p", "", "platform for do view, yt,twitch or kick [required]")
	rootCmd.AddCommand(botCmd)
}

func RunBotView(id string, platform string, numberProxies int) {
	proxies, err := pkg.GetProxies(numberProxies)
	if err != nil {
		log.Fatalf("Error al obtener proxies: %v", err)
	}

	url := ""
	switch platform {
	case "twitch":
		url = "https://www.twitch.tv/" + id
	case "yt":
		url = "https://www.youtube.com/watch?v=" + id
	case "kick":
		url = "https://kick.com/" + id
	default:
		url = "https://www.twitch.tv/bkscode"
	}

	color.Green("id: %s, platform: %s, proxies: %s", id, platform, len(proxies))
	var wg sync.WaitGroup
	for _, proxy := range proxies {
		wg.Add(1) // Incrementa el contador del WaitGroup

		go func(proxy string) {
			defer wg.Done() // Decrementa el contador cuando la goroutine termina

			err := internal.RunBrowser(proxy, url)
			if err != nil {
				color.Red("proxy: %s ,error: %s", proxy, err)
			} else {
				color.Green("proxy: %s completed successfully", proxy)
			}
		}(proxy)
	}

	wg.Wait() // Espera a que todas las goroutines terminen
}
