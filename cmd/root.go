/*
Copyright © 2025 Semen Bakharev <s@bhrv.dev>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "checker",
	Short: "CLI-инструмент для мониторинга DockerSwarm",
	Long: `Лёгкий CLI-инструмент на Go для отображения состояния Docker Swarm сервисов`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

