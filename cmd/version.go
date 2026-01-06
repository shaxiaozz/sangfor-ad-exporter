package cmd

import (
	"fmt"
	"github.com/shaxiaozz/sangfor-ad-exporter/constant"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the version of sangfor-ad-exporter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(constant.Version)
	},
}
