package openvswitch

import (
	"fmt"
	"net"
	"os/exec"
)

const (
	IFNAMSIZ       = 16
)

// Create the actual Open vSwitch device.
func CreateBridge(name string) error {
	if len(name) >= IFNAMSIZ {
		return fmt.Errorf("Interface name %s too long", name)
	}

	cmd := exec.Command("ovs-vsctl", "add-br", name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Failed to create the bridge: %v", err)
	}

	return nil
}

// Delete the actual Open vSwitch bridge device.
func DeleteBridge(name string) error {
	if len(name) >= IFNAMSIZ {
		return fmt.Errorf("Interface name %s too long", name)
	}

	cmd := exec.Command("ovs-vsctl", "del-br", name)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Failed to delete the bridge: %v", err)
	}

	return nil
}

// Add a slave to a bridge device.
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

