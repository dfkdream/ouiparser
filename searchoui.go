package ouiparser

import (
	"fmt"
	"net"
	"strings"
)

//SearchOUI search oui data matches with given net.HardwareAddr
func SearchOUI(oui []OUI, mac net.HardwareAddr) OUI {
	targetOUI := fmt.Sprintf("%x%x%x", mac[0], mac[1], mac[2])
	targetOUI = strings.ToUpper(targetOUI)
	for _, d := range oui {
		if d.CompanyID == targetOUI {
			return d
		}
	}
	return OUI{}
}
