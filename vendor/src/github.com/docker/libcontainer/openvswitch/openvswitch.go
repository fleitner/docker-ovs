package openvswitch

import (
	"fmt"
	"net"
	"os/exec"
)

const (
	IFNAMSIZ       = 16
)

// Add a port to an Open vSwitch bridge.
func AddToBridge(iface, master *net.Interface) error {
	if len(master.Name) >= IFNAMSIZ {
		return fmt.Errorf("Interface name %s too long", master.Name)
	}

	cmd := exec.Command("ovs-vsctl", "add-port", master.Name, iface.Name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Failed to add port to Open vSwitch bridge: %v", err)
	}
	return nil
}

// Remove a port to an Open vSwitch bridge.
func DelFromBridge(iface, master *net.Interface) error {
	if len(master.Name) >= IFNAMSIZ {
		return fmt.Errorf("Interface name %s too long", master.Name)
	}

	cmd := exec.Command("ovs-vsctl", "del-port", master.Name, iface.Name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Failed to remove port %s from Open vSwitch bridge: %v", iface.Name, err)
	}
	return nil
}
