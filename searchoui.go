package ouiparser

import (
	"fmt"
	"net"
)

//OUIMap is map contains OUI data
type OUIMap map[string]OUI

//Find finds OUI data matches with given company ID
func (o OUIMap) Find(companyID string) OUI {
	return o[companyID]
}

//OUISlice is slice contains OUI data
type OUISlice []OUI

//Find finds OUI data matches with given company ID
func (o OUISlice) Find(companyID string) OUI {
	for _, d := range o {
		if d.CompanyID == companyID {
			return d
		}
	}
	return OUI{}
}

//OUIData is the interface that wraps the Find method.
type OUIData interface {
	Find(companyID string) OUI
}

//SearchOUI search oui data matches with given net.HardwareAddr
func SearchOUI(oui OUIData, mac net.HardwareAddr) OUI {
	targetOUI := fmt.Sprintf("%02X%02X%02X", mac[0], mac[1], mac[2])
	return oui.Find(targetOUI)
}
