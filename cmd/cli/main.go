package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"virtbro/pkg/db"
	"virtbro/pkg/tui"
)

var rootCmd = &cobra.Command{
	Use:   "virtbro",
	Short: "VirtBro - KVM Manager",
}

var addCmd = &cobra.Command{
	Use:   "add [name] [address]",
	Short: "Add a new KVM connection",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		address := args[1]
		db.DB.Create(&db.Connection{Name: name, Address: address})
		fmt.Println("Connection added")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all KVM connections",
	Run: func(cmd *cobra.Command, args []string) {
		var connections []db.Connection
		db.DB.Find(&connections)
		for _, conn := range connections {
			fmt.Println(conn.Name, conn.Address)
		}
	},
}

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Start the TUI",
	Run: func(cmd *cobra.Command, args []string) {
		err := tui.StartTUI()
		if err != nil {
			return
		}
	},
}

func main() {
	err := db.InitDB()
	if err != nil {
		return
	}
	rootCmd.AddCommand(addCmd, listCmd, tuiCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
