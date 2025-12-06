package dns

import (
	"context"
	"net"
	"strings"
	"sync"
)

const (
	TypeA     = "A"
	TypeAAAA  = "AAAA"
	TypeCNAME = "CNAME"
)

type HostRecords struct {
	Host    string
	Records []Record
}

func (hostRecords HostRecords) String() string {
	var joinRecords = ""
	for i, record := range hostRecords.Records {
		joinRecords = joinRecords + record.Type + " " + record.Value
		if i < len(hostRecords.Records)-1 {
			joinRecords = joinRecords + " "
		}
	}

	return joinRecords
}

type Record struct {
	Type  string
	Value string
}

func LookupHosts(ctx context.Context, hosts []string) ([]HostRecords, error) {
	hostRecords := make([]HostRecords, len(hosts))

	var wg sync.WaitGroup
	wg.Add(len(hosts))
	for i, host := range hosts {
		hostRecords[i] = HostRecords{
			Host: host,
		}

		go func(i int, host string) {
			defer wg.Done()
			records, _ := LookupHost(host)
			hostRecords[i].Records = records
		}(i, host)
	}
	wg.Wait()

	return hostRecords, nil
}

func LookupHost(host string) ([]Record, error) {
	if !strings.HasSuffix(host, ".") {
		host = host + "."
	}

	records := make([]Record, 0)
	cname, err := net.LookupCNAME(host)
	if err == nil && cname != "" && cname != host {
		records = append(records, Record{
			Type:  TypeCNAME,
			Value: cname,
		})
		return records, nil
	}

	ips, err := net.LookupIP(host)
	if err == nil {
		for _, ip := range ips {
			if ip.To4() != nil {
				records = append(records, Record{
					Type:  TypeA,
					Value: ip.String(),
				})
			} else {
				records = append(records, Record{
					Type:  TypeAAAA,
					Value: ip.String(),
				})
			}
		}
	}

	return records, nil
}
