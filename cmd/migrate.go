package cmd

import (
	"ms/app/models"
	"ms/internal"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "数据表迁移工具",
	Long:  `主要用于将在项目创建时初始化数据表`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.InitConfig()

		internal.InitDb()

		internal.Db.AutoMigrate(&models.User{})
	},
}
