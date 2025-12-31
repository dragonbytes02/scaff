/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/dragonbytes039/scaff/internal/app"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var projectName string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a project",
	Long:  `Create a project by passing the name as a parameter; if you don't pass a positional parameter as the name, it will fail!`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		p := tea.NewProgram(app.InitialApp(projectName))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringVarP(&projectName, "name", "n", "project", "Name for the new project")
}
