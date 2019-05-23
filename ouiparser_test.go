package ouiparser

import (
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
	if err != nil {
		t.FailNow()
	}
	result := SearchOUI(oui, mac)
	t.Log(result.Organization)
}
