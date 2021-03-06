package mail

import (
	"fmt"
	"log"
	"net"
	"strings"
)
// CheckDom takes a domain name as a string, and returns a comma separated string with the domain name,
// whether it has an MX record, whether it has an SPF record, the SPF record, whether it has a DMARC
// record, and the DMARC record
func CheckDom(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var sprRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf") {
			hasSPF = true
			sprRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord)
}
