package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"virtbro/pkg/db"
)

var listHostsCmd = &cobra.Command{
	Use:   "listhosts",
	Short: "List all remote hosts",
	Long:  `List all remote hosts added to the local database.`,
	Run: func(cmd *cobra.Command, args []string) {
		listHosts()
	},
}

func init() {
	rootCmd.AddCommand(listHostsCmd)
}

func listHosts() {
	hosts, err := db.ListHosts()
	if err != nil {
		fmt.Printf("Failed to list hosts: %v\n", err)
		return
	}
	for _, host := range hosts {
		fmt.Printf("ID: %s, Name: %s, URI: %s, UUID: %s\n", host["id"], host["name"], host["uri"], host["uuid"])
	}
}
