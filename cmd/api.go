package cmd

import (
	"fmt"
	"net/http"

	internal "github.com/backsoul/viewer/internal/worker"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Backsoul bot command",
	Long:  logo + ``,
	Run: func(cmd *cobra.Command, args []string) {
		InitializeServer()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func InitializeServer() {
	r := gin.Default()

	r.GET("/api/:trackingID", func(c *gin.Context) {
		trackingID := c.Param("trackingID")
		statuses, err := internal.RunBrowserOndetah("https://dafiti.ondetah.com.br/DF/" + trackingID)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error al obtener status ondetah trackingID: %v, error: %s", trackingID, err)
			return
		}
		c.JSON(http.StatusOK, statuses)
	})
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "¡Hola, mundo!")
	})

	fmt.Println("Iniciando servidor en http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
