package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func checkDomain(domain string) {
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error:%v", err)
	}
	if len(mxRecord) > 0 {
		hasMx = true
	}

	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error:%v", err)
	}
	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error:%v", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Panicf("Error: Could not read the input from user:%v", err)
	}
}
