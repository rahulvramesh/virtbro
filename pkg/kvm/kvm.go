package kvm

import (
    "github.com/libvirt/libvirt-go"
)

func GetConnections() ([]string, error) {
    conn, err := libvirt.NewConnect("qemu:///system")
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    domains, err := conn.ListAllDomains(0)
    if err != nil {
        return nil, err
    }

    var domainNames []string
    for _, domain := range domains {
        name, err := domain.GetName()
        if err != nil {
            return nil, err
        }
        domainNames = append(domainNames, name)
    }

    return domainNames, nil
}
