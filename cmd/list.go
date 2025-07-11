package cmd

import (
	"fmt"

	"github.com/bakharevd/go-swarm-checker/internal/docker"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Показать список сервисов в Docker Swarm",
	Run: func(cmd *cobra.Command, args []string) {
		services, err := docker.GetSwarmServices()
		if err != nil {
			fmt.Println("Ошибка при получении сервисов:", err)
			return
		}

		fmt.Printf("%-25s %-20s %-10s %-10s %-10s\n", "NAME", "IMAGE", "REPLICAS", "MODE", "STATUS")
		fmt.Println("--------------------------------------------------------------------------")
		for _, svc := range services {
			fmt.Printf("%-25s %-20s %-10s %-10s %s%-10s\033[0m\n",
				svc.Name, svc.Image, svc.Replicas, svc.Mode, svc.Color, svc.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
