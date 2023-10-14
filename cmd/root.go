package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-boss",
	Short: "go-boss",
	Long:  `多应用GO语言后台管理系统`,
	Args:  args,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func args(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("至少需要一个参数!")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
