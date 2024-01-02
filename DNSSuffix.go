package main

import (
	"context"
	"log"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/johnstarich/go/dns/scutil"
	"tailscale.com/net/interfaces"
)

/*
Get Interface DNS Search Domains in MAC

# Obtained from scutil package

# scutil --dns
DNS configuration (for scoped queries)

resolver #1
search domain[0] : internal.colortokens.com
nameserver[0] : 8.8.8.8
nameserver[1] : 10.0.102.61
if_index : 11 (en0)
flags    : Scoped, Request A records
reach    : 0x00000002 (Reachable)
*/
func GetMacInterfaceSearchDomain() []string {

	primaryInterface, err := interfaces.DefaultRoute()
	if err != nil {
		log.Println("Error: Failed to Get Default Route : ", err)
		return []string{}
	}

	res, err := scutil.ReadMacOSDNS(context.Background())

	for _, resolver := range res.Resolvers {
		if resolver.InterfaceIndex == primaryInterface.InterfaceIndex {
			// Found SearchDomain/DNSSuffix
			return resolver.SearchDomain
		}
	}

	return []string{}

}

/*
	Get the Primary interface's DNS Suffix(es)

Example below (Connection-specific DNS Suffix)
Wireless LAN adapter Wi-Fi:

	Connection-specific DNS Suffix  . : internal.colortokens.com
	Description . . . . . . . . . . . : Intel(R) Wireless-AC 9260 160MHz
	Physical Address. . . . . . . . . : 5C-87-9C-66-21-FA
	DHCP Enabled. . . . . . . . . . . : Yes
	Autoconfiguration Enabled . . . . : Yes
	IPv4 Address. . . . . . . . . . . : 10.11.50.71(Preferred)
	Subnet Mask . . . . . . . . . . . : 255.255.254.0

....
*/
func GetWindowsInterfaceSearchDomain() []string {

	primaryIfaceAddresses, _, err := GetPrimaryIfaceAddresses()

	if err != nil {
		return []string{}
	}

	hostname, _ := os.Hostname()
	log.Printf(" Default routes = ", primaryIfaceAddresses)

	dnsSuffixes := make(map[string]bool)

	for _, addr := range primaryIfaceAddresses {
		addrStr := addr.String()

		var addrSuffix []string
		// Lookup for network part of full IP
		if strings.Contains(addrStr, "/") {
			addrSuffix, _ = net.LookupAddr(strings.Split(addrStr, "/")[0])
		} else {
			addrSuffix, _ = net.LookupAddr(addrStr)
		}

		for _, suffix := range addrSuffix {
			// Remove the hostname prefix from the DNS Suffix
			final_suffix := strings.TrimPrefix(suffix, hostname)
			final_suffix = strings.TrimPrefix(final_suffix, ".")
			final_suffix = strings.TrimSuffix(final_suffix, ".")
			if len(final_suffix) > 0 {
				dnsSuffixes[final_suffix] = true
			}
		}
	}

	log.Println("RESULT Dns Suffix = ", dnsSuffixes)
	var res []string
	for suffix, _ := range dnsSuffixes {
		res = append(res, suffix)
	}
	return res
}

// Get Primary interface addresses of the machine
func GetPrimaryIfaceAddresses() ([]net.Addr, string, error) {

	var addresses []net.Addr
	interface_name := ""

	primaryInterface, err := interfaces.DefaultRoute()
	if err != nil {
		log.Println("Error: Failed to Get Default Route : ", err)
		return addresses, interface_name, err
	}

	primaryIface, err := net.InterfaceByIndex(primaryInterface.InterfaceIndex)
	if err != nil {
		log.Println("Error: Failed Default Iface index : ", err)
		primaryIface, err = net.InterfaceByName(primaryInterface.InterfaceName)
		interface_name = primaryInterface.InterfaceName
		if err != nil {
			log.Printf("Error: Failed to get Iface By name %s: Err = %v", primaryInterface.InterfaceName, err)
			return addresses, interface_name, err
		}
	}
	interface_name = primaryInterface.InterfaceName

	addresses, err = primaryIface.Addrs()

	if err != nil {
		log.Println("Error: Failed To get primary iface addresses : ", err)
		return addresses, interface_name, err
	}

	log.Println(" Primary Address = %v", addresses)

	return addresses, interface_name, nil
}

func main() {
	var dnsDomains []string
	if runtime.GOOS == "windows" {
		dnsDomains = GetWindowsInterfaceSearchDomain()
	} else if runtime.GOOS == "darwin" {
		dnsDomains = GetMacInterfaceSearchDomain()
	}

	log.Printf("Dns Search Domains = %v", dnsDomains)

	if runtime.GOOS == "windows" {
		CheckWindowsVPN()
	} else if runtime.GOOS == "darwin" {
		CheckMacVPN()
	}
}

func fn() {
	primaryInterface, err := interfaces.DefaultRoute()

	if err != nil {
		log.Println("Didn't get iface xtra")
		return
	}

	log.Println("Primary iface name = ", primaryInterface.InterfaceName)
	log.Println("Primary iface desc = ", primaryInterface.InterfaceDesc)

	xtra, err := interfaces.InterfaceDebugExtras(primaryInterface.InterfaceIndex)
	if err != nil {
		log.Println("Didn't get iface xtra")
		return
	}

	log.Println(" XTra  = ", xtra)

	gw, myip, ok := interfaces.LikelyHomeRouterIP()
	if !ok {
		log.Println("Didn't get likely hmrouter ip")
		return
	}
	log.Println(" Home router IP = ", gw, myip)

	ifaces, err := net.Interfaces()

	for _, iface := range ifaces {
		log.Println("Iface = ", iface.Name)
		log.Println("Iface MAC = ", iface.HardwareAddr)
		log.Println("Iface MAC = ", iface.Index)
		// log.Println(" Iface name = ", iface.InterfaceName)
		// log.Println(" Iface name = ", iface.InterfaceDesc)
	}
}

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}
