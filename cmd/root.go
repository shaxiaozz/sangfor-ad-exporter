package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use:   "sangfor-ad-exporter",
	Short: "sangfor-ad-exporter",
	Long:  `[ sangfor-ad-exporter ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// 没有子命令时打印帮助
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
