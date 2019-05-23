package ouiparser

import (
	"fmt"
	"net"
)

//SearchOUI search oui data matches with given net.HardwareAddr
func SearchOUI(oui []OUI, mac net.HardwareAddr) OUI {
	targetOUI := fmt.Sprintf("%02X%02X%02X", mac[0], mac[1], mac[2])
	for _, d := range oui {
		if d.CompanyID == targetOUI {
			return d
		}
	}
	return OUI{}
}
