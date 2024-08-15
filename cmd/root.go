package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var logo = `
====================================
  ____   ____ _______  __      _______ ________          ________ _____  
 |  _ \ / __ \__   __| \ \    / /_   _|  ____\ \        / /  ____|  __ \ 
 | |_) | |  | | | |     \ \  / /  | | | |__   \ \  /\  / /| |__  | |__) |
 |  _ <| |  | | | |      \ \/ /   | | |  __|   \ \/  \/ / |  __| |  _  / 
 | |_) | |__| | | |       \  /   _| |_| |____   \  /\  /  | |____| | \ \ 
 |____/ \____/  |_|        \/   |_____|______|   \/  \/   |______|_|  \_\
                                                                         
											powered by: backsoul
====================================
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "viewer",
	Aliases: []string{"vw"},
	Short:   "Viewer cli tool",
	Long: logo + `
	Viewer tool
 `,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
