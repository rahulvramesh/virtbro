package cmd

import (
	"fmt"
	"github.com/libvirt/libvirt-go"
	"github.com/spf13/cobra"
	"virtbro/pkg/db"
)

var machineID int

var listKVMCmd = &cobra.Command{
	Use:   "listkvm",
	Short: "Lists KVM machines on a remote libvirt instance",
	Long:  `Connects to a remote libvirt instance and lists all the KVM machines by machine ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		listKVM(machineID)
	},
}

func init() {
	rootCmd.AddCommand(listKVMCmd)
	listKVMCmd.Flags().IntVarP(&machineID, "id", "i", 0, "Machine ID")
	listKVMCmd.MarkFlagRequired("id")
}

func listKVM(id int) {
	uri, err := db.GetHostByID(id)
	if err != nil {
		fmt.Printf("Failed to get URI for machine ID %d: %v\n", id, err)
		return
	}

	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		fmt.Printf("Failed to connect to %s: %v\n", uri, err)
		return
	}
	defer conn.Close()

	domains, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		fmt.Printf("Failed to list domains: %v\n", err)
		return
	}

	fmt.Println("KVM Machines:")
	for _, domain := range domains {
		name, err := domain.GetName()
		if err != nil {
			fmt.Printf("Failed to get domain name: %v\n", err)
			continue
		}
		fmt.Println(name)
	}
}
