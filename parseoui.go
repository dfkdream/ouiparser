package ouiparser

import (
	"bufio"
	"os"
	"strings"
)

//OUI contains ieee oui data
type OUI struct {
	OUI          string
	CompanyID    string
	Organization string
	Address1     string
	Address2     string
	Country      string
}

func generateOUIScanner(path string) (*bufio.Scanner, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	scanner.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := strings.Index(string(data), "\r\n\r\n"); i >= 0 {
			return i + 4, data[0:i], nil
		}

		if atEOF {
			return len(data), data, nil
		}

		return 0, nil, nil
	})

	return scanner, nil
}

func textToOUI(ouitext string) OUI {
	dSpl := strings.Split(ouitext, "\r\n")

	if len(dSpl) < 4 {
		return OUI{}
	}

	ouiData := dSpl[0][:8]
	organization := dSpl[0][18:]
	companyID := dSpl[1][:6]
	address1 := dSpl[2][4:]
	address2 := dSpl[3][4:]
	country := dSpl[4][4:]

	return OUI{
		OUI:          ouiData,
		CompanyID:    companyID,
		Organization: organization,
		Address1:     address1,
		Address2:     address2,
		Country:      country,
	}
}

//ParseOUI parses ieee oui text file into OUI Slice.
func ParseOUI(path string) (OUISlice, error) {
	scanner, err := generateOUIScanner(path)
	if err != nil {
		return nil, err
	}

	result := make(OUISlice, 0)

	scanner.Scan()

	for scanner.Scan() {
		data := scanner.Text()
		oui := textToOUI(data)
		if oui.OUI == "" {
			continue //Skip empty data: Private
		}
		result = append(result, oui)
	}
	return result, nil
}

//ParseOUIMap parses ieee oui text file into OUI Map.
func ParseOUIMap(path string) (OUIMap, error) {
	scanner, err := generateOUIScanner(path)
	if err != nil {
		return nil, err
	}

	result := OUIMap{}

	scanner.Scan()

	for scanner.Scan() {
		data := scanner.Text()
		oui := textToOUI(data)
		if oui.OUI == "" {
			continue
		}
		result[oui.CompanyID] = oui
	}

	return result, nil
}
