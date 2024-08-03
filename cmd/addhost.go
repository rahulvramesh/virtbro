package cmd

import (
	"crypto/rand"
	"fmt"
	"github.com/libvirt/libvirt-go"
	"github.com/spf13/cobra"
	"virtbro/pkg/db"
)

var (
	hostName string
	hostURI  string
)

var addHostCmd = &cobra.Command{
	Use:   "addhost",
	Short: "Add a new remote host",
	Long:  `Add a new remote host by providing the URI and a machine name.`,
	Run: func(cmd *cobra.Command, args []string) {
		addHost(hostName, hostURI)
	},
}

func init() {
	rootCmd.AddCommand(addHostCmd)
	addHostCmd.Flags().StringVarP(&hostName, "name", "n", "", "Name of the remote host")
	addHostCmd.Flags().StringVarP(&hostURI, "uri", "u", "", "URI of the remote host")
	addHostCmd.MarkFlagRequired("name")
	addHostCmd.MarkFlagRequired("uri")
}

func addHost(name, uri string) {
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		fmt.Printf("Failed to connect to %s: %v\n", uri, err)
		return
	}
	defer conn.Close()

	uuid := generateUUID()
	err = db.AddHost(name, uri, uuid)
	if err != nil {
		fmt.Printf("Failed to add host: %v\n", err)
		return
	}
	fmt.Printf("Successfully added host: %s with UUID: %s\n", name, uuid)
}

func generateUUID() string {
	u := make([]byte, 16)
	rand.Read(u)
	u[8] = 0x80
	u[4] = 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}
