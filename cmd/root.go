/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"time"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/yardbirdsax/pomo/pomo"
)

var (
	durationMinutes int16
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pomo",
	Short: "A Pomodoro timer written in Go",
	Long: `A Pomodoro timer written in Go. It will display a countdown timer with a default
length of 25 minutes.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	RunE: func(cmd *cobra.Command, args []string) error {
		m := pomo.NewPomo(time.Duration(durationMinutes) * time.Minute)

		if err := tea.NewProgram(m, tea.WithAltScreen()).Start(); err != nil {
			return err
		}
		return nil
	},
}

func Execute(version string) {
	rootCmd.Version = version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Int16VarP(&durationMinutes, "duration", "d", 25, "Indicates the duration of the countdown in minutes.")
}


