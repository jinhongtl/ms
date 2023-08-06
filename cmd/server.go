package cmd

import (
	"ms/app"
	"ms/internal"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动一个Web服务器",
	Long:  `用于启动一个接收HTTP请求的Web服务器`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.InitConfig()
		internal.InitLogger()
		internal.InitDb()

		server := gin.Default()

		app.InitRouter(server)

		err := server.Run(":" + internal.AppConfig.Server.Port)

		if err != nil {
			panic(err)
		}
	},
}
