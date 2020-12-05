# opendnsmyip

A Go package that returns the public-facing IPv4 address of the client
by querying the Cisco OpenDNS servers.

[![License](http://img.shields.io/badge/license-mit-blue.svg)](https://raw.githubusercontent.com/johnsonjh/opendnsmyip/master/LICENSE)
[![GoReport](https://goreportcard.com/badge/github.com/johnsonjh/opendnsmyip)](https://goreportcard.com/report/github.com/johnsonjh/opendnsmyip)


## Availability

* [Gridfinity GitLab](https://gitlab.gridfinity.com/jeff/go-opendns-myip)
* [GitHub](https://github.com/johnsonjh/opendnsmyip)


## Original Authors

* [James Polera](https://github.com/polera/publicip) <james@uncryptic.com>
* [Dmitriy Kuznetsov](https://github.com/Dikman/publicip) <dmitriy.v.kuznetsov@gmail.com>


## Credits

This package was inspired by:

* [public-ip (nodejs)](https://github.com/sindresorhus/public-ip/blob/master/index.js)
* [OpenDNS::MyIP (Perl)](https://metacpan.org/pod/OpenDNS::MyIP)


## Usage

```go
package main

import (
	"fmt"
	myip "github.com/johnsonjh/opendnsmyip"
)

func main() {
	myIpAddr, err := myip.GetIP()
	if err != nil {
		fmt.Errorf("Error getting IPv4 address: %s\n", err)
	} else {
		fmt.Printf("Public IPv4 address is: %s", myIpAddr)
	}
}
```

