# opendnsmyip

[![GRC](https://goreportcard.com/badge/github.com/johnsonjh/opendnsmyip)](https://goreportcard.com/badge/github.com/johnsonjh/opendnsmyip)

A Go package that returns the public-facing IPv4 address of the client by
querying the Cisco OpenDNS servers.

## Original Authors

- [James Polera](https://github.com/polera/publicip)
  \<[james@uncryptic.com](mailto:james@uncryptic.com)\>
- [Dmitriy Kuznetsov](https://github.com/Dikman/publicip)
  \<[dmitriy.v.kuznetsov@gmail.com](mailto:dmitriy.v.kuznetsov@gmail.com)\>

## Credits

This package was inspired by:

- [public-ip (nodejs)](https://github.com/sindresorhus/public-ip/blob/master/index.js)
- [OpenDNS::MyIP (Perl 5)](https://metacpan.org/pod/OpenDNS::MyIP)

## License

- [MIT License](https://tldrlegal.com/license/mit-license)

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
                fmt.Errorf(
                           "Error getting IPv4 address: %v",
                           err,
                          )
        } else {
                fmt.Printf(
                           "Public IPv4 address is: %s",
                           myIpAddr,
                          )
        }
}
```
