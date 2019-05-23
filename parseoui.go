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

//ParseOUI parses ieee oui text file
func ParseOUI(path string) ([]OUI, error) {
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

	result := make([]OUI, 0)

	scanner.Scan()

	for scanner.Scan() {
		data := scanner.Text()
		dSpl := strings.Split(data, "\r\n")
		if len(dSpl) < 4 {
			continue
		}
		ouiData := dSpl[0][:8]
		organization := dSpl[0][18:]
		companyID := dSpl[1][:6]
		address1 := dSpl[2][4:]
		address2 := dSpl[3][4:]
		country := dSpl[4][4:]
		result = append(result, OUI{
			OUI:          ouiData,
			CompanyID:    companyID,
			Organization: organization,
			Address1:     address1,
			Address2:     address2,
			Country:      country,
		})
	}
	return result, nil
}
