// Copyright (c) 2020 Jeffrey H. Johnson.
// Copyright (c) 2020 Gridfinity, LLC.
// Copyright (c) 2020 Dmitriy Kuznetsov.
// Copyright (c) 2016 James Polera.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package opendnsmyip // import "go.gridfinity.com/opendnsmyip"

import (
	"fmt"

	"github.com/miekg/dns"

	opendnsmyipLegal "go4.org/legal"
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

// init initializes the opendnsmyip package
func init() {
	// Regiser licensing
	opendnsmyipLegal.RegisterLicense(
		"\nThe MIT License (MIT)\n\nCopyright (c) 2020 Jeffrey H. Johnson.\nCopyright (c) 2020 Gridfinity, LLC.\nCopyright (c) 2020 Dmitriy Kuznetsov.\nCopyright (c) 2016 James Polera.\n\nPermission is hereby granted, free of charge, to any person obtaining a copy\nof this software and associated documentation files (the \"Software\"), to deal\nin the Software without restriction, including without limitation the rights\nto use, copy, modify, merge, publish, distribute, sublicense, and/or sell\ncopies of the Software, and to permit persons to whom the Software is\nfurnished to do so, subject to the following conditions:\n\nThe above copyright notice and this permission notice shall be included in all\ncopies or substantial portions of the Software.\n\nTHE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\nIMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\nFITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\nAUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\nLIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\nOUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE\nSOFTWARE.\n",
	)
}
