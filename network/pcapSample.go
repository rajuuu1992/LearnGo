package main

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

type PcapDevice struct {
	Name       string
	IsLoopback bool
	IPs        map[string]bool
}

var (
	pcapMachineIPs map[string]*bool
)

func InitPcapDevices() {
	pcapMachineIPs = make(map[string]*bool, 100)
}

func main() {
	pcapDevices()
	// res, _ := os.Executable()
	// fmt.Printf("\n\n\n***** Exec = %v", res)
}
func pcapDevices() []*PcapDevice {
	InitPcapDevices()
	devs, err := pcap.FindAllDevs()
	if err != nil {
		return nil
	}

	var devices = []*PcapDevice{}

	for i := range devs {
		fmt.Printf("\n\n\n Device: Name =  %v , Flags = %v, Description = %v #Address = %v",
			devs[i].Name, devs[i].Flags, devs[i].Description, len(devs[i].Addresses))
		var newDevice *PcapDevice
		newDevice = nil
		for _, addr := range devs[i].Addresses {
			fmt.Printf("\n\t IP = %v, Netmask = %v, BroadcastAddr = %v, P2P IP = %v", addr.IP, addr.Netmask, addr.Broadaddr, addr.P2P)
			if addr.IP.IsLoopback() {
				newDevice = new(PcapDevice)
				newDevice.Name = devs[i].Name
				newDevice.IsLoopback = true
				fmt.Printf("\n------ Skipped  By CTAGENT PCAP")
				break
			}
			if addr.IP.IsMulticast() ||
				addr.IP.IsUnspecified() ||
				addr.IP.IsLinkLocalUnicast() ||
				addr.Broadaddr == nil {
				fmt.Printf("\n------ Skipped  By CTAGENT PCAP %v addr Bd = %v as Broadcast Addr = nil\n", devs[i].Name, (addr.Broadaddr == nil))
				// if strings.Contains(addr.IP.String(), ":") {
				// 	continue
				// }
				continue
			}
			if newDevice == nil { // first address for the interface
				newDevice = new(PcapDevice)
				newDevice.Name = devs[i].Name
				newDevice.IPs = make(map[string]bool)
				newDevice.IPs[addr.IP.String()] = true
			} else { // additional addresses for the interface
				newDevice.IPs[addr.IP.String()] = true
			}
		}

		if newDevice != nil {
			devices = append(devices, newDevice)
		}

	}

	fmt.Printf("\n\nFinal Interfaces captured by PCAP")
	for _, i := range devices {
		fmt.Printf("\n \n  %v", i)
	}
	return devices
}
