package ouiparser

import (
	"fmt"
	"net"
	"testing"
)

func TestParseOUI(t *testing.T) {
	oui, err := ParseOUI("oui.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(oui[0])
}

func TestSearchOUI(t *testing.T) {
	oui, err := ParseOUI("oui.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	mac, err := net.ParseMAC("d4:38:9c:00:00:00")
	targetOUI := fmt.Sprintf("%02X%02X%02X", mac[0], mac[1], mac[2])
	t.Log(targetOUI)
	if err != nil {
		t.FailNow()
	}
	result := SearchOUI(oui, mac)
	t.Log(result.Organization)
}
