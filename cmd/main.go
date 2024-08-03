package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		tui.StartTUI()
	},
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server",
	Run: func(cmd *cobra.Command, args []string) {
		router := gin.Default()

		router.GET("/connections", func(c *gin.Context) {
			var connections []db.Connection
			db.DB.Find(&connections)
			c.JSON(200, connections)
		})

		router.POST("/connections", func(c *gin.Context) {
			var connection db.Connection
			if err := c.ShouldBindJSON(&connection); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			db.DB.Create(&connection)
			c.JSON(200, connection)
		})

		router.Run(":8080")
	},
}

func main() {
	db.InitDB()
	rootCmd.AddCommand(addCmd, listCmd, tuiCmd, serveCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
