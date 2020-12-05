/*
Package opendnsmyip returns the public-facing IPv4 address
of the requesting client by querying Cisco OpenDNS servers.

Example:

	package main

	import (
		"fmt"
		myip "github.com/johnsonjh/opendnsmyip"
	)

	func main() {
		myIpAddr, err := myip.GetMyIP()
		if err != nil {
			fmt.Errorf("failure getting IP address: %s\n", err)
		} else {
			fmt.Printf("Public IPv4 address is: %s\n", myIpAddr)
		}
	}
*/
package opendnsmyip

import (
	"fmt"

	"github.com/miekg/dns"
)

// GetMyIP returns the public-facing IPv4 address of the
// caller client by querying the Cisco OpenDNS servers.
func GetMyIP() (string, error) {
	config := dns.ClientConfig{
		Servers: []string{
			"208.67.220.220",
			"208.67.222.222",
		},
		Port: "53",
	}
	dnsClient := new(dns.Client)
	message := new(dns.Msg)
	message.SetQuestion("myip.opendns.com.", dns.TypeA)
	message.RecursionDesired = false
	return MyIPDNSLookup(config, dnsClient, message)
}

// MyIPDNSLookup performs a DNS lookup using the "miekg/dns" package.
func MyIPDNSLookup(
	config dns.ClientConfig,
	client *dns.Client,
	message *dns.Msg,
) (string, error) {
	err := fmt.Errorf(
		"failure querying remote DNS server",
	)
	for _, server := range config.Servers {
		serverAddr := fmt.Sprintf(
			"%s:%s",
			server,
			config.Port,
		)
		response, _, cliErr := client.Exchange(message, serverAddr)
		if cliErr != nil {
			return "", fmt.Errorf(
				"DNS lookup failure: %w",
				cliErr,
			)
		}
		if response.Rcode != dns.RcodeSuccess {
			err = fmt.Errorf(
				"DNS request failed with response code: %d",
				response.Rcode,
			)
		} else {
			for _, answer := range response.Answer {
				if aRecord, ok := answer.(*dns.A); ok {
					return aRecord.A.String(), nil
				}
			}
		}
	}
	return "", err
}
