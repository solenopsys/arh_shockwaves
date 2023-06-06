package funcs

import (
	"github.com/miekg/dns"
	"strings"
	"xs/pkg/io"
)

func extractValues(prefix string, strs []string) []string {
	var spValues []string
	for _, str := range strs {
		if strings.HasPrefix(str, prefix) {
			value := strings.TrimPrefix(str, prefix)
			spValues = append(spValues, value)
		}
	}
	return spValues
}

func GetSnRecords(domain string) ([]string, error) {
	records, err := GetTxtRecords(domain)
	if err == nil {
		values := extractValues("sn=", records)
		return values, nil
	}
	return nil, err
}

func GetTxtRecords(domain string) ([]string, error) {
	c := dns.Client{}
	m := dns.Msg{}
	m.SetQuestion(domain+".", dns.TypeTXT)
	r, _, err := c.Exchange(&m, "8.8.8.8:53")
	if err != nil {
		io.Println(err)
		return nil, err
	}
	strings := make([]string, 0)
	for _, a := range r.Answer {
		if t, ok := a.(*dns.TXT); ok {
			strings = append(strings, t.Txt...)
		}
	}
	return strings, nil
}
