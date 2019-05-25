package ouiparser

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"testing"
	"time"
)

var seed int64
var ouislice OUISlice
var ouimap OUIMap

func init() {
	seed = time.Now().UnixNano()
	var err error
	ouislice, err = ParseOUI("oui.txt")
	if err != nil {
		log.Fatal(err)
	}
	ouimap, err = ParseOUIMap("oui.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func TestParseOUI(t *testing.T) {
	oui, err := ParseOUI("oui.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(oui[0])
	for i, d := range oui {
		if d.OUI == "" {
			t.Log(i, d)
			t.Fail()
		}
	}
}

func TestParseOUIMap(t *testing.T) {
	oui, err := ParseOUIMap("oui.txt")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(oui["D4389C"])
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

func TestSearchOUIMap(t *testing.T) {
	oui, err := ParseOUIMap("oui.txt")
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

func BenchmarkParseOUI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseOUI("oui.txt")
	}
}

func BenchmarkParseOUIMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseOUIMap("oui.txt")
	}
}

func BenchmarkSearchOUI(b *testing.B) {
	rand.Seed(seed)
	l := len(ouislice)
	testMAC := make([]net.HardwareAddr, b.N)
	for i := 0; i < b.N; i++ {
		n := rand.Intn(l)
		macstring := strings.Replace(ouislice[n].OUI, "-", ":", -1)
		macstring += ":00:00:00"
		mac, err := net.ParseMAC(macstring)
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
		testMAC[i] = mac
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := SearchOUI(ouislice, testMAC[i])
		b.Log(result.Organization)
	}
}

func BenchmarkSearchOUIMap(b *testing.B) {
	rand.Seed(seed)
	l := len(ouislice)
	testMAC := make([]net.HardwareAddr, b.N)
	for i := 0; i < b.N; i++ {
		n := rand.Intn(l)
		macstring := strings.Replace(ouislice[n].OUI, "-", ":", -1)
		macstring += ":00:00:00"
		mac, err := net.ParseMAC(macstring)
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
		testMAC[i] = mac
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := SearchOUI(ouimap, testMAC[i])
		b.Log(result.Organization)
	}
}
