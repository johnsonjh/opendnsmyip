// Copyright © 2021 Jeffrey H. Johnson. <trnsz@pobox.com>
// Copyright © 2021 Gridfinity, LLC.
// Copyright © 2020 Dmitriy Kuznetsov.
// Copyright © 2016 James Polera.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package opendnsmyip_test

import (
	"testing"

	"github.com/miekg/dns"
	u "github.com/johnsonjh/leaktestfe"
	myip "github.com/johnsonjh/opendnsmyip"
)

func TestGetMyIPAddress(t *testing.T) {
	defer u.Leakplug(
		t,
	)
	_, err := myip.GetMyIP()
	if err != nil {
		t.Errorf("Error: %s\n", err)
	}
}

func TestGenerateClientError(t *testing.T) {
	defer u.Leakplug(
		t,
	)
	config := dns.ClientConfig{Servers: []string{"0.0.0.0"}, Port: "53"}
	dnsClient := new(dns.Client)
	message := new(dns.Msg)
	message.SetQuestion("myip.opendns.com.", dns.TypeA)
	message.RecursionDesired = false
	_, err := myip.MyIPDNSLookup(config, dnsClient, message)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestGenerateLookupError(t *testing.T) {
	defer u.Leakplug(
		t,
	)
	config := dns.ClientConfig{
		Servers: []string{"208.67.220.220", "208.67.222.222"},
		Port:    "53",
	}
	dnsClient := new(dns.Client)
	message := new(dns.Msg)
	message.SetQuestion("notmyip.opendns.com.", dns.TypeA)
	message.RecursionDesired = false
	_, err := myip.MyIPDNSLookup(config, dnsClient, message)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
