package scout

import (
	"context"

	"github.com/emmanuelgautier/domain-scout/dns"
	"github.com/emmanuelgautier/domain-scout/http"
)

type Availability struct {
	Domain string

	Records             dns.HostRecords
	IsRootHTTPReachable *http.Reachable
}

func CheckAvailability(ctx context.Context, domains []string) ([]Availability, error) {
	availabilities := make([]Availability, len(domains))
	for i := range availabilities {
		availabilities[i] = Availability{
			Domain: domains[i],
		}
	}

	records, err := dns.LookupHosts(ctx, domains)
	if err != nil {
		return availabilities, err
	}

	for _, record := range records {
		for i, availability := range availabilities {
			if record.Host == availability.Domain {
				availabilities[i].Records = record
			}
		}
	}

	for i, availability := range availabilities {
		if len(availability.Records.Records) > 0 {
			availabilities[i].IsRootHTTPReachable, _ = http.IsAddrReachable(availability.Domain)
		}
	}

	return availabilities, nil
}
