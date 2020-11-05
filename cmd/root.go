package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "web-cli",
	Short: "Read Atom feeds",
	Long: `Just a small CLI application.
			Read Atom feeds`,
}

var cmdLs = &cobra.Command{
	Use:   "ls",
	Short: "List news",
	Long:  `List first 5 news`,
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

var cmdDescribe = &cobra.Command{
	Use:   "describe [id]",
	Short: "Show details for an article",
	Long:  `Details for an article`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		describe(args[0])
	},
}

var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Serve news",
	Long:  `Serve news on :8080`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func Exec() {
	rootCmd.AddCommand(cmdLs)
	rootCmd.AddCommand(cmdDescribe)
	rootCmd.AddCommand(cmdServe)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

////TODO understand this
//func (n News) Abstract() string {
//	return n.Text[:100]
//}
